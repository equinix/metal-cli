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
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func (c *Client) Retrieve() *cobra.Command {
	var projectID string

	// retrieveVirtualNetworksCmd represents the retrieveVirtualNetworks command
	retrieveVirtualNetworksCmd := &cobra.Command{
		Use:     `get -p <project_UUID>`,
		Aliases: []string{"list"},
		Short:   "Lists virtual networks.",
		Long:    "Retrieves a list of all VLANs for the specified project.",
		Example: `  # Lists virtual networks for project 3b0795ba-ec9a-4a9e-83a7-043e7e11407c:
  metal virtual-network get -p 3b0795ba-ec9a-4a9e-83a7-043e7e11407c`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			request := c.Service.FindVirtualNetworks(cmd.Context(), projectID).Include(c.Servicer.Includes(nil)).Exclude(c.Servicer.Excludes(nil))
			filters := c.Servicer.Filters()
			if filters["facility"] != "" {
				request = request.Facility(filters["facility"])
			}

			if filters["metro"] != "" {
				request = request.Metro(filters["metro"])
			}
			virtualNetworksList, _, err := request.Execute()
			if err != nil {
				return fmt.Errorf("Could not list Project Virtual Networks: %w", err)
			}
			virtualNetworks := virtualNetworksList.GetVirtualNetworks()
			data := make([][]string, len(virtualNetworks))

			for i, n := range virtualNetworks {
				data[i] = []string{n.GetId(), n.GetDescription(), strconv.Itoa(int(n.GetVxlan())), n.GetMetroCode(), n.GetCreatedAt().String()}
			}
			header := []string{"ID", "Description", "VXLAN", "Metro", "Created"}

			return c.Out.Output(virtualNetworksList, header, &data)
		},
	}
	retrieveVirtualNetworksCmd.Flags().StringVarP(&projectID, "project-id", "p", "", "The project's UUID. This flag is required, unless specified in the config created by metal init or set as METAL_PROJECT_ID environment variable.")
	_ = retrieveVirtualNetworksCmd.MarkFlagRequired("project-id")

	return retrieveVirtualNetworksCmd
}
