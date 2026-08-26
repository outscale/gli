## octl iaas vmgroup create

alias for api CreateVmGroup

### Synopsis

> alias for api CreateVmGroup

> [WARNING]

> This feature is currently under development and may not function properly.


Creates a group of virtual machines (VMs) containing the same characteristics as a specified VM template, and then launches them.

You can create up to 100 VM groups in your OUTSCALE account.

```
octl iaas vmgroup create [flags]
```

### Options

```
      --description string            A description for the VM group.
  -h, --help                          help for create
      --name string                   [REQUIRED] The name of the VM group.
      --positioning-strategy string   The positioning strategy of VMs on hypervisors.
      --security-group-id strings     One or more IDs of security groups for the VM group.
      --subnet-id string              [REQUIRED] The ID of the Subnet in which you want to create the VM group.
      --tag-key string                The key of the tag, between 1 and 255 characters.
      --tag-value string              The value of the tag, between 0 and 255 characters.
      --vm-count int                  [REQUIRED] The number of VMs deployed in the VM group.
      --vm-template-id string         [REQUIRED] The ID of the VM template used to launch VMs in the VM group.
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
  -o, --output string              output format (raw, json, yaml, table, csv, none, base64, text)
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

* [octl iaas vmgroup](octl_iaas_vmgroup.md)	 - vmgroup commands

