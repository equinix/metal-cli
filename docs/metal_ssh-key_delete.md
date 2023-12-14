## metal ssh-key delete

Deletes an SSH key.

### Synopsis

Deletes an SSH key with a confirmation prompt. To skip the confirmation use --force. Does not remove the SSH key from existing servers.

```
metal ssh-key delete --id <SSH-key_UUID> [--force] [flags]
```

### Examples

```
  # Deletes an SSH key, with confirmation:
  metal ssh-key delete -i 5cb96463-88fd-4d68-94ba-2c9505ff265e
  >
  ✔ Are you sure you want to delete SSH Key 5cb96463-88fd-4d68-94ba-2c9505ff265e [Y/N]: y
  
  # Deletes an SSH key, skipping confirmation:
  metal ssh-key delete -i 5cb96463-88fd-4d68-94ba-2c9505ff265e -f
```

### Options

```
  -f, --force       Skips confirmation for the deletion of the SSH key.
  -h, --help        help for delete
  -i, --id string   The UUID of the SSH key.
```

### Options inherited from parent commands

```
      --config string         Path to JSON or YAML configuration file
      --exclude strings       Comma separated Href references to collapse in results, may be dotted three levels deep
      --filter stringArray    Filter 'get' actions with name value pairs. Filter is not supported by all resources and is implemented as request query parameters.
      --http-header strings   Headers to add to requests (in format key=value)
      --include strings       Comma separated Href references to expand in results, may be dotted three levels deep
  -o, --output string         Output format (*table, json, yaml). env output formats are (*sh, terraform, capp).
      --search string         Search keyword for use in 'get' actions. Search is not supported by all resources.
      --sort-by string        Sort fields for use in 'get' actions. Sort is not supported by all resources.
      --sort-dir string       Sort field direction for use in 'get' actions. Sort is not supported by all resources.
      --token string          Metal API Token (METAL_AUTH_TOKEN)
```

### SEE ALSO

* [metal ssh-key](metal_ssh-key.md)	 - SSH key operations: create, get, update, and delete.

