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
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func (c *Client) Retrieve() *cobra.Command {
	return &cobra.Command{
		Use:     `get`,
		Aliases: []string{"list"},
		Short:   "Retrieves a list of facilities.",
		Long:    "Retrieves a list of facilities available to the current user.",
		Example: `  # Lists facilities for current user:
  metal facilities get`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			facilities, _, err := c.Service.List(c.Servicer.ListOptions(nil, nil))
			if err != nil {
				return fmt.Errorf("Could not list Facilities: %w", err)
			}
			data := make([][]string, len(facilities))

			for i, facility := range facilities {
				var metro string
				if facility.Metro != nil {
					metro = facility.Metro.Code
				}
				data[i] = []string{facility.Name, facility.Code, metro, strings.Join(facility.Features, ",")}
			}
			header := []string{"Name", "Code", "Metro", "Features"}

			return c.Out.Output(facilities, header, &data)
		},
	}
}
