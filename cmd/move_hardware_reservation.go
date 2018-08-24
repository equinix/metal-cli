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

	"github.com/spf13/cobra"
)

// moveHardwareReservationCmd represents the moveHardwareReservation command
var moveHardwareReservationCmd = &cobra.Command{
	Use:   "move",
	Short: "Move hardware reservation to another project",
	Long: `Example:

packet hardware_reservation move -i [hardware_reservation_UUID] -p [project_UUID]
`,
	Run: func(cmd *cobra.Command, args []string) {
		header := []string{"ID", "Facility", "Plan", "Created"}
		r, _, err := PacknGo.HardwareReservations.Move(hardwareReservationID, projectID)
		if err != nil {
			fmt.Println("Client error:", err)
			return
		}

		data := make([][]string, 1)

		data[0] = []string{r.ID, r.Facility.Code, r.Plan.Name, r.CreatedAt.String()}

		output(r, header, &data)
	},
}

func init() {
	hardwareReservationsCmd.AddCommand(moveHardwareReservationCmd)
	moveHardwareReservationCmd.Flags().StringVarP(&hardwareReservationID, "id", "i", "", "UUID of the hardware reservation")
	moveHardwareReservationCmd.Flags().StringVarP(&projectID, "project-id", "p", "", "UUID of the project")
	moveHardwareReservationCmd.MarkFlagRequired("project-id")
	moveHardwareReservationCmd.MarkFlagRequired("id")
}
