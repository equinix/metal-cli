## packet env

Generate environment variables

### Synopsis

Currently emitted variables:
	- METAL_AUTH_TOKEN

	To load environment variables:

	Bash, Zsh:

	$ source <(packet env)

	Bash (3.2.x):

	$ eval "$(packet env)"

	Fish:

	$ packet env | source
	

```
packet env
```

### Options

```
  -h, --help   help for env
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

