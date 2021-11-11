## metal project create

Creates a project

### Synopsis

Example:

metal project create --name [project_name]
  
  

```
metal project create [flags]
```

### Options

```
  -h, --help                       help for create
  -n, --name string                Name of the project
  -O, --organization-id string     UUID of the organization
  -m, --payment-method-id string   UUID of the payment method
```

### Options inherited from parent commands

```
      --config string     Path to JSON or YAML configuration file
      --exclude strings   Comma seperated Href references to collapse in results, may be dotted three levels deep
      --include strings   Comma seperated Href references to expand in results, may be dotted three levels deep
  -o, --output string     Output format (*table, json, yaml)
      --search string     Search keyword for use in 'get' actions. Search is not supported by all resources.
      --sort-by string    Sort fields for use in 'get' actions. Sort is not supported by all resources.
      --sort-dir string   Sort field direction for use in 'get' actions. Sort is not supported by all resources.
      --token string      Metal API Token (METAL_AUTH_TOKEN)
```

### SEE ALSO

* [metal project](metal_project.md)	 - Project operations

