## packet project get

Retrieves all available projects or a single project

### Synopsis

Example:

Retrieve all projects:
packet project get
  
Retrieve a specific project:
packet project get -i [project_UUID]
packet project get -n [project_name]
	

```
packet project get [flags]
```

### Options

```
  -h, --help                help for get
  -j, --json                JSON output
  -n, --project string      Name of the project
  -i, --project-id string   UUID of the project
  -y, --yaml                YAML output
```

### Options inherited from parent commands

```
      --config string   Path to JSON or YAML configuration file
```

### SEE ALSO

* [packet project](packet_project.md)	 - Project operations

