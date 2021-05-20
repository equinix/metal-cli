## packet completion

Generate completion script

### Synopsis

To load completions:

	Bash:

	$ source <(packet completion bash)

	Bash (3.2.x):

	$ eval "$(packet completion bash)"

	# To load completions for each session, execute once:
	Linux:
	  $ packet completion bash > /etc/bash_completion.d/packet-cli
	MacOS:
	  $ packet completion bash > /usr/local/etc/bash_completion.d/packet-cli

	Zsh:

	$ source <(packet completion zsh)

	# To load completions for each session, execute once:
	$ packet completion zsh > "${fpath[1]}/_packet-cli"

	Fish:

	$ packet completion fish | source

	# To load completions for each session, execute once:
	$ packet completion fish > ~/.config/fish/completions/packet-cli.fish
	

```
packet completion [bash|zsh|fish|powershell]
```

### Options

```
  -h, --help   help for completion
```

### Options inherited from parent commands

```
      --config string     Path to JSON or YAML configuration file
      --exclude strings   Comma seperated Href references to collapse in results, may be dotted three levels deep
      --include strings   Comma seperated Href references to expand in results, may be dotted three levels deep
  -j, --json              JSON output
      --search string     Search keyword for use in 'get' actions. Search is not supported by all resources.
  -y, --yaml              YAML output
```

### SEE ALSO

* [packet](packet.md)	 - Command line interface for Equinix Metal

