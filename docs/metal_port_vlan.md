## metal port vlan

Modifies VLAN assignments on a port

### Synopsis

Modifies the VLANs of the specified port to the desired state. Existing state can be restated without error.

```
metal port vlan -i <port_UUID> [--native <vlan>] [--unassign <vlan>]... [--assign <vlan>]... [flags]
```

### Examples

```
  # Assigns VLANs 1234 and 5678 to the port:
  metal port vlans -i 3bd5bf07-6094-48ad-bd03-d94e8712fdc8 -a 1234 -a 5678

  # Unassigns VXLAN 1234 from the port:
  metal port vlans -i 3bd5bf07-6094-48ad-bd03-d94e8712fdc8 -u 1234

  # Assigns VXLAN 1234 to the port and makes it the Native VLAN:
  metal port vlans -i 3bd5bf07-6094-48ad-bd03-d94e8712fdc8 --native=1234
```

### Options

```
  -a, --assign strings     A VXLAN to assign to the port. May also be used to change a Native VLAN assignment to tagged (non-native).
  -h, --help               help for vlan
  -n, --native string      The VXLAN to make assign as the Native VLAN
  -i, --port-id string     The UUID of a port.
  -u, --unassign strings   A VXLAN to unassign from a port.
```

### Options inherited from parent commands

```
      --config string        Path to JSON or YAML configuration file
      --exclude strings      Comma separated Href references to collapse in results, may be dotted three levels deep
      --filter stringArray   Filter 'get' actions with name value pairs. Filter is not supported by all resources and is implemented as request query parameters.
      --include strings      Comma separated Href references to expand in results, may be dotted three levels deep
  -o, --output string        Output format (*table, json, yaml)
      --search string        Search keyword for use in 'get' actions. Search is not supported by all resources.
      --sort-by string       Sort fields for use in 'get' actions. Sort is not supported by all resources.
      --sort-dir string      Sort field direction for use in 'get' actions. Sort is not supported by all resources.
      --token string         Metal API Token (METAL_AUTH_TOKEN)
```

### SEE ALSO

* [metal port](metal_port.md)	 - Port operations: get, convert, vlans.

