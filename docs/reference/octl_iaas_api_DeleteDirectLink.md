## octl iaas api DeleteDirectLink

Deletes a specified DirectLink.

### Synopsis

Deletes a specified DirectLink.

Before deleting a DirectLink, ensure that all your DirectLink interfaces related to this DirectLink are deleted.

```
octl iaas api DeleteDirectLink [flags]
```

### Options

```
      --DirectLinkId string   The ID of the DirectLink you want to delete.
      --DryRun                If true, checks whether you have the required permissions to perform the action.
  -h, --help                  help for DeleteDirectLink
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

* [octl iaas api](octl_iaas_api.md)	 - iaas api calls

