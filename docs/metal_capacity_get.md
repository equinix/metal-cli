## metal capacity get

Returns a list of facilities or metros and plans with their current capacity, with filtering.

```
metal capacity get [[-m | -f] | [--metros metros,... | --facilities facilities,...]] [-P plans,...] [flags]
```

### Examples

```
metal capacity get -m sv,ny,da -P c3.large.arm,c3.medium.x86
```

### Options

```
      --facilities strings   Codes of the facilities (client side filtering)
  -f, --facility             Report all facilites (default true)
  -h, --help                 help for get
  -m, --metro                Report all metros
      --metros strings       Codes of the metros (client side filtering)
  -P, --plans strings        Names of the plans
```

### Options inherited from parent commands

```
      --config string        Path to JSON or YAML configuration file
      --exclude strings      Comma separated Href references to collapse in results, may be dotted three levels deep
      --filter stringArray   Filter 'get' actions with name value pairs. Filter is not supported by all resources and is implemented as request query parameters.
      --include strings      Comma separated Href references to expand in results, may be dotted three levels deep
  -o, --output string        Output format (*table, json, yaml)
      --search string        Search keyword for use in 'get' actions. Search is not supported by all resources.
      --sort-by string       Sort fields for use in 'get' actions. Sort is not supported by all resources.
      --sort-dir string      Sort field direction for use in 'get' actions. Sort is not supported by all resources.
      --token string         Metal API Token (METAL_AUTH_TOKEN)
```

### SEE ALSO

* [metal capacity](metal_capacity.md)	 - Capacities operations

