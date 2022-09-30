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

package ports

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/packethost/packngo"
	"github.com/spf13/cobra"
)

func (c *Client) Vlans() *cobra.Command {
	var portID, native string
	var assignments, unassignments []string

	// retrievePortCmd represents the retrievePort command
	retrievePortCmd := &cobra.Command{
		Use:     `vlan -i <port_UUID> [--native <vlan>] [--unassign <vlan>]... [--assign <vlan>]...`,
		Aliases: []string{"vlans"},
		Short:   "Modifies VLAN assignments on a port",
		Long:    "Modifies the VLANs of the specified port to the desired state. Existing state can be restated without error.",
		Example: `  # Assigns VLANs 1234 and 5678 to the port:
  metal port vlans -i 3bd5bf07-6094-48ad-bd03-d94e8712fdc8 -a 1234 -a 5678

  # Unassigns VXLAN 1234 from the port:
  metal port vlans -i 3bd5bf07-6094-48ad-bd03-d94e8712fdc8 -u 1234

  # Assigns VXLAN 1234 to the port and makes it the Native VLAN:
  metal port vlans -i 3bd5bf07-6094-48ad-bd03-d94e8712fdc8 --native=1234`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			listOpts := c.Servicer.ListOptions([]string{"port"}, nil)

			getOpts := &packngo.GetOptions{Includes: listOpts.Includes, Excludes: listOpts.Excludes}
			req := &packngo.VLANAssignmentBatchCreateRequest{}
			f := false
			t := true
			for _, vlan := range assignments {
				assignment := packngo.VLANAssignmentCreateRequest{VLAN: vlan, State: packngo.VLANAssignmentAssigned, Native: &f}
				req.VLANAssignments = append(req.VLANAssignments, assignment)
			}
			for _, vlan := range unassignments {
				assignment := packngo.VLANAssignmentCreateRequest{VLAN: vlan, State: packngo.VLANAssignmentUnassigned}
				req.VLANAssignments = append(req.VLANAssignments, assignment)
			}
			if native != "" {
				assignment := packngo.VLANAssignmentCreateRequest{VLAN: native, State: packngo.VLANAssignmentAssigned, Native: &t}
				req.VLANAssignments = append(req.VLANAssignments, assignment)
			}
			if len(req.VLANAssignments) == 0 {
				return errors.New("no VLAN assignments specified")
			}
			batch, _, err := c.VLANService.CreateBatch(portID, req, getOpts)
			if err != nil {
				return fmt.Errorf("Could not update port VLAN assignments: %w", err)
			}

			// TODO: should we return the batch?
			port := batch.Port

			data := make([][]string, 1)

			data[0] = []string{port.ID, port.Name, port.Type, port.NetworkType, port.Data.MAC, strconv.FormatBool(port.Data.Bonded)}
			header := []string{"ID", "Name", "Type", "Network Type", "MAC", "Bonded"}

			return c.Out.Output(port, header, &data)
		},
	}

	retrievePortCmd.Flags().StringVarP(&portID, "port-id", "i", "", "The UUID of a port.")
	retrievePortCmd.Flags().StringVarP(&native, "native", "n", "", "The VXLAN to make assign as the Native VLAN")

	retrievePortCmd.Flags().StringSliceVarP(&assignments, "assign", "a", []string{}, "A VXLAN to assign to the port. May also be used to change a Native VLAN assignment to tagged (non-native).")
	retrievePortCmd.Flags().StringSliceVarP(&unassignments, "unassign", "u", []string{}, "A VXLAN to unassign from a port.")

	_ = retrievePortCmd.MarkFlagRequired("port-id")

	return retrievePortCmd
}
