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

func (c *Client) Stop() *cobra.Command {
	var deviceID string
	stopDeviceCmd := &cobra.Command{
		Use:   `stop -i <device_id>`,
		Short: "Stops a device.",
		Long:  "Stops or powers off a device that is currently started or powered on.",
		Example: `  # Stops the specified device:
  metal device stop --id [device_UUID]`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			DeviceAction := metal.NewDeviceActionInput("power_off")
			_, err := c.Service.PerformAction(context.Background(), deviceID).DeviceActionInput(*DeviceAction).Execute()
			if err != nil {
				return fmt.Errorf("Could not stop Device: %w", err)
			}

			fmt.Println("Device", deviceID, "successfully stopped.")
			return nil
		},
	}

	stopDeviceCmd.Flags().StringVarP(&deviceID, "id", "i", "", "The UUID of the device.")
	_ = stopDeviceCmd.MarkFlagRequired("id")
	return stopDeviceCmd
}
