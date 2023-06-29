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

package devices

import (
	metal "github.com/equinix-labs/metal-go/metal/v1"
	"github.com/equinix/metal-cli/internal/outputs"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Client struct {
	Servicer Servicer
	Service  metal.DevicesApiService
	Out      outputs.Outputer
}

func (c *Client) NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     `device`,
		Aliases: []string{"server", "servers", "devices"},
		Short:   "Device operations: create, get, update, delete, reinstall, start, stop, and reboot.",
		Long:    "Device operations that control server provisioning, metadata, and basic operations.",

		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if root := cmd.Root(); root != nil {
				if root.PersistentPreRun != nil {
					root.PersistentPreRun(cmd, args)
				}
			}

			c.Service = *c.Servicer.MetalAPI(cmd).DevicesApi
		},
	}

	cmd.AddCommand(
		c.Retrieve(),
		c.Create(),
		c.Delete(),
		c.Update(),
		c.Reboot(),
		c.Reinstall(),
		c.Start(),
		c.Stop(),
	)
	return cmd
}

type Servicer interface {
	MetalAPI(*cobra.Command) *metal.APIClient
	Config(cmd *cobra.Command) *viper.Viper
	Filters() map[string]string
	Includes(defaultIncludes []string) (incl []string)
	Excludes(defaultExcludes []string) (excl []string)
}

func NewClient(s Servicer, out outputs.Outputer) *Client {
	return &Client{
		Servicer: s,
		Out:      out,
	}
}
