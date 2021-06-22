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

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// unassignIpCmd represents the unassignIp command
var unassignIPCmd = &cobra.Command{
	Use:   "unassign",
	Short: "Unassigns an IP address.",
	Long: `Example:

metal ip unassign --id [assignment-UUID]

	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		_, err := apiClient.DeviceIPs.Unassign(assignmentID)
		if err != nil {
			return errors.Wrap(err, "Could not unassign IP address")
		}

		fmt.Println("IP address unassigned successfully.")
		return nil
	},
}

func init() {
	unassignIPCmd.Flags().StringVarP(&assignmentID, "id", "i", "", "UUID of the assignment")
	_ = unassignIPCmd.MarkFlagRequired("id")
}
