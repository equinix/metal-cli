## metal organization update

Updates the specified organization.

### Synopsis

Updates the specified organization. You can update the name, website, Twitter, or logo.

```
metal organization update -i <organization_UUID> [-n <name>] [-d <description>] [-w <website_URL>] [-t <twitter_URL>] [-l <logo_URL>] [flags]
```

### Examples

```
  # Updates the name of an organization:
  metal organization update -i 3bd5bf07-6094-48ad-bd03-d94e8712fdc8 --name test-cluster02
```

### Options

```
  -d, --description string   User-friendly description of the organization.
  -h, --help                 help for update
  -i, --id string            An organization UUID.
  -l, --logo string          A logo image URL for the organization.
  -n, --name string          New name for the organization.
  -t, --twitter string       A Twitter URL of the organization.
  -w, --website string       A website URL for the organization.
```

### Options inherited from parent commands

```
      --config string         Path to JSON or YAML configuration file (METAL_CONFIG)
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

