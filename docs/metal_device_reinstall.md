## metal device reinstall

Reinstalls a device.

### Synopsis

Reinstalls the provided device. The ID of the device to reinstall is required.

```
metal device reinstall -d <device-id> [flags]
```

### Options

```
  -h, --help                      help for reinstall
  -d, --id string                 ID of device to be reinstalled
  -O, --operating-system string   Operating system name for the device
      --preserve-data             Avoid wiping data on disks where the os is *not* to be installed into
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

* [metal device](metal_device.md)	 - Device operations: create, get, update, delete, reinstall, start, stop, and reboot.

