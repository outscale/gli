## octl iaas api ReadVolumeUpdateTasks

Lists one or more specified tasks of volume update.

### Synopsis

Lists one or more specified tasks of volume update.

```
octl iaas api ReadVolumeUpdateTasks [flags]
```

### Options

```
      --DryRun                      If true, checks whether you have the required permissions to perform the action.
      --Filters.TaskIds strings     The IDs of the snapshot export tasks.
      --Filters.VolumeIds strings   The IDs of the volumes used for the snapshot export tasks.
      --NextPageToken string        The token to request the next page of results.
      --ResultsPerPage int          The maximum number of logs returned in a single response (between 1 and 1000, both included).
  -h, --help                        help for ReadVolumeUpdateTasks
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

