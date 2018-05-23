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
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var force bool
var deleteDeviceCmd = &cobra.Command{
	Use:   "device",
	Short: "Deletes a device",
	// Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if !force {
			reader := bufio.NewReader(os.Stdin)
			fmt.Printf("Are you sure you want to delete device %s: ", deviceID)
			text, _ := reader.ReadString('\n')
			text = strings.TrimSuffix(text, "\n")
			if text == "" {
				err := deleteDevice(deviceID)
				if err != nil {
					fmt.Println("Client error:", err)
					return
				}
			} else if strings.ToUpper(text) == "N" || strings.ToUpper(text) == "NO" {
				return
			} else if strings.ToUpper(text) == "Y" || strings.ToUpper(text) == "YES" {
				err := deleteDevice(deviceID)
				if err != nil {
					fmt.Println("Client error:", err)
					return
				}
			}
		} else {
			err := deleteDevice(deviceID)
			if err != nil {
				fmt.Println("Client error:", err)
				return
			}
		}
	},
}

func deleteDevice(id string) error {
	_, err := PacknGo.Devices.Delete(id)
	if err != nil {
		return err
	}
	fmt.Println("Device deletion initiated. Please check 'packet get device", deviceID, "' for status")
	return nil
}
func init() {
	deleteDeviceCmd.Flags().StringVarP(&deviceID, "id", "i", "", "--id or -i [UUID]")
	deleteDeviceCmd.Flags().BoolVarP(&force, "force", "f", false, "--force or -f")
	deleteDeviceCmd.MarkFlagRequired("id")
}
