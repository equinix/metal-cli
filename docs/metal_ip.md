## metal ip

IP operations

### Synopsis

IP address, reservations and assignment operations: assign, unassign, remove, available, request and get 

### Options

```
  -h, --help   help for ip
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

* [metal](metal.md)	 - Command line interface for Equinix Metal
* [metal ip assign](metal_ip_assign.md)	 - Assigns an IP address to a given device
* [metal ip available](metal_ip_available.md)	 - Retrieves a list of IP resevations for a single project.
* [metal ip get](metal_ip_get.md)	 - Retrieves information about IP addresses, IP reservations and IP assignments
* [metal ip remove](metal_ip_remove.md)	 - Command to remove IP reservation.
* [metal ip request](metal_ip_request.md)	 - Request an IP block
* [metal ip unassign](metal_ip_unassign.md)	 - Unassigns an IP address.

