## octl iaas snapshotexporttask list

alias for api ReadSnapshotExportTasks

### Synopsis

> *alias for api ReadSnapshotExportTasks*

Lists one or more snapshot export tasks.

```
octl iaas snapshotexporttask list [flags]
```

### Options

```
  -h, --help                  help for list
      --snapshot-id strings   The IDs of the snapshots used with the snapshot export tasks
      --task-id strings       The IDs of the snapshot export tasks.
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

* [octl iaas snapshotexporttask](octl_iaas_snapshotexporttask.md)	 - snapshotexporttask commands

