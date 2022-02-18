## metal env

Generate environment variables

### Synopsis

Currently emitted variables:
	- METAL_AUTH_TOKEN
	- METAL_PROJECT_ID

	To load environment variables:

	Bash, Zsh:

	$ source <(metal env)

	Bash (3.2.x):

	$ eval "$(metal env)"

	Fish:

	$ metal env | source
	

```
metal env
```

### Options

```
  -h, --help                help for env
  -p, --project-id string   Project ID (METAL_PROJECT_ID)
```

### Options inherited from parent commands

```
      --config string        Path to JSON or YAML configuration file
      --exclude strings      Comma separated Href references to collapse in results, may be dotted three levels deep
      --filter stringArray   Filter 'get' actions with name value pairs. Filter is not supported by all resources and is implemented as request query parameters.
      --include strings      Comma separated Href references to expand in results, may be dotted three levels deep
  -o, --output string        Output format (*table, json, yaml)
      --search string        Search keyword for use in 'get' actions. Search is not supported by all resources.
      --sort-by string       Sort fields for use in 'get' actions. Sort is not supported by all resources.
      --sort-dir string      Sort field direction for use in 'get' actions. Sort is not supported by all resources.
      --token string         Metal API Token (METAL_AUTH_TOKEN)
```

### SEE ALSO

* [metal](metal.md)	 - Command line interface for Equinix Metal

