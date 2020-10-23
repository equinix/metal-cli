## packet 2fa disable

Disables two factor authentication

### Synopsis

Example:

Disable two factor authentication via SMS
packet 2fa disable -s -t [token]

Disable two factor authentication via APP
packet 2fa disable -a -t [token]


```
packet 2fa disable [flags]
```

### Options

```
  -a, --app            Issues otp uri for auth application
  -h, --help           help for disable
  -s, --sms            Issues SMS otp token to user's phone
  -t, --token string   Two factor authentication token
```

### Options inherited from parent commands

```
      --config string     Path to JSON or YAML configuration file
      --exclude strings   Comma seperated Href references to collapse in results, may be dotted three levels deep
      --include strings   Comma seperated Href references to expand in results, may be dotted three levels deep
  -j, --json              JSON output
      --search string     Search keyword for use in 'get' actions. Search is not supported by all resources.
  -y, --yaml              YAML output
```

### SEE ALSO

* [packet 2fa](packet_2fa.md)	 - Two Factor Authentication operations

