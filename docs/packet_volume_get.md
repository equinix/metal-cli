## packet volume get

Retrieves a volume list or volume details.

### Synopsis

Retrieves a volume list or volume details.

```
packet volume get [flags]
```

### Examples

```

Retrieve the list of volumes:
packet volume get --project-id [project_UUID]
Retrieve a specific volume:
packet volume get --id [volume_UUID]
```

### Options

```
  -h, --help                help for get
  -i, --id string           UUID of the volume
  -j, --json                JSON output
  -p, --project-id string   UUID of the project
  -y, --yaml                YAML output
```

### Options inherited from parent commands

```
      --config string   Path to JSON or YAML configuration file
```

### SEE ALSO

* [packet volume](packet_volume.md)	 - Volume operations

