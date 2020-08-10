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

	"github.com/manifoldco/promptui"
	"github.com/packethost/packet-cli/internal/output"
	"github.com/packethost/packngo"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// deleteVolumeCmd represents the deleteVolume command
func deleteCmd(svc packngo.VolumeService, out output.Outputer, volumeID *string, force *bool) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		if !*force {
			prompt := promptui.Prompt{
				Label:     fmt.Sprintf("Are you sure you want to delete the volume %s: ", *volumeID),
				IsConfirm: true,
			}

			_, err := prompt.Run()
			if err != nil {
				return err
			}
		}
		return errors.Wrap(deleteVolume(svc, *volumeID), "Could not delete Volume")
	}
}

func deleteVolume(svc packngo.VolumeService, volumeID string) error {
	_, err := svc.Delete(volumeID)
	if err != nil {
		return err
	}
	fmt.Println("Volume deletion initiated. Please check 'packet volume get -i", volumeID, "' for status")
	return nil
}

func Delete(client *VolumeClient, out output.Outputer) *cobra.Command {
	var (
		volumeID string
		force    bool
	)
	deleteVolumeCmd := deleteCmd(client.VolumeService, out, &volumeID, &force)
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Deletes a volume",
		Example: `
packet volume delete --id [volume_UUID]`,
		RunE: deleteVolumeCmd,
	}
	cmd.Flags().StringVarP(&volumeID, "id", "i", "", "UUID of volume")
	cmd.Flags().BoolVarP(&force, "force", "f", false, "Force removal of the volume")

	_ = cmd.MarkFlagRequired("id")
	return cmd
}
