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

package cmd

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// retrieveCapacitiesCmd represents the retrieveCapacity command
var retrieveCapacityCmd = &cobra.Command{
	Use:   "get",
	Short: "Returns a list of facilities and plans with their current capacity.",
	Long: `Example:
Retrieve capacities:
packet capacity get
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error
		capacities, _, err := PacknGo.CapacityService.List()
		if err != nil {
			return errors.Wrap(err, "Could not get Capacity")
		}

		header := []string{"Facility", "Plan", "Level"}
		requiredDataFormat := [][]string{}

		for facilityCode, capacity := range *capacities {
			for plan, bm := range capacity {
				facility := []string{}
				facility = append(facility, facilityCode, plan, bm.Level)
				requiredDataFormat = append(requiredDataFormat, facility)
			}
		}

		return outputMergingCells(capacities, header, &requiredDataFormat)
	},
}

func init() {
	capacityCmd.AddCommand(retrieveCapacityCmd)
}
