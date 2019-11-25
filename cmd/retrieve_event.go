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

var eventID string

// retrieveEventsCmd represents the retrieveEvents command
var retrieveEventCmd = &cobra.Command{
	Use:   "get",
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
`,
	Run: func(cmd *cobra.Command, args []string) {
		events := []packngo.Event{}
		var err error
		header := []string{"ID", "Body", "Type", "Created"}
		listOpt := &packngo.ListOptions{Includes: []string{"relationships"}}

		if deviceID != "" && projectID != "" && organizationID != "" && eventID != "" {
			fmt.Println("The id, project-id, device-id, and organization-id parameters are mutually exclusive")
			return
		} else if deviceID != "" {
			events, _, err = PacknGo.Devices.ListEvents(deviceID, listOpt)
			if err != nil {
				fmt.Println("Client error:", err)
				return
			}
		} else if projectID != "" {
			events, _, err = PacknGo.Projects.ListEvents(projectID, listOpt)
			if err != nil {
				fmt.Println("Client error:", err)
				return
			}
		} else if organizationID != "" {
			events, _, err = PacknGo.Organizations.ListEvents(organizationID, listOpt)
			if err != nil {
				fmt.Println("Client error:", err)
				return
			}
		} else if eventID != "" {
			getOpt := &packngo.GetOptions{Includes: listOpt.Includes}
			event, _, err := PacknGo.Events.Get(eventID, getOpt)
			if err != nil {
				fmt.Println("Client error:", err)
				return
			}
			data := make([][]string, 1)

			data[0] = []string{event.ID, event.Body, event.Type, event.CreatedAt.String()}
			output(event, header, &data)
			return
		} else {
			events, _, err = PacknGo.Events.List(listOpt)
			if err != nil {
				fmt.Println("Client error:", err)
				return
			}
		}

		data := make([][]string, len(events))

		for i, event := range events {
			data[i] = []string{event.ID, event.Body, event.Type, event.CreatedAt.String()}
		}

		output(events, header, &data)
	},
}

func init() {
	eventCmd.AddCommand(retrieveEventCmd)
	retrieveEventCmd.Flags().StringVarP(&eventID, "id", "i", "", "UUID of the event")
	retrieveEventCmd.Flags().StringVarP(&projectID, "project-id", "p", "", "UUID of the project")
	retrieveEventCmd.Flags().StringVarP(&deviceID, "device-id", "d", "", "UUID of the device")
	retrieveEventCmd.Flags().StringVarP(&volumeID, "organization-id", "o", "", "UUID of the organization")

	retrieveEventCmd.PersistentFlags().BoolVarP(&isJSON, "json", "j", false, "JSON output")
	retrieveEventCmd.PersistentFlags().BoolVarP(&isYaml, "yaml", "y", false, "YAML output")
}
