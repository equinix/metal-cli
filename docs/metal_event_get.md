## metal event get

Retrieves one or more events for organizations, projects, or devices.

### Synopsis

Example:
Retrieve all events:
metal event get

Retrieve a specific event:
metal event get -i [event_UUID]

Retrieve all events of an organization:
metal event get -o [organization_UUID]

Retrieve all events of a project:
metal event get -p [project_UUID]

Retrieve all events of a device:
metal event get -d [device_UUID]

Retrieve all events of a current user:
metal event get

When using "--json" or "--yaml", "--include=relationships" is implied.


```
metal event get [flags]
```

### Options

```
  -d, --device-id string         UUID of the device
  -h, --help                     help for get
  -i, --id string                UUID of the event
  -o, --organization-id string   UUID of the organization
  -p, --project-id string        UUID of the project
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

* [metal event](metal_event.md)	 - Events operations

