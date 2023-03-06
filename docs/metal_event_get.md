## metal event get

Retrieves events for the current user, an organization, a project, a device, or the details of a specific event.

### Synopsis

Retrieves events for the current user, an organization, a project, a device, or the details of a specific event. The current user's events includes all events in all projects and devices that the user has access to. When using --json or --yaml flags, the --include=relationships flag is implied.

```
metal event get [-p <project_id>] | [-d <device_id>] | [-i <event_id>] | [-O <organization_id>] [flags]
```

### Examples

```
  # Retrieve all events of a current user:
  metal event get

  # Retrieve the details of a specific event:
  metal event get -i e9a969b3-8911-4667-9d99-57cd3dd4ef6f

  # Retrieve all the events of an organization:
  metal event get -o c079178c-9557-48f2-9ce7-cfb927b81928

  # Retrieve all events of a project:
  metal event get -p 1867ee8f-6a11-470a-9505-952d6a324040

  # Retrieve all events of a device:
  metal event get -d ca614540-fbd4-4dbb-9689-457c6ccc8353
```

### Options

```
  -d, --device-id string         UUID of the device
  -h, --help                     help for get
  -i, --id string                UUID of the event
  -O, --organization-id string   UUID of the organization
  -p, --project-id string        Project ID (METAL_PROJECT_ID)
```

### Options inherited from parent commands

```
      --config string        Path to JSON or YAML configuration file
      --exclude strings      Comma separated Href references to collapse in results, may be dotted three levels deep
      --filter stringArray   Filter 'get' actions with name value pairs. Filter is not supported by all resources and is implemented as request query parameters.
      --include strings      Comma separated Href references to expand in results, may be dotted three levels deep
  -o, --output string        Output format (*table, json, yaml). env output formats are (*sh, terraform, capp).
      --search string        Search keyword for use in 'get' actions. Search is not supported by all resources.
      --sort-by string       Sort fields for use in 'get' actions. Sort is not supported by all resources.
      --sort-dir string      Sort field direction for use in 'get' actions. Sort is not supported by all resources.
      --token string         Metal API Token (METAL_AUTH_TOKEN)
```

### SEE ALSO

* [metal event](metal_event.md)	 - Events operations: get.

