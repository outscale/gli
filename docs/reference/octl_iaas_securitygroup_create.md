## octl iaas securitygroup create

alias for api CreateSecurityGroup

### Synopsis

> *alias for api CreateSecurityGroup*

Creates a security group.

This action creates a security group either in the public Cloud or in a specified Net. By default, a default security group for use in the public Cloud and a default security group for use in a Net are created.

When launching a virtual machine (VM), if no security group is explicitly specified, the appropriate default security group is assigned to the VM. Default security groups include a default rule granting VMs network access to each other.

When creating a security group, you specify a name. Two security groups for use in the public Cloud or for use in a Net cannot have the same name.

You can have up to 500 security groups in the public Cloud. You can create up to 500 security groups per Net.

To add or remove rules, use the [CreateSecurityGroupRule](#createsecuritygrouprule) method.

For more information, see [About Security Groups](https://docs.outscale.com/en/userguide/About-Security-Groups.html).

```
octl iaas securitygroup create [flags]
```

### Options

```
      --description string   A description for the security group.
  -h, --help                 help for create
      --name string          The name of the security group.
      --net-id string        The ID of the Net for the security group.
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

* [octl iaas securitygroup](octl_iaas_securitygroup.md)	 - securitygroup commands

