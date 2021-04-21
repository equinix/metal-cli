## packet capacity get

Returns a list of facilities or metros and plans with their current capacity.

### Synopsis

Example:
Retrieve capacities:
packet capacity get { --metro | --facility }


```
packet capacity get [flags]
```

### Options

```
  -f, --facility   Facility code (sv15) (default true)
  -h, --help       help for get
  -m, --metro      Metro code (sv)
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

* [packet capacity](packet_capacity.md)	 - Capacities operations

