## metal 2fa disable

Disables two-factor authentication.

### Synopsis

Disables two-factor authentication. Requires the current OTP code from either SMS or application. If you no longer have access to your two-factor authentication device, please contact support.

```
metal 2fa disable (-a | -s) --code <OTP_code>  [flags]
```

### Examples

```
  # Disable two-factor authentication via SMS
  metal 2fa disable -s -c <OTP_code>

  # Disable two-factor authentication via APP
  metal 2fa disable -a -c <OTP_code>
```

### Options

```
  -a, --app           The OTP code is issued from an application.
  -c, --code string   The two-factor authentication OTP code.
  -h, --help          help for disable
  -s, --sms           The OTP code is issued to you via SMS.
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

