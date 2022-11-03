## metal hardware-reservation get

Lists a Project's hardware reservations or the details of a specified hardware reservation.

### Synopsis

Lists a Project's hardware reservations or the details of a specified hardware reservation. When using --json or --yaml flags, the --include=project,facility,device flag is implied.

```
metal hardware-reservation get [-p <project_id>] | [-i <hardware_reservation_id>] [flags]
```

### Examples

```
  # Retrieve all hardware reservations of a project:
  metal hardware_reservations get -p $METAL_PROJECT_ID
  
  # Retrieve the details of a specific hardware reservation:
  metal hardware_reservations get -i 8404b73c-d18f-4190-8c49-20bb17501f88
```

### Options

```
  -h, --help                help for get
  -i, --id string           The UUID of a hardware reservation.
  -p, --project-id string   A project's UUID.
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

