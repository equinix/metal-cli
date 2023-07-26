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

	metal "github.com/equinix-labs/metal-go/metal/v1"
	"github.com/spf13/cobra"
)

func (c *Client) Create() *cobra.Command {
	var projectID, vnID, reservationID string
	var netSize int

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

			nSize := int32(netSize)
			gatewayRequestInput := &metal.MetalGatewayCreateInput{
				VirtualNetworkId:      vnID,
				IpReservationId:       &reservationID,
				PrivateIpv4SubnetSize: &nSize,
			}

			gatewayRequest := metal.CreateMetalGatewayRequest{MetalGatewayCreateInput: gatewayRequestInput}
			n, _, err := c.Service.CreateMetalGateway(context.Background(), projectID).CreateMetalGatewayRequest(gatewayRequest).Execute()
			if err != nil {
				return fmt.Errorf("Could not create Metal Gateway: %w", err)
			}

			data := make([][]string, 1)
			address := ""

			if n.MetalGateway.IpReservation != nil {
				address = n.MetalGateway.IpReservation.GetAddress() + "/" + strconv.Itoa(int(n.MetalGateway.IpReservation.GetCidr()))
			}

			data[0] = []string{n.MetalGateway.GetId(), n.MetalGateway.VirtualNetwork.GetMetroCode(), strconv.Itoa(int(n.MetalGateway.VirtualNetwork.GetVxlan())), address, string(n.MetalGateway.GetState()), n.MetalGateway.CreatedAt.String()}

			header := []string{"ID", "Metro", "VXLAN", "Addresses", "State", "Created"}

			return c.Out.Output(n, header, &data)
		},
	}

	createMetalGatewayCmd.Flags().StringVarP(&projectID, "project-id", "p", "", "The project's UUID. This flag is required, unless specified in the config created by metal init or set as METAL_PROJECT_ID environment variable.")
	createMetalGatewayCmd.Flags().StringVarP(&reservationID, "ip-reservation-id", "r", "", "UUID of the Public or VRF IP Reservation to assign.")
	createMetalGatewayCmd.Flags().StringVarP(&vnID, "virtual-network", "v", "", "UUID of the Virtual Network to assign.")
	createMetalGatewayCmd.Flags().IntVarP(&netSize, "private-subnet-size", "s", 0, "Size of the private subnet to request (8 for /29)")

	_ = createMetalGatewayCmd.MarkFlagRequired("project-id")
	_ = createMetalGatewayCmd.MarkFlagRequired("virtual-network")
	return createMetalGatewayCmd
}
