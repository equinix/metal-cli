## metal device get

Retrieves device list or device details.

### Synopsis

Retrieves a list of devices in the project, or the details of the specified device. Either a project ID or a device ID is required.

```
metal device get [-p <project_id>] | [-i <device_id>] [flags]
```

### Examples

```
  # Gets the details of the specified device:
  metal device get -i 52b60ca7-1ae2-4875-846b-4e4635223471
  
  # Gets a list of devices in the specified project:
  metal device get -p 5ad070a5-62e8-4cfe-a0b9-3b79e59f1cfe

  # Get a list of devices with the hostname foo and a default project configured:
  metal device get --filter hostname=foo
```

### Options

```
  -h, --help                help for get
  -i, --id string           The UUID of a device.
  -p, --project-id string   The project's UUID. This flag is required, unless specified in the config created by metal init or set as METAL_PROJECT_ID environment variable.
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

* [metal device](metal_device.md)	 - Device operations: create, get, update, delete, reinstall, start, stop, and reboot.

