## metal env

Prints or generates environment variables.

### Synopsis

Prints or generates environment variables. Currently emitted variables: METAL_AUTH_TOKEN, METAL_ORGANIZATION_ID, METAL_PROJECT_ID, METAL_CONFIG. Use the --project-id flag to set the METAL_PROJECT_ID variable. Use the --organization-id flag to set the METAL_ORGANIZATION_ID variable.

```
metal env [-p <project_id>]
```

### Examples

```
  # Print the current environment variables:
  metal env
  
  # Load environment variables in Bash, Zsh:
  source <(metal env)
  
  # Load environment variables in Bash 3.2.x:
  eval "$(metal env)"
  
  # Load environment variables in Fish:
  metal env | source
```

### Options

```
  -h, --help                     help for env
  -O, --organization-id string   A organization UUID to set as an environment variable.
  -p, --project-id string        A project UUID to set as an environment variable.
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

