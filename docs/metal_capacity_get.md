## metal capacity get

Returns capacity of metros or facilities, with optional filtering.

### Synopsis

Returns the capacity of metros or facilities. Filters for metros, facilities, plans are available. Metro flags and facility flags are mutually exclusive. If no flags are included, returns capacity for all plans in all facilities.

```
metal capacity get [-m | -f] | [--metros <list> | --facilities <list>] [-P <list>] [flags]
```

### Examples

```
  # Returns the capacity of all plans in all facilities:
  metal capacity get 

  # Returns the capacity of the c3.small.x86 in all metros:
  metal capacity get -m -P c3.small.x86

  # Returns c3.large.arm and c3.medium.x86 capacity in the Silicon Valley, New York, and Dallas metros:
  metal capacity get --metros sv,ny,da -P c3.large.arm,c3.medium.x86
```

### Options

```
      --facilities strings   A facility or list of facilities for client-side filtering. Will only return the capacity for the specified facilities. Can not be used with --metros.
  -f, --facility             Return the capacity for all facilities. Can not be used with -m. (default true)
  -h, --help                 help for get
  -m, --metro                Return the capacity for all metros. Can not be used with -f.
      --metros strings       A metro or list of metros for client-side filtering. Will only return the capacity for the specified metros. Can not be used with --facilities.
  -P, --plans strings        Return only the capacity for the specified plans.
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

* [metal capacity](metal_capacity.md)	 - Capacity operations. For more information on capacity in metros, visit https://metal.equinix.com/developers/docs/locations/metros/. For more information on capacity in facilities, visit https://metal.equinix.com/developers/docs/locations/facilities/.

