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
	"strconv"

	"github.com/packethost/packngo"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) Assign() *cobra.Command {
	var (
		address  string
		deviceID string
	)

	// assignIpCmd represents the assignIp command
	var assignIPCmd = &cobra.Command{
		Use:   "assign",
		Short: "Assigns an IP address to a given device",
		Long: `Example:

metal ip assign -d [device-id] -a [ip-address]

	`,
		RunE: func(cmd *cobra.Command, args []string) error {
			assignment, _, err := c.DeviceService.Assign(deviceID, &packngo.AddressStruct{Address: address})
			if err != nil {
				return errors.Wrap(err, "Could not assign Device IP address")
			}

			data := make([][]string, 1)

			data[0] = []string{assignment.ID, assignment.Address, strconv.FormatBool(assignment.Public), assignment.Created}
			header := []string{"ID", "Address", "Public", "Created"}

			return c.Out.Output(assignment, header, &data)
		},
	}

	assignIPCmd.Flags().StringVarP(&deviceID, "device-id", "d", "", "UUID of the device")
	assignIPCmd.Flags().StringVarP(&address, "address", "a", "", "IP address")

	_ = assignIPCmd.MarkFlagRequired("device-id")
	_ = assignIPCmd.MarkFlagRequired("address")
	return assignIPCmd
}
