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

package projects

import (
	"context"
	"fmt"

	"github.com/equinix/metal-cli/internal/outputs"
	"github.com/spf13/cobra"
)

func (c *Client) Retrieve() *cobra.Command {
	var projectID, projectName string
	// retrieveProjectCmd represents the retriveProject command
	retrieveProjectCmd := &cobra.Command{
		Use:     `get [-i <project_UUID> | -n <project_name>]`,
		Aliases: []string{"list"},
		Short:   "Retrieves all the current user's projects or the details of a specified project.",
		Long:    "Retrieves all the current user's projects or the details of a specified project. You can specify which project by UUID or name. When using `--json` or `--yaml` flags, the `--include=members` flag is implied.",
		Example: `  # Retrieve all projects:
  metal project get
  
  # Retrieve a specific project by UUID: 
  metal project get -i 2008f885-1aac-406b-8d99-e6963fd21333

  # Retrieve a specific project by name:
  metal project get -n dev-cluster03`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			if projectID != "" && projectName != "" {
				return fmt.Errorf("Must specify only one of project-id and project name")
			}
			inc := c.Servicer.Includes(nil)
			exc := c.Servicer.Excludes(nil)

			// only fetch extra details when rendered
			switch c.Servicer.Format() {
			case outputs.FormatJSON, outputs.FormatYAML:
				inc = append(inc, "members")
			}

			if projectID == "" {
				resp, err := c.ProjectService.FindProjects(context.Background()).Include(inc).Exclude(exc).ExecuteWithPagination()
				if err != nil {
					return fmt.Errorf("Could not list Projects: %w", err)
				}

				var data [][]string
				projects := resp.Projects
				if projectName == "" {
					data = make([][]string, len(projects))
					for i, p := range projects {
						data[i] = []string{p.GetId(), p.GetName(), p.GetCreatedAt().String()}
					}
				} else {
					data = make([][]string, 0)
					for _, p := range projects {
						if p.GetName() == projectName {
							data = append(data, []string{p.GetId(), p.GetName(), p.GetCreatedAt().String()})
							break
						}
					}
					if len(data) == 0 {
						return fmt.Errorf("Could not find project with name %q", projectName)
					}
				}

				header := []string{"ID", "Name", "Created"}
				return c.Out.Output(projects, header, &data)
			} else {
				p, _, err := c.ProjectService.FindProjectById(context.Background(), projectID).Include(inc).Exclude(exc).Execute()
				if err != nil {
					return fmt.Errorf("Could not get Project: %w", err)
				}

				data := make([][]string, 1)

				data[0] = []string{p.GetId(), p.GetName(), p.GetCreatedAt().String()}
				header := []string{"ID", "Name", "Created"}
				return c.Out.Output(p, header, &data)
			}
		},
	}

	retrieveProjectCmd.Flags().StringVarP(&projectName, "project", "n", "", "The name of the project.")
	retrieveProjectCmd.Flags().StringVarP(&projectID, "id", "i", "", "The project's UUID, which can be specified in the config created by metal init or set as METAL_PROJECT_ID environment variable.")
	return retrieveProjectCmd
}
