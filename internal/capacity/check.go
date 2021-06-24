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
	"strconv"

	"github.com/packethost/packngo"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// createOrganizationCmd represents the createOrganization command
func (c *Client) Check() *cobra.Command {
	var (
		metro, plan, facility string
		quantity              int
	)
	var checkCapacityCommand = &cobra.Command{
		Use:   "check",
		Short: "Validates if a deploy can be fulfilled.",
		Long: `Example:

metal capacity check {-m [metro] | -f [facility]} -p [plan] -q [quantity]

	`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if facility != "" && metro != "" {
				return errors.New("Either facility or metro should be set")
			}
			checker := c.Service.Check
			req := &packngo.CapacityInput{
				Servers: []packngo.ServerInfo{
					{
						Facility: facility,
						Plan:     plan,
						Quantity: quantity},
				},
			}
			locationField := "Facility"
			locationer := func(si packngo.ServerInfo) string {
				return si.Facility
			}

			if metro != "" {
				req.Servers[0].Metro = metro
				checker = c.Service.CheckMetros
				locationer = func(si packngo.ServerInfo) string {
					return si.Metro
				}
				locationField = "Metro"
			}

			availability, _, err := checker(req)
			if err != nil {
				return errors.Wrap(err, "Could not check capacity")
			}

			data := make([][]string, 1)

			data[0] = []string{locationer(availability.Servers[0]), availability.Servers[0].Plan,
				strconv.Itoa(availability.Servers[0].Quantity), strconv.FormatBool(availability.Servers[0].Available)}
			header := []string{locationField, "Plan", "Quantity", "Availability"}
			return c.Out.Output(availability, header, &data)
		},
	}

	checkCapacityCommand.Flags().StringVarP(&metro, "metro", "m", "", "Code of the metro")
	checkCapacityCommand.Flags().StringVarP(&facility, "facility", "f", "", "Code of the facility")
	checkCapacityCommand.Flags().StringVarP(&plan, "plan", "p", "", "Name of the plan")
	checkCapacityCommand.Flags().IntVarP(&quantity, "quantity", "q", 0, "Number of devices wanted")

	_ = checkCapacityCommand.MarkFlagRequired("plan")
	_ = checkCapacityCommand.MarkFlagRequired("quantity")
	return checkCapacityCommand
}
