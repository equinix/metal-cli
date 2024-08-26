## metal device update

Updates a device.

### Synopsis

Updates the hostname of a device. Updates or adds a description, tags, userdata, custom data, and iPXE settings for an already provisioned device. Can also lock or unlock future changes to the device.

```
metal device update -i <device_id> [-H <hostname>] [-d <description>] [--locked=<true|false>] [-t <tags>] [-u <userdata> | --userdata-file <filepath>] [-c <customdata>] [-s <ipxe_script_url>] [--always-pxe=<true|false>] [flags]
```

### Examples

```
  # Updates the hostname of a device:
  metal device update --id 30c15082-a06e-4c43-bfc3-252616b46eba --hostname renamed-staging04
```

### Options

```
  -a, --always-pxe               Updates the always_pxe toggle for the device (<true|false>).
  -c, --customdata string        Adds or updates custom data to be included with your device's metadata.
  -d, --description string       Adds or updates the description for the device.
  -h, --help                     help for update
  -H, --hostname string          The new hostname of the device.
  -i, --id string                The UUID of the device.
  -s, --ipxe-script-url string   Add or update the URL of the iPXE script.
  -l, --locked bools             Locks or unlocks the device for future changes (<true|false>). (default [])
  -t, --tags strings             Adds or updates the tags for the device --tags="tag1,tag2".
  -u, --userdata string          Adds or updates the userdata for the device.
      --userdata-file string     Path to a userdata file for device initialization. Can not be used with --userdata.
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

