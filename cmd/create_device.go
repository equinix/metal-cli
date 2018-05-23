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

var (
	projectID       string
	facility        string
	plan            string
	hostname        string
	operatingSystem string
	billingCycle    string
)

var createDeviceCmd = &cobra.Command{
	Use:   "device",
	Short: "Create a device",
	// Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		request := &packngo.DeviceCreateRequest{
			HostName:     hostname,
			ProjectID:    projectID,
			Facility:     facility,
			Plan:         plan,
			OS:           operatingSystem,
			BillingCycle: billingCycle,
		}

		device, _, err := PacknGo.Devices.Create(request)
		if err != nil {
			fmt.Println("Client error:", err)
			return
		}

		header := []string{"ID", "Hostname", "OS", "State"}
		data := make([][]string, 1)
		data[0] = []string{device.ID, device.Hostname, device.OS.Name, device.State}

		output(device, header, &data)

	},
}

func init() {
	createDeviceCmd.Flags().StringVarP(&projectID, "project-id", "p", "", "--project-id or -P [UUID]")
	createDeviceCmd.Flags().StringVarP(&facility, "facility", "f", "", "--facility or -f [facility_code]")
	createDeviceCmd.Flags().StringVarP(&plan, "plan", "P", "", "--plan or -p [plan_name]")
	createDeviceCmd.Flags().StringVarP(&hostname, "hostname", "H", "", "--hostname or -H [hostname]")
	createDeviceCmd.Flags().StringVarP(&operatingSystem, "operating-system", "o", "", "--operating-system or -o [operating_system]")
	createDeviceCmd.Flags().StringVarP(&billingCycle, "billing-cycle", "b", "hourly", "--billing-cycle or -b [billing_cycle]")
	createDeviceCmd.MarkFlagRequired("project-id")
	createDeviceCmd.MarkFlagRequired("facility")
	createDeviceCmd.MarkFlagRequired("plan")
	createDeviceCmd.MarkFlagRequired("hostname")
	createDeviceCmd.MarkFlagRequired("operating-system")
}
