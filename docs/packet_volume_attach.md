## packet volume attach

Attaches a volume to a device.

### Synopsis

Attaches a volume to a device.

```
packet volume attach [flags]
```

### Examples

```

packet volume attach --id [volume_UUID] --device-id [device_UUID]
```

### Options

```
  -d, --device-id string   UUID of the device
  -h, --help               help for attach
  -i, --id string          UUID of the volume
```

### Options inherited from parent commands

```
      --config string   Path to JSON or YAML configuration file
```

### SEE ALSO

* [packet volume](packet_volume.md)	 - Volume operations

