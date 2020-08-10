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
	"github.com/packethost/packet-cli/internal/output"
	"github.com/packethost/packngo"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func attachCmd(svc packngo.VolumeAttachmentService, out output.Outputer, volumeID, deviceID *string) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		a, _, err := svc.Create(*volumeID, *deviceID)
		if err != nil {
			return errors.Wrap(err, "Could not create volume attachment")
		}

		header := []string{"ID"}
		data := make([][]string, 1)
		data[0] = []string{a.ID}

		return out.Output(a, header, &data)
	}
}

func Attach(client *VolumeClient, out output.Outputer) *cobra.Command {
	var volumeID, deviceID string
	attachVolumeCmd := attachCmd(client.VolumeAttachmentService, out, &volumeID, &deviceID)

	cmd := &cobra.Command{
		Use:   "attach",
		Short: "Attaches a volume to a device.",
		Example: `
packet volume attach --id [volume_UUID] --device-id [device_UUID]`,
		RunE: attachVolumeCmd,
	}

	cmd.Flags().StringVarP(&volumeID, "id", "i", "", "UUID of the volume")
	cmd.Flags().StringVarP(&deviceID, "device-id", "d", "", "UUID of the device")

	_ = cmd.MarkFlagRequired("id")
	_ = cmd.MarkFlagRequired("device-id")

	return cmd
}
