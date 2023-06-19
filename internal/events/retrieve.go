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

package events

import (
	"context"
	"fmt"

	metal "github.com/equinix-labs/metal-go/metal/v1"
	"github.com/equinix/metal-cli/internal/outputs"
	pager "github.com/equinix/metal-cli/internal/pagination"
	"github.com/spf13/cobra"
)

func (c *Client) Retrieve() *cobra.Command {
	var eventID, deviceID, projectID, organizationID string

	retrieveEventCmd := &cobra.Command{
		Use:     `get [-p <project_id>] | [-d <device_id>] | [-i <event_id>] | [-O <organization_id>]`,
		Aliases: []string{"list"},
		Short:   "Retrieves events for the current user, an organization, a project, a device, or the details of a specific event.",
		Long:    "Retrieves events for the current user, an organization, a project, a device, or the details of a specific event. The current user's events includes all events in all projects and devices that the user has access to. When using --json or --yaml flags, the --include=relationships flag is implied.",
		Example: `  # Retrieve all events of a current user:
  metal event get

  # Retrieve the details of a specific event:
  metal event get -i e9a969b3-8911-4667-9d99-57cd3dd4ef6f

  # Retrieve all the events of an organization:
  metal event get -o c079178c-9557-48f2-9ce7-cfb927b81928

  # Retrieve all events of a project:
  metal event get -p 1867ee8f-6a11-470a-9505-952d6a324040

  # Retrieve all events of a device:
  metal event get -d ca614540-fbd4-4dbb-9689-457c6ccc8353`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			var events []metal.Event
			var err error
			header := []string{"ID", "Body", "Type", "Created"}
			inc := []string{}
			exc := []string{}

			// only fetch extra details when rendered
			switch c.Servicer.Format() {
			case outputs.FormatJSON, outputs.FormatYAML:
				inc = append(inc, "relationship")
			}

			if deviceID != "" && projectID != "" && organizationID != "" && eventID != "" {
				return fmt.Errorf("id, project-id, device-id, and organization-id parameters are mutually exclusive")
			} else if deviceID != "" {
				events, err = pager.GetDeviceEvents(c.EventService, deviceID, inc, exc)
				if err != nil {
					return fmt.Errorf("Could not list Device Events: %w", err)
				}
			} else if projectID != "" {
				events, err = pager.GetProjectEvents(c.EventService, projectID, inc, exc)
				if err != nil {
					return fmt.Errorf("Could not list Project Events: %w", err)
				}
			} else if organizationID != "" {
				events, err = pager.GetOrganizationEvents(c.EventService, organizationID, inc, exc)
				if err != nil {
					return fmt.Errorf("Could not list Organization Events: %w", err)
				}
			} else if eventID != "" {
				event, _, err := c.EventService.FindEventById(context.Background(), eventID).Include(inc).Exclude(exc).Execute()
				if err != nil {
					return fmt.Errorf("Could not get Event: %w", err)
				}

				data := make([][]string, 1)

				data[0] = []string{event.GetId(), event.GetBody(), event.GetType(), event.GetCreatedAt().String()}
				return c.Out.Output(event, header, &data)
			} else {
				events, err = pager.GetAllEvents(c.EventService, inc, exc)
				if err != nil {
					return fmt.Errorf("Could not list Events: %w", err)
				}
			}

			data := make([][]string, len(events))

			for i, event := range events {
				data[i] = []string{event.GetId(), event.GetBody(), event.GetType(), event.GetCreatedAt().String()}
			}

			return c.Out.Output(events, header, &data)
		},
	}
	retrieveEventCmd.Flags().StringVarP(&eventID, "id", "i", "", "UUID of the event")
	retrieveEventCmd.Flags().StringVarP(&projectID, "project-id", "p", "", "Project ID (METAL_PROJECT_ID)")
	retrieveEventCmd.Flags().StringVarP(&deviceID, "device-id", "d", "", "UUID of the device")
	retrieveEventCmd.Flags().StringVarP(&organizationID, "organization-id", "O", "", "UUID of the organization")
	return retrieveEventCmd
}
