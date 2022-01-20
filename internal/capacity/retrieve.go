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
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) Retrieve() *cobra.Command {
	var (
		checkFacility, checkMetro       bool
		metros, facilities, plans, locs []string
	)
	// retrieveCapacitiesCmd represents the retrieveCapacity command
	var retrieveCapacityCmd = &cobra.Command{
		Use:     `get [[-m | -f] | [--metros metros,... | --facilities facilities,...]] [-P plans,...]`,
		Aliases: []string{"list"},
		Short:   "Returns a list of facilities or metros and plans with their current capacity, with filtering.",
		Example: `metal capacity get -m sv,ny,da -P c3.large.arm,c3.medium.x86`,

		RunE: func(cmd *cobra.Command, args []string) error {
			var err error
			var locationField string
			lister := c.Service.List // Default to facilities

			fs := cmd.Flags()
			if fs.Changed("metros") && fs.Changed("facilities") {
				return errors.New("Either facilities or metros, but not both, can be set")
			}
			if fs.Changed("facility") && fs.Changed("metro") {
				return errors.New("Either --facility (-f) or --metro (-m), but not both, can be set")
			}
			if fs.Changed("facility") && fs.Changed("metros") || fs.Changed("facilities") && fs.Changed("metro") {
				return errors.New("Cannot specify both facility and metro filtering")
			}
			if fs.Changed("metro") && fs.Changed("metros") || fs.Changed("facility") && fs.Changed("facilities") {
				return errors.New("Cannot use both --metro (-m) and --metros or --facility (-f) and --facilities")
			}
			// add other bad combos

			cmd.SilenceUsage = true

			if len(facilities) > 0 {
				locationField = "Facility"
				locs = append(locs, facilities...)
			} else if len(metros) > 0 || checkMetro {
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
				// If the list of locations isn't empty and contains this location code
				if !(len(locs) > 0 && !contains(locs, locCode)) {
					for plan, bm := range capacity {
						loc := []string{}
						// If the list of plans isn't empty and contains this plan
						if !(len(plans) > 0 && !contains(plans, plan)) {
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
	fs.BoolVarP(&checkFacility, "facility", "f", true, "Report all facilites")
	fs.BoolVarP(&checkMetro, "metro", "m", false, "Report all metros")
	fs.StringSliceVar(&metros, "metros", []string{}, "Codes of the metros (client side filtering)")
	fs.StringSliceVar(&facilities, "facilities", []string{}, "Codes of the facilities (client side filtering)")
	fs.StringSliceVarP(&plans, "plans", "P", []string{}, "Names of the plans")
	return retrieveCapacityCmd
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
