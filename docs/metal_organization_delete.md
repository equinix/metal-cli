## metal organization delete

Deletes an organization.

### Synopsis

Deletes an organization. You can not delete an organization that contains projects or has outstanding charges. Only organization owners can delete an organization.

```
metal organization delete -i <organization_UUID> [flags]
```

### Examples

```
  # Deletes an organization, with confirmation: 
  metal organization delete -i 3bd5bf07-6094-48ad-bd03-d94e8712fdc8
  >
  ✔ Are you sure you want to delete organization 3bd5bf07-6094-48ad-bd03-d94e8712fdc8: y
  
  # Deletes an organization, skipping confirmation:
  metal organization delete -i 3bd5bf07-6094-48ad-bd03-d94e8712fdc8 -f
```

### Options

```
  -f, --force                    Skips confirmation for the removal of the organization.
  -h, --help                     help for delete
  -i, --organization-id string   The UUID of the organization.
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

* [metal organization](metal_organization.md)	 - Organization operations: create, get, update, payment-methods, and delete.

