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

	"github.com/spf13/cobra"
)

// paymentMethodsCmd represents the paymentMethods command
var paymentMethodsCmd = &cobra.Command{
	Use:     "payment-methods",
	Aliases: []string{"payment-method"},
	Short:   "Retrieves a list of payment methods for the organization",
	Long: `Example:

packet organization get payment-methods --id [organization_UUID]

`,
	Run: func(cmd *cobra.Command, args []string) {
		paymentMethods, _, err := PacknGo.Organizations.ListPaymentMethods(organizationID)
		if err != nil {
			fmt.Println("Client error:", err)
			return
		}

		data := make([][]string, len(paymentMethods))

		for i, p := range paymentMethods {
			data[i] = []string{p.ID, p.CardholderName, p.ExpMonth, p.ExpYear, p.Created}
		}
		header := []string{"ID", "Cardholder", "Exp. Month", "Exp. Year", "Created"}

		output(paymentMethods, header, &data)
	},
}

func init() {
	retrieveOrganizationCmd.AddCommand(paymentMethodsCmd)
	paymentMethodsCmd.Flags().StringVarP(&organizationID, "id", "i", "", "UUID of the organization")
	paymentMethodsCmd.MarkFlagRequired("id")
}
