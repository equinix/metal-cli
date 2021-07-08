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

package hardware

import (
	"fmt"

	"github.com/equinix/metal-cli/internal/outputs"
	"github.com/packethost/packngo"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) Retrieve() *cobra.Command {
	var retrieveHardwareReservationsCmd = &cobra.Command{
		Use:     "get",
		Aliases: []string{"list"},
		Short:   "Retrieves all hardware reservations of a project or a single hardware reservation",
		Long: `Example:

Retrieve all hardware reservations of a project:
metal hardware_reservations get -p [project_id]

When using "--json" or "--yaml", "--include=project,facility,device" is implied.
	`,

		RunE: func(cmd *cobra.Command, args []string) error {
			projectID, _ := cmd.Flags().GetString("project-id")
			hardwareReservationID, _ := cmd.Flags().GetString("id")

			header := []string{"ID", "Facility", "Plan", "Created"}

			inc := []string{}

			// only fetch extra details when rendered
			switch c.Servicer.Format() {
			case outputs.FormatJSON, outputs.FormatYAML:
				inc = append(inc, "project", "facility", "device")
			}

			listOpt := c.Servicer.ListOptions(inc, nil)

			if hardwareReservationID == "" && projectID == "" {
				return fmt.Errorf("Either id or project-id should be set.")
			} else if hardwareReservationID != "" && projectID != "" {
				return fmt.Errorf("Either id or project-id can be set.")
			} else if projectID != "" {
				reservations, _, err := c.Service.List(projectID, listOpt)
				if err != nil {
					return errors.Wrap(err, "Could not list Hardware Reservations")
				}

				data := make([][]string, len(reservations))

				for i, r := range reservations {
					data[i] = []string{r.ID, r.Facility.Code, r.Plan.Name, r.CreatedAt.String()}
				}

				return c.Out.Output(reservations, header, &data)
			} else if hardwareReservationID != "" {
				getOpts := &packngo.GetOptions{Includes: listOpt.Includes, Excludes: listOpt.Excludes}
				r, _, err := c.Service.Get(hardwareReservationID, getOpts)
				if err != nil {
					return errors.Wrap(err, "Could not get Hardware Reservation")
				}

				data := make([][]string, 1)

				data[0] = []string{r.ID, r.Facility.Code, r.Plan.Name, r.CreatedAt.String()}

				return c.Out.Output(r, header, &data)
			}
			return nil
		},
	}

	retrieveHardwareReservationsCmd.Flags().StringP("project-id", "p", "", "Project ID (METAL_PROJECT_ID)")
	retrieveHardwareReservationsCmd.Flags().StringP("id", "i", "", "UUID of the hardware reservation")

	return retrieveHardwareReservationsCmd
}
