## metal organization update

Updates an organization

### Synopsis

Example:

metal organization update --id [organization_UUID] --name [new_name]



```
metal organization update [flags]
```

### Options

```
  -d, --description string   Description of the organization
  -h, --help                 help for update
  -i, --id string            Organization ID
  -l, --logo string          Logo URL of the organization
  -n, --name string          Name of the organization
  -t, --twitter string       Twitter URL of the organization
  -w, --website string       Website of the organization
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

* [metal organization](metal_organization.md)	 - Organization operations

