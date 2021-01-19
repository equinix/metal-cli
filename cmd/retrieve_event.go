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
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var eventID string

// retrieveEventsCmd represents the retrieveEvents command
var retrieveEventCmd = &cobra.Command{
	Use:   "get",
	Aliases: []string{"list"},
	Short: "Retrieves one or more events for organizations, projects, or devices.",
	Long: `Example:
Retrieve all events:
packet event get

Retrieve a specific event:
packet event get -i [event_UUID]

Retrieve all events of an organization:
packet event get -o [organization_UUID]

Retrieve all events of a project:
packet event get -p [project_UUID]

Retrieve all events of a device:
packet event get -d [device_UUID]

Retrieve all events of a current user:
packet event get

When using "--json" or "--yaml", "--include=relationships" is implied.
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var events []packngo.Event
		var err error
		header := []string{"ID", "Body", "Type", "Created"}

		inc := []string{"relationships"}

		// don't fetch extra details that won't be rendered
		if !isYaml && !isJSON {
			inc = nil
		}
		listOpt := listOptions(inc, nil)

		if deviceID != "" && projectID != "" && organizationID != "" && eventID != "" {
			return fmt.Errorf("The id, project-id, device-id, and organization-id parameters are mutually exclusive")
		} else if deviceID != "" {
			events, _, err = apiClient.Devices.ListEvents(deviceID, listOpt)
			if err != nil {
				return errors.Wrap(err, "Could not list Device Events")
			}
		} else if projectID != "" {
			events, _, err = apiClient.Projects.ListEvents(projectID, listOpt)
			if err != nil {
				return errors.Wrap(err, "Could not list Project Events")
			}
		} else if organizationID != "" {
			events, _, err = apiClient.Organizations.ListEvents(organizationID, listOpt)
			if err != nil {
				return errors.Wrap(err, "Could not list Organization Events")
			}
		} else if eventID != "" {
			getOpt := &packngo.GetOptions{Includes: listOpt.Includes}
			event, _, err := apiClient.Events.Get(eventID, getOpt)
			if err != nil {
				return errors.Wrap(err, "Could not get Event")
			}
			data := make([][]string, 1)

			data[0] = []string{event.ID, event.Body, event.Type, event.CreatedAt.String()}
			return output(event, header, &data)
		} else {
			events, _, err = apiClient.Events.List(listOpt)
			if err != nil {
				return errors.Wrap(err, "Could not list Events")
			}
		}

		data := make([][]string, len(events))

		for i, event := range events {
			data[i] = []string{event.ID, event.Body, event.Type, event.CreatedAt.String()}
		}

		return output(events, header, &data)
	},
}

func init() {
	eventCmd.AddCommand(retrieveEventCmd)
	retrieveEventCmd.Flags().StringVarP(&eventID, "id", "i", "", "UUID of the event")
	retrieveEventCmd.Flags().StringVarP(&projectID, "project-id", "p", "", "UUID of the project")
	retrieveEventCmd.Flags().StringVarP(&deviceID, "device-id", "d", "", "UUID of the device")
	retrieveEventCmd.Flags().StringVarP(&volumeID, "organization-id", "o", "", "UUID of the organization")
}
