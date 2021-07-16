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
	"io/ioutil"
	"time"

	"github.com/packethost/packngo"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func (c *Client) Create() *cobra.Command {

	var (
		projectID       string
		metro           string
		facility        string
		plan            string
		hostname        string
		operatingSystem string
		billingCycle    string

		userdata              string
		userdataFile          string
		customdata            string
		tags                  []string
		ipxescripturl         string
		publicIPv4SubnetSize  int
		alwaysPXE             bool
		hardwareReservationID string
		spotInstance          bool
		spotPriceMax          float64
		terminationTime       string
	)

	var createDeviceCmd = &cobra.Command{
		Use:   "create",
		Short: "Creates a device",
		Long: `Example:

metal device create --hostname [hostname] --plan [plan] --metro [metro_code] --facility [facility_code] --operating-system [operating_system] --project-id [project_UUID]

`,
		RunE: func(cmd *cobra.Command, args []string) error {
			var endDt *packngo.Timestamp

			if userdata != "" && userdataFile != "" {
				return fmt.Errorf("Either userdata or userdata-file should be set")
			}

			if userdataFile != "" {
				userdataRaw, readErr := ioutil.ReadFile(userdataFile)
				if readErr != nil {
					return errors.Wrap(readErr, "Could not read userdata-file")
				}
				userdata = string(userdataRaw)
			}

			if terminationTime != "" {
				parsedTime, err := time.Parse(time.RFC3339, terminationTime)
				if err != nil {
					return errors.Wrap(err, fmt.Sprintf("Could not parse time %q", terminationTime))
				}
				endDt = &packngo.Timestamp{Time: parsedTime}
			}

			var facilityArgs []string
			if facility != "" {
				facilityArgs = append(facilityArgs, facility)
			}

			request := &packngo.DeviceCreateRequest{
				Hostname:              hostname,
				Plan:                  plan,
				Facility:              facilityArgs,
				Metro:                 metro,
				OS:                    operatingSystem,
				BillingCycle:          billingCycle,
				ProjectID:             projectID,
				UserData:              userdata,
				CustomData:            customdata,
				IPXEScriptURL:         ipxescripturl,
				Tags:                  tags,
				PublicIPv4SubnetSize:  publicIPv4SubnetSize,
				AlwaysPXE:             alwaysPXE,
				HardwareReservationID: hardwareReservationID,
				SpotInstance:          spotInstance,
				SpotPriceMax:          spotPriceMax,
				TerminationTime:       endDt,
			}

			device, _, err := c.Service.Create(request)
			if err != nil {
				return errors.Wrap(err, "Could not create Device")
			}

			header := []string{"ID", "Hostname", "OS", "State", "Created"}
			data := make([][]string, 1)
			data[0] = []string{device.ID, device.Hostname, device.OS.Name, device.State, device.Created}

			return c.Out.Output(device, header, &data)
		},
	}
	createDeviceCmd.Flags().StringVarP(&projectID, "project-id", "p", "", "Project ID (METAL_PROJECT_ID) where the device will be created")
	_ = viper.BindPFlag("project-id", createDeviceCmd.Flags().Lookup("project-id"))

	createDeviceCmd.Flags().StringVarP(&facility, "facility", "f", "", "Code of the facility where the device will be created")
	createDeviceCmd.Flags().StringVarP(&metro, "metro", "m", "", "Code of the metro where the device will be created")
	createDeviceCmd.Flags().StringVarP(&plan, "plan", "P", "", "Name of the plan")
	createDeviceCmd.Flags().StringVarP(&hostname, "hostname", "H", "", "Hostname")
	createDeviceCmd.Flags().StringVarP(&operatingSystem, "operating-system", "O", "", "Operating system name for the device")
	_ = createDeviceCmd.MarkFlagRequired("project-id")
	_ = createDeviceCmd.MarkFlagRequired("plan")
	_ = createDeviceCmd.MarkFlagRequired("hostname")
	_ = createDeviceCmd.MarkFlagRequired("operating-system")

	createDeviceCmd.Flags().StringVarP(&ipxescripturl, "ipxe-script-url", "i", "", "URL to the iPXE script")
	createDeviceCmd.Flags().StringVarP(&userdata, "userdata", "u", "", "Userdata for device initialization (can not be used with --userdata-file)")
	createDeviceCmd.Flags().StringVarP(&userdataFile, "userdata-file", "", "", "Path to a userdata file for device initialization (can not be used with --userdata)")
	createDeviceCmd.Flags().StringVarP(&customdata, "customdata", "c", "", "Custom data")
	createDeviceCmd.Flags().StringSliceVarP(&tags, "tags", "t", []string{}, `Tags for the device: --tags="tag1,tag2"`)
	createDeviceCmd.Flags().IntVarP(&publicIPv4SubnetSize, "public-ipv4-subnet-size", "S", 0, "Size of the public IPv4 subnet")
	createDeviceCmd.Flags().StringVarP(&hardwareReservationID, "hardware-reservation-id", "r", "", "UUID of the hardware reservation")
	createDeviceCmd.Flags().StringVarP(&billingCycle, "billing-cycle", "b", "hourly", "Billing cycle")
	createDeviceCmd.Flags().BoolVarP(&alwaysPXE, "always-pxe", "a", false, ``)
	createDeviceCmd.Flags().BoolVarP(&spotInstance, "spot-instance", "I", false, `Set the device as a spot instance`)
	createDeviceCmd.Flags().Float64VarP(&spotPriceMax, "spot-price-max", "", 0, `--spot-price-max=1.2 or -m=1.2`)
	createDeviceCmd.Flags().StringVarP(&terminationTime, "termination-time", "T", "", `Device termination time: --termination-time="15:04:05"`)

	return createDeviceCmd
}
