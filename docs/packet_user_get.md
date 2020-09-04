## packet user get

Retrieves information about the current user or a specified user

### Synopsis

Example:

Retrieve the current user:
packet user get
  
Retrieve a specific user:
packet user get --id [user_UUID]

  

```
packet user get [flags]
```

### Options

```
  -h, --help        help for get
  -i, --id string   UUID of the user
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

* [packet user](packet_user.md)	 - User operations

