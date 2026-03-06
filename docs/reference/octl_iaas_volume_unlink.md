## octl iaas volume unlink

alias for api UnlinkVolume --VolumeId link_volume_id

### Synopsis

> *alias for api UnlinkVolume --VolumeId link_volume_id*

Detaches a Block Storage Unit (BSU) volume from a virtual machine (VM).

To detach the root device of a VM, this VM must be stopped.

```
octl iaas volume unlink volume_id [flags]
```

### Options

```
      --force   Forces the detachment of the volume in case of previous failure.
  -h, --help    help for unlink
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

* [octl iaas volume](octl_iaas_volume.md)	 - volume commands

