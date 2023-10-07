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
	"strconv"

	lbaas "github.com/equinix/metal-cli/internal/loadbalancers/api/v1"
	"github.com/spf13/cobra"
)

func (c *Client) Create() *cobra.Command {
	var (
		name       string
		locationId string
		portIds    []string
		providerId string
		projectID  string
	)

	// loadbalancerCreateCmd represents the loadbalancerCreate command
	createLoadBalancerCmd := &cobra.Command{
		Use:   `create -n <loadbalancer_name> [-O <organization_UUID>] [-m <payment_method_UUID>]`,
		Short: "Creates a loadbalancer.",
		Long:  "Creates a loadbalancer with the specified name. If no organization is specified, the loadbalancer is created in the current user's default organization. If no payment method is specified the organization's default payment method is used.",
		Example: `  # Creates a new loadbalancer named dev-cluster02: 
  metal loadbalancer create --name dev-cluster02
  
  # Creates a new loadbalancer named dev-cluster03 in the specified organization with a payment method:
  metal loadbalancer create -n dev-cluster03 -O 814b09ca-0d0c-4656-9de0-4ce65c6faf70 -m ab1fbdaa-8b25-4c3e-8360-e283852e3747`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true

			// handle metro aliases for locations
			if LBMetros[locationId] != "" {
				locationId = LBMetros[locationId]
			}

			req := c.projectService.CreateLoadBalancer(context.Background(), projectID)
			req.LoadBalancerCreate(*lbaas.NewLoadBalancerCreate(name, locationId, portIds, providerId))
			lb, _, err := req.Execute()
			if err != nil {
				return fmt.Errorf("Could not create LoadBalancer: %w", err)
			}

			data := make([][]string, 1)

			data[0] = []string{lb.GetId(), lb.GetMessage(), lb.GetVersion(), strconv.Itoa(int(lb.GetStatus()))}
			header := []string{"ID", "Message", "Version", "Status"}
			return c.Out.Output(lb, header, &data)
		},
	}

	createLoadBalancerCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the loadbalancer")
	createLoadBalancerCmd.Flags().StringVarP(&projectID, "project-id", "p", "", "The project's UUID. This flag is required, unless specified in the config created by metal init or set as METAL_PROJECT_ID environment variable.")
	createLoadBalancerCmd.Flags().StringVarP(&locationId, "location", "l", "", "The location's ID.")
	createLoadBalancerCmd.Flags().StringVarP(&providerId, "provider", "r", ProviderID, "The provider ID.")
	createLoadBalancerCmd.Flags().StringSliceVarP(&portIds, "port", "o", []string{}, "The port's UUID. This flag is required, unless specified in the config created by metal init or set as METAL_PORT_ID environment variable.")

	// TODO(displague) Not sure if this is needed
	_ = createLoadBalancerCmd.MarkFlagRequired("port")
	_ = createLoadBalancerCmd.MarkFlagRequired("provider")
	_ = createLoadBalancerCmd.MarkFlagRequired("location")
	_ = createLoadBalancerCmd.MarkFlagRequired("project-id")
	_ = createLoadBalancerCmd.MarkFlagRequired("name")
	return createLoadBalancerCmd
}
