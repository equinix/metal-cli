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
	"strconv"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/MakeNowJust/heredoc"
)

func (c *Client) Retrieve() *cobra.Command {
	var projectID string

	// retrieveVirtualNetworksCmd represents the retrieveVirtualNetworks command
	var retrieveVirtualNetworksCmd = &cobra.Command{
		Use: `get -p <project_UUID> [global_options]`,
		Aliases: []string{"list"},
		Short:   "Retrieves a list of virtual networks for the specified project.",
		Example: heredoc.Doc(`
			# Lists virtual networks for project 3b0795ba-ec9a-4a9e-83a7-043e7e11407c:
			virtual-network get -p 3b0795ba-ec9a-4a9e-83a7-043e7e11407c
		`),
		
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			vnets, _, err := c.Service.List(projectID, c.Servicer.ListOptions(nil, nil))
			if err != nil {
				return errors.Wrap(err, "Could not list Project Virtual Networks")
			}

			data := make([][]string, len(vnets.VirtualNetworks))

			for i, n := range vnets.VirtualNetworks {
				data[i] = []string{n.ID, n.Description, strconv.Itoa(n.VXLAN), n.FacilityCode, n.CreatedAt}
			}
			header := []string{"ID", "Description", "VXLAN", "Facility", "Created"}

			return c.Out.Output(vnets, header, &data)
		},
	}
	retrieveVirtualNetworksCmd.Flags().StringVarP(&projectID, "project-id", "p", "", "Project ID (METAL_PROJECT_ID)")
	_ = retrieveVirtualNetworksCmd.MarkFlagRequired("project-id")

	return retrieveVirtualNetworksCmd
}
