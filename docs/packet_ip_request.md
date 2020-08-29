## packet ip request

Request an IP block

### Synopsis

Example:

packet ip request --quantity [quantity] --facility [facility_code] --type [address_type]

	

```
packet ip request [flags]
```

### Options

```
  -c, --comments string     General comments
  -f, --facility string     Code of the facility
  -h, --help                help for request
  -p, --project-id string   UUID of the project
  -q, --quantity int        Number of IP addresses to reserve
  -t, --type string         Address type public_ipv4 or global_ipv6
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

