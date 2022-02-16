## metal capacity check

Validates if a deploy can be fulfilled with the given quantity in any of the given locations and plans

```
metal capacity check {-m [metros,...] | -f [facilities,...]} -P [plans,...] -q [quantity] [flags]
```

### Examples

```
metal capacity check -m sv,ny,da -P c3.large.arm,c3.medium.x86 -q 10
```

### Options

```
  -f, --facilities strings   Codes of the facilities
  -h, --help                 help for check
  -m, --metros strings       Codes of the metros
  -P, --plans strings        Names of the plans
  -q, --quantity int         Number of devices wanted
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

