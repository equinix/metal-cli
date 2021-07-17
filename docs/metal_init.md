## metal init

Create a configuration file

### Synopsis

Init will prompt for account settings and store the values as defaults in a configuration file that may be shared with other Equinix Metal tools. This file is typically stored in $HOME/.config/equinix/metal.yaml.

		Any Metal CLI command line argument can be specified in the config file. Be careful not to define options that you do not intend to use as defaults.

		This action may request additional settings in the future.

		The configuration file written to can be changed with METAL_CONFIG and --config.

	Example config:

	--
	token: foo
	project-id: uuid
	organization-id: uuid
	

```
metal init
```

### Options

```
  -h, --help   help for init
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

* [metal](metal.md)	 - Command line interface for Equinix Metal

