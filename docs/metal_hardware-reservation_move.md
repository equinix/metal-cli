## metal hardware-reservation move

Move hardware reservation to another project

### Synopsis

Example:

metal hardware_reservation move -i [hardware_reservation_UUID] -p [project_UUID]


```
metal hardware-reservation move [flags]
```

### Options

```
  -h, --help                help for move
  -i, --id string           UUID of the hardware reservation
  -p, --project-id string   Project ID (METAL_PROJECT_ID)
```

### Options inherited from parent commands

```
      --config string        Path to JSON or YAML configuration file
      --exclude strings      Comma seperated Href references to collapse in results, may be dotted three levels deep
      --filter stringArray   Filter 'get' actions with name value pairs. Filter is not supported by all resources and is implemented as request query parameters.
      --include strings      Comma seperated Href references to expand in results, may be dotted three levels deep
  -o, --output string        Output format (*table, json, yaml)
      --search string        Search keyword for use in 'get' actions. Search is not supported by all resources.
      --sort-by string       Sort fields for use in 'get' actions. Sort is not supported by all resources.
      --sort-dir string      Sort field direction for use in 'get' actions. Sort is not supported by all resources.
      --token string         Metal API Token (METAL_AUTH_TOKEN)
```

### SEE ALSO

* [metal hardware-reservation](metal_hardware-reservation.md)	 - Hardware reservation operations

