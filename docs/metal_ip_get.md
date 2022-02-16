## metal ip get

Retrieves information about IP addresses, IP reservations and IP assignments

### Synopsis

Example:
	
To get all IP addresses under a project:

metal ip get --project-id [project_UUID] 

To get IP addresses by assignment id:

metal ip get --assignment-id [assignment_UUID]

To get IP addresses by reservation ID:

metal ip get --reservation-id [reservation_UUID]

	

```
metal ip get [flags]
```

### Options

```
  -a, --assignment-id string    UUID of the assignment
  -h, --help                    help for get
  -p, --project-id string       Project ID (METAL_PROJECT_ID)
  -r, --reservation-id string   UUID of the reservation
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

* [metal ip](metal_ip.md)	 - IP operations

