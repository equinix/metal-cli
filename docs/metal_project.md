## metal project

Project operations: create, get, update, delete, and bgp-enable, bgp-config, bgp-sessions.

### Synopsis

Information and management for Projects and Project-level BGP. Documentation on Projects is on https://metal.equinix.com/developers/docs/accounts/projects/, and documentation on BGP is on https://metal.equinix.com/developers/docs/bgp/bgp-on-equinix-metal/.

### Options

```
  -h, --help   help for project
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

* [metal](metal.md)	 - Command line interface for Equinix Metal
* [metal project bgp-config](metal_project_bgp-config.md)	 - Gets BGP Config for a project.
* [metal project bgp-enable](metal_project_bgp-enable.md)	 - Enables BGP on a project.
* [metal project bgp-sessions](metal_project_bgp-sessions.md)	 - Gets BGP Sessions for a project.
* [metal project create](metal_project_create.md)	 - Creates a project.
* [metal project delete](metal_project_delete.md)	 - Deletes a project.
* [metal project get](metal_project_get.md)	 - Retrieves all the current user's projects or the details of a specified project.
* [metal project update](metal_project_update.md)	 - Updates a project.

