## packet organization get

Retrieves an organization or list of organizations

### Synopsis

Example:
	
To retrieve list of all available organizations:
packet organization get

To retrieve a single organization:
packet organization get -i [organization-id]

	

```
packet organization get [flags]
```

### Options

```
  -h, --help                     help for get
  -i, --organization-id string   UUID of the organization
```

### Options inherited from parent commands

```
      --config string     Path to JSON or YAML configuration file
      --exclude strings   Comma seperated Href references to collapse in results, may be dotted three levels deep
      --include strings   Comma seperated Href references to expand in results, may be dotted three levels deep
  -j, --json              JSON output
  -y, --yaml              YAML output
```

### SEE ALSO

* [packet organization](packet_organization.md)	 - Organization operations
* [packet organization get payment-methods](packet_organization_get_payment-methods.md)	 - Retrieves a list of payment methods for the organization

