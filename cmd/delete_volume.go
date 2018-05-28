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

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// deleteVolumeCmd represents the deleteVolume command
var deleteVolumeCmd = &cobra.Command{
	Use:   "volume",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if !force {
			prompt := promptui.Prompt{
				Label:     fmt.Sprintf("Are you sure you want to delete volume %s: ", volumeID),
				IsConfirm: true,
			}

			_, err := prompt.Run()
			if err != nil {
				return
			}

			err = deleteVolume(volumeID)
			if err != nil {
				fmt.Println("Client error:", err)
				return
			}
		} else {
			err := deleteVolume(volumeID)
			if err != nil {
				fmt.Println("Client error:", err)
				return
			}
		}
	},
}

func deleteVolume(id string) error {
	_, err := PacknGo.Volumes.Delete(id)
	if err != nil {
		return err
	}
	fmt.Println("Volume deletion initiated. Please check 'packet get volume -i", volumeID, "' for status")
	return nil
}

func init() {
	deleteCmd.AddCommand(deleteVolumeCmd)

	deleteVolumeCmd.Flags().StringVarP(&volumeID, "volume-id", "i", "", "--volume-id or -i [UUID]")
	deleteVolumeCmd.Flags().BoolVarP(&force, "force", "f", false, "--force or -f")
	deleteVolumeCmd.MarkFlagRequired("id")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteVolumeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteVolumeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
