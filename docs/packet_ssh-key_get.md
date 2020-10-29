## packet ssh-key get

Retrieves a list of available SSH keys or a single SSH key

### Synopsis

Example:

Retrieve all SSH keys: 
packet ssh-key get
  
Retrieve a specific SSH key:
packet ssh-key get --id [ssh-key_UUID] 



```
packet ssh-key get [flags]
```

### Options

```
  -h, --help        help for get
  -i, --id string   UUID of the SSH key
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

* [packet ssh-key](packet_ssh-key.md)	 - SSH key operations

