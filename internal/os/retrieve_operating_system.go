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
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// operatingSystemCmd represents the operatingSystem command
var retrieveOperatingSystemCmd = &cobra.Command{
	Use:   "get",
	Aliases: []string{"list"},
	Short: "Retrieves a list of available operating systems.",
	Long: `Example:
  metal operating-systems get`,
	RunE: func(cmd *cobra.Command, args []string) error {
		oss, _, err := apiClient.OperatingSystems.List()
		if err != nil {
			return errors.Wrap(err, "Could not list OperatingSystems")
		}

		data := make([][]string, len(oss))

		for i, os := range oss {
			data[i] = []string{os.Name, os.Slug, os.Distro, os.Version}
		}
		header := []string{"Name", "Slug", "Distro", "Version"}

		return output(oss, header, &data)
	},
}
