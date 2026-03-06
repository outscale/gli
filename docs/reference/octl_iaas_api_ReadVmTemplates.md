## octl iaas api ReadVmTemplates

> [WARNING] > This feature is currently under development and may not function properly.

### Synopsis

> [WARNING]

> This feature is currently under development and may not function properly.

Lists one or more virtual machine (VM) templates.

```
octl iaas api ReadVmTemplates [flags]
```

### Options

```
      --DryRun                            If true, checks whether you have the required permissions to perform the action.
      --Filters.CpuCores ints             The number of vCores.
      --Filters.CpuGenerations strings    The processor generations (for example, v4).
      --Filters.CpuPerformances strings   The performances of the VMs.
      --Filters.Descriptions strings      The descriptions of the VM templates.
      --Filters.ImageIds strings          The IDs of the OMIs.
      --Filters.KeypairNames strings      The names of the keypairs.
      --Filters.Rams ints                 The amount of RAM.
      --Filters.TagKeys strings           The keys of the tags associated with the VM templates.
      --Filters.TagValues strings         The values of the tags associated with the VM templates.
      --Filters.Tags strings              The key/value combination of the tags associated with the VM templates, in the following format: "Filters":{"Tags":["TAGKEY=TAGVALUE"]}.
      --Filters.VmTemplateIds strings     The IDs of the VM templates.
      --Filters.VmTemplateNames strings   The names of the VM templates.
  -h, --help                              help for ReadVmTemplates
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

