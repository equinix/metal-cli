## metal ssh-key

SSH key operations

### Synopsis

SSH key operations: create, delete, update and get

### Options

```
  -h, --help   help for ssh-key
```

### Options inherited from parent commands

```
      --config string        Path to JSON or YAML configuration file
      --exclude strings      Comma seperated Href references to collapse in results, may be dotted three levels deep
      --filter stringArray   Filter 'get' actions with name value pairs. Filter is not supported by all resources and is implemented as request query parameters.
      --include strings      Comma seperated Href references to expand in results, may be dotted three levels deep
  -o, --output string        Output format (*table, json, yaml)
      --search string        Search keyword for use in 'get' actions. Search is not supported by all resources.
      --sort-by string       Sort fields for use in 'get' actions. Sort is not supported by all resources.
      --sort-dir string      Sort field direction for use in 'get' actions. Sort is not supported by all resources.
      --token string         Metal API Token (METAL_AUTH_TOKEN)
```

### SEE ALSO

* [metal](metal.md)	 - Command line interface for Equinix Metal
* [metal ssh-key create](metal_ssh-key_create.md)	 - Creates an SSH key
* [metal ssh-key delete](metal_ssh-key_delete.md)	 - Deletes an SSH key
* [metal ssh-key get](metal_ssh-key_get.md)	 - Retrieves a list of available SSH keys or a single SSH key
* [metal ssh-key update](metal_ssh-key_update.md)	 - Updates a project

