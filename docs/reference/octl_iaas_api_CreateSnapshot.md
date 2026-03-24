## octl iaas api CreateSnapshot

Creates a snapshot.

### Synopsis

Creates a snapshot. Snapshots are point-in-time images of a volume that you can use to back up your data or to create replicas of this volume.

You can use this method in three different ways:
* **Creating from a volume**: You create a snapshot from one of your volumes.

* **Copying a snapshot**: You copy an existing snapshot. The source snapshot can be one of your own snapshots, or a snapshot owned by another OUTSCALE account that has granted you permission via the [UpdateSnapshot](#updatesnapshot) method.

* **Importing from a bucket**: You import a snapshot located in an OUTSCALE Object Storage (OOS) bucket. First, the owner of the source snapshot must export it to a bucket by using the [CreateSnapshotExportTask](#createsnapshotexporttask) method. Then, they must grant you permission to read the snapshot via a pre-signed URL. For more information, see [Creating a Pre-Signed URL](https://docs.outscale.com/en/userguide/Creating-a-Pre-Signed-URL.html).

**[NOTE]**

In case of excessive use of the snapshot creation feature on the same volume over a short period of time, 3DS OUTSCALE reserves the right to temporarily block the feature.

For more information, see [About Snapshots](https://docs.outscale.com/en/userguide/About-Snapshots.html).

```
octl iaas api CreateSnapshot [flags]
```

### Options

```
      --ClientToken string        A unique identifier which enables you to manage the idempotency.
      --Description string        A description for the snapshot.
      --DryRun                    If true, checks whether you have the required permissions to perform the action.
      --FileLocation string       **(when importing from a bucket)** The pre-signed URL of the snapshot you want to import.
      --SnapshotSize int          **(when importing from a bucket)** The size of the snapshot you want to create in your account, in bytes.
      --SourceRegionName string   **(when copying a snapshot)** The name of the source Region, which must be the same as the Region of your account.
      --SourceSnapshotId string   **(when copying a snapshot)** The ID of the snapshot you want to copy.
      --VolumeId string           **(when creating from a volume)** The ID of the volume you want to create a snapshot of.
  -h, --help                      help for CreateSnapshot
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

