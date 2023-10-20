## metal

Command line interface for Equinix Metal

### Synopsis

Command line interface for Equinix Metal

### Options

```
      --config string         Path to JSON or YAML configuration file (METAL_CONFIG)
      --exclude strings       Comma separated Href references to collapse in results, may be dotted three levels deep
      --filter stringArray    Filter 'get' actions with name value pairs. Filter is not supported by all resources and is implemented as request query parameters.
  -h, --help                  help for metal
      --http-header strings   Headers to add to requests (in format key=value)
      --include strings       Comma separated Href references to expand in results, may be dotted three levels deep
  -o, --output string         Output format (*table, json, yaml). env output formats are (*sh, terraform, capp).
      --search string         Search keyword for use in 'get' actions. Search is not supported by all resources.
      --sort-by string        Sort fields for use in 'get' actions. Sort is not supported by all resources.
      --sort-dir string       Sort field direction for use in 'get' actions. Sort is not supported by all resources.
      --token string          Metal API Token (METAL_AUTH_TOKEN)
```

### SEE ALSO

* [metal 2fa](metal_2fa.md)	 - Two-factor Authentication operations: receive, enable, disable.
* [metal capacity](metal_capacity.md)	 - Capacity operations: get, check
* [metal completion](metal_completion.md)	 - Generates completion scripts.
* [metal device](metal_device.md)	 - Device operations: create, get, update, delete, reinstall, start, stop, and reboot.
* [metal docs](metal_docs.md)	 - Generate command markdown documentation.
* [metal emdocs](metal_emdocs.md)	 - Generates single-page reference documentation.
* [metal env](metal_env.md)	 - Prints or generates environment variables.
* [metal event](metal_event.md)	 - Events operations: get.
* [metal facilities](metal_facilities.md)	 - Facility operations: get.
* [metal gateway](metal_gateway.md)	 - Metal Gateway operations: create, delete, and retrieve.
* [metal hardware-reservation](metal_hardware-reservation.md)	 - Hardware reservation operations: get, move.
* [metal init](metal_init.md)	 - Create a configuration file.
* [metal interconnections](metal_interconnections.md)	 - interconnections operations: create, get, update, delete
* [metal ip](metal_ip.md)	 - IP address, reservations, and assignment operations: assign, unassign, remove, available, request and get.
* [metal metros](metal_metros.md)	 - Metro operations: get.
* [metal operating-systems](metal_operating-systems.md)	 - Operating system operations: get.
* [metal organization](metal_organization.md)	 - Organization operations: create, get, update, payment-methods, and delete.
* [metal plan](metal_plan.md)	 - Plan operations: get.
* [metal port](metal_port.md)	 - Port operations: get, convert, vlans.
* [metal project](metal_project.md)	 - Project operations: create, get, update, delete, and bgp-enable, bgp-config, bgp-sessions.
* [metal ssh-key](metal_ssh-key.md)	 - SSH key operations: create, get, update, and delete.
* [metal user](metal_user.md)	 - User operations: get and add.
* [metal virtual-circuit](metal_virtual-circuit.md)	 - virtual-circuit operations: create, get, update, delete
* [metal virtual-network](metal_virtual-network.md)	 - Virtual network (VLAN) operations : create, get, delete.
* [metal vrf](metal_vrf.md)	 - VRF operations : create, get, delete

