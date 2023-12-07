// Copyright © 2022 Equinix Metal Developers <support@equinixmetal.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package ports

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	metal "github.com/equinix-labs/metal-go/metal/v1"
	"github.com/rivo/tview"
	"github.com/spf13/cobra"
)

func (c *Client) Convert() *cobra.Command {
	var portID string
	var bonded, layer2, bulk, force, ipv4, ipv6 bool
	convertPortCmd := &cobra.Command{
		Use:     `convert -i <port_UUID> [--bonded] [--bulk] --layer2 [--force] [--public-ipv4] [--public-ipv6]`,
		Aliases: []string{},
		Short:   "Converts a list of ports or the details of the specified port.",
		Long:    "Converts a list of ports or the details of the specified port. Details of an port are only available to its members.",
		Example: `  # Converts list of the current user's ports:
  metal port convert -i <port_UUID> [--bonded] [--bulk] [--layer2] [--force] [--public-ipv4] [--public-ipv6]

  # Converts port to layer-2 unbonded:
  metal port convert -i 3bd5bf07-6094-48ad-bd03-d94e8712fdc8 --layer2 --bonded=false

  # Converts port to layer-2 bonded:
  metal port convert -i 3bd5bf07-6094-48ad-bd03-d94e8712fdc8 --layer2 --bonded

  # Converts port to layer-3 bonded with public IPv4 and public IPv6:
  metal port convert -i 3bd5bf07-6094-48ad-bd03-d94e8712fdc8 -2=false -b -4 -6`,
		// TODO: can we add ip-reservation-id?
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true

			if cmd.Flag("bonded").Changed {
				if err := portBondingHandler(bonded, c, portID); err != nil {
					return fmt.Errorf("failed to change port bonding: %w", err)
				}
			}

			convToL2 := func(portID string) (*metal.Port, *http.Response, error) {
				var (
					port *metal.Port
					resp *http.Response
					rerr error
				)
				if !force {
					app := tview.NewApplication()

					confirmModal := tview.NewModal().
						SetText(fmt.Sprintf("Are you sure you want to convert Port %s to Layer2 and remove assigned IP addresses?", portID)).
						AddButtons([]string{"Yes", "No"}).
						SetDoneFunc(func(buttonIndex int, buttonLabel string) {
							if buttonLabel == "Yes" {
								port, resp, rerr = c.PortService.ConvertLayer2(context.Background(), portID).PortAssignInput(*metal.NewPortAssignInput()).Execute()
							}
							app.Stop()
						})

					if err := app.SetRoot(confirmModal, false).Run(); err != nil {
						return nil, nil, nil
					}
				}
				return port, resp, rerr
			}

			addressFamily := int32(metal.IPADDRESSADDRESSFAMILY__4)
			public := false
			addrs := []metal.PortConvertLayer3InputRequestIpsInner{{AddressFamily: &addressFamily, Public: &public}}

			if f := cmd.Flag("public-ipv4"); f.Changed {
				addressFamily = int32(metal.IPADDRESSADDRESSFAMILY__4)
				public = true
				addrs = append(addrs, metal.PortConvertLayer3InputRequestIpsInner{AddressFamily: &addressFamily, Public: &public})
			}
			if f := cmd.Flag("public-ipv6"); f.Changed {
				addressFamily = int32(metal.IPADDRESSADDRESSFAMILY__6)
				public = true
				addrs = append(addrs, metal.PortConvertLayer3InputRequestIpsInner{AddressFamily: &addressFamily, Public: &public})
			}

			convToL3 := func(portID string) (*metal.Port, *http.Response, error) {
				log.Printf("Converting port %s to layer-3", portID)
				return c.PortService.
					ConvertLayer3(context.Background(), portID).
					PortConvertLayer3Input(metal.PortConvertLayer3Input{
						RequestIps: addrs,
					}).
					Execute()
			}
			if f := cmd.Flag("layer2"); f.Changed {
				_, _, err := map[bool]func(string) (*metal.Port, *http.Response, error){
					true:  convToL2,
					false: convToL3,
				}[layer2](portID)
				if err != nil {
					return fmt.Errorf("failed to change port network mode: %w", err)
				}
			}

			port, _, err := c.PortService.FindPortById(context.Background(), portID).
				Include(c.Servicer.Includes(nil)).
				Execute()
			if err != nil {
				return fmt.Errorf("Could not get Port: %w", err)
			}

			data := make([][]string, 1)

			data[0] = []string{port.GetId(), port.GetName(), string(port.GetType()), string(port.GetNetworkType()), port.Data.GetMac(), strconv.FormatBool(port.Data.GetBonded())}
			header := []string{"ID", "Name", "Type", "Network Type", "MAC", "Bonded"}

			return c.Out.Output(port, header, &data)
		},
	}

	convertPortCmd.Flags().StringVarP(&portID, "port-id", "i", "", "The UUID of a port.")
	convertPortCmd.Flags().BoolVarP(&bonded, "bonded", "b", false, "Convert to layer-2 bonded.")
	convertPortCmd.Flags().BoolVarP(&bulk, "bulk", "", false, "Affect both ports in a bond.")
	convertPortCmd.Flags().BoolVarP(&layer2, "layer2", "2", false, "Convert to layer-2 unbonded.")
	convertPortCmd.Flags().BoolVarP(&force, "force", "f", false, "Force conversion to layer-2 bonded.")
	convertPortCmd.Flags().BoolVarP(&ipv4, "public-ipv4", "4", false, "Convert to layer-2 bonded with public IPv4.")
	convertPortCmd.Flags().BoolVarP(&ipv6, "public-ipv6", "6", false, "Convert to layer-2 bonded with public IPv6.")

	return convertPortCmd
}

func portBondingHandler(bonded bool, c *Client, portId string) error {
	if bonded {
		_, _, err := c.PortService.BondPort(context.Background(), portId).
			Include(c.Servicer.Includes(nil)).
			Execute()
		return err
	}

	_, _, err := c.PortService.DisbondPort(context.Background(), portId).
		Include(c.Servicer.Includes(nil)).
		Execute()
	return err
}
