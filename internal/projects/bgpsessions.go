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
	"strings"

	"github.com/packethost/packngo"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) BGPSessions() *cobra.Command {
	var projectID string

	// bgpSessionsProjectCmd represents the updateProject command
	bgpSessionsProjectCmd := &cobra.Command{
		Use:   `bgp-sessions --id <project_UUID>`,
		Short: "Gets BGP Sessions for a project.",
		Long:  `Gets BGP Sessions for a project.`,
		Example: `  # Get BGP Sessions for project 50693ba9-e4e4-4d8a-9eb2-4840b11e9375:
  metal project bgp-sessions --id 50693ba9-e4e4-4d8a-9eb2-4840b11e9375`,
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			listOpt := c.Servicer.ListOptions(nil, nil)
			getOpts := &packngo.GetOptions{Includes: listOpt.Includes, Excludes: listOpt.Excludes}
			p, _, err := c.BGPConfigService.Get(projectID, getOpts)
			if err != nil {
				return errors.Wrap(err, "Could not get Project BGP Sessions")
			}

			data := make([][]string, len(p.Sessions))
			for i, s := range p.Sessions {
				data[i] = []string{
					s.ID,
					s.Status,
					strings.Join(s.LearnedRoutes, ","),
					strconv.FormatBool(s.DefaultRoute != nil && *s.DefaultRoute),
				}
			}
			header := []string{"ID", "Status", "Learned Routes", "Default Route"}
			return c.Out.Output(p, header, &data)
		},
	}

	bgpSessionsProjectCmd.Flags().StringVarP(&projectID, "id", "i", "", "Project ID (METAL_PROJECT_ID)")

	_ = bgpSessionsProjectCmd.MarkFlagRequired("id")
	return bgpSessionsProjectCmd
}
