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
		Use:     `get [-m | -f] | [--metros <list> | --facilities <list>] [-P <list>]`,
		Aliases: []string{"list"},
		Short:   "Returns capacity of metros or facilities, with optional filtering.",
		Long:    "Returns the capacity of metros or facilities. Filters for metros, facilities, plans are available. Metro flags and facility flags are mutually exclusive. If no flags are included, returns capacity for all plans in all facilities.",
		Example: `  # Returns the capacity of all plans in all facilities:
  metal capacity get 

  # Returns the capacity of the c3.small.x86 in all metros:
  metal capacity get -m -P c3.small.x86

  # Returns c3.large.arm and c3.medium.x86 capacity in the Silicon Valley, New York, and Dallas metros:
  metal capacity get --metros sv,ny,da -P c3.large.arm,c3.medium.x86`,

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

			filtered := map[string]map[string]map[string]string{}

			for locCode, capacity := range *capacities {
				if len(locs) > 0 && !contains(locs, locCode) {
					continue
				}
				for plan, bm := range capacity {
					if len(plans) > 0 && !contains(plans, plan) {
						continue
					}
					loc := []string{locCode, plan, bm.Level}
					data = append(data, loc)
					if len(filtered[locCode]) == 0 {
						filtered[locCode] = map[string]map[string]string{}
					}
					filtered[locCode][plan] = map[string]string{"levels": bm.Level}
				}
			}

			return c.Out.Output(filtered, header, &data)
		},
	}

	fs := retrieveCapacityCmd.Flags()
	fs.BoolVarP(&checkFacility, "facility", "f", true, "Return the capacity for all facilities. Can not be used with -m.")
	fs.BoolVarP(&checkMetro, "metro", "m", false, "Return the capacity for all metros. Can not be used with -f.")
	fs.StringSliceVar(&metros, "metros", []string{}, "A metro or list of metros for client-side filtering. Will only return the capacity for the specified metros. Can not be used with --facilities.")
	fs.StringSliceVar(&facilities, "facilities", []string{}, "A facility or list of facilities for client-side filtering. Will only return the capacity for the specified facilities. Can not be used with --metros.")
	fs.StringSliceVarP(&plans, "plans", "P", []string{}, "Return only the capacity for the specified plans.")
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
