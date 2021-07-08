## metal device delete

Deletes a device

### Synopsis

Example:	

  metal device delete -i [device_UUID]
  
	

```
metal device delete [flags]
```

### Options

```
  -f, --force       Force removal of the device
  -h, --help        help for delete
  -i, --id string   UUID of the device
```

### Options inherited from parent commands

```
      --config string     Path to JSON or YAML configuration file
      --exclude strings   Comma seperated Href references to collapse in results, may be dotted three levels deep
      --include strings   Comma seperated Href references to expand in results, may be dotted three levels deep
  -j, --json              JSON output
      --search string     Search keyword for use in 'get' actions. Search is not supported by all resources.
      --token string      Metal API Token (METAL_AUTH_TOKEN)
  -y, --yaml              YAML output
```

### SEE ALSO

* [metal device](metal_device.md)	 - Device operations

