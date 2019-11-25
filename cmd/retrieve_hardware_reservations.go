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

	"github.com/packethost/packngo"
	"github.com/spf13/cobra"
)

// retrieveHardwareReservationsCmd represents the retrieveHardwareReservations command
var retrieveHardwareReservationsCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieves all hardware reservations of a project or a single hardware reservation",
	Long: `Example:

Retrieve all hardware reservations of a project:
packet hardware_reservations get -p [project_id]
	`,
	Run: func(cmd *cobra.Command, args []string) {
		header := []string{"ID", "Facility", "Plan", "Created"}
		listOpt := &packngo.ListOptions{Includes: []string{"project,facility,device"}}

		if hardwareReservationID == "" && projectID == "" {
			fmt.Println("Either id or project-id should be set.")
			return
		} else if hardwareReservationID != "" && projectID != "" {
			fmt.Println("Either id or project-id can be set.")
			return
		} else if projectID != "" {
			reservations, _, err := PacknGo.HardwareReservations.List(projectID, listOpt)
			if err != nil {
				fmt.Println("Client error:", err)
				return
			}

			data := make([][]string, len(reservations))

			for i, r := range reservations {
				data[i] = []string{r.ID, r.Facility.Code, r.Plan.Name, r.CreatedAt.String()}
			}

			output(reservations, header, &data)
		} else if hardwareReservationID != "" {
			getOpt := &packngo.GetOptions{Includes: listOpt.Includes}
			r, _, err := PacknGo.HardwareReservations.Get(hardwareReservationID, getOpt)
			if err != nil {
				fmt.Println("Client error:", err)
				return
			}

			data := make([][]string, 1)

			data[0] = []string{r.ID, r.Facility.Code, r.Plan.Name, r.CreatedAt.String()}

			output(r, header, &data)
		}
	},
}

func init() {
	hardwareReservationsCmd.AddCommand(retrieveHardwareReservationsCmd)
	retrieveHardwareReservationsCmd.Flags().StringVarP(&projectID, "project-id", "p", "", "UUID of the project")
	retrieveHardwareReservationsCmd.Flags().StringVarP(&hardwareReservationID, "id", "i", "", "UUID of the hardware reservation")
	retrieveHardwareReservationsCmd.PersistentFlags().BoolVarP(&isJSON, "json", "j", false, "JSON output")
	retrieveHardwareReservationsCmd.PersistentFlags().BoolVarP(&isYaml, "yaml", "y", false, "YAML output")
}
