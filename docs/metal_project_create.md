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
  -o, --organization-id string     UUID of the organization
  -m, --payment-method-id string   UUID of the payment method
```

### Options inherited from parent commands

```
      --config string     Path to JSON or YAML configuration file
      --exclude strings   Comma seperated Href references to collapse in results, may be dotted three levels deep
      --include strings   Comma seperated Href references to expand in results, may be dotted three levels deep
  -j, --json              JSON output
      --search string     Search keyword for use in 'get' actions. Search is not supported by all resources.
  -y, --yaml              YAML output
```

### SEE ALSO

* [metal project](metal_project.md)	 - Project operations

