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
		checkFacility bool
		checkMetro    bool
	)
	// retrieveCapacitiesCmd represents the retrieveCapacity command
	var retrieveCapacityCmd = &cobra.Command{
		Use:     "get",
		Aliases: []string{"list"},
		Short:   "Returns a list of facilities or metros and plans with their current capacity.",
		Long: `Example:
Retrieve capacities:
metal capacity get { --metro | --facility }
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			var err error
			lister := c.Service.List
			fieldName := "Facility"

			if checkMetro {
				fieldName = "Metro"
				lister = c.Service.ListMetros
			}

			capacities, _, err := lister()
			if err != nil {
				return errors.Wrap(err, "Could not get Capacity")
			}

			header := []string{fieldName, "Plan", "Level"}
			requiredDataFormat := [][]string{}

			for locCode, capacity := range *capacities {
				for plan, bm := range capacity {
					loc := []string{}
					loc = append(loc, locCode, plan, bm.Level)
					requiredDataFormat = append(requiredDataFormat, loc)
				}
			}

			return c.Out.Output(capacities, header, &requiredDataFormat)
		},
	}
	retrieveCapacityCmd.Flags().BoolVarP(&checkFacility, "facility", "f", true, "Facility code (sv15)")
	retrieveCapacityCmd.Flags().BoolVarP(&checkMetro, "metro", "m", false, "Metro code (sv)")
	return retrieveCapacityCmd
}
