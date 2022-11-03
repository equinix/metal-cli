## metal organization get

Retrieves a list of organizations or the details of the specified organization.

### Synopsis

Retrieves a list of organizations or the details of the specified organization. Details of an organization are only available to its members.

```
metal organization get -i <organization_UUID> [flags]
```

### Examples

```
  # Retrieves list of the current user's organizations:
  metal organization get

  # Retrieves details of an organization:
  metal organization get -i 3bd5bf07-6094-48ad-bd03-d94e8712fdc8
```

### Options

```
  -h, --help                     help for get
  -i, --organization-id string   The UUID of an organization.
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

* [metal organization](metal_organization.md)	 - Organization operations: create, get, update, payment-methods, and delete.

