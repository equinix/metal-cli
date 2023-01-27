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
	"os"
	"time"

	"github.com/packethost/packngo"
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
		//IPAddress string

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

	createDeviceCmd := &cobra.Command{
		Use:   `create -p <project_id> (-m <metro> | -f <facility>) -P <plan> -H <hostname> -O <operating_system> [-u <userdata> | --userdata-file <filepath>] [-c <customdata>] [-t <tags>] [-r <hardware_reservation_id>] [-I <ipxe_script_url>] [--always-pxe] [--spot-instance] [--spot-price-max=<max_price>]`,
		Short: "Creates a device.",
		Long:  "Creates a device in the specified project. A plan, hostname, operating system, and either metro or facility is required.",
		Example: `  # Provisions a c3.small.x86 in the Dallas metro running Ubuntu 20.04:
  metal device create -p $METAL_PROJECT_ID -P c3.small.x86 -m da -H test-staging-2 -O ubuntu_20_04

  # Provisions a c3.medium.x86 in Silicon Valley, running Rocky Linux, from a hardware reservation:
  metal device create -p $METAL_PROJECT_ID -P c3.medium.x86 -m sv -H test-rocky -O rocky_8 -r 47161704-1715-4b45-8549-fb3f4b2c32c7`,

		RunE: func(cmd *cobra.Command, args []string) error {
			var endDt *packngo.Timestamp

			if userdata != "" && userdataFile != "" {
				return fmt.Errorf("either userdata or userdata-file should be set")
			}
			cmd.SilenceUsage = true

			if userdataFile != "" {
				userdataRaw, readErr := os.ReadFile(userdataFile)
				if readErr != nil {
					return fmt.Errorf("Could not read userdata-file: %w", readErr)
				}
				userdata = string(userdataRaw)
			}

			if terminationTime != "" {
				parsedTime, err := time.Parse(time.RFC3339, terminationTime)
				if err != nil {
					return fmt.Errorf("Could not parse time %q: %w", terminationTime, err)
				}
				endDt = &packngo.Timestamp{Time: parsedTime}
			}

			var facilityArgs []string
			if facility != "" {
				facilityArgs = append(facilityArgs, facility)
			}

			request := &packngo.DeviceCreateRequest{
				Hostname: hostname,
				Plan:     plan,
				Facility: facilityArgs,
				Metro:    metro,
				OS:       operatingSystem,
				//IPAddresses:           ipaddress,
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
				return fmt.Errorf("Could not create Device: %w", err)
			}

			header := []string{"ID", "Hostname", "OS", "State", "Created"}
			data := make([][]string, 1)
			data[0] = []string{device.ID, device.Hostname, device.OS.Name, device.State, device.Created}

			return c.Out.Output(device, header, &data)
		},
	}
	createDeviceCmd.Flags().StringVarP(&projectID, "project-id", "p", "", "The project's UUID. This flag is required, unless specified in the config created by metal init or set as METAL_PROJECT_ID environment variable.")
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

	createDeviceCmd.Flags().StringVarP(&ipxescripturl, "ipxe-script-url", "I", "", "The URL of an iPXE script.")
	createDeviceCmd.Flags().StringVarP(&userdata, "userdata", "u", "", "Userdata for device initialization. Can not be used with --userdata-file.")
	createDeviceCmd.Flags().StringVarP(&userdataFile, "userdata-file", "", "", "Path to a userdata file for device initialization. Can not be used with --userdata.")
	createDeviceCmd.Flags().StringVarP(&customdata, "customdata", "c", "", "Custom data to be included with your device's metadata.")
	createDeviceCmd.Flags().StringSliceVarP(&tags, "tags", "t", []string{}, `Tag or list of tags for the device: --tags="tag1,tag2".`)
	createDeviceCmd.Flags().IntVarP(&publicIPv4SubnetSize, "public-ipv4-subnet-size", "S", 0, "Size of the public IPv4 subnet.")
	createDeviceCmd.Flags().StringVarP(&hardwareReservationID, "hardware-reservation-id", "r", "", "The UUID of a hardware reservation, if you are provisioning a server from your reserved hardware.")
	createDeviceCmd.Flags().StringVarP(&billingCycle, "billing-cycle", "b", "hourly", "Billing cycle")
	createDeviceCmd.Flags().BoolVarP(&alwaysPXE, "always-pxe", "a", false, "Sets whether the device always PXE boots on reboot.")
	createDeviceCmd.Flags().BoolVarP(&spotInstance, "spot-instance", "s", false, "Provisions the device as a spot instance.")
	createDeviceCmd.Flags().Float64VarP(&spotPriceMax, "spot-price-max", "", 0, `Sets the maximum spot market price for the device: --spot-price-max=1.2`)
	createDeviceCmd.Flags().StringVarP(&terminationTime, "termination-time", "T", "", `Device termination time: --termination-time="15:04:05"`)

	return createDeviceCmd
}
