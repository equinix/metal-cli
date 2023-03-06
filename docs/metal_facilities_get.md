## metal facilities get

Retrieves a list of facilities.

### Synopsis

Retrieves a list of facilities available to the current user.

```
metal facilities get [flags]
```

### Examples

```
  # Lists facilities for current user:
  metal facilities get
```

### Options

```
  -h, --help   help for get
```

### Options inherited from parent commands

```
      --config string        Path to JSON or YAML configuration file
      --exclude strings      Comma separated Href references to collapse in results, may be dotted three levels deep
      --filter stringArray   Filter 'get' actions with name value pairs. Filter is not supported by all resources and is implemented as request query parameters.
      --include strings      Comma separated Href references to expand in results, may be dotted three levels deep
  -o, --output string        Output format (*table, json, yaml). env output formats are (*sh, terraform, capp).
      --search string        Search keyword for use in 'get' actions. Search is not supported by all resources.
      --sort-by string       Sort fields for use in 'get' actions. Sort is not supported by all resources.
      --sort-dir string      Sort field direction for use in 'get' actions. Sort is not supported by all resources.
      --token string         Metal API Token (METAL_AUTH_TOKEN)
```

### SEE ALSO

* [metal facilities](metal_facilities.md)	 - Facility operations: get.

