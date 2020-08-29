## packet hardware-reservation move

Move hardware reservation to another project

### Synopsis

Example:

packet hardware_reservation move -i [hardware_reservation_UUID] -p [project_UUID]


```
packet hardware-reservation move [flags]
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
  -y, --yaml              YAML output
```

### SEE ALSO

* [packet hardware-reservation](packet_hardware-reservation.md)	 - Hardware reservation operations

