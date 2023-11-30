## metal project delete

Deletes a project.

### Synopsis

Deletes the specified project with a confirmation prompt. To skip the confirmation use --force. You can't delete a project that has active resources. You have to deprovision all servers and other infrastructure from a project in order to delete it.

```
metal project delete --id <project_UUID> [--force] [flags]
```

### Examples

```
  # Deletes project 50693ba9-e4e4-4d8a-9eb2-4840b11e9375:
  metal project delete -i 50693ba9-e4e4-4d8a-9eb2-4840b11e9375
  >
  ✔ Are you sure you want to delete project 50693ba9-e4e4-4d8a-9eb2-4840b11e9375: y
  
  # Deletes project 50693ba9-e4e4-4d8a-9eb2-4840b11e9375, skipping confirmation:
  metal project delete -i 50693ba9-e4e4-4d8a-9eb2-4840b11e9375 -f
```

### Options

```
  -f, --force       Force removal of the project
  -h, --help        help for delete
  -i, --id string   The project's UUID. This flag is required, unless specified in the config created by metal init or set as METAL_PROJECT_ID environment variable.
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

