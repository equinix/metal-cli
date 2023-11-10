## metal interconnections

interconnections operations: create, get, update, delete

### Synopsis

Get information on Metro locations. For more information on https://deploy.equinix.com/developers/docs/metal/interconnections.

### Options

```
  -h, --help   help for interconnections
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

* [metal](metal.md)	 - Command line interface for Equinix Metal
* [metal interconnections create](metal_interconnections_create.md)	 - Creates an interconnection.
* [metal interconnections delete](metal_interconnections_delete.md)	 - Deletes a interconnection.
* [metal interconnections get](metal_interconnections_get.md)	 - Retrieves interconnections for the current user, an organization, a project or the details of a specific interconnection.
* [metal interconnections update](metal_interconnections_update.md)	 - Updates a connection.

