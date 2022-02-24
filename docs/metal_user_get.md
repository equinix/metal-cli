## metal user get

Retrieves information about the current user or a specified user.

### Synopsis

Returns either information about the current user or information about a specified user. Specified user information is only available if that user shares a project with the current user.

```
metal user get [-i <user_UUID>] [flags]
```

### Examples

```
  # Retrieves the current user's information:
  metal user get
  
  # Returns information on user 3b0795ba-fd0b-4a9e-83a7-063e5e12409d:
  metal user get --i 3b0795ba-fd0b-4a9e-83a7-063e5e12409d
```

### Options

```
  -h, --help        help for get
  -i, --id string   UUID of the user.
```

### Options inherited from parent commands

```
      --config string        Path to JSON or YAML configuration file
      --exclude strings      Comma separated Href references to collapse in results, may be dotted three levels deep
      --filter stringArray   Filter 'get' actions with name value pairs. Filter is not supported by all resources and is implemented as request query parameters.
      --include strings      Comma separated Href references to expand in results, may be dotted three levels deep
  -o, --output string        Output format (*table, json, yaml)
      --search string        Search keyword for use in 'get' actions. Search is not supported by all resources.
      --sort-by string       Sort fields for use in 'get' actions. Sort is not supported by all resources.
      --sort-dir string      Sort field direction for use in 'get' actions. Sort is not supported by all resources.
      --token string         Metal API Token (METAL_AUTH_TOKEN)
```

### SEE ALSO

* [metal user](metal_user.md)	 - User operations. For more information on user and account management, visit https://metal.equinix.com/developers/docs/accounts/users/ in the Equinix Metal documentation.

