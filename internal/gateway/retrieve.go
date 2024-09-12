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

package gateway

import (
	"context"
	"fmt"
	"strconv"
	"time"

	metal "github.com/equinix/equinix-sdk-go/services/metalv1"

	"github.com/spf13/cobra"
)

type ipreservationIsh interface {
	GetAddress() string
	GetCidr() int32
}

var (
	_ ipreservationIsh = (*metal.VrfIpReservation)(nil)
	_ ipreservationIsh = (*metal.IPReservation)(nil)
)

type gatewayIsh interface {
	GetId() string
	GetCreatedAt() time.Time
	GetState() metal.MetalGatewayState
	GetVirtualNetwork() metal.VirtualNetwork
	GetIpReservation() ipreservationIsh
	GetVrfName() string
}

var (
	_ gatewayIsh = (*metalGatewayWrapper)(nil)
	_ gatewayIsh = (*vrfMetalGatewayWrapper)(nil)
)

type metalGatewayWrapper struct {
	*metal.MetalGateway
}

func (m metalGatewayWrapper) GetIpReservation() ipreservationIsh {
	if m.MetalGateway.IpReservation == nil {
		return nil
	}
	return m.MetalGateway.IpReservation
}

func (m metalGatewayWrapper) GetVrfName() string {
	return ""
}

type vrfMetalGatewayWrapper struct {
	*metal.VrfMetalGateway
}

func newGatewayishFromAny(gateways ...any) gatewayIsh {
	for _, gateway := range gateways {
		switch g := gateway.(type) {
		case *metal.MetalGateway:
			if g == nil {
				continue
			}
			return metalGatewayWrapper{g}
		case *metal.VrfMetalGateway:
			if g == nil {
				continue
			}
			return vrfMetalGatewayWrapper{g}
		default:
			panic(fmt.Sprintf("unexpected gateway type: %T", gateway))
		}
	}

	return nil
}

func (v vrfMetalGatewayWrapper) GetIpReservation() ipreservationIsh {
	if v.VrfMetalGateway.IpReservation == nil {
		return nil
	}
	return v.VrfMetalGateway.IpReservation
}

func (v vrfMetalGatewayWrapper) GetVrfName() string {
	vrf := v.GetVrf()
	return vrf.GetName()
}

func (c *Client) Retrieve() *cobra.Command {
	// retrieveMetalGatewaysCmd represents the retrieveMetalGateways command
	retrieveMetalGatewaysCmd := &cobra.Command{
		Use:     `get [-p <project_UUID>] [-i <gateway_UUID>]`,
		Aliases: []string{"list"},
		Short:   "Lists Metal Gateways.",
		Long:    "Retrieves a list of all VLANs for the specified project.",
		Example: `
  # Lists Metal Gateways for project 3b0795ba-ec9a-4a9e-83a7-043e7e11407c:
  metal gateways get -p 3b0795ba-ec9a-4a9e-83a7-043e7e11407c
  
  # Gets a Metal Gateway with ID 8404b73c-d18f-4190-8c49-20bb17501f88:
  metal gateways get -i 8404b73c-d18f-4190-8c49-20bb17501f88`,

		RunE: func(cmd *cobra.Command, args []string) error {
			gatewayID, _ := cmd.Flags().GetString("id")
			projectID, _ := cmd.Flags().GetString("project-id")

			if gatewayID == "" && projectID == "" {
				return fmt.Errorf("either id or project-id should be set")
			}

			cmd.SilenceUsage = true
			includes := []string{"virtual_network", "ip_reservation", "vrf"}
			gateways := make([]gatewayIsh, 0)

			if gatewayID != "" {
				gateway, _, err := c.Service.FindMetalGatewayById(context.Background(), gatewayID).
					Include(c.Servicer.Includes(includes)).
					Exclude(c.Servicer.Excludes(nil)).
					Execute()
				if err != nil {
					return fmt.Errorf("could not get Metal Gateway: %w", err)
				}
				gateways = append(gateways, newGatewayishFromAny(gateway.GetActualInstance()))
			} else {
				gwayList, _, err := c.Service.
					FindMetalGatewaysByProject(context.Background(), projectID).
					Include(c.Servicer.Includes(includes)).
					Exclude(c.Servicer.Excludes(nil)).
					Execute()
				if err != nil {
					return fmt.Errorf("could not list Project Metal Gateways: %w", err)
				}

				gways := gwayList.GetMetalGateways()

				for _, gwaysInner := range gways {
					gway := newGatewayishFromAny(gwaysInner.MetalGateway, gwaysInner.VrfMetalGateway)
					if gway == nil {
						continue
					}
					gateways = append(gateways, gway)
				}
			}

			data := make([][]string, 0, len(gateways))

			for _, gway := range gateways {
				address := ""
				ipReservation := gway.GetIpReservation()
				if ipReservation != nil {
					address = ipReservation.GetAddress() + "/" + strconv.Itoa(int(ipReservation.GetCidr()))
				}

				vnet := gway.GetVirtualNetwork()
				metroCode := vnet.GetMetroCode()
				vxlan := strconv.Itoa(int(vnet.GetVxlan()))

				data = append(data, []string{
					gway.GetId(),
					metroCode,
					vxlan,
					address,
					string(gway.GetState()),
					gway.GetCreatedAt().String(),
					gway.GetVrfName(),
				})
			}

			header := []string{"ID", "Metro", "VXLAN", "Addresses", "State", "Created", "VRF"}

			return c.Out.Output(gateways, header, &data)
		},
	}

	retrieveMetalGatewaysCmd.Flags().StringP("project-id", "p", "", "The project's UUID. This flag is required, unless specified in the config created by metal init or set as METAL_PROJECT_ID environment variable.")
	retrieveMetalGatewaysCmd.Flags().StringP("id", "i", "", "The UUID of a gateway.")

	return retrieveMetalGatewaysCmd
}
