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

func newGatewayish(gwaysInner metal.MetalGatewayListMetalGatewaysInner) gatewayIsh {
	if gwaysInner.MetalGateway != nil {
		return metalGatewayWrapper{gwaysInner.MetalGateway}
	}
	if gwaysInner.VrfMetalGateway != nil {
		return vrfMetalGatewayWrapper{gwaysInner.VrfMetalGateway}
	}
	return nil // Return nil if neither type is present
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
	var projectID string

	// retrieveMetalGatewaysCmd represents the retrieveMetalGateways command
	retrieveMetalGatewaysCmd := &cobra.Command{
		Use:     `get -p <project_UUID>`,
		Aliases: []string{"list"},
		Short:   "Lists Metal Gateways.",
		Long:    "Retrieves a list of all VLANs for the specified project.",
		Example: `
  # Lists Metal Gateways for project 3b0795ba-ec9a-4a9e-83a7-043e7e11407c:
  metal gateways get -p 3b0795ba-ec9a-4a9e-83a7-043e7e11407c`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			includes := []string{"virtual_network", "ip_reservation", "vrf"}

			gwayList, _, err := c.Service.
				FindMetalGatewaysByProject(context.Background(), projectID).
				Include(c.Servicer.Includes(includes)).
				Exclude(c.Servicer.Excludes(nil)).
				Execute()
			if err != nil {
				return fmt.Errorf("Could not list Project Metal Gateways: %w", err)
			}

			gways := gwayList.GetMetalGateways()

			data := make([][]string, 0, len(gways))
			var gateways []gatewayIsh

			for _, gwaysInner := range gways {
				gway := newGatewayish(gwaysInner)
				if gway == nil {
					continue
				}

				gateways = append(gateways, gway)

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

	retrieveMetalGatewaysCmd.Flags().StringVarP(&projectID, "project-id", "p", "", "The project's UUID. This flag is required, unless specified in the config created by metal init or set as METAL_PROJECT_ID environment variable.")
	_ = retrieveMetalGatewaysCmd.MarkFlagRequired("project-id")

	return retrieveMetalGatewaysCmd
}
