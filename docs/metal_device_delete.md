## metal device delete

Deletes a device.

### Synopsis

Deletes the specified device with a confirmation prompt. To skip the confirmation use --force.

```
metal device delete -i <device_id> [-f] [flags]
```

### Examples

```
  # Deletes the specified device:
  metal device delete -i 7ec86e23-8dcf-48ed-bd9b-c25c20958277
  >
  ✔ Are you sure you want to delete device 7ec86e23-8dcf-48ed-bd9b-c25c20958277: y
		
  # Deletes a VLAN, skipping confirmation:
  metal device delete -f -i 7ec86e23-8dcf-48ed-bd9b-c25c20958277
```

### Options

```
  -f, --force       Skips confirmation for the device deletion.
  -h, --help        help for delete
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

