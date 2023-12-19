## metal vrf ips

Retrieves the list of VRF IP Reservations for the VRF.

### Synopsis

Retrieves the list of VRF IP Reservations for the VRF.

```
metal vrf ips [-v <vrf-id] [-i <ip-id>] [flags]
```

### Examples

```
 # Retrieves the list of VRF IP Reservations for the VRF.
	
	metal vrf ips [-v <vrf-id] 

	# Retrieve a specific IP Reservation for a VRF
	metal vrf ips [-v <vrf-id] [-i <ip-id>]
```

### Options

```
  -h, --help            help for ips
  -i, --id string       Specify the IP UUID to retrieve the details of a VRF IP reservation.
  -v, --vrf-id string   Specify the VRF UUID to list its associated IP reservations.
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

