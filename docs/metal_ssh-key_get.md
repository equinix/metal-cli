## metal ssh-key get

Retrieves a list of SSH keys or a specified SSH key.

### Synopsis

Retrieves a list of SSH keys associated with the current user's account or the details of single SSH key.

```
metal ssh-key get [-i <SSH-key_UUID>] [-P] [-p <project_id>] [flags]
```

### Examples

```
  # Retrieves the SSH keys of the current user: 
  metal ssh-key get
  
  # Returns the details of SSH key 5cb96463-88fd-4d68-94ba-2c9505ff265e:
  metal ssh-key get --id 5cb96463-88fd-4d68-94ba-2c9505ff265e

  # Retrieve all project SSH keys
  metal ssh-key get --project-ssh-keys --project-id [project_UUID]
```

### Options

```
  -h, --help                help for get
  -i, --id string           The UUID of an SSH key.
  -p, --project-id string   List SSH Keys for the project identified by Project ID (ignored without -P)
  -P, --project-ssh-keys    List SSH Keys for projects
```

### Options inherited from parent commands

```
      --config string         Path to JSON or YAML configuration file (METAL_CONFIG)
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

