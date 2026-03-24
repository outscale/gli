## octl iaas api CreateVolume

Creates a Block Storage Unit (BSU) volume in a specified Region.

### Synopsis

Creates a Block Storage Unit (BSU) volume in a specified Region.

BSU volumes can be attached to a virtual machine (VM) in the same Subregion. You can create an empty volume or restore a volume from an existing snapshot.

You can create the following volume types: Enterprise (`io1`) for provisioned IOPS SSD volumes, Performance (`gp2`) for general purpose SSD volumes, or Magnetic (`standard`) volumes.

For more information, see [About Volumes](https://docs.outscale.com/en/userguide/About-Volumes.html).

```
octl iaas api CreateVolume [flags]
```

### Options

```
      --ClientToken string     A unique identifier which enables you to manage the idempotency.
      --DryRun                 If true, checks whether you have the required permissions to perform the action.
      --Iops int               The number of I/O operations per second (IOPS).
      --Size int               The size of the volume, in gibibytes (GiB).
      --SnapshotId string      The ID of the snapshot from which you want to create the volume.
      --SubregionName string   The Subregion in which you want to create the volume.
      --VolumeType string      The type of volume you want to create (io1 | gp2 | standard).
  -h, --help                   help for CreateVolume
```

### Options inherited from parent commands

```
  -c, --columns string              columns to display - [+]<title>:<jq query for content>||<title>:<jq query for content>
      --config string               Path of profile file (by default, ~/.osc/config.json)
      --filter strings              comma separated list of filters for results - name:value,name:value, alias for jq filter 'select(.name | test("value"))'
      --jq string                   jq filter
      --no-upgrade                  do not check for new versions
  -O, --out-file string             redirect output to file
  -o, --output string               output format (raw, json, yaml, table, csv, none, base64) (default "raw")
      --payload string              JSON content for query body
      --profile string              Profile to use in profile file (by default, "default")
      --single                      convert single entry lists to a single object
      --template string             JSON template file for query body
  -v, --verbose                     Verbose output
      --waitfor string              jq expression to wait for - octl will query every waitfor-interval until the expression returns 1/true or a non empty result
      --waitfor-interval duration   interval between two waitfor iterations (default 5s)
      --waitfor-timeout duration    maximum duration of a wait (default 10m0s)
  -y, --yes                         answer yes to all prompts
```

### SEE ALSO

* [octl iaas api](octl_iaas_api.md)	 - iaas api calls

