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

package ips

import (
	"fmt"
	"strconv"

	"github.com/packethost/packngo"
	"github.com/spf13/cobra"
)

func (c *Client) Assign() *cobra.Command {
	var (
		address  string
		deviceID string
	)

	// assignIpCmd represents the assignIp command
	assignIPCmd := &cobra.Command{
		Use:   `assign -a <IP_address> -d <device_UUID>`,
		Short: "Assigns an IP address to a specified device.",
		Long:  "Assigns an IP address and subnet to a specified device. Returns an assignment ID.",
		Example: `  # Assigns an IP address to a server:
  metal ip assign -d 060d1626-2481-475a-9789-c6f4bb927303  -a 198.51.100.3/31`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			assignment, _, err := c.DeviceService.Assign(deviceID, &packngo.AddressStruct{Address: address})
			if err != nil {
				return fmt.Errorf("Could not assign Device IP address: %w", err)
			}

			data := make([][]string, 1)

			data[0] = []string{assignment.ID, assignment.Address, strconv.FormatBool(assignment.Public), assignment.Created}
			header := []string{"ID", "Address", "Public", "Created"}

			return c.Out.Output(assignment, header, &data)
		},
	}

	assignIPCmd.Flags().StringVarP(&deviceID, "device-id", "d", "", "The UUID of the device.")
	assignIPCmd.Flags().StringVarP(&address, "address", "a", "", "IP address and CIDR you would like to assign.")

	_ = assignIPCmd.MarkFlagRequired("device-id")
	_ = assignIPCmd.MarkFlagRequired("address")
	return assignIPCmd
}
