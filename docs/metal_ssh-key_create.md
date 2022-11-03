## metal ssh-key create

Adds an SSH key for the current user's account.

### Synopsis

Adds an SSH key for the current user's account. The key will then be added to the user's servers at provision time.

```
metal ssh-key create --key <public_key> --label <label> [flags]
```

### Examples

```
 # Adds a key labled "example-key" to the current user account.
  metal ssh-key create --key ssh-rsa AAAAB3N...user@domain.com --label example-key
```

### Options

```
  -h, --help           help for create
  -k, --key string     User's full SSH public key string.
  -l, --label string   Name or other user-friendly description of the SSH key.
```

### Options inherited from parent commands

```
      --config string        Path to JSON or YAML configuration file
      --exclude strings      Comma separated Href references to collapse in results, may be dotted three levels deep
      --filter stringArray   Filter 'get' actions with name value pairs. Filter is not supported by all resources and is implemented as request query parameters.
      --include strings      Comma separated Href references to expand in results, may be dotted three levels deep
  -o, --output string        Output format (*table, json, yaml)
      --search string        Search keyword for use in 'get' actions. Search is not supported by all resources.
      --sort-by string       Sort fields for use in 'get' actions. Sort is not supported by all resources.
      --sort-dir string      Sort field direction for use in 'get' actions. Sort is not supported by all resources.
      --token string         Metal API Token (METAL_AUTH_TOKEN)
```

### SEE ALSO

* [metal ssh-key](metal_ssh-key.md)	 - SSH key operations: create, get, update, and delete.

