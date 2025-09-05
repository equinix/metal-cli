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
	"context"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func (c *Client) Delete() *cobra.Command {
	var (
		vnetID       string
		force        bool
		confirmation string
	)

	deleteVnet := func(id string) error {
		_, err := c.Service.DeleteVirtualNetwork(context.Background(), vnetID).Execute()
		if err != nil {
			return err
		}
		fmt.Println("Virtual Network", id, "successfully deleted.")
		return err
	}

	// deleteVirtualNetworkCmd represents the deleteVirtualNetwork command
	deleteVirtualNetworkCmd := &cobra.Command{
		Use:   `delete -i <virtual_network_UUID> [-f]`,
		Short: "Deletes a virtual network.",
		Long:  "Deletes the specified VLAN with a confirmation prompt. To skip the confirmation use --force. You are not able to delete a VLAN that is attached to any ports.",
		Example: `  # Deletes a VLAN, with confirmation.
  metal virtual-network delete -i 77e6d57a-d7a4-4816-b451-cf9b043444e2
  >
  ✔ Are you sure you want to delete virtual network 77e6d57a-d7a4-4816-b451-cf9b043444e2 [Y/N]: y
		
  # Deletes a VLAN, skipping confirmation.
  metal virtual-network delete -f -i 77e6d57a-d7a4-4816-b451-cf9b043444e2`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true

			if !force {
				fmt.Printf("Are you sure you want to delete virtual network %s [Y/N]: ", vnetID)

				_, err := fmt.Scanln(&confirmation)
				if err != nil {
					fmt.Println("Error reading confirmation:", err)
					return nil
				}

				confirmation = strings.TrimSpace(strings.ToLower(confirmation))
				if confirmation != "yes" && confirmation != "y" {
					fmt.Println("virtual network deletion cancelled.")
					return nil
				}
			}
			if err := deleteVnet(vnetID); err != nil {
				return fmt.Errorf("could not delete Virtual Network: %w", err)
			}
			return nil
		},
	}

	deleteVirtualNetworkCmd.Flags().StringVarP(&vnetID, "id", "i", "", "UUID of the VLAN.")
	_ = deleteVirtualNetworkCmd.MarkFlagRequired("id")
	deleteVirtualNetworkCmd.Flags().BoolVarP(&force, "force", "f", false, "Skips confirmation for the removal of the virtual network.")

	return deleteVirtualNetworkCmd
}
