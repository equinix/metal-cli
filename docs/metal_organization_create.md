## metal organization create

Creates an organization.

### Synopsis

Creates a new organization with the current user as the organization's owner. 

```
metal organization create -n <name> [-d <description>] [-w <website_URL>] [-t <twitter_URL>] [-l <logo_URL>] [flags]
```

### Examples

```
  # Creates a new organization named "it-backend-infra": 
  metal organization create -n it-backend-infra
  
  # Creates a new organization with name, website, and twitter:
  metal organization create -n test-org -w www.metal.equinix.com -t https://twitter.com/equinixmetal 
```

### Options

```
  -d, --description string   Description of the organization.
  -h, --help                 help for create
  -l, --logo string          A Logo image URL.]
  -n, --name string          Name of the organization.
  -t, --twitter string       Twitter URL of the organization.
  -w, --website string       Website URL of the organization.
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

