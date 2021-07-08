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
      --config string     Path to JSON or YAML configuration file
      --exclude strings   Comma seperated Href references to collapse in results, may be dotted three levels deep
      --include strings   Comma seperated Href references to expand in results, may be dotted three levels deep
  -j, --json              JSON output
      --search string     Search keyword for use in 'get' actions. Search is not supported by all resources.
      --token string      Metal API Token (METAL_AUTH_TOKEN)
  -y, --yaml              YAML output
```

### SEE ALSO

* [metal ip](metal_ip.md)	 - IP operations

