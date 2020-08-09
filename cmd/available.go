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
	"github.com/packethost/packngo"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	cidr int
)

// availableCmd represents the available command
var availableCmd = &cobra.Command{
	Use:   "available",
	Short: "Retrieves a list of IP resevations for a single project.",
	Long: `Example:

packet ip available --reservation-id [reservation_id] --cidr [size_of_subnet]

  `,
	RunE: func(cmd *cobra.Command, args []string) error {
		result, _, err := PacknGo.ProjectIPs.AvailableAddresses(reservationID, &packngo.AvailableRequest{CIDR: cidr})

		if err != nil {
			return errors.Wrap(err, "Could not get available IP addresses")
		}
		data := make([][]string, len(result))
		for i, r := range result {
			data[i] = []string{r}
		}
		header := []string{"Available IPs"}

		return output(result, header, &data)
	},
}

func init() {
	availableCmd.Flags().StringVarP(&reservationID, "reservation-id", "r", "", "UUID of the reservation")
	availableCmd.Flags().IntVarP(&cidr, "cidr", "c", 0, "Size of subnet")

	_ = availableCmd.MarkFlagRequired("reservation-id")
	_ = availableCmd.MarkFlagRequired("cidr")
}
