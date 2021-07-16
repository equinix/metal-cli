## metal 2fa disable

Disables two factor authentication

### Synopsis

Example:

Disable two factor authentication via SMS
metal 2fa disable -s -c [code]

Disable two factor authentication via APP
metal 2fa disable -a -c [code]


```
metal 2fa disable [flags]
```

### Options

```
  -a, --app           Issues otp uri for auth application
  -c, --code string   Two factor authentication code
  -h, --help          help for disable
  -s, --sms           Issues SMS otp token to user's phone
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

* [metal 2fa](metal_2fa.md)	 - Two Factor Authentication operations

