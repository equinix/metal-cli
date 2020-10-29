## packet 2fa enable

Enables two factor authentication

### Synopsis

Example:

Enable two factor authentication via SMS
packet 2fa enable -s -t [token]

Enable two factor authentication via APP
packet 2fa enable -a -t [token]


```
packet 2fa enable [flags]
```

### Options

```
  -a, --app            Issues otp uri for auth application
  -h, --help           help for enable
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

