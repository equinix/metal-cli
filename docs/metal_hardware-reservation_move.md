## metal hardware-reservation move

Moves a hardware reservation.

### Synopsis

Moves a hardware reservation to a specified project. Both the hardware reservation ID and the Project ID for the destination project are required.

```
metal hardware-reservation move -i <hardware_reservation_id> -p <project_id> [flags]
```

### Examples

```
  # Moves a hardware reservation to the specified Project:
  metal hardware_reservation move -i 8404b73c-d18f-4190-8c49-20bb17501f88 -p 278bca90-f6b2-4659-b1a4-1bdffa0d80b7
```

### Options

```
  -h, --help                help for move
  -i, --id string           The UUID of the hardware reservation.
  -p, --project-id string   The Project ID of the Project you are moving the hardware reservation to.
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

* [metal hardware-reservation](metal_hardware-reservation.md)	 - Hardware reservation operations: get, move.

