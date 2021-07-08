## metal ssh-key get

Retrieves a list of available SSH keys or a single SSH key

### Synopsis

Example:

Retrieve all SSH keys: 
metal ssh-key get
  
Retrieve a specific SSH key:
metal ssh-key get --id [ssh-key_UUID] 



```
metal ssh-key get [flags]
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
      --token string      Metal API Token (METAL_AUTH_TOKEN)
  -y, --yaml              YAML output
```

### SEE ALSO

* [metal ssh-key](metal_ssh-key.md)	 - SSH key operations

