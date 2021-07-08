## metal capacity check

Validates if a deploy can be fulfilled.

### Synopsis

Example:

metal capacity check {-m [metro] | -f [facility]} -p [plan] -q [quantity]

	

```
metal capacity check [flags]
```

### Options

```
  -f, --facility string   Code of the facility
  -h, --help              help for check
  -m, --metro string      Code of the metro
  -p, --plan string       Name of the plan
  -q, --quantity int      Number of devices wanted
```

### Options inherited from parent commands

```
      --config string     Path to JSON or YAML configuration file
      --exclude strings   Comma seperated Href references to collapse in results, may be dotted three levels deep
      --include strings   Comma seperated Href references to expand in results, may be dotted three levels deep
  -j, --json              JSON output
      --search string     Search keyword for use in 'get' actions. Search is not supported by all resources.
      --token string      Metal API Token (METAL_AUTH_TOKEN)
  -y, --yaml              YAML output
```

### SEE ALSO

* [metal capacity](metal_capacity.md)	 - Capacities operations

