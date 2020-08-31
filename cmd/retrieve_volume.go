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
	"strconv"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	volumeID string
)

// retriveVolumeCmd represents the retriveVolume command
var retriveVolumeCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieves a volume list or volume details.",
	Long: `Example:
	
Retrieve the list of volumes:
packet volume get --project-id [project_UUID]
  
Retrieve a specific volume:
packet volume get --id [volume_UUID]

`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if projectID != "" && volumeID != "" {
			return fmt.Errorf("Either id or project-id can be set.")
		} else if projectID == "" && volumeID == "" {
			return fmt.Errorf("Either id or project-id should be set.")
		} else if projectID != "" {
			volumes, _, err := PacknGo.Volumes.List(projectID, listOptions(nil, nil))
			if err != nil {
				return errors.Wrap(err, "Could not list Volumes")
			}
			data := make([][]string, len(volumes))

			for i, v := range volumes {
				data[i] = []string{v.ID, v.Name, strconv.Itoa(v.Size), v.State, v.Created}
			}
			header := []string{"ID", "Name", "Size", "State", "Created"}

			return output(volumes, header, &data)
		} else if volumeID != "" {

			v, _, err := PacknGo.Volumes.Get(volumeID, nil)
			if err != nil {
				return errors.Wrap(err, "Could not get Volume")
			}

			header := []string{"ID", "Name", "Size", "State", "Created"}
			data := make([][]string, 1)
			data[0] = []string{v.ID, v.Name, strconv.Itoa(v.Size), v.State, v.Created}

			return output(v, header, &data)
		}
		return nil
	},
}

func init() {
	retriveVolumeCmd.Flags().StringVarP(&projectID, "project-id", "p", "", "UUID of the project")
	viper.BindPFlag("project-id", retriveVolumeCmd.Flags().Lookup("project-id"))

	retriveVolumeCmd.Flags().StringVarP(&volumeID, "id", "i", "", "UUID of the volume")
}
