## metal ip remove

Command to remove IP reservation.

### Synopsis

Example:	

metal ip remove --id [reservation-UUID]



```
metal ip remove [flags]
```

### Options

```
  -h, --help        help for remove
  -i, --id string   UUID of the reservation
```

### Options inherited from parent commands

```
      --config string     Path to JSON or YAML configuration file
      --exclude strings   Comma seperated Href references to collapse in results, may be dotted three levels deep
      --include strings   Comma seperated Href references to expand in results, may be dotted three levels deep
  -o, --output string     Output format (*table, json, yaml)
      --search string     Search keyword for use in 'get' actions. Search is not supported by all resources.
      --token string      Metal API Token (METAL_AUTH_TOKEN)
```

### SEE ALSO

* [metal ip](metal_ip.md)	 - IP operations

