## metal volume detach

Detaches a volume from a device

### Synopsis

Example:

metal volume detach --id [attachment_UUID]



```
metal volume detach [flags]
```

### Options

```
  -h, --help        help for detach
  -i, --id string   UUID of the attached volume
```

### Options inherited from parent commands

```
      --config string     Path to JSON or YAML configuration file
      --exclude strings   Comma seperated Href references to collapse in results, may be dotted three levels deep
      --include strings   Comma seperated Href references to expand in results, may be dotted three levels deep
  -j, --json              JSON output
      --search string     Search keyword for use in 'get' actions. Search is not supported by all resources.
  -y, --yaml              YAML output
```

### SEE ALSO

* [metal volume](metal_volume.md)	 - Volume operations

