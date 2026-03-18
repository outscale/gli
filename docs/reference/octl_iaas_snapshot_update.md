## octl iaas snapshot update

alias for api UpdateSnapshot --SnapshotId snapshot_id

### Synopsis

> *alias for api UpdateSnapshot --SnapshotId snapshot_id*

Modifies the permissions for a specified snapshot.

You must specify either the `Additions` or the `Removals` parameter.

After sharing a snapshot with an OUTSCALE account, the other account can create a copy of it that they own. For more information about copying snapshots, see [CreateSnapshot](#createsnapshot).

```
octl iaas snapshot update snapshot_id [snapshot_id]... [flags]
```

### Options

```
  -h, --help                                                      help for update
      --permission-to-create-volume-addition-account-id strings   One or more OUTSCALE account IDs that the permission is associated with.
      --permission-to-create-volume-addition-global-permission    A global permission for all accounts.
      --permission-to-create-volume-removal-account-id strings    One or more OUTSCALE account IDs that the permission is associated with.
      --permission-to-create-volume-removal-global-permission     A global permission for all accounts.
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

* [octl iaas snapshot](octl_iaas_snapshot.md)	 - snapshot commands

