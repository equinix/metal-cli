## metal interconnections get

Retrieves interconnections for the current user, an organization, a project or the details of a specific interconnection.

### Synopsis

Retrieves interconnections for the current user, an organization, a project or the details of a specific interconnection.

```
metal interconnections get [flags]
```

### Examples

```
  # Retrieve all interconnections of a current user::
  
  # Retrieve the details of a specific interconnection:
  metal interconnections get -i e9a969b3-8911-4667-9d99-57cd3dd4ef6f

  # Retrieve all the interconnection of an organization:
  metal interconnections get -O c079178c-9557-48f2-9ce7-cfb927b81928

  # Retrieve all interconnection of a project:
  metal interconnections get -p 1867ee8f-6a11-470a-9505-952d6a324040 
```

### Options

```
  -i, --connID string           UUID of the interconnection
  -h, --help                    help for get
  -O, --organizationID string   UUID of the organization
  -p, --projectID string        Project ID (METAL_PROJECT_ID)
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

* [metal interconnections](metal_interconnections.md)	 - interconnections operations: create, get, update, delete

