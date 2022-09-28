// Copyright © 2022 Equinix Metal Developers <support@equinixmetal.com>
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

package users

import (
	"strconv"
	"strings"

	"github.com/packethost/packngo"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) Add() *cobra.Command {
	var email string
	var roles []string
	var projectsIDs []string
	var organizationID string

	// addUserCmd represents the addUser command
	addUserCmd := &cobra.Command{
		Use:   `add --email <email> --roles <roles> [--organization-id <organization_id>] [--project-id <project_id>]`,
		Short: "Adds a user to an organization or project",
		Long:  "Adds a user, by email, to the organization or project specified by the --organization-id or --project-id flag. The user will be assigned the roles specified by the --roles flag.",
		Example: `  # Adds a user to a project with admin role:
  metal user add --email user@example.org --roles admin --project-id 3b0795ba-fd0b-4a9e-83a7-063e5e12409d
`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			getOpts := c.Servicer.ListOptions(nil, nil)
			createRequest := &packngo.InvitationCreateRequest{
				Invitee:     email,
				Roles:       roles,
				ProjectsIDs: projectsIDs,
			}
			invitation, _, err := c.InvitationService.Create(organizationID, createRequest, getOpts)
			if err != nil {
				return errors.Wrap(err, "Could not add Users")
			}

			data := make([][]string, 1)

			data[0] = []string{invitation.ID, invitation.Nonce, invitation.Invitee, invitation.Organization.Href, strconv.Itoa(len(invitation.Projects)), strings.Join(invitation.Roles, ", ")}
			header := []string{"ID", "Nonce", "Email", "Organization", "Projects", "Roles"}

			return c.Out.Output(invitation, header, &data)
		},
	}

	addUserCmd.Flags().StringVar(&email, "email", "", "Email of the user.")
	addUserCmd.Flags().StringVar(&organizationID, "organization-id", "", "Organization to invite the user to.")
	addUserCmd.Flags().StringSliceVar(&roles, "roles", []string{""}, "Roles to assign to the user.")
	addUserCmd.Flags().StringSliceVarP(&projectsIDs, "project-id", "p", []string{""}, "Projects to invite the user to with the specified roles.")

	_ = addUserCmd.MarkFlagRequired("email")

	return addUserCmd
}
