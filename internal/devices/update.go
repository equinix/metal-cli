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
	"context"
	"encoding/json"
	"fmt"

	metal "github.com/equinix-labs/metal-go/metal/v1"

	"github.com/spf13/cobra"
)

func (c *Client) Update() *cobra.Command {
	var (
		description   string
		locked        bool
		userdata      string
		hostname      string
		tags          []string
		alwaysPXE     bool
		ipxescripturl string
		customdata    string
		deviceID      string
	)

	// updateDeviceCmd represents the updateDevice command
	updateDeviceCmd := &cobra.Command{
		Use:   `update -i <device_id> [-H <hostname>] [-d <description>] [--locked <boolean>] [-t <tags>] [-u <userdata>] [-c <customdata>] [-s <ipxe_script_url>] [--always-pxe]`,
		Short: "Updates a device.",
		Long:  "Updates the hostname of a device. Updates or adds a description, tags, userdata, custom data, and iPXE settings for an already provisioned device. Can also lock or unlock future changes to the device.",
		Example: `  # Updates the hostname of a device:
  metal device update --id 30c15082-a06e-4c43-bfc3-252616b46eba --hostname renamed-staging04`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			deviceUpdate := metal.NewDeviceUpdateInput()

			if hostname != "" {
				deviceUpdate.Hostname = &hostname
			}

			if description != "" {
				deviceUpdate.Description = &description
			}

			if userdata != "" {
				deviceUpdate.Userdata = &userdata
			}

			if locked {
				deviceUpdate.Locked = &locked
			}

			if len(tags) > 0 {
				deviceUpdate.Tags = tags
			}

			if alwaysPXE {
				deviceUpdate.AlwaysPxe = &alwaysPXE
			}

			if ipxescripturl != "" {
				deviceUpdate.IpxeScriptUrl = &ipxescripturl
			}

			if customdata != "" {
				var customdataIntr map[string]interface{}
				err := json.Unmarshal([]byte(customdata), &customdataIntr)
				if err != nil {
					panic(err)
				}

				deviceUpdate.Customdata = customdataIntr
			}
			device, _, err := c.Service.UpdateDevice(context.Background(), deviceID).DeviceUpdateInput(*deviceUpdate).Execute()
			if err != nil {
				return fmt.Errorf("Could not update Device: %w", err)
			}

			header := []string{"ID", "Hostname", "OS", "State"}
			data := make([][]string, 1)
			data[0] = []string{device.GetId(), device.GetHostname(), device.OperatingSystem.GetName(), device.GetState()}

			return c.Out.Output(device, header, &data)
		},
	}

	updateDeviceCmd.Flags().StringVarP(&deviceID, "id", "i", "", "The UUID of the device.")
	updateDeviceCmd.Flags().StringVarP(&hostname, "hostname", "H", "", "The new hostname of the device.")
	updateDeviceCmd.Flags().StringVarP(&description, "description", "d", "", "Adds or updates the description for the device.")
	updateDeviceCmd.Flags().StringVarP(&userdata, "userdata", "u", "", "Adds or updates the userdata for the device.")
	updateDeviceCmd.Flags().BoolVarP(&locked, "locked", "l", false, "Locks or unlocks the device for future changes.")
	updateDeviceCmd.Flags().StringSliceVarP(&tags, "tags", "t", []string{}, `Adds or updates the tags for the device --tags="tag1,tag2".`)
	updateDeviceCmd.Flags().BoolVarP(&alwaysPXE, "always-pxe", "a", false, "Sets the device to always iPXE on reboot.")
	updateDeviceCmd.Flags().StringVarP(&ipxescripturl, "ipxe-script-url", "s", "", "Add or update the URL of the iPXE script.")
	updateDeviceCmd.Flags().StringVarP(&customdata, "customdata", "c", "", "Adds or updates custom data to be included with your device's metadata.")

	_ = updateDeviceCmd.MarkFlagRequired("id")
	return updateDeviceCmd
}
