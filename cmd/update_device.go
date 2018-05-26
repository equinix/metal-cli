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
	description string
	locked      bool
)

// updateDeviceCmd represents the updateDevice command
var updateDeviceCmd = &cobra.Command{
	Use:   "device",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		req := &packngo.DeviceUpdateRequest{}

		if hostname != "" {
			req.Hostname = &hostname
		}

		if description != "" {
			req.Description = &description
		}

		if userdata != "" {
			req.UserData = &userdata
		}

		if locked != false {
			req.Locked = &locked
		}

		if len(tags) > 0 {
			req.Tags = &tags
		}

		if alwaysPXE != false {
			req.AlwaysPXE = &alwaysPXE
		}

		if ipxescripturl != "" {
			req.IPXEScriptURL = &ipxescripturl
		}

		if customdata != "" {
			req.CustomData = &customdata
		}

		device, _, err := PacknGo.Devices.Update(deviceID, req)
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
	updateCmd.AddCommand(updateDeviceCmd)
	updateDeviceCmd.Flags().StringVarP(&deviceID, "id", "i", "", "--id or -i [UUID]")
	updateDeviceCmd.MarkFlagRequired("id")

	updateDeviceCmd.Flags().StringVarP(&hostname, "hostname", "H", "", "--hostname or -H [hostname]")
	updateDeviceCmd.Flags().StringVarP(&description, "description", "d", "", "--description or -d [description]")
	updateDeviceCmd.Flags().StringVarP(&userdata, "userdata", "u", "", "--userdata or -u [userdata]")
	updateDeviceCmd.Flags().BoolVarP(&locked, "locked", "l", false, "--locked or -l")
	updateDeviceCmd.Flags().StringSliceVarP(&tags, "tags", "t", []string{}, `--tags="tag1,tag2" or -t="tag1,tag2"`)
	updateDeviceCmd.Flags().BoolVarP(&alwaysPXE, "always-pxe", "a", false, `--alaways-pxe or -a`)
	updateDeviceCmd.Flags().StringVarP(&ipxescripturl, "ipxe-script-url", "s", "", "--ipxe-script-url or -i [ipxe_script_url]")
	updateDeviceCmd.Flags().StringVarP(&customdata, "customdata", "c", "", "--customdata or -c [customdata]")

}
