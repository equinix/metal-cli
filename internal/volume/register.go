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
	"github.com/spf13/cobra"
)

type VolumeClient struct {
	VolumeService           packngo.VolumeService
	VolumeAttachmentService packngo.VolumeAttachmentService
}

func (client *VolumeClient) Register(rootCmd *cobra.Command, out output.Outputer) {
	// volumeCmd represents the volume command
	var volumeCmd = &cobra.Command{
		Use:     "volume",
		Aliases: []string{"volumes"},
		Short:   "Volume operations",
		Long:    `Volume operations: create, delete, attach, detach and get`,
	}

	rootCmd.AddCommand(volumeCmd)

	for _, makeCmd := range []func(*VolumeClient, output.Outputer) *cobra.Command{
		Create,
		Retrieve,
		Delete,
		Attach,
		Detach,
	} {
		volumeCmd.AddCommand(makeCmd(client, out))
	}
}
