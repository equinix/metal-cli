## metal vrf delete

Deletes a VRF.

### Synopsis

Deletes the specified VRF with a confirmation prompt. To skip the confirmation, use --force.

```
metal vrf delete vrf -i <metal_vrf_UUID> [-f] [flags]
```

### Examples

```
# Deletes a VRF, with confirmation.
  metal delete vrf -i 77e6d57a-d7a4-4816-b451-cf9b043444e2
  >
  ✔ Are you sure you want to delete device 7ec86e23-8dcf-48ed-bd9b-c25c20958277 [Y/N]: y

  # Deletes a VRF, skipping confirmation.
  metal delete vrf -f -i 77e6d57a-d7a4-4816-b451-cf9b043444e2
```

### Options

```
  -f, --force       Skips confirmation for the removal of the VRF.
  -h, --help        help for delete
  -i, --id string   Specify the UUID of the VRF
```

### Options inherited from parent commands

```
      --config string         Path to JSON or YAML configuration file (METAL_CONFIG)
      --exclude strings       Comma separated Href references to collapse in results, may be dotted three levels deep
      --filter stringArray    Filter 'get' actions with name value pairs. Filter is not supported by all resources and is implemented as request query parameters.
      --http-header strings   Headers to add to requests (in format key=value)
      --include strings       Comma separated Href references to expand in results, may be dotted three levels deep
  -o, --output string         Output format (*table, json, yaml). env output formats are (*sh, terraform, capp).
      --search string         Search keyword for use in 'get' actions. Search is not supported by all resources.
      --sort-by string        Sort fields for use in 'get' actions. Sort is not supported by all resources.
      --sort-dir string       Sort field direction for use in 'get' actions. Sort is not supported by all resources.
      --token string          Metal API Token (METAL_AUTH_TOKEN)
```

### SEE ALSO

* [metal vrf](metal_vrf.md)	 - VRF operations : create, get, delete

