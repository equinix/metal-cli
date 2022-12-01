## metal project get

Retrieves all the current user's projects or the details of a specified project.

### Synopsis

Retrieves all the current user's projects or the details of a specified project. You can specify which project by UUID or name. When using `--json` or `--yaml` flags, the `--include=members` flag is implied.

```
metal project get [-i <project_UUID> | -n <project_name>] [flags]
```

### Examples

```
  # Retrieve all projects:
  metal project get
  
  # Retrieve a specific project by UUID: 
  metal project get -i 2008f885-1aac-406b-8d99-e6963fd21333

  # Retrieve a specific project by name:
  metal project get -n dev-cluster03
```

### Options

```
  -h, --help             help for get
  -i, --id string        The project's UUID, which can be specified in the config created by metal init or set as METAL_PROJECT_ID environment variable.
  -n, --project string   The name of the project.
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

* [metal project](metal_project.md)	 - Project operations: create, get, update, delete, and bgpenable, bgpconfig, bgpsessions.

