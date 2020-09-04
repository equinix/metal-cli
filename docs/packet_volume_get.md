## packet volume get

Retrieves a volume list or volume details.

### Synopsis

Example:
	
Retrieve the list of volumes:
packet volume get --project-id [project_UUID]
  
Retrieve a specific volume:
packet volume get --id [volume_UUID]



```
packet volume get [flags]
```

### Options

```
  -h, --help                help for get
  -i, --id string           UUID of the volume
  -p, --project-id string   UUID of the project
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

