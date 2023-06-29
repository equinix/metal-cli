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
	"fmt"

	metal "github.com/equinix-labs/metal-go/metal/v1"
	"github.com/spf13/cobra"
)

func (c *Client) Reboot() *cobra.Command {
	var deviceID string

	rebootDeviceCmd := &cobra.Command{
		Use:   `reboot -i <device_id>`,
		Short: "Reboots a device.",
		Long:  "Reboots the specified device.",
		Example: `  # Reboots the specified device:
  metal device reboot --id 26a9da5f-a0db-41f6-8467-827e144e59a7`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			DeviceAction := metal.NewDeviceActionInput("reboot")
			_, err := c.Service.PerformAction(context.Background(), deviceID).DeviceActionInput(*DeviceAction).Execute()
			if err != nil {
				return fmt.Errorf("Could not reboot Device: %w", err)
			}

			fmt.Println("Device", deviceID, "successfully rebooted.")
			return nil
		},
	}

	rebootDeviceCmd.Flags().StringVarP(&deviceID, "id", "i", "", "The device UUID.")

	_ = rebootDeviceCmd.MarkFlagRequired("id")
	return rebootDeviceCmd
}
