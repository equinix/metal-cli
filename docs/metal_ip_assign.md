## metal ip assign

Assigns an IP address to a given device

### Synopsis

Example:

metal ip assign -d [device-id] -a [ip-address]

	

```
metal ip assign [flags]
```

### Options

```
  -a, --address string     IP address
  -d, --device-id string   UUID of the device
  -h, --help               help for assign
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

