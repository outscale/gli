## octl iaas api ReadUserGroups

Lists all the user groups of the OUTSCALE account.

### Synopsis

Lists all the user groups of the OUTSCALE account.

The response can be filtered using either the PathPrefix or the UserGroupIds.

```
octl iaas api ReadUserGroups [flags]
```

### Options

```
      --DryRun                         If true, checks whether you have the required permissions to perform the action.
      --Filters.PathPrefix string      The path prefix of the groups.
      --Filters.UserGroupIds strings   The IDs of the user groups.
      --FirstItem int                  The item starting the list of groups requested.
      --ResultsPerPage int             The maximum number of items that can be returned in a single response (by default, 100).
  -h, --help                           help for ReadUserGroups
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
      --payload string    JSON content for query body
      --profile string    Profile to use in profile file (by default, "default")
      --single            convert single entry lists to a single object
      --template string   JSON template file for query body
  -v, --verbose           Verbose output
  -y, --yes               answer yes to all prompts
```

### SEE ALSO

* [octl iaas api](octl_iaas_api.md)	 - iaas api calls

