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
	size int
)

// createVolumeCmd represents the createVolume command
var createVolumeCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a volume",
	Long: `Example:

  packet volume create --size [size_in_GB] --plan [plan_UUID] --project-id [project_UUID] --facility [facility_code]
  
  `,
	Run: func(cmd *cobra.Command, args []string) {
		req := &packngo.VolumeCreateRequest{
			BillingCycle: billingCycle,
			PlanID:       plan,
			FacilityID:   facility,
			Size:         size,
		}

		if description != "" {
			req.Description = description
		}
		if locked != false {
			req.Locked = locked
		}

		v, _, err := PacknGo.Volumes.Create(req, projectID)
		if err != nil {
			fmt.Println("Client error:", err)
			return
		}

		header := []string{"ID", "Name", "Size", "State", "Created"}
		data := make([][]string, 1)
		data[0] = []string{v.ID, v.Name, strconv.Itoa(v.Size), v.State, v.Created}

		output(v, header, &data)
	},
}

func init() {
	createVolumeCmd.Flags().StringVarP(&projectID, "project-id", "p", "", "UUID of the project")
	createVolumeCmd.Flags().StringVarP(&plan, "plan", "P", "", "Name of the plan")
	createVolumeCmd.Flags().StringVarP(&facility, "facility", "f", "", "Code of the facility where the volume will be created")
	createVolumeCmd.Flags().IntVarP(&size, "size", "s", 0, "Size in GB]")

	createVolumeCmd.MarkFlagRequired("size")
	createVolumeCmd.MarkFlagRequired("facility")
	createVolumeCmd.MarkFlagRequired("plan")
	createVolumeCmd.MarkFlagRequired("project-id")

	createVolumeCmd.Flags().StringVarP(&billingCycle, "billing-cycle", "b", "hourly", "Billing cycle")
	createVolumeCmd.Flags().StringVarP(&description, "description", "d", "", "Description of the volume")
	createVolumeCmd.Flags().BoolVarP(&locked, "locked", "l", false, "Set the volume to be locked")

}
