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
	"github.com/equinix/metal-cli/internal/outputs"
	"github.com/packethost/packngo"
	"github.com/spf13/cobra"
)

type Client struct {
	Servicer Servicer
	Service  packngo.ProjectVirtualNetworkService
	Out      outputs.Outputer
}

func (c *Client) NewCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "virtual-network",
		Aliases: []string{"vlan", "vlans"},
		Short:   "Virtual network operations",
		Long:    `Virtual network operations: create, delete and get`,

		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			c.Service = c.Servicer.API().ProjectVirtualNetworks
		},
	}

	cmd.AddCommand(
		c.Retrieve(),
		c.Create(),
		c.Delete(),
	)
	return cmd
}

type Servicer interface {
	API() *packngo.Client
	ListOptions(defaultIncludes, defaultExcludes []string) *packngo.ListOptions
}

func NewClient(s Servicer, out outputs.Outputer) *Client {
	return &Client{
		Servicer: s,
		Out:      out,
	}
}
