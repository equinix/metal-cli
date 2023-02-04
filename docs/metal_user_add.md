## metal user add

Adds a user to an organization or project

### Synopsis

Adds a user, by email, to the organization or project specified by the --organization-id or --project-id flag. The user will be assigned the roles specified by the --roles flag.

```
metal user add --email <email> --roles <roles> [--organization-id <organization_id>] [--project-id <project_id>] [flags]
```

### Examples

```
  # Adds a user to a project with admin role:
  metal user add --email user@example.org --roles admin --project-id 3b0795ba-fd0b-4a9e-83a7-063e5e12409d

```

### Options

```
      --email string             Email of the user.
  -h, --help                     help for add
      --organization-id string   Organization to invite the user to.
  -p, --project-id strings       Projects to invite the user to with the specified roles.
      --roles strings            Roles to assign to the user.
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

* [metal user](metal_user.md)	 - User operations: get and add.

