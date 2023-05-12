## metal device reinstall

Reinstalls a device.

### Synopsis

Reinstalls the provided device with the current operating system or a new operating system with optional flags to preserve data or skip disk clean-up. The ID of the device to reinstall is required.

```
metal device reinstall --id <device-id> [--operating-system <os_slug>] [--deprovision-fast] [--preserve-data] [flags]
```

### Examples

```
  # Reinstalls a device with the current OS:
  metal device reinstall -d 50382f72-02b7-4b40-ac8d-253713e1e174
  
  # Reinstalls a device with Ubuntu 22.04 while preserving the data on non-OS disks:
  metal device reinstall -d 50382f72-02b7-4b40-ac8d-253713e1e174 -O ubuntu_22_04 --preserve-data
```

### Options

```
      --deprovision-fast          Avoid optional potentially slow clean-up tasks.
  -h, --help                      help for reinstall
  -d, --id string                 ID of device to be reinstalled
  -O, --operating-system string   Operating system install on the device. If omitted the current OS will be reinstalled.
      --preserve-data             Avoid wiping data on disks where the OS is *not* being installed.
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

