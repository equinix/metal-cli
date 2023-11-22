## metal vrf get

Lists VRFs.

### Synopsis

Retrieves a list of all VRFs for the specified project or the details of the specified VRF ID. Either a project ID or a VRF ID is required.

```
metal vrf get -p <project_Id>  [flags]
```

### Examples

```
 # Gets the details of the specified device
  metal vrf get -v 3b0795ba-ec9a-4a9e-83a7-043e7e11407c

  # Lists VRFs for project 3b0795ba-ec9a-4a9e-83a7-043e7e11407c:
  metal vrf list -p 3b0795ba-ec9a-4a9e-83a7-043e7e11407c
```

### Options

```
  -h, --help                help for get
  -m, --metro string        Filter by Metro ID (uuid) or Metro Code
  -p, --project-id string   The project's UUID. This flag is required, unless specified in the config created by metal init or set as METAL_PROJECT_ID environment variable.
  -v, --vrf-id string       Specify the VRF UUID.
```

### Options inherited from parent commands

```
      --config string         Path to JSON or YAML configuration file (METAL_CONFIG)
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

* [metal vrf](metal_vrf.md)	 - VRF operations : create, get, delete

