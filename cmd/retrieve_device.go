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

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	organizationID string
	deviceID       string
)

var retriveDeviceCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieves device list or device details",
	Long: `Example:
	
packet device get --id [device_UUID]

	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// projectID, _ = cmd.Flags().GetString("x")
		fmt.Println(projectID, "viper", viper.GetString("project-id"))
		panic("as")

		if deviceID != "" && projectID != "" {
			return fmt.Errorf("Either id or project-id can be set.")
		} else if deviceID == "" && projectID == "" {
			return fmt.Errorf("Either id or project-id should be set.")
		} else if projectID != "" {
			devices, _, err := PacknGo.Devices.List(projectID, listOptions(nil, nil))
			if err != nil {
				return errors.Wrap(err, "Could not list Devices")
			}
			data := make([][]string, len(devices))

			for i, dc := range devices {
				data[i] = []string{dc.ID, dc.Hostname, dc.OS.Name, dc.State, dc.Created}
			}
			header := []string{"ID", "Hostname", "OS", "State", "Created"}

			return output(devices, header, &data)
		} else if deviceID != "" {
			device, _, err := PacknGo.Devices.Get(deviceID, nil)
			if err != nil {
				return errors.Wrap(err, "Could not get Devices")
			}
			header := []string{"ID", "Hostname", "OS", "State", "Created"}

			data := make([][]string, 1)
			data[0] = []string{device.ID, device.Hostname, device.OS.Name, device.State, device.Created}

			return output(device, header, &data)
		}
		return nil
	},
}

func init() {
	retriveDeviceCmd.Flags().StringVarP(&projectID, "project-id", "p", "", "UUID of the project")
	viper.BindPFlag("project-id", retriveDeviceCmd.Flags().Lookup("project-id"))

	retriveDeviceCmd.Flags().StringVarP(&deviceID, "id", "i", "", "UUID of the device")
}
