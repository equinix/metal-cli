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

* [metal organization](metal_organization.md)	 - Organization operations

