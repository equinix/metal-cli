## metal vrf

VRF operations : create, get, delete

### Synopsis

VRF operations : It defines a collection of customer-managed IP blocks that can be used in BGP peering on one or more virtual networks and basic operations

### Options

```
  -h, --help   help for vrf
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

* [metal](metal.md)	 - Command line interface for Equinix Metal
* [metal vrf create](metal_vrf_create.md)	 - Creates a Virtual Routing and Forwarding(VRF) for a specified project.
* [metal vrf create-route](metal_vrf_create-route.md)	 - Create a route in a VRF. Currently only static default routes are supported.
* [metal vrf delete](metal_vrf_delete.md)	 - Deletes a VRF.
* [metal vrf delete-route](metal_vrf_delete-route.md)	 - Delete a VRF Route
* [metal vrf get](metal_vrf_get.md)	 - Lists VRFs.
* [metal vrf get-route](metal_vrf_get-route.md)	 - Retrieve all routes in the VRF
* [metal vrf ips](metal_vrf_ips.md)	 - Retrieves the list of VRF IP Reservations for the VRF.
* [metal vrf update-route](metal_vrf_update-route.md)	 - Requests a VRF Route be redeployed/update across the network.

