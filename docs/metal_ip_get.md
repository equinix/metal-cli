## metal ip get

Retrieves information about IP addresses, IP address reservations, and IP address assignments.

### Synopsis

Retrieves information about the IP addresses in a project, the IP addresses that are in a specified assignment, or the IP addresses that are in a specified reservation.

```
metal ip get -p <project_UUID> | -a <assignment_UUID> | -r <reservation_UUID> [flags]
```

### Examples

```
  # Lists all IP addresses in a project:
  metal ip get -p bb73aa19-c216-4ce2-a613-e5ca93732722 

  # Gets information about the IP addresses from an assignment ID:
  metal ip get -a bb526d47-8536-483c-b436-116a5fb72235

  # Gets the IP addresses from a reservation ID:
  metal ip get -r da1bb048-ea6e-4911-8ab9-b95635ca127a
```

### Options

```
  -a, --assignment-id string    UUID of an IP address assignment. When you assign an IP address to a server, it gets an assignment UUID.
  -h, --help                    help for get
  -p, --project-id string       A Project UUID (METAL_PROJECT_ID).
  -r, --reservation-id string   UUID of an IP address reservation.
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

