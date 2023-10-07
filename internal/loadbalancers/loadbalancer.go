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
	lbaas "github.com/equinix/metal-cli/internal/loadbalancers/api/v1"
	v1 "github.com/equinix/metal-cli/internal/loadbalancers/api/v1"
	"github.com/equinix/metal-cli/internal/outputs"
	"github.com/spf13/cobra"
)

type Client struct {
	Servicer            Servicer
	loadBalancerService lbaas.LoadBalancersApiService
	projectService      lbaas.ProjectsApiService
	portsService        lbaas.PortsApiService
	poolsService        lbaas.PoolsApiService
	originsService      lbaas.OriginsApiService

	Out outputs.Outputer
}

const ProviderID = "loadpvd-gOB_-byp5ebFo7A3LHv2B"

var LBMetros = map[string]string{
	"da": "lctnloc--uxs0GLeAELHKV8GxO_AI",
	"ny": "lctnloc-Vy-1Qpw31mPi6RJQwVf9A",
	"sv": "lctnloc-H5rl2M2VL5dcFmdxhbEKx",
}

func (c *Client) NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     `loadbalancer`,
		Aliases: []string{"loadbalancers"},
		Short:   "LoadBalancer operations: create, get, update, delete, and bgpconfig.",
		Long:    "Information and management for LoadBalancers and LoadBalancer-level BGP. Documentation on LoadBalancers is on https://lbaas.equinix.com/developers/docs/accounts/loadbalancers/, and documentation on BGP is on https://lbaas.equinix.com/developers/docs/bgp/bgp-on-equinix-metal/.",

		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if root := cmd.Root(); root != nil {
				if root.PersistentPreRun != nil {
					root.PersistentPreRun(cmd, args)
				}
			}
			c.loadBalancerService = *c.Servicer.LoadbalancerAPI(cmd).LoadBalancersApi
			c.originsService = *c.Servicer.LoadbalancerAPI(cmd).OriginsApi
			c.portsService = *c.Servicer.LoadbalancerAPI(cmd).PortsApi
			c.poolsService = *c.Servicer.LoadbalancerAPI(cmd).PoolsApi
			c.projectService = *c.Servicer.LoadbalancerAPI(cmd).ProjectsApi
		},
	}

	cmd.AddCommand(
		c.Retrieve(),
		c.Create(),
		c.Delete(),
		// c.Update(),
		// c.BGPConfig(),
	)
	return cmd
}

type Servicer interface {
	LoadbalancerAPI(cmd *cobra.Command) *v1.APIClient
	Format() outputs.Format
}

func NewClient(s Servicer, out outputs.Outputer) *Client {
	return &Client{
		Servicer: s,
		Out:      out,
	}
}
