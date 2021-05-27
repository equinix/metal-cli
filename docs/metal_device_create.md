## metal device create

Creates a device

### Synopsis

Example:

packet device create --hostname [hostname] --plan [plan] --metro [metro_code] --facility [facility_code] --operating-system [operating_system] --project-id [project_UUID]



```
metal device create [flags]
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
  -m, --metro string                     Code of the metro where the device will be created
  -o, --operating-system string          Operating system name for the device
  -P, --plan string                      Name of the plan
  -p, --project-id string                UUID of the project where the device will be created
  -v, --public-ipv4-subnet-size int      Size of the public IPv4 subnet
  -I, --spot-instance                    Set the device as a spot instance
      --spot-price-max float             --spot-price-max=1.2 or -m=1.2
  -t, --tags strings                     Tags for the device: --tags="tag1,tag2"
  -T, --termination-time string          Device termination time: --termination-time="15:04:05"
  -u, --userdata string                  Userdata for device initialization (can not be used with --userdata-file)
      --userdata-file string             Path to a userdata file for device initialization (can not be used with --userdata)
```

### Options inherited from parent commands

```
      --config string     Path to JSON or YAML configuration file
      --exclude strings   Comma seperated Href references to collapse in results, may be dotted three levels deep
      --include strings   Comma seperated Href references to expand in results, may be dotted three levels deep
  -j, --json              JSON output
      --search string     Search keyword for use in 'get' actions. Search is not supported by all resources.
  -y, --yaml              YAML output
```

### SEE ALSO

* [metal device](metal_device.md)	 - Device operations

