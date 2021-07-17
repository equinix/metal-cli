## metal device update

Updates a device

### Synopsis

Example:

metal device update --id [device_UUID] --hostname [new_hostname]



```
metal device update [flags]
```

### Options

```
  -a, --always-pxe               --alaways-pxe or -a
  -c, --customdata string        Custom data
  -d, --description string       Description for the device
  -h, --help                     help for update
  -H, --hostname string          Hostname
  -i, --id string                UUID of the device
  -s, --ipxe-script-url string   URL to the iPXE script
  -l, --locked                   Lock device
  -t, --tags strings             Tags for the device --tags="tag1,tag2"
  -u, --userdata string          User data
```

### Options inherited from parent commands

```
      --config string     Path to JSON or YAML configuration file
      --exclude strings   Comma seperated Href references to collapse in results, may be dotted three levels deep
      --include strings   Comma seperated Href references to expand in results, may be dotted three levels deep
  -o, --output string     Output format (*table, json, yaml)
      --search string     Search keyword for use in 'get' actions. Search is not supported by all resources.
      --token string      Metal API Token (METAL_AUTH_TOKEN)
```

### SEE ALSO

* [metal device](metal_device.md)	 - Device operations

