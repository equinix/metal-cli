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
	"fmt"

	"github.com/packethost/packngo"
	"github.com/spf13/cobra"
)

// updateProjectCmd represents the updateProject command
var updateProjectCmd = &cobra.Command{
	Use:   "project",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		req := &packngo.ProjectUpdateRequest{}
		if name != "" {
			req.Name = &name
		}

		if paymentMethodID != "" {
			req.PaymentMethodID = &paymentMethodID
		}
		p, _, err := PacknGo.Projects.Update(projectID, req)
		if err != nil {
			fmt.Println("Client error:", err)
			return
		}

		data := make([][]string, 1)

		data[0] = []string{p.ID, p.Name, p.Created}
		header := []string{"ID", "Name", "Created"}
		output(p, header, &data)
	},
}

func init() {
	updateCmd.AddCommand(updateProjectCmd)
	updateProjectCmd.Flags().StringVarP(&projectID, "project-id", "i", "", "--project-id or -i [UUID]")
	updateProjectCmd.MarkFlagRequired("project-id")

	updateProjectCmd.Flags().StringVarP(&name, "name", "n", "", "--name or -n [name]")

	updateProjectCmd.Flags().StringVarP(&paymentMethodID, "payment-method-id", "m", "", "--payment-method-id or -m [payment_method_UUID]")
}
