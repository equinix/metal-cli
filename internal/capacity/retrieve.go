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

package capacity

import (
	"github.com/packethost/packngo"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) Retrieve() *cobra.Command {
	var (
		metros, facilities, plans, locs []string
		metro, facility, plan           string
	)
	// retrieveCapacitiesCmd represents the retrieveCapacity command
	var retrieveCapacityCmd = &cobra.Command{
		Use:     `get {-m [metros,...] | -f [facilities,...]} -P [plans,...]`,
		Aliases: []string{"list"},
		Short:   "Returns a list of facilities or metros and plans with their current capacity, optionally filtered by given locations and plans.",
		Example: `metal capacity get -m sv,ny,da -P c3.large.arm,c3.medium.x86`,

		RunE: func(cmd *cobra.Command, args []string) error {
			var err error
			var lister func() (*packngo.CapacityReport, *packngo.Response, error)
			var locationField string

			if metro != "" {
				metros = append(metros, metro)
			}
			if facility != "" {
				facilities = append(facilities, facility)
			}

			if (len(facilities) > 0) == (len(metros) > 0) {
				return errors.New("Either facilities or metros should be set")
			}
			cmd.SilenceUsage = true

			if plan != "" {
				plans = append(plans, plan)
			}

			if len(facilities) > 0 {
				lister = c.Service.List
				locationField = "Facility"
				locs = append(locs, facilities...)
			} else if len(metros) > 0 {
				lister = c.Service.ListMetros
				locationField = "Metro"
				locs = append(locs, metros...)
			}

			capacities, _, err := lister()
			if err != nil {
				return errors.Wrap(err, "Could not get Capacity")
			}

			header := []string{locationField, "Plan", "Level"}
			data := [][]string{}

			for locCode, capacity := range *capacities {
				for plan, bm := range capacity {
					if len(locs) > 0 {
						for _, location := range locs {
							if location == locCode {
								if len(plans) > 0 {
									for _, p := range plans {
										if plan == p {
											loc := []string{}
											loc = append(loc, locCode, plan, bm.Level)
											data = append(data, loc)
										}
									}
								} else {
									loc := []string{}
									loc = append(loc, locCode, plan, bm.Level)
									data = append(data, loc)
								}
							}
						}
					} else {
						if len(plans) > 0 {
							for _, p := range plans {
								if plan == p {
									loc := []string{}
									loc = append(loc, locCode, plan, bm.Level)
									data = append(data, loc)
								}
							}
						} else {
							loc := []string{}
							loc = append(loc, locCode, plan, bm.Level)
							data = append(data, loc)
						}
					}
				}
			}
			return c.Out.Output(capacities, header, &data)
		},
	}

	fs := retrieveCapacityCmd.Flags()

	fs.StringSliceVarP(&metros, "metros", "m", []string{}, "Codes of the metros")
	fs.StringSliceVarP(&facilities, "facilities", "f", []string{}, "Codes of the facilities")
	fs.StringSliceVarP(&plans, "plans", "P", []string{}, "Names of the plans")

	fs.StringVar(&metro, "metro", "", "Code of the metro")
	fs.StringVar(&facility, "facility", "", "Code of the facility")
	fs.StringVar(&plan, "plan", "", "Name of the plan")
	_ = fs.MarkDeprecated("metro", "use --metros instead")
	_ = fs.MarkDeprecated("plan", "use --plans instead")
	_ = fs.MarkDeprecated("facility", "use --facilities instead")

	return retrieveCapacityCmd
}
