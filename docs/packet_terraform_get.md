## packet terraform get

Retrieves all available terraform modules or a single terraform modules

### Synopsis

Example:

Retrieve all projects:
packet terraform get
  
Retrieve a specific project:
packet terraform get -n [project_name]
	

```
packet terraform get [flags]
```

### Options

```
  -h, --help          help for get
  -n, --name string   Name of the terraform module
```

### Options inherited from parent commands

```
      --config string   Path to JSON or YAML configuration file
```

### SEE ALSO

* [packet terraform](packet_terraform.md)	 - Terraform operations

