## octl iaas vmgroup create

alias for api CreateVmGroup

### Synopsis

> *alias for api CreateVmGroup*

> [WARNING]

> This feature is currently under development and may not function properly.

Creates a group of virtual machines (VMs) containing the same characteristics as a specified VM template, and then launches them.

You can create up to 100 VM groups in your account.

```
octl iaas vmgroup create [flags]
```

### Options

```
      --description string            A description for the VM group.
  -h, --help                          help for create
      --name string                   The name of the VM group.
      --positioning-strategy string   The positioning strategy of VMs on hypervisors.
      --security-group-id strings     One or more IDs of security groups for the VM group.
      --subnet-id string              The ID of the Subnet in which you want to create the VM group.
      --tag-key string                The key of the tag, between 1 and 255 characters.
      --tag-value string              The value of the tag, between 0 and 255 characters.
      --vm-count int                  The number of VMs deployed in the VM group.
      --vm-template-id string         The ID of the VM template used to launch VMs in the VM group.
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

* [octl iaas vmgroup](octl_iaas_vmgroup.md)	 - vmgroup commands

