## metal interconnections update

Updates a connection.

### Synopsis

Updates a specified connection.

```
metal interconnections update -i <connection_id> [flags]
```

### Examples

```
  # Updates a specified connection.:
  metal interconnections update --id 30c15082-a06e-4c43-bfc3-252616b46eba -n [<name>] -d [<description>] -r [<'redundant'|'primary'>]-m [<standard|tunnel>] -e [<E-mail>] --tags="tag1,tag2"
```

### Options

```
  -e, --contactEmail string   adds or updates the Email
  -d, --description string    Adds or updates the description for the interconnection.
  -h, --help                  help for update
  -i, --id string             The UUID of the interconnection.
  -m, --mode string           Adds or updates the mode for the interconnection.
  -n, --name string           The new name of the interconnection.
  -r, --redundancy string     Updating from 'redundant' to 'primary' will remove a secondary port, while updating from 'primary' to 'redundant' will add one.
  -t, --tags strings          Adds or updates the tags for the connection --tags="tag1,tag2".
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

