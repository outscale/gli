## octl iaas user list

alias for api ReadUsers

### Synopsis

> *alias for api ReadUsers*

Lists all EIM users in the account.

The response can be filtered using the UserIds.

```
octl iaas user list [flags]
```

### Options

```
      --first-item int   The item starting the list of users requested.
  -h, --help             help for list
      --id strings       The IDs of the users.
```

### Options inherited from parent commands

```
  -c, --columns string    columns to display - [+]<title>:<jq query for content>||<title>:<jq query for content>
      --config string     Path of profile file (by default, ~/.osc/config.json)
      --filter strings    comma separated list of filters for results - name:value,name:value, alias for jq filter 'select(.name | test("value"))'
      --jq string         jq filter
      --no-upgrade        do not check for new versions
  -O, --out-file string   redirect output to file
  -o, --output string     output format (raw, json, yaml, table, csv, none, base64) (default "raw")
      --profile string    Profile to use in profile file (by default, "default")
      --single            convert single entry lists to a single object
      --template string   JSON template for query body
  -v, --verbose           Verbose output
  -y, --yes               answer yes to all prompts
```

### SEE ALSO

* [octl iaas user](octl_iaas_user.md)	 - user commands

