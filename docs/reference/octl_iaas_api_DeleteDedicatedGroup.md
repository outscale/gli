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

* [octl iaas api](octl_iaas_api.md)	 - iaas api calls

