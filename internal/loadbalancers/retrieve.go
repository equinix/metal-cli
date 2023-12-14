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

	v1 "github.com/equinix/metal-cli/internal/loadbalancers/api/v1"
	"github.com/spf13/cobra"
)

func (c *Client) Retrieve() *cobra.Command {
	var loadbalancerID, loadbalancerName, projectID string
	// retrieveLoadBalancerCmd represents the retriveLoadBalancer command
	retrieveLoadBalancerCmd := &cobra.Command{
		Use:     `get [-i <loadbalancer_UUID> | -n <loadbalancer_name>]`,
		Aliases: []string{"list"},
		Short:   "Retrieves all the project loadbalancers or the details of a specified loadbalancer.",
		Long:    "Retrieves all the project loadbalancers or the details of a specified loadbalancer. You can specify which loadbalancer by UUID or name.",
		Example: `  # Retrieve all loadbalancers:
  metal loadbalancer get
  
  # Retrieve a specific loadbalancer by UUID: 
  metal loadbalancer get -i 2008f885-1aac-406b-8d99-e6963fd21333

  # Retrieve a specific loadbalancer by name:
  metal loadbalancer get -n dev-cluster03`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			if loadbalancerID != "" && loadbalancerName != "" {
				return fmt.Errorf("Must specify only one of loadbalancer-id and loadbalancer name")
			}

			lbs := []v1.LoadBalancer{}
			if loadbalancerID == "" {
				req := c.projectService.ListLoadBalancers(context.Background(), projectID)

				ret, _, err := req.Execute()
				if err != nil {
					return fmt.Errorf("Could not list LoadBalancers: %w", err)
				}
				lbs = append(lbs, ret.Loadbalancers...)
			} else {
				req := c.loadBalancerService.GetLoadBalancer(context.Background(), loadbalancerID)
				ret, _, err := req.Execute()
				if err != nil {
					return fmt.Errorf("Could not find LoadBalancer: %w", err)
				}
				lbs = append(lbs, *ret)
			}

			var data [][]string
			for _, p := range lbs {
				row := []string{p.GetId(), p.GetName(), p.GetCreatedAt().String()}

				if loadbalancerName > "" {
					if p.GetName() == loadbalancerName {
						data = append(data, row)
						break
					}
					continue
				}
				data = append(data, row)
			}
			if loadbalancerName > "" && len(data) == 0 {
				return fmt.Errorf("Could not find loadbalancer with name %q", loadbalancerName)
			}

			header := []string{"ID", "Name", "Created"}
			return c.Out.Output(lbs, header, &data)
		},
	}

	retrieveLoadBalancerCmd.Flags().StringVarP(&projectID, "project-id", "p", "", "The project's UUID. This flag is required, unless specified in the config created by metal init or set as METAL_PROJECT_ID environment variable.")

	retrieveLoadBalancerCmd.Flags().StringVarP(&loadbalancerName, "loadbalancer", "n", "", "The name of the loadbalancer.")
	retrieveLoadBalancerCmd.Flags().StringVarP(&loadbalancerID, "id", "i", "", "The loadbalancer's UUID, which can be specified in the config created by metal init or set as METAL_PROJECT_ID environment variable.")
	return retrieveLoadBalancerCmd
}
