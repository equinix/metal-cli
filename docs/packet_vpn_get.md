## packet vpn get

Retrieves VPN service

### Synopsis

Example:
	
Enable VPN service: 
packet vpn get --faciliy ewr1


```
packet vpn get [flags]
```

### Options

```
  -f, --facility string   Code of the facility for which VPN config is to be retrieved
  -h, --help              help for get
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

* [packet vpn](packet_vpn.md)	 - VPN service operations

