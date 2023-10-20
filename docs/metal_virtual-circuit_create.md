## metal virtual-circuit create

Creates an create-virtual-circuit for specific interconnection.

### Synopsis

Creates an create-virtual-circuit for specific interconnection

```
metal virtual-circuit create  [-c connection_id] [-p port_id] [-P <project_id> ] -n <name> [-d <description>] [--vnid <vnid> ] [-V <vlan> ] [-s <speed> ] [-t <tags> ] [flags]
```

### Examples

```
  # Creates a new virtual-circuit named "interconnection": 
  metal vc create [-c connection_id] [-p port_id] [-P <project_id> ] [-n <name>] [-d <description>] [--vnid <vnid> ] [-V <vlan> ] [-s <speed> ] [-t <tags> ]

  metal vc create -c 81c9cb9e-b02f-4c73-9e04-06702f1380a0 -p 9c8f0c71-591d-42fe-9519-2f632761e2da -P b4673e33-0f48-4948-961a-c31d6edf64f8 -n test-inter  -d test-interconnection -v 15315810-2fda-48b8-b8cd-441ebab684b5 -V 1010 -s 100
  
  metal vc create [-c connection_id] [-p port_id] [-P <project_id> ] [-n <name>] [-d <description>] [-v <vrf-id>] [-M <md5sum>] [-a <peer-asn>] [-S <subnet>] [-c <customer_ip>] [-m <metal_ip>]
```

### Options

```
  -c, --connection-id string   Specify the UUID of the interconnection.
      --customer-ip string     An IP address from the subnet that will be used on the Customer side
  -d, --description string     Description for a Virtual Circuit
  -h, --help                   help for create
  -M, --md5 string             The plaintext BGP peering password shared by neighbors as an MD5 checksum
  -m, --metal-ip string        An IP address from the subnet that will be used on the Metal side. 
  -n, --name string            Name of the Virtual Circuit
  -a, --peer-asn int           The peer ASN that will be used with the VRF on the Virtual Circuit.
  -p, --port-id string         Specify the UUID of the port.
  -P, --project-id string      The project's UUID. This flag is required, unless specified in the config created by metal init or set as METAL_PROJECT_ID environment variable.
  -s, --speed int              bps speed or string (e.g. 52 - '52m' or '100g' or '4 gbps')
  -S, --subnet string          The /30 or /31 subnet of one of the VRF IP Blocks that will be used with the VRF for the Virtual Circuit. 
  -t, --tags strings           Adds the tags for the virtual-circuit --tags "tag1,tag2" OR --tags "tag1" --tags "tag2"
  -V, --vlan int               Adds or updates vlan  Must be between 2 and 4094
      --vnid string            Specify the UUID  of the VLAN.
  -v, --vrf-id string          The UUID of the VRF that will be associated with the Virtual Circuit.
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

