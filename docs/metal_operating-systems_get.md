## metal operating-systems get

Retrieves a list of operating systems.

### Synopsis

Retrieves a list of operating systems available to the current user. Response includes the operating system's slug, distro, version, and name.

```
metal operating-systems get [flags]
```

### Examples

```
  # Lists the operating systems available to the current user:
  metal operating-systems get
```

### Options

```
  -h, --help   help for get
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

* [metal operating-systems](metal_operating-systems.md)	 - Operating system operations: get.

