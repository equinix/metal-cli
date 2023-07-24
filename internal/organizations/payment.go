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

package organizations

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

func (c *Client) PaymentMethods() *cobra.Command {
	var organizationID string
	// paymentMethodsCmd represents the paymentMethods command
	paymentMethodsCmd := &cobra.Command{
		Use:     `payment-methods -i <organization_UUID>`,
		Aliases: []string{"payment-method"},
		Short:   "Retrieves a list of payment methods.",
		Long:    "Retrieves a list of payment methods for the specified organization if the current user is a member with the proper role.",
		Example: `  # Lists the payment methods for an organization:
  metal organization payment-methods --id 3bd5bf07-6094-48ad-bd03-d94e8712fdc8`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true

			paymentMethodsList, _, err := c.Service.FindOrganizationPaymentMethods(context.Background(), organizationID).Include(c.Servicer.Includes(nil)).Exclude(c.Servicer.Excludes(nil)).Execute()
			if err != nil {
				return fmt.Errorf("could not list Payment Methods: %w", err)
			}
			paymentMethods := paymentMethodsList.GetPaymentMethods()
			data := make([][]string, len(paymentMethods))

			for i, p := range paymentMethods {
				data[i] = []string{p.GetId(), p.GetCardholderName(), p.GetExpirationMonth(), p.GetExpirationYear(), p.GetCreatedAt().String()}
			}
			header := []string{"ID", "Cardholder", "Exp. Month", "Exp. Year", "Created"}

			return c.Out.Output(paymentMethods, header, &data)
		},
	}

	paymentMethodsCmd.Flags().StringVarP(&organizationID, "id", "i", "", "The UUID of the organization.")
	_ = paymentMethodsCmd.MarkFlagRequired("id")
	return paymentMethodsCmd
}
