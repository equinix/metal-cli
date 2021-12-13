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

package devices

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) Retrieve() *cobra.Command {
	var retrieveDeviceCmd = &cobra.Command{
		Use:     "get",
		Aliases: []string{"list"},
		Short:   "Retrieves device list or device details",

		Long: `Example:
	
metal device get --id [device_UUID]

	`,
		RunE: func(cmd *cobra.Command, args []string) error {
			deviceID, _ := cmd.Flags().GetString("id")
			projectID, _ := cmd.Flags().GetString("project-id")

			if deviceID == "" && projectID == "" {
				return fmt.Errorf("Either id or project-id should be set.")
			}
			cmd.SilenceUsage = true

			if deviceID != "" {
				device, _, err := c.Service.Get(deviceID, nil)
				if err != nil {
					return errors.Wrap(err, "Could not get Devices")
				}
				header := []string{"ID", "Hostname", "OS", "State", "Created"}

				data := make([][]string, 1)
				data[0] = []string{device.ID, device.Hostname, device.OS.Name, device.State, device.Created}

				return c.Out.Output(device, header, &data)
			}

			devices, _, err := c.Service.List(projectID, c.Servicer.ListOptions(nil, nil))
			if err != nil {
				return errors.Wrap(err, "Could not list Devices")
			}
			data := make([][]string, len(devices))

			for i, dc := range devices {
				data[i] = []string{dc.ID, dc.Hostname, dc.OS.Name, dc.State, dc.Created}
			}
			header := []string{"ID", "Hostname", "OS", "State", "Created"}

			return c.Out.Output(devices, header, &data)
		},
	}

	retrieveDeviceCmd.Flags().StringP("project-id", "p", "", "Project ID (METAL_PROJECT_ID)")
	retrieveDeviceCmd.Flags().StringP("id", "i", "", "UUID of the device")

	return retrieveDeviceCmd
}
