## packet ip get

Retrieves information about IP addresses, IP reservations and IP assignments

### Synopsis

Example:
	
To get all IP addresses under a project:

packet ip get --project-id [project_UUID] 

To get IP addresses by assignment id:

packet ip get --assignment-id [assignment_UUID]

To get IP addresses by reservation ID:

packet ip get --reservation-id [reservation_UUID]

	

```
packet ip get [flags]
```

### Options

```
  -a, --assignment-id string    UUID of the assignment
  -h, --help                    help for get
  -p, --project-id string       UUID of the project
  -r, --reservation-id string   UUID of the reservation
```

### Options inherited from parent commands

```
      --config string     Path to JSON or YAML configuration file
      --exclude strings   Comma seperated Href references to collapse in results, may be dotted three levels deep
      --include strings   Comma seperated Href references to expand in results, may be dotted three levels deep
  -j, --json              JSON output
  -y, --yaml              YAML output
```

### SEE ALSO

* [packet ip](packet_ip.md)	 - IP operations

