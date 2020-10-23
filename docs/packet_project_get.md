## packet project get

Retrieves all available projects or a single project

### Synopsis

Example:

Retrieve all projects:
packet project get
  
Retrieve a specific project:
packet project get -i [project_UUID]
packet project get -n [project_name]

When using "--json" or "--yaml", "--include=members" is implied.
	

```
packet project get [flags]
```

### Options

```
  -h, --help                help for get
  -n, --project string      Name of the project
  -i, --project-id string   UUID of the project
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

* [packet project](packet_project.md)	 - Project operations

