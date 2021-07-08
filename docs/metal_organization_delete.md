## metal organization delete

Deletes an organization

### Synopsis

Example:
	
metal organization delete -i [organization_UUID]

	

```
metal organization delete [flags]
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
      --token string      Metal API Token (METAL_AUTH_TOKEN)
  -y, --yaml              YAML output
```

### SEE ALSO

* [metal organization](metal_organization.md)	 - Organization operations

