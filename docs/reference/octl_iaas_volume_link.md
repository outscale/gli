## octl iaas volume link

alias for api LinkVolume --VolumeId volume_id

### Synopsis

> *alias for api LinkVolume --VolumeId volume_id*

Attaches a Block Storage Unit (BSU) volume to a virtual machine (VM).

The volume and the VM must be in the same Subregion. The VM can be running or stopped. The volume is attached to the specified VM device.

```
octl iaas volume link volume_id [flags]
```

### Options

```
      --device_name string   The name of the device.
  -h, --help                 help for link
      --vm-id string         The ID of the VM you want to attach the volume to.
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

* [octl iaas volume](octl_iaas_volume.md)	 - volume commands

