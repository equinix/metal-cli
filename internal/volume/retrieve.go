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

package volume

import (
	"fmt"
	"strconv"

	"github.com/packethost/packet-cli/internal/output"
	"github.com/packethost/packngo"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func retrieveCmd(svc packngo.VolumeService, out output.Outputer, projectID, volumeID *string, isJSON, isYaml *bool) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		out.SetFormat(output.FormatSwitch(*isJSON, *isYaml))
		if *projectID != "" && *volumeID != "" {
			return fmt.Errorf("Either id or project-id can be set.")
		} else if *projectID == "" && *volumeID == "" {
			return fmt.Errorf("Either id or project-id  should be set.")
		} else if *projectID != "" {
			volumes, _, err := svc.List(*projectID, nil)
			if err != nil {
				return errors.Wrap(err, "Could not list Volumes")
			}
			data := make([][]string, len(volumes))

			for i, v := range volumes {
				data[i] = []string{v.ID, v.Name, strconv.Itoa(v.Size), v.State, v.Created}
			}
			header := []string{"ID", "Name", "Size", "State", "Created"}

			return out.Output(volumes, header, &data)
		} else if *volumeID != "" {

			v, _, err := svc.Get(*volumeID, nil)
			if err != nil {
				return errors.Wrap(err, "Could not get Volume")
			}

			header := []string{"ID", "Name", "Size", "State", "Created"}
			data := make([][]string, 1)
			data[0] = []string{v.ID, v.Name, strconv.Itoa(v.Size), v.State, v.Created}

			return out.Output(v, header, &data)
		}
		return nil
	}
}

func Retrieve(client *VolumeClient, out output.Outputer) *cobra.Command {
	var (
		volumeID, projectID string
		isJSON, isYaml      bool
	)
	retriveVolumeCmd := retrieveCmd(client.VolumeService, out, &projectID, &volumeID, &isJSON, &isYaml)
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Retrieves a volume list or volume details.",
		Example: `
Retrieve the list of volumes:
packet volume get --project-id [project_UUID]
Retrieve a specific volume:
packet volume get --id [volume_UUID]`,
		RunE: retriveVolumeCmd,
	}
	cmd.Flags().StringVarP(&projectID, "project-id", "p", "", "UUID of the project")
	cmd.Flags().StringVarP(&volumeID, "id", "i", "", "UUID of the volume")

	cmd.PersistentFlags().BoolVarP(&isJSON, "json", "j", false, "JSON output")
	cmd.PersistentFlags().BoolVarP(&isYaml, "yaml", "y", false, "YAML output")
	return cmd
}
