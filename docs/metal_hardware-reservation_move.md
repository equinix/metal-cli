## metal hardware-reservation move

Move hardware reservation to another project

### Synopsis

Example:

packet hardware_reservation move -i [hardware_reservation_UUID] -p [project_UUID]


```
metal hardware-reservation move [flags]
```

### Options

```
  -h, --help                help for move
  -i, --id string           UUID of the hardware reservation
  -p, --project-id string   UUID of the project
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

* [metal hardware-reservation](metal_hardware-reservation.md)	 - Hardware reservation operations

