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

func (c *Client) Create() *cobra.Command {
	var (
		website     string
		twitter     string
		name        string
		description string
	)

	// createOrganizationCmd represents the createOrganization command
	createOrganizationCmd := &cobra.Command{
		Use:   `create -n <name> [-d <description>] [-w <website_URL>] [-t <twitter_URL>]`,
		Short: "Creates an organization.",
		Long:  "Creates a new organization with the current user as the organization's owner. ",
		Example: `  # Creates a new organization named "it-backend-infra": 
  metal organization create -n it-backend-infra
  
  # Creates a new organization with name, website, and twitter:
  metal organization create -n test-org -w www.metal.equinix.com -t https://twitter.com/equinixmetal `,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true

			// API spec says organization address.address is required,
			// but the address is not included by default
			defaultIncludes := []string{"address", "billing_address"}

			include := c.Servicer.Includes(defaultIncludes)
			exclude := c.Servicer.Excludes(nil)

			req := metal.NewOrganizationInput()
			req.Name = &name

			if description != "" {
				req.Description = &description
			}

			if twitter != "" {
				req.Twitter = &twitter
			}

			if website != "" {
				req.Website = &website
			}

			org, _, err := c.Service.CreateOrganization(context.Background()).OrganizationInput(*req).Include(include).Exclude(exclude).Execute()
			if err != nil {
				return fmt.Errorf("Could not create Organization: %w", err)
			}

			data := make([][]string, 1)

			data[0] = []string{org.GetId(), org.GetName(), org.GetCreatedAt().String()}
			header := []string{"ID", "Name", "Created"}

			return c.Out.Output(org, header, &data)
		},
	}

	createOrganizationCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the organization.")
	createOrganizationCmd.Flags().StringVarP(&description, "description", "d", "", "Description of the organization.")
	createOrganizationCmd.Flags().StringVarP(&website, "website", "w", "", "Website URL of the organization.")
	createOrganizationCmd.Flags().StringVarP(&twitter, "twitter", "t", "", "Twitter URL of the organization.")

	_ = createOrganizationCmd.MarkFlagRequired("name")
	return createOrganizationCmd
}
