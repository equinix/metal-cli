## metal 2fa receive

Generates a two-factor authentication token for use in enabling two-factor authentication on the current user's account.

### Synopsis

Generates a two-factor authentication token for use in enabling two-factor authentication on the current user's account. In order to use SMS, a phone number must be associated with the account to receive the code. If you are using an app, a URI for the application is returned.

```
metal 2fa receive (-s | -a) [flags]
```

### Examples

```
  # Issue the token via SMS:
  metal 2fa receive -s 

  # Issue the token via app:
  metal 2fa receive -a
```

### Options

```
  -a, --app    Issues an OTP URI for an authentication application.
  -h, --help   help for receive
  -s, --sms    Issues SMS OTP token to the phone number associated with the current user account.
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

* [metal 2fa](metal_2fa.md)	 - Two-factor Authentication operations. More information is available at https://metal.equinix.com/developers/docs/accounts/two-factor-authentication/.

