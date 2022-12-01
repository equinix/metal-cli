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

package facilities

import (
	"github.com/equinix/metal-cli/internal/outputs"
	"github.com/packethost/packngo"
	"github.com/spf13/cobra"
)

type Client struct {
	Servicer Servicer
	Service  packngo.FacilityService
	Out      outputs.Outputer
}

func (c *Client) NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     `facilities`,
		Aliases: []string{"facility"},
		Short:   "Facility operations: get.",
		Long:    "Information about specific facilities. Facility-level operations have mostly been replaced by Metros, but remains for backwards-compatibility. Documentation about facilities is available at https://metal.equinix.com/developers/docs/locations/facilities/.",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if root := cmd.Root(); root != nil {
				if root.PersistentPreRun != nil {
					root.PersistentPreRun(cmd, args)
				}
			}
			c.Service = c.Servicer.API(cmd).Facilities
		},
	}

	cmd.AddCommand(
		c.Retrieve(),
	)
	return cmd
}

type Servicer interface {
	API(*cobra.Command) *packngo.Client
	ListOptions(defaultIncludes, defaultExcludes []string) *packngo.ListOptions
}

func NewClient(s Servicer, out outputs.Outputer) *Client {
	return &Client{
		Servicer: s,
		Out:      out,
	}
}
