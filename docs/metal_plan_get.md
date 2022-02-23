## metal plan get

Retrieves a list plans.

### Synopsis

Retrieves a list of plans available to the current user. Response includes plan UUID, slug, and name.

```
metal plan get [flags]
```

### Examples

```
  # Lists the plans available to the current user:
  metal plans get
```

### Options

```
  -h, --help   help for get
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

* [metal plan](metal_plan.md)	 - Plan operations. For more information on the different Equinix Metal plans, visit https://metal.equinix.com/developers/docs/servers/server-specs/.

