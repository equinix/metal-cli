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

package metros

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

func (c *Client) Retrieve() *cobra.Command {
	// metrosCmd represents the metros command
	retrieveMetrosCmd := &cobra.Command{
		Use:     "get",
		Aliases: []string{"list"},
		Short:   "Retrieves a list of metros.",
		Long:    "Retrieves a list of metros available to the current user.",
		Example: `  # Lists metros available to the current user:	
  metal metros get`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			metrosList, _, err := c.Service.FindMetros(context.Background()).Execute()
			if err != nil {
				return fmt.Errorf("could not list metros: %w", err)
			}
			metros := metrosList.GetMetros()
			data := make([][]string, len(metros))

			for i, metro := range metros {
				data[i] = []string{metro.GetId(), metro.GetName(), metro.GetCode()}
			}
			header := []string{"ID", "Name", "Code"}

			return c.Out.Output(metros, header, &data)
		},
	}
	return retrieveMetrosCmd
}
