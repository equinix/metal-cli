## metal organization

Organization operations: create, get, update, payment-methods, and delete.

### Synopsis

Information and management of Organization-level settings. Documentation on organizations is in https://metal.equinix.com/developers/docs/accounts/organizations/.

### Options

```
  -h, --help   help for organization
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
* [metal organization create](metal_organization_create.md)	 - Creates an organization.
* [metal organization delete](metal_organization_delete.md)	 - Deletes an organization.
* [metal organization get](metal_organization_get.md)	 - Retrieves a list of organizations or the details of the specified organization.
* [metal organization payment-methods](metal_organization_payment-methods.md)	 - Retrieves a list of payment methods.
* [metal organization update](metal_organization_update.md)	 - Updates the specified organization.

