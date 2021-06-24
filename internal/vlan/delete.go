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

package vlan

import (
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) Delete() *cobra.Command {
	var (
		vnetID string
		force  bool
	)

	deleteVnet := func(id string) error {
		_, err := c.Service.Delete(id)
		if err != nil {
			return err
		}
		fmt.Println("Virtual Network", id, "successfully deleted.")
		return err
	}

	// deleteVirtualNetworkCmd represents the deleteVirtualNetwork command
	var deleteVirtualNetworkCmd = &cobra.Command{
		Use:   "delete",
		Short: "Deletes a Virtual Network",
		Long: `Example:

metal virtual-network delete -i [virtual_network_UUID]

	`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if !force {
				prompt := promptui.Prompt{
					Label:     fmt.Sprintf("Are you sure you want to delete virual network %s", vnetID),
					IsConfirm: true,
				}

				_, err := prompt.Run()
				if err != nil {
					return nil
				}
			}
			return errors.Wrap(deleteVnet(vnetID), "Could not delete Virtual Network")
		},
	}

	deleteVirtualNetworkCmd.Flags().StringVarP(&vnetID, "id", "i", "", "UUID of the vlan")
	_ = deleteVirtualNetworkCmd.MarkFlagRequired("id")
	deleteVirtualNetworkCmd.Flags().BoolVarP(&force, "force", "f", false, "Force removal of the virtual network")

	return deleteVirtualNetworkCmd
}
