## packet project delete

Deletes a project

### Synopsis

Example:

packet project delete --id [project_UUID]



```
packet project delete [flags]
```

### Options

```
  -f, --force       Force removal of the project
  -h, --help        help for delete
  -i, --id string   UUID of the project
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

