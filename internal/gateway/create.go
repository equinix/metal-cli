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
	"errors"
	"fmt"
	"strconv"

	metal "github.com/equinix-labs/metal-go/metal/v1"
	"github.com/spf13/cobra"
)

func (c *Client) Create() *cobra.Command {
	var projectID, vnID, reservationID string
	var netSize int32

	// createMetalGatewayCmd represents the createMetalGateway command
	createMetalGatewayCmd := &cobra.Command{
		Use:   `create -p <project_UUID> --virtual-network <virtual_network_UUID> (--ip-reservation-id <ip_reservation_UUID> | --private-subnet-size <size>)`,
		Short: "Creates a Metal Gateway.",
		Long:  "Creates a Metal Gateway on the VLAN. Either an IP Reservation ID or a Private Subnet Size must be specified.",
		Example: `  # Creates a Metal Gateway on the VLAN with a given IP Reservation ID:
  metal gateway create -p $METAL_PROJECT_ID -v 77e6d57a-d7a4-4816-b451-cf9b043444e2 -r 50052f72-02b7-4b40-ac9d-253713e1e178

  # Creates a Metal Gateway on the VLAN with a Private 10.x.x.x/28 subnet:
  metal gateway create -p $METAL_PROJECT_ID --virtual-network 77e6d57a-d7a4-4816-b451-cf9b043444e2 --private-subnet-size 16`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			includes := []string{"virtual_network", "ip_reservation"}

			if reservationID == "" && netSize == 0 {
				return errors.New("Invalid input. Provide either 'private-subnet-size' or 'ip-reservation-id'")
			}

			req := metal.CreateMetalGatewayRequest{
				MetalGatewayCreateInput: &metal.MetalGatewayCreateInput{
					VirtualNetworkId: vnID,
				},
			}
			if reservationID != "" {
				req.MetalGatewayCreateInput.SetIpReservationId(reservationID)
			} else {
				req.MetalGatewayCreateInput.SetPrivateIpv4SubnetSize(netSize)
			}

			n, _, err := c.Service.
				CreateMetalGateway(context.Background(), projectID).
				Include(c.Servicer.Includes(includes)).
				Exclude(c.Servicer.Excludes(nil)).
				CreateMetalGatewayRequest(req).
				Execute()
			if err != nil {
				return fmt.Errorf("Could not create Metal Gateway: %w", err)
			}

			data := make([][]string, 1)
			address := ""

			gway := n.MetalGateway
			ipReservation := gway.IpReservation
			if ipReservation != nil {
				address = ipReservation.GetAddress() + "/" + strconv.Itoa(int(ipReservation.GetCidr()))
			}

			data[0] = []string{gway.GetId(), gway.VirtualNetwork.GetMetroCode(),
				strconv.Itoa(int(gway.VirtualNetwork.GetVxlan())), address, string(gway.GetState()), gway.GetCreatedAt().String()}

			header := []string{"ID", "Metro", "VXLAN", "Addresses", "State", "Created"}

			return c.Out.Output(gway, header, &data)
		},
	}

	createMetalGatewayCmd.Flags().StringVarP(&projectID, "project-id", "p", "", "The project's UUID. This flag is required, unless specified in the config created by metal init or set as METAL_PROJECT_ID environment variable.")
	createMetalGatewayCmd.Flags().StringVarP(&reservationID, "ip-reservation-id", "r", "", "UUID of the Public or VRF IP Reservation to assign.")
	createMetalGatewayCmd.Flags().StringVarP(&vnID, "virtual-network", "v", "", "UUID of the Virtual Network to assign.")
	createMetalGatewayCmd.Flags().Int32VarP(&netSize, "private-subnet-size", "s", 0, "Size of the private subnet to request (8 for /29)")

	_ = createMetalGatewayCmd.MarkFlagRequired("project-id")
	_ = createMetalGatewayCmd.MarkFlagRequired("virtual-network")
	return createMetalGatewayCmd
}
