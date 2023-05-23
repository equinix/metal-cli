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

func (c *Client) Reinstall() *cobra.Command {
	var (
		id string

		operatingSystem string
		deprovisionFast bool
		preserveData    bool
	)

	reinstallDeviceCmd := &cobra.Command{
		Use:   `reinstall --id <device-id> [--operating-system <os_slug>] [--deprovision-fast] [--preserve-data]`,
		Short: "Reinstalls a device.",
		Long:  "Reinstalls the provided device with the current operating system or a new operating system with optional flags to preserve data or skip disk clean-up. The ID of the device to reinstall is required.",
		Example: `  # Reinstalls a device with the current OS:
  metal device reinstall -d 50382f72-02b7-4b40-ac8d-253713e1e174
  
  # Reinstalls a device with Ubuntu 22.04 while preserving the data on non-OS disks:
  metal device reinstall -d 50382f72-02b7-4b40-ac8d-253713e1e174 -O ubuntu_22_04 --preserve-data`,
		RunE: func(cmd *cobra.Command, args []string) error {

			DeviceAction := metal.NewDeviceActionInput("reinstall")

			if preserveData {
				DeviceAction.PreserveData = &preserveData
			}
			if deprovisionFast {
				DeviceAction.DeprovisionFast = &deprovisionFast
			}
			device, _, Err := c.Service.FindDeviceById(context.Background(), id).Execute()
			if Err != nil {
				fmt.Printf("Error when calling `DevicesApiService.FindDeviceByID``: %v\n", Err)
				Err = fmt.Errorf("Could not reinstall Device: %w", Err)
				return Err
			}

			if operatingSystem == "" || operatingSystem == "null" {
				fmt.Print("operatingSystem : ")
				fmt.Print(operatingSystem)
				fmt.Print("\n")
				DeviceAction.SetOperatingSystem(device.OperatingSystem.GetSlug())
			} else {
				DeviceAction.OperatingSystem = &operatingSystem
			}

			_, err := c.Service.PerformAction(context.Background(), id).DeviceActionInput(*DeviceAction).Execute()
			if err != nil {
				fmt.Printf("Error when calling `DevicesApiService.PerformAction``: %v\n", err)
				err = fmt.Errorf("Could not reinstall Device: %w", err)
			}
			return err
		},
	}
	reinstallDeviceCmd.Flags().StringVarP(&id, "id", "d", "", "ID of device to be reinstalled")
	_ = reinstallDeviceCmd.MarkFlagRequired("id")

	reinstallDeviceCmd.Flags().StringVarP(&operatingSystem, "operating-system", "O", "", "Operating system install on the device. If omitted the current OS will be reinstalled.")
	reinstallDeviceCmd.Flags().BoolVarP(&deprovisionFast, "deprovision-fast", "", false, "Avoid optional potentially slow clean-up tasks.")
	reinstallDeviceCmd.Flags().BoolVarP(&preserveData, "preserve-data", "", false, "Avoid wiping data on disks where the OS is *not* being installed.")

	return reinstallDeviceCmd
}
