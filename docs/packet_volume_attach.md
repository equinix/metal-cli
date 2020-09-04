## packet volume attach

Attaches a volume to a device.

### Synopsis

Example:

packet volume attach --id [volume_UUID] --device-id [device_UUID]

	

```
packet volume attach [flags]
```

### Options

```
  -d, --device-id string   UUID of the device
  -h, --help               help for attach
  -i, --id string          UUID of the volume
```

### Options inherited from parent commands

```
      --config string     Path to JSON or YAML configuration file
      --exclude strings   Comma seperated Href references to collapse in results, may be dotted three levels deep
      --include strings   Comma seperated Href references to expand in results, may be dotted three levels deep
  -j, --json              JSON output
  -y, --yaml              YAML output
```

### SEE ALSO

* [packet volume](packet_volume.md)	 - Volume operations

