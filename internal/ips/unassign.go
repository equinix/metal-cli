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

	"github.com/spf13/cobra"
)

func (c *Client) Unassign() *cobra.Command {
	var assignmentID string
	// unassignIpCmd represents the unassignIp command
	unassignIPCmd := &cobra.Command{
		Use:   `unassign -i <assignment_UUID> `,
		Short: "Unassigns an IP address assignment.",
		Long:  "Unassigns an subnet and IP address assignment from a device by its assignment ID. ",
		Example: `  # Unassigns an IP address assignment:
  metal ip unassign --id abd8674b-96c4-4271-92f5-2eaf5944c86f`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			_, err := c.DeviceService.Unassign(assignmentID)
			if err != nil {
				return fmt.Errorf("Could not unassign IP address: %w", err)
			}

			fmt.Println("IP address unassigned successfully.")
			return nil
		},
	}

	unassignIPCmd.Flags().StringVarP(&assignmentID, "id", "i", "", "The UUID of the assignment.")
	_ = unassignIPCmd.MarkFlagRequired("id")
	return unassignIPCmd
}
