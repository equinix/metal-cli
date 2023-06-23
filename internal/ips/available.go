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
	"context"
	"fmt"
	"strconv"

	metal "github.com/equinix-labs/metal-go/metal/v1"
	"github.com/spf13/cobra"
)

var cidr int

func (c *Client) Available() *cobra.Command {
	var reservationID string
	// availableCmd represents the available command
	availableCmd := &cobra.Command{
		Use:   `available -r <reservation_UUID> -c <size_of_subnet>`,
		Short: "Lists available IP addresses from a reservation.",
		Long:  "Lists available IP addresses in a specified reservation for the desired subnet size.",
		Example: `  # Lists available IP addresses in a reservation for a /31 subnet:
  metal ip available --reservation-id da1bb048-ea6e-4911-8ab9-b95635ca127a --cidr 31`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			Cidr := metal.FindIPAvailabilitiesCidrParameter(strconv.Itoa(cidr))
			resultList, _, err := c.IPService.FindIPAvailabilities(context.Background(), reservationID).Cidr(Cidr).Execute()
			if err != nil {
				return fmt.Errorf("Could not get available IP addresses: %w", err)
			}
			result := resultList.GetAvailable()
			data := make([][]string, len(result))
			for i, r := range result {
				data[i] = []string{r}
			}
			header := []string{"Available IPs"}

			return c.Out.Output(result, header, &data)
		},
	}

	availableCmd.Flags().StringVarP(&reservationID, "reservation-id", "r", "", "The UUID of the IP address reservation.")
	availableCmd.Flags().IntVarP(&cidr, "cidr", "c", 0, "The size of the desired subnet in bits.")

	_ = availableCmd.MarkFlagRequired("reservation-id")
	_ = availableCmd.MarkFlagRequired("cidr")
	return availableCmd
}
