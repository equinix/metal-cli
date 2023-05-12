## metal device stop

Stops a device.

### Synopsis

Stops or powers off a device that is currently started or powered on.

```
metal device stop -i <device_id> [flags]
```

### Examples

```
  # Stops the specified device:
  metal device stop --id [device_UUID]
```

### Options

```
  -h, --help        help for stop
  -i, --id string   The UUID of the device.
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

* [metal device](metal_device.md)	 - Device operations: create, get, update, delete, reinstall, start, stop, and reboot.

