/*
Copyright © 2020 Equinix Metal Developers <support@equinixmetal.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package docs

import (
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use:                   "docs [DESTINATION]",
		Short:                 "Generate command documentation",
		Long:                  "To generate documentation in the ./docs directory: docs ./docs",
		DisableFlagsInUseLine: true,
		Args:                  cobra.ExactValidArgs(1),
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			dest := args[0]
			return doc.GenMarkdownTree(cmd.Parent(), dest)
		},
	}
}
