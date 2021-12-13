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
package completion

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	// completionCmd represents the completion command that, when run, generates a
	// bash or zsh completion script for the CLI
	var completionCmd = &cobra.Command{
		Use:   "completion [bash|zsh|fish|powershell]",
		Short: "Generate completion script",
		Long: `To load completions:

	Bash:

	$ source <(metal completion bash)

	Bash (3.2.x):

	$ eval "$(metal completion bash)"

	# To load completions for each session, execute once:
	Linux:
	  $ metal completion bash > /etc/bash_completion.d/metal-cli
	MacOS:
	  $ metal completion bash > /usr/local/etc/bash_completion.d/metal-cli

	Zsh:

	$ source <(metal completion zsh)

	# To load completions for each session, execute once:
	$ metal completion zsh > "${fpath[1]}/_metal-cli"

	Fish:

	$ metal completion fish | source

	# To load completions for each session, execute once:
	$ metal completion fish > ~/.config/fish/completions/metal-cli.fish
	`,
		DisableFlagsInUseLine: true,
		ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
		Args:                  cobra.ExactValidArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			switch args[0] {
			case "bash":
				return cmd.Root().GenBashCompletion(os.Stdout)
			case "zsh":
				return cmd.Root().GenZshCompletion(os.Stdout)
			case "fish":
				return cmd.Root().GenFishCompletion(os.Stdout, true)
			case "powershell":
				return cmd.Root().GenPowerShellCompletion(os.Stdout)
			}
			// ValidArgs make this error response dead-code
			return fmt.Errorf("unknown shell: %q", args[0])
		},
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			return nil
		},
	}

	return completionCmd
}
