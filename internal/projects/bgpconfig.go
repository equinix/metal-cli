// Copyright © 2022 Equinix Metal Developers <support@equinixmetal.com>
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
	"strconv"

	"github.com/packethost/packngo"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) BGPConfig() *cobra.Command {
	var projectID string

	// bgpConfigProjectCmd represents the updateProject command
	bgpConfigProjectCmd := &cobra.Command{
		Use:   `bgp-config --id <project_UUID>`,
		Short: "Gets BGP Config for a project.",
		Long:  `Gets BGP Config for a project.`,
		Example: `  # Get BGP config for project 50693ba9-e4e4-4d8a-9eb2-4840b11e9375:
  metal project bgp-config --id 50693ba9-e4e4-4d8a-9eb2-4840b11e9375`,
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			listOpt := c.Servicer.ListOptions(nil, nil)
			getOpts := &packngo.GetOptions{Includes: listOpt.Includes, Excludes: listOpt.Excludes}
			p, _, err := c.BGPConfigService.Get(projectID, getOpts)
			if err != nil {
				return errors.Wrap(err, "Could not get Project BGP Config")
			}

			data := make([][]string, 1)

			data[0] = []string{projectID, p.Status, strconv.Itoa(p.Asn), p.DeploymentType, strconv.Itoa(p.MaxPrefix)}
			header := []string{"ID", "Status", "Sessions", "ASN", "DeploymentType", "MaxPrefix"}
			return c.Out.Output(p, header, &data)
		},
	}

	bgpConfigProjectCmd.Flags().StringVarP(&projectID, "id", "i", "", "Project ID (METAL_PROJECT_ID)")

	_ = bgpConfigProjectCmd.MarkFlagRequired("id")
	return bgpConfigProjectCmd
}
