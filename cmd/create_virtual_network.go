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

package cmd

import (
	"strconv"

	"github.com/packethost/packngo"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// createVirtualNetworkCmd represents the createVirtualNetwork command
var createVirtualNetworkCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a virtual network",
	Long: `Example:

packet virtual-network create --project-id [project_UUID] --facility [facility_code]

`,
	RunE: func(cmd *cobra.Command, args []string) error {
		req := &packngo.VirtualNetworkCreateRequest{
			ProjectID: projectID,
			Facility:  facility,
		}
		if description != "" {
			req.Description = description
		}

		n, _, err := PacknGo.ProjectVirtualNetworks.Create(req)
		if err != nil {
			return errors.Wrap(err, "Could not create ProjectVirtualNetwork")
		}

		data := make([][]string, 1)

		data[0] = []string{n.ID, n.Description, strconv.Itoa(n.VXLAN), n.FacilityCode, n.CreatedAt}

		header := []string{"ID", "Description", "VXLAN", "Facility", "Created"}

		return output(n, header, &data)
	},
}

func init() {
	createVirtualNetworkCmd.Flags().StringVarP(&projectID, "project-id", "p", "", "UUID of the project")
	createVirtualNetworkCmd.Flags().StringVarP(&facility, "facility", "f", "", "Code of the facility")
	createVirtualNetworkCmd.Flags().StringVarP(&description, "description", "d", "", "Description of the virtual network")

	_ = createVirtualNetworkCmd.MarkFlagRequired("project-id")
	_ = createVirtualNetworkCmd.MarkFlagRequired("facility")
	createVirtualNetworkCmd.PersistentFlags().BoolVarP(&isJSON, "json", "j", false, "JSON output")
	createVirtualNetworkCmd.PersistentFlags().BoolVarP(&isYaml, "yaml", "y", false, "YAML output")
}
