## metal 2fa

Two-factor Authentication operations: receive, enable, disable.

### Synopsis

Enable or disable two-factor authentication on your user account or receive an OTP token. More information is available at https://metal.equinix.com/developers/docs/accounts/users/#multi-factor-authentication.

### Options

```
  -h, --help   help for 2fa
```

### Options inherited from parent commands

```
      --config string        Path to JSON or YAML configuration file
      --exclude strings      Comma separated Href references to collapse in results, may be dotted three levels deep
      --filter stringArray   Filter 'get' actions with name value pairs. Filter is not supported by all resources and is implemented as request query parameters.
      --include strings      Comma separated Href references to expand in results, may be dotted three levels deep
  -o, --output string        Output format (*table, json, yaml). env output formats are (*sh, terraform, capp).
      --search string        Search keyword for use in 'get' actions. Search is not supported by all resources.
      --sort-by string       Sort fields for use in 'get' actions. Sort is not supported by all resources.
      --sort-dir string      Sort field direction for use in 'get' actions. Sort is not supported by all resources.
      --token string         Metal API Token (METAL_AUTH_TOKEN)
```

### SEE ALSO

* [metal](metal.md)	 - Command line interface for Equinix Metal
* [metal 2fa disable](metal_2fa_disable.md)	 - Disables two-factor authentication.
* [metal 2fa enable](metal_2fa_enable.md)	 - Enables two factor authentication.
* [metal 2fa receive](metal_2fa_receive.md)	 - Generates a two-factor authentication token for use in enabling two-factor authentication on the current user's account.

