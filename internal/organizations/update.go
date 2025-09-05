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

	metal "github.com/equinix/equinix-sdk-go/services/metalv1"
	"github.com/spf13/cobra"
)

func (c *Client) Update() *cobra.Command {
	var organizationID, name, description, twitter, website string
	// createOrganizationCmd represents the createOrganization command
	updateOrganizationCmd := &cobra.Command{
		Use:   `update -i <organization_UUID> [-n <name>] [-d <description>] [-w <website_URL>] [-t <twitter_URL>]`,
		Short: "Updates the specified organization.",
		Long:  "Updates the specified organization. You can update the name, website, or Twitter.",
		Example: `  # Updates the name of an organization:
  metal organization update -i 3bd5bf07-6094-48ad-bd03-d94e8712fdc8 --name test-cluster02`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			req := metal.NewOrganizationInput()

			// API spec says organization address.address is required,
			// but the address is not included by default
			defaultIncludes := []string{"address", "billing_address"}

			include := c.Servicer.Includes(defaultIncludes)
			exclude := c.Servicer.Excludes(nil)

			if name != "" {
				req.Name = &name
			}

			if description != "" {
				req.Description = &description
			}

			if twitter != "" {
				req.Twitter = &twitter
			}

			if website != "" {
				req.Website = &website
			}

			org, _, err := c.Service.UpdateOrganization(context.Background(), organizationID).OrganizationInput(*req).Include(include).Exclude(exclude).Execute()
			if err != nil {
				return fmt.Errorf("Could not update Organization: %w", err)
			}

			data := make([][]string, 1)

			data[0] = []string{org.GetId(), org.GetName(), org.GetCreatedAt().String()}
			header := []string{"ID", "Name", "Created"}

			return c.Out.Output(org, header, &data)
		},
	}

	updateOrganizationCmd.Flags().StringVarP(&organizationID, "id", "i", "", "An organization UUID.")
	updateOrganizationCmd.Flags().StringVarP(&name, "name", "n", "", "New name for the organization.")
	updateOrganizationCmd.Flags().StringVarP(&description, "description", "d", "", "User-friendly description of the organization.")
	updateOrganizationCmd.Flags().StringVarP(&website, "website", "w", "", "A website URL for the organization.")
	updateOrganizationCmd.Flags().StringVarP(&twitter, "twitter", "t", "", "A Twitter URL of the organization.")

	_ = updateOrganizationCmd.MarkFlagRequired("id")
	return updateOrganizationCmd
}
