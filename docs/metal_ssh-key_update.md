## metal ssh-key update

Updates a project

### Synopsis

Example:

packet ssh-key update --id [ssh-key_UUID] --key [new_key]



```
metal ssh-key update [flags]
```

### Options

```
  -h, --help           help for update
  -i, --id string      UUID of the SSH key
  -k, --key string     Public SSH key string
  -l, --label string   Name of the SSH key
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

* [metal ssh-key](metal_ssh-key.md)	 - SSH key operations

