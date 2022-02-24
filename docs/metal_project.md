## metal project

Project operations. For more information on Equinix Metal Projects, visit https://metal.equinix.com/developers/docs/accounts/projects/.

### Synopsis

Project operations: create, get, update, and delete.

### Options

```
  -h, --help   help for project
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
* [metal project create](metal_project_create.md)	 - Creates a project.
* [metal project delete](metal_project_delete.md)	 - Deletes a project.
* [metal project get](metal_project_get.md)	 - Retrieves all the current user's projects or the details of a specified project.
* [metal project update](metal_project_update.md)	 - Updates a project.

