## packet ssh-key create

Creates an SSH key

### Synopsis

Example:

packet ssh-key create --key [public_key] --label [label]

	

```
packet ssh-key create [flags]
```

### Options

```
  -h, --help           help for create
  -k, --key string     Public SSH key string
  -l, --label string   Name of the SSH key
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

* [packet ssh-key](packet_ssh-key.md)	 - SSH key operations

