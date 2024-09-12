// Copyright © 2018 Jasmin Gacic <jasmin@stackpointcloud.com>
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

package vlan

import (
	"context"
	"fmt"
	"strconv"

	metal "github.com/equinix/equinix-sdk-go/services/metalv1"
	"github.com/spf13/cobra"
)

func (c *Client) Create() *cobra.Command {
	var vxlan int
	var projectID, metro, facility, description string

	// createVirtualNetworkCmd represents the createVirtualNetwork command
	createVirtualNetworkCmd := &cobra.Command{
		Use:   `create -p <project_UUID>  [-m <metro_code> -vxlan <vlan> | -f <facility_code>] [-d <description>]`,
		Short: "Creates a virtual network.",
		Long:  "Creates a VLAN in the specified project. If you are creating a VLAN in a metro, you can optionally specify the VXLAN ID otherwise it is auto-assigned. If you are creating a VLAN in a facility, the VXLAN ID is auto-assigned.",
		Example: `  # Creates a VLAN with vxlan ID 1999 in the Dallas metro:
  metal virtual-network create -p $METAL_PROJECT_ID -m da --vxlan 1999

  # Creates a VLAN in the sjc1 facility
  metal virtual-network create -p $METAL_PROJECT_ID -f sjc1`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			virtualNetworkCreateInput := metal.NewVirtualNetworkCreateInput()
			if metro != "" {
				virtualNetworkCreateInput.SetMetro(metro)
			}
			if facility != "" {
				virtualNetworkCreateInput.SetFacility(facility)
			}
			virtualNetworkCreateInput.SetVxlan(int32(vxlan))

			if description != "" {
				virtualNetworkCreateInput.Description = &description
			}

			n, _, err := c.Service.CreateVirtualNetwork(context.Background(), projectID).VirtualNetworkCreateInput(*virtualNetworkCreateInput).Execute()
			if err != nil {
				return fmt.Errorf("could not create ProjectVirtualNetwork: %w", err)
			}

			data := make([][]string, 1)

			// TODO(displague) metro is not in the response
			data[0] = []string{n.GetId(), n.GetDescription(), strconv.Itoa(int(n.GetVxlan())), n.GetMetroCode(), n.CreatedAt.String()}

			header := []string{"ID", "Description", "VXLAN", "Metro", "Created"}

			return c.Out.Output(n, header, &data)
		},
	}

	createVirtualNetworkCmd.Flags().StringVarP(&projectID, "project-id", "p", "", "The project's UUID. This flag is required, unless specified in the config created by metal init or set as METAL_PROJECT_ID environment variable.")
	createVirtualNetworkCmd.Flags().StringVarP(&facility, "facility", "f", "", "Code of the facility.")
	createVirtualNetworkCmd.Flags().StringVarP(&metro, "metro", "m", "", "Code of the metro.")
	createVirtualNetworkCmd.Flags().StringVarP(&description, "description", "d", "", "A user-friendly description of the virtual network.")
	createVirtualNetworkCmd.Flags().IntVarP(&vxlan, "vxlan", "", 0, "Optional VXLAN ID. Must be between 2 and 3999 and can only be used with --metro.")

	_ = createVirtualNetworkCmd.MarkFlagRequired("project-id")
	return createVirtualNetworkCmd
}
