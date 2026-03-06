## octl iaas flexiblegpu update

alias for api UpdateFlexibleGpu --FlexibleGpuId flexible_gpu_id

### Synopsis

> *alias for api UpdateFlexibleGpu --FlexibleGpuId flexible_gpu_id*

Modifies a flexible GPU (fGPU) behavior.

```
octl iaas flexiblegpu update flexible_gpu_id [flexible_gpu_id]... [flags]
```

### Options

```
      --delete-on-vm-deletion   If true, the fGPU is deleted when the VM is terminated.
  -h, --help                    help for update
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

* [octl iaas flexiblegpu](octl_iaas_flexiblegpu.md)	 - flexiblegpu commands

