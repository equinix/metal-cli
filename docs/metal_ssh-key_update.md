## metal ssh-key update

Updates an SSH key.

### Synopsis

Updates an SSH key with either a new public key, a new label, or both.

```
metal ssh-key update -i <SSH-key_UUID> [-k <public_key>] [-l <label>] [flags]
```

### Examples

```
  # Updates SSH key 5cb96463-88fd-4d68-94ba-2c9505ff265e with a new public key: 
  metal ssh-key update -i 5cb96463-88fd-4d68-94ba-2c9505ff265e -k AAAAB3N...user@domain.com
  
  # Updates SSH key 5cb96463-88fd-4d68-94ba-2c9505ff265e with a new label:
  metal ssh-key update -i 5cb96463-88fd-4d68-94ba-2c9505ff265e -l test-machine-2
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
      --config string        Path to JSON or YAML configuration file
      --exclude strings      Comma separated Href references to collapse in results, may be dotted three levels deep
      --filter stringArray   Filter 'get' actions with name value pairs. Filter is not supported by all resources and is implemented as request query parameters.
      --include strings      Comma separated Href references to expand in results, may be dotted three levels deep
  -o, --output string        Output format (*table, json, yaml). env output formats are (*sh, terraform, capp).
      --search string        Search keyword for use in 'get' actions. Search is not supported by all resources.
      --sort-by string       Sort fields for use in 'get' actions. Sort is not supported by all resources.
      --sort-dir string      Sort field direction for use in 'get' actions. Sort is not supported by all resources.
      --token string         Metal API Token (METAL_AUTH_TOKEN)
```

### SEE ALSO

* [metal ssh-key](metal_ssh-key.md)	 - SSH key operations: create, get, update, and delete.

