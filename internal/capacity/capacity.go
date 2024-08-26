// Copyright © 2019 Ali Bazlamit <ali.bazlamit@hotmail.com>
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

package capacity

import (
	metal "github.com/equinix/equinix-sdk-go/services/metalv1"
	"github.com/equinix/metal-cli/internal/outputs"
	"github.com/spf13/cobra"
)

type Client struct {
	Servicer Servicer
	Service  metal.CapacityApiService
	Out      outputs.Outputer
}

func (c *Client) NewCommand() *cobra.Command {
	// capacityCmd represents the capacity command
	cmd := &cobra.Command{
		Use:   "capacity",
		Short: `Capacity operations: get, check`,
		Long:  "Capacity operations. For more information on capacity in metros, visit https://deploy.equinix.com/developers/docs/metal/locations/metros/ For more information on capacity in facilities, visit https://deploy.equinix.com/developers/docs/metal/locations/facilities/.",
		Args:  cobra.NoArgs,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if root := cmd.Root(); root != nil {
				if root.PersistentPreRun != nil {
					root.PersistentPreRun(cmd, args)
				}
			}
			c.Service = *c.Servicer.MetalAPI(cmd).CapacityApi
		},
	}

	cmd.AddCommand(
		c.Check(),
		c.Retrieve(),
	)
	return cmd
}

type Servicer interface {
	MetalAPI(*cobra.Command) *metal.APIClient
	Format() outputs.Format
}

func NewClient(s Servicer, out outputs.Outputer) *Client {
	return &Client{
		Servicer: s,
		Out:      out,
	}
}
