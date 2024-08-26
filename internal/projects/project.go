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

package projects

import (
	metal "github.com/equinix/equinix-sdk-go/services/metalv1"
	"github.com/equinix/metal-cli/internal/outputs"
	"github.com/spf13/cobra"
)

type Client struct {
	Servicer         Servicer
	ProjectService   metal.ProjectsApiService
	BGPConfigService metal.BGPApiService

	Out outputs.Outputer
}

func (c *Client) NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     `project`,
		Aliases: []string{"projects"},
		Short:   "Project operations: create, get, update, delete, and bgp-enable, bgp-config, bgp-sessions.",
		Long:    "Information and management for Projects and Project-level BGP. Documentation on Projects is on https://deploy.equinix.com/developers/docs/metal/projects/creating-a-project/, and documentation on BGP is on https://deploy.equinix.com/developers/docs/metal/bgp/bgp-on-equinix-metal/.",
		Args:    cobra.NoArgs,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if root := cmd.Root(); root != nil {
				if root.PersistentPreRun != nil {
					root.PersistentPreRun(cmd, args)
				}
			}
			c.ProjectService = *c.Servicer.MetalAPI(cmd).ProjectsApi
			c.BGPConfigService = *c.Servicer.MetalAPI(cmd).BGPApi
		},
	}

	cmd.AddCommand(
		c.Retrieve(),
		c.Create(),
		c.Delete(),
		c.Update(),
		c.BGPEnable(),
		c.BGPConfig(),
		c.BGPSessions(),
	)
	return cmd
}

type Servicer interface {
	MetalAPI(*cobra.Command) *metal.APIClient
	Format() outputs.Format
	Includes(defaultIncludes []string) (incl []string)
	Excludes(defaultExcludes []string) (excl []string)
	Filters() map[string]string
}

func NewClient(s Servicer, out outputs.Outputer) *Client {
	return &Client{
		Servicer: s,
		Out:      out,
	}
}
