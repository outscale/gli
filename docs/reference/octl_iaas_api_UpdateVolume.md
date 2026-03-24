## octl iaas api UpdateVolume

Modifies the specified attributes of a volume.

### Synopsis

Modifies the specified attributes of a volume.

**[WARNING]**

- We recommend creating a snapshot of your volume before updating it, in case any issue occurs during the process. For more information, see [Creating a Snapshot of a Volume](https://docs.outscale.com/en/userguide/Creating-a-Snapshot-of-a-Volume.html).
- Do not shut down or restart the virtual machine (VM) from within the guest operating system while a volume update is in progress. This interrupts the process and compromises the integrity of the volume.
- When the modification is not instantaneous, the response displays the previous value. You can use the [ReadVolumeUpdateTasks](#readvolumeupdatetasks) method to see the progression of the update.

```
octl iaas api UpdateVolume [flags]
```

### Options

```
      --DryRun              If true, checks whether you have the required permissions to perform the action.
      --Iops int            The new number of I/O operations per second (IOPS).
      --Size int            The new size of the volume, in gibibytes (GiB).
      --VolumeId string     The ID of the volume you want to update.
      --VolumeType string   The new type of the volume (standard | io1 | gp2).
  -h, --help                help for UpdateVolume
```

### Options inherited from parent commands

```
  -c, --columns string              columns to display - [+]<title>:<jq query for content>||<title>:<jq query for content>
      --config string               Path of profile file (by default, ~/.osc/config.json)
      --filter strings              comma separated list of filters for results - name:value,name:value, alias for jq filter 'select(.name | test("value"))'
      --jq string                   jq filter
      --no-upgrade                  do not check for new versions
  -O, --out-file string             redirect output to file
  -o, --output string               output format (raw, json, yaml, table, csv, none, base64, text) (default "raw")
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

