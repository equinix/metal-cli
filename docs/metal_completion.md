## metal completion

Generates completion scripts.

### Synopsis

Generates shell completion scripts for different shells.

```
metal completion [bash | zsh | fish | powershell]
```

### Examples

```
  # To load completions in Bash:
  source <(metal completion bash)
  
  # To load completions in Bash (3.2.x):
  eval "$(metal completion bash)"

  # To load completions in Bash for each session, on Linux execute once:
  metal completion bash > /etc/bash_completion.d/metal-cli
  
  # To load completions in Bash for each session, on Mac execute once:
  metal completion bash > /usr/local/etc/bash_completion.d/metal-cli

  # To load completions in Zsh:
  source <(metal completion zsh)
  
  # To load completions in Zsh for each session, execute once:
  metal completion zsh > "${fpath[1]}/_metal-cli"

  # To load completions in Fish:
  metal completion fish | source

  # To load completions in Fish for each session, execute once:
  metal completion fish > ~/.config/fish/completions/metal-cli.fish
```

### Options

```
  -h, --help   help for completion
```

### Options inherited from parent commands

```
      --config string         Path to JSON or YAML configuration file
      --exclude strings       Comma separated Href references to collapse in results, may be dotted three levels deep
      --filter stringArray    Filter 'get' actions with name value pairs. Filter is not supported by all resources and is implemented as request query parameters.
      --http-header strings   Headers to add to requests (in format key=value)
      --include strings       Comma separated Href references to expand in results, may be dotted three levels deep
  -o, --output string         Output format (*table, json, yaml). env output formats are (*sh, terraform, capp).
      --search string         Search keyword for use in 'get' actions. Search is not supported by all resources.
      --sort-by string        Sort fields for use in 'get' actions. Sort is not supported by all resources.
      --sort-dir string       Sort field direction for use in 'get' actions. Sort is not supported by all resources.
      --token string          Metal API Token (METAL_AUTH_TOKEN)
```

### SEE ALSO

* [metal](metal.md)	 - Command line interface for Equinix Metal

