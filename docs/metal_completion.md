## metal completion

Generate completion script

### Synopsis

To load completions:

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
	

```
metal completion [bash|zsh|fish|powershell]
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
      --token string      Metal API Token (METAL_AUTH_TOKEN)
  -y, --yaml              YAML output
```

### SEE ALSO

* [metal](metal.md)	 - Command line interface for Equinix Metal

