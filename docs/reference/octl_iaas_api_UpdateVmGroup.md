## octl iaas api UpdateVmGroup

> [WARNING] > This feature is currently under development and may not function properly.

### Synopsis

> [WARNING]

> This feature is currently under development and may not function properly.

Modifies the specified attributes of a group of virtual machines (VMs).

```
octl iaas api UpdateVmGroup [flags]
```

### Options

```
      --Description string    A new description for the VM group.
      --DryRun                If true, checks whether you have the required permissions to perform the action.
      --Tags.0.Key string     The key of the tag, between 1 and 255 characters.
      --Tags.0.Value string   The value of the tag, between 0 and 255 characters.
      --VmGroupId string      The ID of the VM group you want to update.
      --VmGroupName string    A new name for your VM group.
      --VmTemplateId string   A new VM template ID for your VM group.
  -h, --help                  help for UpdateVmGroup
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

