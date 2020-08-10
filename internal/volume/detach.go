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

	"github.com/packethost/packet-cli/internal/output"
	"github.com/packethost/packngo"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// detachCmd represents the detachVolume command
func detachCmd(svc packngo.VolumeAttachmentService, out output.Outputer, attachmentID *string) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		_, err := svc.Delete(*attachmentID)
		if err != nil {
			return errors.Wrap(err, "Could not detach Volume")
		}
		fmt.Println("Volume detachment initiated.")
		return nil
	}
}

func Detach(client *VolumeClient, out output.Outputer) *cobra.Command {
	var attachmentID string

	cmd := &cobra.Command{
		Use:   "detach",
		Short: "Detaches a volume from a device",
		Example: `
packet volume detach --id [attachment_UUID]`,
		RunE: detachCmd(client.VolumeAttachmentService, out, &attachmentID),
	}
	cmd.Flags().StringVarP(&attachmentID, "id", "i", "", "UUID of the attached volume")

	_ = cmd.MarkFlagRequired("id")
	return cmd
}
