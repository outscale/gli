## octl iaas api CreateDedicatedGroup

Creates a dedicated group for virtual machines (VMs).

### Synopsis

Creates a dedicated group for virtual machines (VMs).

For more information, see [About Dedicated Groups](https://docs.outscale.com/en/userguide/About-Dedicated-Groups.html).

```
octl iaas api CreateDedicatedGroup [flags]
```

### Options

```
      --CpuGeneration int      The processor generation for the VMs in the dedicated group (for example, 4).
      --DryRun                 If true, checks whether you have the required permissions to perform the action.
      --Name string            A name for the dedicated group.
      --SubregionName string   The Subregion in which you want to create the dedicated group.
  -h, --help                   help for CreateDedicatedGroup
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

