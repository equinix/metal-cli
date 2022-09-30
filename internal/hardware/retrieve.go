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
	"github.com/spf13/cobra"
)

func (c *Client) Retrieve() *cobra.Command {
	retrieveHardwareReservationsCmd := &cobra.Command{
		Use:     `get [-p <project_id>] | [-i <hardware_reservation_id>]`,
		Aliases: []string{"list"},
		Short:   "Lists a Project's hardware reservations or the details of a specified hardware reservation.",
		Long:    "Lists a Project's hardware reservations or the details of a specified hardware reservation. When using --json or --yaml flags, the --include=project,facility,device flag is implied.",
		Example: `  # Retrieve all hardware reservations of a project:
  metal hardware_reservations get -p $METAL_PROJECT_ID
  
  # Retrieve the details of a specific hardware reservation:
  metal hardware_reservations get -i 8404b73c-d18f-4190-8c49-20bb17501f88`,

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
				return fmt.Errorf("either id or project-id should be set")
			}

			cmd.SilenceUsage = true
			if hardwareReservationID != "" {
				getOpts := &packngo.GetOptions{Includes: listOpt.Includes, Excludes: listOpt.Excludes}
				r, _, err := c.Service.Get(hardwareReservationID, getOpts)
				if err != nil {
					return fmt.Errorf("Could not get Hardware Reservation: %w", err)
				}

				data := make([][]string, 1)

				data[0] = []string{r.ID, r.Facility.Code, r.Plan.Name, r.CreatedAt.String()}

				return c.Out.Output(r, header, &data)
			}

			reservations, _, err := c.Service.List(projectID, listOpt)
			if err != nil {
				return fmt.Errorf("Could not list Hardware Reservations: %w", err)
			}

			data := make([][]string, len(reservations))

			for i, r := range reservations {
				data[i] = []string{r.ID, r.Facility.Code, r.Plan.Name, r.CreatedAt.String()}
			}

			return c.Out.Output(reservations, header, &data)
		},
	}

	retrieveHardwareReservationsCmd.Flags().StringP("project-id", "p", "", "A project's UUID.")
	retrieveHardwareReservationsCmd.Flags().StringP("id", "i", "", "The UUID of a hardware reservation.")

	return retrieveHardwareReservationsCmd
}
