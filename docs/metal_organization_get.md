## metal organization get

Retrieves an organization or list of organizations

### Synopsis

Example:
	
To retrieve list of all available organizations:
metal organization get

To retrieve a single organization:
metal organization get -i [organization-id]

	

```
metal organization get [flags]
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
      --search string     Search keyword for use in 'get' actions. Search is not supported by all resources.
  -y, --yaml              YAML output
```

### SEE ALSO

* [metal organization](metal_organization.md)	 - Organization operations

