## metal project bgp-enable

Enables BGP on a project

### Synopsis

Example:

metal project bgp-enable --id [project_UUID] --asn [asn] --md5 [md5_secret] --use-case [use_case] --deployment-type [deployment_type]



```
metal project bgp-enable [flags]
```

### Options

```
      --asn int                 Local ASN (default 65000)
      --deploymentType string   Deployment type (local, global)
  -h, --help                    help for bgp-enable
  -i, --id string               Project ID (METAL_PROJECT_ID)
      --md5 string              BGP Password
      --useCase string          Use case for BGP
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

* [metal project](metal_project.md)	 - Project operations. For more information on Equinix Metal Projects, visit https://metal.equinix.com/developers/docs/accounts/projects/.

