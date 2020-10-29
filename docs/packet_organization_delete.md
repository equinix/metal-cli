## packet organization delete

Deletes an organization

### Synopsis

Example:
	
packet organization delete -i [organization_UUID]

	

```
packet organization delete [flags]
```

### Options

```
  -f, --force                    Force removal of the organization
  -h, --help                     help for delete
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

* [packet organization](packet_organization.md)	 - Organization operations

