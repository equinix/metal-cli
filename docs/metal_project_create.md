## metal project create

Creates a project.

### Synopsis

Creates a project with the specified name. If no organization is specified, the project is created in the current user's default organization. If no payment method is specified the organization's default payment method is used.

```
metal project create -n <project_name> [-O <organization_UUID>] [-m <payment_method_UUID>] [flags]
```

### Examples

```
  # Creates a new project named dev-cluster02: 
  metal project create --name dev-cluster02
  
  # Creates a new project named dev-cluster03 in the specified organization with a payment method:
  metal project create -n dev-cluster03 -O 814b09ca-0d0c-4656-9de0-4ce65c6faf70 -m ab1fbdaa-8b25-4c3e-8360-e283852e3747
```

### Options

```
  -h, --help                       help for create
  -n, --name string                Name of the project
  -O, --organization-id string     The UUID of the organization.
  -m, --payment-method-id string   The UUID of the payment method.
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

* [metal project](metal_project.md)	 - Project operations: create, get, update, delete, and bgpenable, bgpconfig, bgpsessions.

