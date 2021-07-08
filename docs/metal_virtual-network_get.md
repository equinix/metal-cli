## metal virtual-network get

Retrieves a list of virtual networks for a single project.

### Synopsis

Example:

metal virtual-network get -p [project_UUID]

	

```
metal virtual-network get [flags]
```

### Options

```
  -h, --help                help for get
  -p, --project-id string   Project ID (METAL_PROJECT_ID)
```

### Options inherited from parent commands

```
      --config string     Path to JSON or YAML configuration file
      --exclude strings   Comma seperated Href references to collapse in results, may be dotted three levels deep
      --include strings   Comma seperated Href references to expand in results, may be dotted three levels deep
  -j, --json              JSON output
      --search string     Search keyword for use in 'get' actions. Search is not supported by all resources.
      --token string      Metal API Token (METAL_AUTH_TOKEN)
  -y, --yaml              YAML output
```

### SEE ALSO

* [metal virtual-network](metal_virtual-network.md)	 - Virtual network operations

