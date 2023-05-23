## metal capacity get

Returns capacity of metros, with optional filtering.

### Synopsis

Returns the capacity of metros. Filters for metros, plans are available. Metro flags are mutually exclusive. If no flags are included, returns capacity for all plans in all metros.

```
metal capacity get [-m | -f] | [--metros <list> ] [-P <list>] [flags]
```

### Examples

```
  # Returns the capacity of all plans in all metros:
  metal capacity get 

  # Returns the capacity of the c3.small.x86 in all metros:
  metal capacity get -m -P c3.small.x86

  # Returns c3.large.arm and c3.medium.x86 capacity in the Silicon Valley, New York, and Dallas metros:
  metal capacity get --metros sv,ny,da -P c3.large.arm,c3.medium.x86
```

### Options

```
  -h, --help                 help for get
  -m, --metro                Return the capacity for all metros. Can not be used with -f.
      --metros strings       A metro or list of metros for client-side filtering. Will only return the capacity for the specified metros.
  -P, --plans strings        Return only the capacity for the specified plans.
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

* [metal capacity](metal_capacity.md)	 - Capacity operations: get, check

