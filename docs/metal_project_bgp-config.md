## metal project bgp-config

Gets BGP Config for a project.

### Synopsis

Gets BGP Config for a project.

```
metal project bgp-config --id <project_UUID> [flags]
```

### Examples

```
  # Get BGP config for project 50693ba9-e4e4-4d8a-9eb2-4840b11e9375:
  metal project bgp-config --id 50693ba9-e4e4-4d8a-9eb2-4840b11e9375
```

### Options

```
  -h, --help        help for bgp-config
  -i, --id string   Project ID (METAL_PROJECT_ID)
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

* [metal project](metal_project.md)	 - Project operations. For more information on Equinix Metal Projects, visit https://metal.equinix.com/developers/docs/accounts/projects/.

