## octl iaas api CreateVmGroup

> [WARNING] > This feature is currently under development and may not function properly.

### Synopsis

> [WARNING]

> This feature is currently under development and may not function properly.


Creates a group of virtual machines (VMs) containing the same characteristics as a specified VM template, and then launches them.

You can create up to 100 VM groups in your OUTSCALE account.

```
octl iaas api CreateVmGroup [flags]
```

### Options

```
      --Description string           A description for the VM group.
      --DryRun                       If true, checks whether you have the required permissions to perform the action.
      --PositioningStrategy string   The positioning strategy of VMs on hypervisors.
      --SecurityGroupIds strings     One or more IDs of security groups for the VM group.
      --SubnetId string              [REQUIRED] The ID of the Subnet in which you want to create the VM group.
      --Tags.0.Key string            The key of the tag, between 1 and 255 characters.
      --Tags.0.Value string          The value of the tag, between 0 and 255 characters.
      --VmCount int                  [REQUIRED] The number of VMs deployed in the VM group.
      --VmGroupName string           [REQUIRED] The name of the VM group.
      --VmTemplateId string          [REQUIRED] The ID of the VM template used to launch VMs in the VM group.
  -h, --help                         help for CreateVmGroup
```

### Options inherited from parent commands

```
  -c, --columns string             columns to display - [+]<title>:<jq query for content>||<title>:<jq query for content>
      --config string              Path of profile file (by default, ~/.osc/config.json)
      --dry-run                    Display the request payload that would be sent to the API without sending it
      --elapsed                    add elapsed time column when using --watch (default true)
      --filter strings             comma separated list of filters for results - name:value,name:value, alias for jq filter 'select(.name | tostring | test("value"))'
      --interval duration          interval between two watch/waitfor iterations (default 5s)
      --jq string                  jq filter
      --max-pages int              maximum number of pages a command can fetch (default 20)
      --no-upgrade                 do not check for new versions
  -O, --out-file string            redirect output to file
  -o, --output string              output format (raw, json, yaml, table, csv, none, text)
      --payload string             JSON content for query body
      --profile string             Profile to use in profile file (by default, "default")
  -s, --silent                     Hides all information messages
      --single                     convert single entry lists to a single object
      --template string            JSON template file for query body
  -v, --verbose                    Verbose output
      --waitfor string             repeatedly call the API until the specified jq expression returns 1/true or a non empty result
      --waitfor-timeout duration   maximum duration of a wait (default 10m0s)
      --watch                      repeatedly call the API and display changes
  -y, --yes                        answer yes to all prompts
```

### SEE ALSO

* [octl iaas api](octl_iaas_api.md)	 - Call iaas API

