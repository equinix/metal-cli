## packet 2fa receive

Receive two factor authentication token

### Synopsis

Example:
Issue the token via SMS:
packet 2fa receive -s 

Issue the token via app:
packet 2fa receive -a



```
packet 2fa receive [flags]
```

### Options

```
  -a, --app    Issues otp uri for auth application
  -h, --help   help for receive
  -s, --sms    Issues SMS otp token to user's phone
```

### Options inherited from parent commands

```
      --config string     Path to JSON or YAML configuration file
      --exclude strings   Comma seperated Href references to collapse in results, may be dotted three levels deep
      --include strings   Comma seperated Href references to expand in results, may be dotted three levels deep
  -j, --json              JSON output
  -y, --yaml              YAML output
```

### SEE ALSO

* [packet 2fa](packet_2fa.md)	 - Two Factor Authentication operations

