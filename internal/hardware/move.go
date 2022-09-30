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

	"github.com/spf13/cobra"
)

func (c *Client) Move() *cobra.Command {
	var projectID, hardwareReservationID string

	moveHardwareReservationCmd := &cobra.Command{
		Use:   `move -i <hardware_reservation_id> -p <project_id>`,
		Short: "Moves a hardware reservation.",
		Long:  "Moves a hardware reservation to a specified project. Both the hardware reservation ID and the Project ID for the destination project are required.",
		Example: `  # Moves a hardware reservation to the specified Project:
  metal hardware_reservation move -i 8404b73c-d18f-4190-8c49-20bb17501f88 -p 278bca90-f6b2-4659-b1a4-1bdffa0d80b7`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			header := []string{"ID", "Facility", "Plan", "Created"}
			r, _, err := c.Service.Move(hardwareReservationID, projectID)
			if err != nil {
				return fmt.Errorf("Could not move Hardware Reservation: %w", err)
			}

			data := make([][]string, 1)

			data[0] = []string{r.ID, r.Facility.Code, r.Plan.Name, r.CreatedAt.String()}

			return c.Out.Output(r, header, &data)
		},
	}

	moveHardwareReservationCmd.Flags().StringVarP(&hardwareReservationID, "id", "i", "", "The UUID of the hardware reservation.")
	moveHardwareReservationCmd.Flags().StringVarP(&projectID, "project-id", "p", "", "The Project ID of the Project you are moving the hardware reservation to.")
	_ = moveHardwareReservationCmd.MarkFlagRequired("project-id")
	_ = moveHardwareReservationCmd.MarkFlagRequired("id")

	return moveHardwareReservationCmd
}
