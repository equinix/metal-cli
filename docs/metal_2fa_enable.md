## metal 2fa enable

Enables two factor authentication

### Synopsis

Example:

Enable two factor authentication via SMS
metal 2fa enable -s -c [code]

Enable two factor authentication via APP
metal 2fa enable -a -c [code]


```
metal 2fa enable [flags]
```

### Options

```
  -a, --app           Issues otp uri for auth application
  -c, --code string   Two factor authentication code
  -h, --help          help for enable
  -s, --sms           Issues SMS otp token to user's phone
```

### Options inherited from parent commands

```
      --config string     Path to JSON or YAML configuration file
      --exclude strings   Comma seperated Href references to collapse in results, may be dotted three levels deep
      --include strings   Comma seperated Href references to expand in results, may be dotted three levels deep
  -o, --output string     Output format (*table, json, yaml)
      --search string     Search keyword for use in 'get' actions. Search is not supported by all resources.
      --sort-by string    Sort fields for use in 'get' actions. Sort is not supported by all resources.
      --sort-dir string   Sort field direction for use in 'get' actions. Sort is not supported by all resources.
      --token string      Metal API Token (METAL_AUTH_TOKEN)
```

### SEE ALSO

* [metal 2fa](metal_2fa.md)	 - Two Factor Authentication operations

