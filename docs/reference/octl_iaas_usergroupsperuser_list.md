## octl iaas usergroupsperuser list

alias for api ReadUserGroupsPerUser

### Synopsis

> *alias for api ReadUserGroupsPerUser*

Lists the groups a specified user belongs to.

```
octl iaas usergroupsperuser list [flags]
```

### Options

```
  -h, --help               help for list
      --user-name string   The name of the user.
      --user-path string   The path to the user (by default, /).
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

* [octl iaas usergroupsperuser](octl_iaas_usergroupsperuser.md)	 - usergroupsperuser commands

