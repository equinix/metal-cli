## metal project get

Retrieves all available projects or a single project

### Synopsis

Example:

Retrieve all projects:
metal project get
  
Retrieve a specific project:
metal project get -i [project_UUID]
metal project get -n [project_name]

When using "--json" or "--yaml", "--include=members" is implied.
	

```
metal project get [flags]
```

### Options

```
  -h, --help             help for get
  -i, --id string        Project ID (METAL_PROJECT_ID)
  -n, --project string   Name of the project
```

### Options inherited from parent commands

```
      --config string     Path to JSON or YAML configuration file
      --exclude strings   Comma seperated Href references to collapse in results, may be dotted three levels deep
      --include strings   Comma seperated Href references to expand in results, may be dotted three levels deep
  -o, --output string     Output format (*table, json, yaml)
      --search string     Search keyword for use in 'get' actions. Search is not supported by all resources.
      --token string      Metal API Token (METAL_AUTH_TOKEN)
```

### SEE ALSO

* [metal project](metal_project.md)	 - Project operations

