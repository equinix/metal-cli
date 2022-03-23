## metal capacity check

Validates if the number of the specified server plan is available in the specified metro or facility.

### Synopsis

Validates if the number of the specified server plan is available in the specified metro or facility. Metro and facility are mutally exclusive. At least one metro (or facility), one plan, and quantity of 1 or more is required.

```
metal capacity check (-m <metro> | -f <facility>) -P <plan> -q <quantity> [flags]
```

### Examples

```
  # Checks if 10 c3.medium.x86 servers are available in NY or Dallas:
  metal capacity check -m ny,da -P c3.medium.x86 -q 10
  
  # Checks if Silicon Valley or Dallas has either 4 c3.medium.x86 or m3.large.x86
  metal capacity check -m sv,da -P c3.medium.x86,m3.large.x86 -q 4
```

### Options

```
  -f, --facilities strings   A facility or list of facilities.
  -h, --help                 help for check
  -m, --metros strings       A metro or list of metros.
  -P, --plans strings        A plan or list of plans.
  -q, --quantity int         The number of devices wanted.
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

