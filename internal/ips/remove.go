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

	"github.com/spf13/cobra"
)

func (c *Client) Remove() *cobra.Command {
	var reservationID string
	// removeIPCmd represents the removeIp command
	removeIPCmd := &cobra.Command{
		Use:   `remove -i <reservation_UUID>`,
		Short: "Removes an IP address reservation from a project.",
		Long:  "Removes an IP address reservation from a project. Any subnets and IP addresses in the reservation will no longer be able to be used by your devices.",
		Example: `  # Removes an IP address reservation:
  metal ip remove --id a9dfc9d5-ba1a-4d11-8cfc-6e30b9630876`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			_, err := c.IPService.DeleteIPAddress(context.Background(), reservationID).Execute()
			if err != nil {
				return fmt.Errorf("Could not remove IP address Reservation: %w", err)
			}

			fmt.Println("IP reservation removed successfully.")
			return nil
		},
	}

	removeIPCmd.Flags().StringVarP(&reservationID, "id", "i", "", "UUID of the reservation")

	_ = removeIPCmd.MarkFlagRequired("id")
	return removeIPCmd
}
