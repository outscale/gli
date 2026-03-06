## octl iaas api DeleteDedicatedGroup

Deletes a specified dedicated group of virtual machines (VMs).

### Synopsis

Deletes a specified dedicated group of virtual machines (VMs).

**[WARNING]**

A dedicated group can be deleted only if no VM or Net is in the dedicated group. Otherwise, you need to force the deletion.

If you force the deletion:

- all VMs are terminated.

- all Nets are deleted, and all resources associated with Nets are detached.

```
octl iaas api DeleteDedicatedGroup [flags]
```

### Options

```
      --DedicatedGroupId string   The ID of the dedicated group you want to delete.
      --DryRun                    If true, checks whether you have the required permissions to perform the action.
      --Force                     If true, forces the deletion of the dedicated group and all its dependencies.
  -h, --help                      help for DeleteDedicatedGroup
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

