## metal virtual-circuit update

Updates a virtualcircuit.

### Synopsis

Updates a specified virtualcircuit etiher of vlanID OR vrfID

```
metal virtual-circuit update -i <id> [-v <vlan UUID>] [-d <description>] [-n <name>] [-s <speed>] [-t <tags>] [flags]
```

### Examples

```
  # Updates a specified virtualcircuit etiher of vlanID OR vrfID:

	metal vc update [-i <id>] [-n <name>] [-d <description>] [-v <vnid> ] [-s <speed> ] [-t <tags> ]

	metal vc update -i e2edb90b-a8ef-47cb-a577-63b0ba129c29 -d "test-inter-fri-dedicated"

	metal vc update [-i <id>] [-n <name>] [-d <description>] [-M <md5sum>] [-a <peer-asn>] [-S <subnet>] [-c <customer-ip>] [-m <metal-ip>] [-t <tags> ]
```

### Options

```
  -c, --customer-ip string   An IP address from the subnet that will be used on the Customer side
  -d, --description string   Description for a Virtual Circuit
  -h, --help                 help for update
  -i, --id string            Specify the UUID of the virtual-circuit.
  -M, --md5 string           The plaintext BGP peering password shared by neighbors as an MD5 checksum
  -m, --metal-ip string      An IP address from the subnet that will be used on the Metal side. 
  -n, --name string          Name of the Virtual Circuit
  -a, --peer-asn int         The peer ASN that will be used with the VRF on the Virtual Circuit.
  -s, --speed string         Adds or updates Speed can be changed only if it is an interconnection on a Dedicated Port
  -S, --subnet string        The /30 or /31 subnet of one of the VRF IP Blocks that will be used with the VRF for the Virtual Circuit. 
  -t, --tags strings         updates the tags for the virtual circuit --tags "tag1,tag2" OR --tags "tag1" --tags "tag2" (NOTE: --tags "" will remove all tags from the virtual circuit
  -v, --vnid string          A Virtual Network record UUID or the VNID of a Metro Virtual Network in your project.
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

* [metal virtual-circuit](metal_virtual-circuit.md)	 - virtual-circuit operations: create, get, update, delete

