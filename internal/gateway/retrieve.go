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

	"github.com/spf13/cobra"
)

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
  metal virtual-network get -p 3b0795ba-ec9a-4a9e-83a7-043e7e11407c`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			inc := []string{"virtual_network", "ip_reservation"}
			exc := []string{}
			gwaysList, err := c.Service.FindMetalGatewaysByProject(context.Background(), projectID).Include(c.Servicer.Includes(inc)).Exclude(c.Servicer.Excludes(exc)).ExecuteWithPagination()
			if err != nil {
				return fmt.Errorf("Could not list Project Metal Gateways: %w", err)
			}
			gways := gwaysList.GetMetalGateways()
			data := make([][]string, len(gways))

			for i, n := range gways {
				address := ""

				if n.MetalGateway.IpReservation != nil {
					address = n.MetalGateway.IpReservation.GetAddress() + "/" + strconv.Itoa(int(n.MetalGateway.IpReservation.GetCidr()))
				}

				data[i] = []string{n.MetalGateway.GetId(), n.MetalGateway.VirtualNetwork.GetMetroCode(), strconv.Itoa(int(n.MetalGateway.VirtualNetwork.GetVxlan())), address, string(n.MetalGateway.GetState()), n.MetalGateway.CreatedAt.String()}
			}
			header := []string{"ID", "Metro", "VXLAN", "Addresses", "State", "Created"}

			return c.Out.Output(gways, header, &data)
		},
	}
	retrieveMetalGatewaysCmd.Flags().StringVarP(&projectID, "project-id", "p", "", "The project's UUID. This flag is required, unless specified in the config created by metal init or set as METAL_PROJECT_ID environment variable.")
	_ = retrieveMetalGatewaysCmd.MarkFlagRequired("project-id")

	return retrieveMetalGatewaysCmd
}
