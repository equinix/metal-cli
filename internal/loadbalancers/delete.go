// Copyright © 2022 Equinix Metal Developers <support@equinixlbaas.com>
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

package loadbalancers

import (
	"context"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func (c *Client) Delete() *cobra.Command {
	var (
		force          bool
		loadbalancerID string
		confirmation   string
	)
	deleteLoadBalancer := func(id string) error {
		_, err := c.loadBalancerService.DeleteLoadBalancer(context.Background(), id).Execute()
		if err != nil {
			return err
		}
		fmt.Println("LoadBalancer", loadbalancerID, "successfully deleted.")
		return nil
	}
	// deleteLoadBalancerCmd represents the deleteLoadBalancer command
	deleteLoadBalancerCmd := &cobra.Command{
		Use:   `delete --id <loadbalancer_UUID> [--force]`,
		Short: "Deletes a loadbalancer.",
		Long:  "Deletes the specified loadbalancer with a confirmation prompt. To skip the confirmation use --force.",
		Example: `  # Deletes loadbalancer 50693ba9-e4e4-4d8a-9eb2-4840b11e9375:
  metal loadbalancer delete -i 50693ba9-e4e4-4d8a-9eb2-4840b11e9375
  >
  ✔ Are you sure you want to delete loadbalancer 50693ba9-e4e4-4d8a-9eb2-4840b11e9375: y
  
  # Deletes loadbalancer 50693ba9-e4e4-4d8a-9eb2-4840b11e9375, skipping confirmation:
  metal loadbalancer delete -i 50693ba9-e4e4-4d8a-9eb2-4840b11e9375 -f`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			if !force {
				fmt.Printf("Are you sure you want to delete Metal Loadbalancer %s [Y/N]: ", loadbalancerID)
				_, err := fmt.Scanln(&confirmation)
				if err != nil {
					return nil
				}
			}
			confirmation = strings.TrimSpace(strings.ToLower(confirmation))
			if confirmation != "yes" && confirmation != "y" {
				fmt.Println("Metal Loadbalancer deletion cancelled.")
				return nil
			}

			if err := deleteLoadBalancer(loadbalancerID); err != nil {
				return fmt.Errorf("Could not delete LoadBalancer: %w", err)
			}
			return nil
		},
	}

	deleteLoadBalancerCmd.Flags().StringVarP(&loadbalancerID, "id", "i", "", "The loadbalancer's ID. This flag is required.")
	_ = deleteLoadBalancerCmd.MarkFlagRequired("id")

	deleteLoadBalancerCmd.Flags().BoolVarP(&force, "force", "f", false, "Force removal of the loadbalancer")
	return deleteLoadBalancerCmd
}
