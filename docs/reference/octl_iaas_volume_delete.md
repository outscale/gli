## octl iaas volume delete

alias for api DeleteVolume --VolumeId volume_id

### Synopsis

> *alias for api DeleteVolume --VolumeId volume_id*

Deletes a specified Block Storage Unit (BSU) volume.

You can delete available volumes only, that is, volumes that are not attached to a virtual machine (VM).

```
octl iaas volume delete volume_id [volume_id]... [flags]
```

### Options

```
  -h, --help   help for delete
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

