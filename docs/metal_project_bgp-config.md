## metal project bgp-config

Gets BGP Config for a project.

### Synopsis

Gets BGP Config for a project.

```
metal project bgp-config --project-id <project_UUID> [flags]
```

### Examples

```
  metal project bgp-config --project-id 50693ba9-e4e4-4d8a-9eb2-4840b11e9375 -d 
```

### Options

```
  -h, --help                help for bgp-config
  -p, --project-id string   Project ID (METAL_PROJECT_ID)
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

* [metal project](metal_project.md)	 - Project operations: create, get, update, delete, and bgp-enable, bgp-config, bgp-sessions.

