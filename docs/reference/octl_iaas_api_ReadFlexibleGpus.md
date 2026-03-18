## octl iaas api ReadFlexibleGpus

Lists one or more flexible GPUs (fGPUs) allocated to your OUTSCALE account.

### Synopsis

Lists one or more flexible GPUs (fGPUs) allocated to your OUTSCALE account.

```
octl iaas api ReadFlexibleGpus [flags]
```

### Options

```
      --DryRun                               If true, checks whether you have the required permissions to perform the action.
      --Filters.DeleteOnVmDeletion           Indicates whether the fGPU is deleted when terminating the VM.
      --Filters.FlexibleGpuIds strings       One or more IDs of fGPUs.
      --Filters.Generations strings          The processor generations that the fGPUs are compatible with.
      --Filters.ModelNames strings           One or more models of fGPUs.
      --Filters.States strings               The states of the fGPUs (allocated | attaching | attached | detaching).
      --Filters.SubregionNames strings       The Subregions where the fGPUs are located.
      --Filters.Tags.0.Key string            The key of the tag, between 1 and 255 characters.
      --Filters.Tags.0.ResourceId string     The ID of the resource.
      --Filters.Tags.0.ResourceType string   The type of the resource.
      --Filters.Tags.0.Value string          The value of the tag, between 0 and 255 characters.
      --Filters.VmIds strings                One or more IDs of VMs.
  -h, --help                                 help for ReadFlexibleGpus
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

