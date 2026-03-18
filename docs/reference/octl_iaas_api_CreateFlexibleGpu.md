## octl iaas api CreateFlexibleGpu

Allocates a flexible GPU (fGPU) to your OUTSCALE account.

### Synopsis

Allocates a flexible GPU (fGPU) to your OUTSCALE account.

You can then attach this fGPU to a virtual machine (VM).

For more information, see [About Flexible GPUs](https://docs.outscale.com/en/userguide/About-Flexible-GPUs.html).

```
octl iaas api CreateFlexibleGpu [flags]
```

### Options

```
      --DeleteOnVmDeletion     If true, the fGPU is deleted when the VM is terminated.
      --DryRun                 If true, checks whether you have the required permissions to perform the action.
      --Generation string      The processor generation that the fGPU must be compatible with.
      --ModelName string       The model of fGPU you want to allocate.
      --SubregionName string   The Subregion in which you want to create the fGPU.
  -h, --help                   help for CreateFlexibleGpu
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

