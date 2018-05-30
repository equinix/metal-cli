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
	"fmt"
	"strconv"

	"github.com/packethost/packngo"
	"github.com/spf13/cobra"
)

var (
	vlan  int
	vxlan int
)

// createVirtualNetworkCmd represents the createVirtualNetwork command
var createVirtualNetworkCmd = &cobra.Command{
	Use:   "virtual-network",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		req := &packngo.VirtualNetworkCreateRequest{
			ProjectID: projectID,
			Facility:  facility,
		}

		if vxlan != 0 {
			req.VXLAN = vxlan
		}

		if vlan != 0 {
			req.VLAN = vlan
		}

		if description != "" {
			req.Description = description
		}

		n, _, err := PacknGo.ProjectVirtualNetworks.Create(req)
		if err != nil {
			fmt.Println("Client error:", err)
		}

		data := make([][]string, 1)

		data[0] = []string{n.ID, n.Description, strconv.Itoa(n.VXLAN), n.FacilityCode, n.CreatedAt}

		header := []string{"ID", "Description", "VXLAN", "Facility", "Created"}

		output(n, header, &data)
	},
}

func init() {
	createCmd.AddCommand(createVirtualNetworkCmd)
	createVirtualNetworkCmd.Flags().StringVarP(&projectID, "project-id", "p", "", "--project-id or -i [UUID]")
	createVirtualNetworkCmd.Flags().StringVarP(&facility, "facility", "f", "", "--facility or -f [facility_code]")
	createVirtualNetworkCmd.Flags().IntVarP(&vlan, "vlan", "l", 0, "--vlan or -l [vlan]")
	createVirtualNetworkCmd.Flags().IntVarP(&vxlan, "vxlan", "x", 0, "--vxlan or -x [vxlan]")
	createVirtualNetworkCmd.Flags().StringVarP(&description, "description", "d", "", "--description or -d [description]")

	createVirtualNetworkCmd.MarkFlagRequired("project-id")
	createVirtualNetworkCmd.MarkFlagRequired("facility")
}
