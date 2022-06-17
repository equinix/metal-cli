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
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) Delete() *cobra.Command {
	var (
		force     bool
		projectID string
	)
	deleteProject := func(id string) error {
		_, err := c.ProjectService.Delete(id)
		if err != nil {
			return err
		}
		fmt.Println("Project", projectID, "successfully deleted.")
		return nil
	}
	// deleteProjectCmd represents the deleteProject command
	deleteProjectCmd := &cobra.Command{
		Use:   `delete --id <project_UUID> [--force]`,
		Short: "Deletes a project.",
		Long:  "Deletes the specified project with a confirmation prompt. To skip the confirmation use --force. You can't delete a project that has active resources. You have to deprovision all servers and other infrastructure from a project in order to delete it.",
		Example: `  # Deletes project 50693ba9-e4e4-4d8a-9eb2-4840b11e9375:
  metal project delete -i 50693ba9-e4e4-4d8a-9eb2-4840b11e9375
  >
  ✔ Are you sure you want to delete project 50693ba9-e4e4-4d8a-9eb2-4840b11e9375: y
  
  # Deletes project 50693ba9-e4e4-4d8a-9eb2-4840b11e9375, skipping confirmation:
  metal project delete -i 50693ba9-e4e4-4d8a-9eb2-4840b11e9375 -f`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			if !force {
				prompt := promptui.Prompt{
					Label:     fmt.Sprintf("Are you sure you want to delete project %s: ", projectID),
					IsConfirm: true,
				}

				_, err := prompt.Run()
				if err != nil {
					return nil
				}
			}
			return errors.Wrap(deleteProject(projectID), "Could not delete Project")
		},
	}

	deleteProjectCmd.Flags().StringVarP(&projectID, "id", "i", "", "The project's UUID. This flag is required, unless specified in the config created by metal init or set as METAL_PROJECT_ID environment variable.")
	_ = deleteProjectCmd.MarkFlagRequired("id")

	deleteProjectCmd.Flags().BoolVarP(&force, "force", "f", false, "Force removal of the project")
	return deleteProjectCmd
}
