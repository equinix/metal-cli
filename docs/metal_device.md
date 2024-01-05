## metal device

Device operations: create, get, update, delete, reinstall, start, stop, and reboot.

### Synopsis

Device operations that control server provisioning, metadata, and basic operations.

### Options

```
  -h, --help   help for device
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

* [metal](metal.md)	 - Command line interface for Equinix Metal
* [metal device create](metal_device_create.md)	 - Creates a device.
* [metal device delete](metal_device_delete.md)	 - Deletes a device.
* [metal device get](metal_device_get.md)	 - Retrieves device list or device details.
* [metal device reboot](metal_device_reboot.md)	 - Reboots a device.
* [metal device reinstall](metal_device_reinstall.md)	 - Reinstalls a device.
* [metal device start](metal_device_start.md)	 - Starts a device.
* [metal device stop](metal_device_stop.md)	 - Stops a device.
* [metal device update](metal_device_update.md)	 - Updates a device.

