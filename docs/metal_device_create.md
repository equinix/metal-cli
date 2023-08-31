## metal device create

Creates a device.

### Synopsis

Creates a device in the specified project. A plan, hostname, operating system, and either metro or facility is required.

```
metal device create -p <project_id> (-m <metro> | -f <facility>) -P <plan> -H <hostname> -O <operating_system> [-u <userdata> | --userdata-file <filepath>] [-c <customdata>] [-t <tags>] [-r <hardware_reservation_id>] [-I <ipxe_script_url>] [--always-pxe] [--spot-instance] [--spot-price-max=<max_price>] [flags]
```

### Examples

```
  # Provisions a c3.small.x86 in the Dallas metro running Ubuntu 20.04:
  metal device create -p $METAL_PROJECT_ID -P c3.small.x86 -m da -H test-staging-2 -O ubuntu_20_04

  # Provisions a c3.medium.x86 in Silicon Valley, running Rocky Linux, from a hardware reservation:
  metal device create -p $METAL_PROJECT_ID -P c3.medium.x86 -m sv -H test-rocky -O rocky_8 -r 47161704-1715-4b45-8549-fb3f4b2c32c7
```

### Options

```
  -a, --always-pxe string                Sets whether the device always PXE boots on reboot.
  -b, --billing-cycle string             Billing cycle  (default "hourly")
  -c, --customdata string                Custom data to be included with your device's metadata.
  -f, --facility string                  Code of the facility where the device will be created
  -r, --hardware-reservation-id string   The UUID of a hardware reservation, if you are provisioning a server from your reserved hardware.
  -h, --help                             help for create
  -H, --hostname string                  Hostname
  -I, --ipxe-script-url string           The URL of an iPXE script.
  -m, --metro string                     Code of the metro where the device will be created
  -O, --operating-system string          Operating system name for the device
  -P, --plan string                      Name of the plan
  -p, --project-id string                The project's UUID. This flag is required, unless specified in the config created by metal init or set as METAL_PROJECT_ID environment variable.
  -S, --public-ipv4-subnet-size int      Size of the public IPv4 subnet.
  -s, --spot-instance                    Provisions the device as a spot instance.
      --spot-price-max float             Sets the maximum spot market price for the device: --spot-price-max=1.2
  -t, --tags strings                     Tag or list of tags for the device: --tags="tag1,tag2".
  -T, --termination-time string          Device termination time: --termination-time="2023-08-24T15:04:05Z"
  -u, --userdata string                  Userdata for device initialization. Can not be used with --userdata-file.
      --userdata-file string             Path to a userdata file for device initialization. Can not be used with --userdata.
```

### Options inherited from parent commands

```
      --config string         Path to JSON or YAML configuration file
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

* [metal device](metal_device.md)	 - Device operations: create, get, update, delete, reinstall, start, stop, and reboot.

