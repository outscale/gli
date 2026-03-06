## octl iaas api CreateSnapshotExportTask

Exports a snapshot to an OUTSCALE Object Storage (OOS) bucket that belongs to you.

### Synopsis

Exports a snapshot to an OUTSCALE Object Storage (OOS) bucket that belongs to you. This action enables you to create a backup of your snapshot.

You can share this snapshot with others accounts by granting permission to read it via pre-signed URLs. For more information, see [Creating a Pre-Signed URL](https://docs.outscale.com/en/userguide/Creating-a-Pre-Signed-URL.html).

For more information, see [About Snapshots](https://docs.outscale.com/en/userguide/About-Snapshots.html).

```
octl iaas api CreateSnapshotExportTask [flags]
```

### Options

```
      --DryRun                                 If true, checks whether you have the required permissions to perform the action.
      --OsuExport.DiskImageFormat string       The format of the export disk (qcow2 | raw).
      --OsuExport.OsuApiKey.ApiKeyId string    The API key of the OOS account that enables you to access the bucket.
      --OsuExport.OsuApiKey.SecretKey string   The secret key of the OOS account that enables you to access the bucket.
      --OsuExport.OsuBucket string             The name of the OOS bucket where you want to export the object.
      --OsuExport.OsuManifestUrl string        The URL of the manifest file.
      --OsuExport.OsuPrefix string             The prefix for the key of the OOS object.
      --SnapshotId string                      The ID of the snapshot to export.
  -h, --help                                   help for CreateSnapshotExportTask
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

