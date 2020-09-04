## packet device create

Creates a device

### Synopsis

Example:

packet device create --hostname [hostname] --plan [plan] --facility [facility_code] --operating-system [operating_system] --project-id [project_UUID]



```
packet device create [flags]
```

### Options

```
  -a, --always-pxe                       
  -b, --billing-cycle string             Billing cycle (default "hourly")
  -c, --customdata string                Custom data
  -f, --facility string                  Code of the facility where the device will be created
  -r, --hardware-reservation-id string   UUID of the hardware reservation
  -h, --help                             help for create
  -H, --hostname string                  Hostname
  -i, --ipxe-script-url string           URL to the iPXE script
  -o, --operating-system string          Operating system name for the device
  -P, --plan string                      Name of the plan
  -p, --project-id string                UUID of the project where the device will be created
  -v, --public-ipv4-subnet-size int      Size of the public IPv4 subnet
  -I, --spot-instance                    Set the device as a spot instance
  -m, --spot-price-max float             --spot-price-max=1.2 or -m=1.2
  -s, --storage string                   UUID of the storage
  -t, --tags strings                     Tags for the device: --tags="tag1,tag2"
  -T, --termination-time string          Device termination time: --termination-time="15:04:05"
  -u, --userdata string                  User data
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

* [packet device](packet_device.md)	 - Device operations

