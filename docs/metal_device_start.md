## metal device start

Starts a device.

### Synopsis

Starts or powers on a device that is currently stopped or powered off.

```
metal device start -i <device_id> [flags]
```

### Examples

```
  # Starts the specified device:
  metal device start --id 26a9da5f-a0db-41f6-8467-827e144e59a7
```

### Options

```
  -h, --help        help for start
  -i, --id string   The UUID of the device.
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

