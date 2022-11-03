## metal ip unassign

Unassigns an IP address assignment.

### Synopsis

Unassigns an subnet and IP address assignment from a device by its assignment ID. 

```
metal ip unassign -i <assignment_UUID>  [flags]
```

### Examples

```
  # Unassigns an IP address assignment:
  metal ip unassign --id abd8674b-96c4-4271-92f5-2eaf5944c86f
```

### Options

```
  -h, --help        help for unassign
  -i, --id string   The UUID of the assignment.
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

* [metal ip](metal_ip.md)	 - IP address, reservations, and assignment operations: assign, unassign, remove, available, request and get.

