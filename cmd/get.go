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
	"encoding/json"
	"fmt"
	"os"

	"github.com/ghodss/yaml"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get command",
	Args:  cobra.MinimumNArgs(1),
}

func init() {

	rootCmd.AddCommand(getCmd, createCmd)

	getCmd.AddCommand(retriveDeviceCmd, facilitiesCmd)
}

func output(in interface{}, header []string, data *[][]string) {
	// if header != nil && data != nil {
	if !isJSON && !isYaml {

		table := tablewriter.NewWriter(os.Stdout)
		table.SetAutoWrapText(false)
		table.SetAlignment(tablewriter.ALIGN_LEFT)
		table.SetHeader(header)
		table.AppendBulk(*data)
		table.Render()
	} else if isJSON {
		output, err := json.MarshalIndent(in, "", "  ")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(output))
	} else if isYaml {
		fmt.Println("*****isYaml", isYaml)
		output, err := yaml.Marshal(in)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(output))
	}
}
