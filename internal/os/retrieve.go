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

package os

import (
	"sort"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) Retrieve() *cobra.Command {
	return &cobra.Command{
		Use:     `get`,
		Aliases: []string{"list"},
		Short:   "Retrieves a list of operating systems.",
		Long:    "Retrieves a list of operating systems available to the current user. Response includes the operating system's slug, distro, version, and name.",
		Example: `  # Lists the operating systems available to the current user:
  metal operating-systems get`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			oss, _, err := c.Service.List()
			if err != nil {
				return errors.Wrap(err, "Could not list OperatingSystems")
			}

			sort.Slice(oss, func(a, b int) bool {
				return oss[a].Name < oss[b].Name
			})

			data := make([][]string, len(oss))

			for i, os := range oss {
				data[i] = []string{os.Name, os.Slug, os.Distro, os.Version}
			}
			header := []string{"Name", "Slug", "Distro", "Version"}

			return c.Out.Output(oss, header, &data)
		},
	}
}
