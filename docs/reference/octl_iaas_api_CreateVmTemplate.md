## octl iaas api CreateVmTemplate

> [WARNING] > This feature is currently under development and may not function properly.

### Synopsis

> [WARNING]

> This feature is currently under development and may not function properly.

Creates a virtual machine (VM) template. You can then use the VM template to create VM groups.

You can create up to 50 VM templates in your OUTSCALE account.

```
octl iaas api CreateVmTemplate [flags]
```

### Options

```
      --CpuCores int            The number of vCores to use for each VM.
      --CpuGeneration string    The processor generation to use for each VM (for example, v4).
      --CpuPerformance string   The performance of the VMs.
      --Description string      A description for the VM template.
      --DryRun                  If true, checks whether you have the required permissions to perform the action.
      --ImageId string          The ID of the OMI to use for each VM.
      --KeypairName string      The name of the keypair to use for each VM.
      --Ram int                 The amount of RAM to use for each VM.
      --Tags.0.Key string       The key of the tag, between 1 and 255 characters.
      --Tags.0.Value string     The value of the tag, between 0 and 255 characters.
      --VmTemplateName string   The name of the VM template.
  -h, --help                    help for CreateVmTemplate
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

