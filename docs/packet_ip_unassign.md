## packet ip unassign

Unassigns an IP address.

### Synopsis

Example:

packet ip unassign --id [assignment-UUID]

	

```
packet ip unassign [flags]
```

### Options

```
  -h, --help        help for unassign
  -i, --id string   UUID of the assignment
```

### Options inherited from parent commands

```
      --config string     Path to JSON or YAML configuration file
      --exclude strings   Comma seperated Href references to collapse in results, may be dotted three levels deep
      --include strings   Comma seperated Href references to expand in results, may be dotted three levels deep
  -j, --json              JSON output
      --search string     Search keyword for use in 'get' actions. Search is not supported by all resources.
  -y, --yaml              YAML output
```

### SEE ALSO

* [packet ip](packet_ip.md)	 - IP operations

