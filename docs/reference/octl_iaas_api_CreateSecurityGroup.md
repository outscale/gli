## octl iaas api CreateSecurityGroup

Creates a security group.

### Synopsis

Creates a security group.

This action creates a security group either in the public Cloud or in a specified Net. By default, a default security group for use in the public Cloud and a default security group for use in a Net are created.

When launching a virtual machine (VM), if no security group is explicitly specified, the appropriate default security group is assigned to the VM. Default security groups include a default rule granting VMs network access to each other.

When creating a security group, you specify a name. Two security groups for use in the public Cloud or for use in a Net cannot have the same name.

You can have up to 500 security groups in the public Cloud. You can create up to 500 security groups per Net.

To add or remove rules, use the [CreateSecurityGroupRule](#createsecuritygrouprule) method.

For more information, see [About Security Groups](https://docs.outscale.com/en/userguide/About-Security-Groups.html).

```
octl iaas api CreateSecurityGroup [flags]
```

### Options

```
      --Description string         A description for the security group.
      --DryRun                     If true, checks whether you have the required permissions to perform the action.
      --NetId string               The ID of the Net for the security group.
      --SecurityGroupName string   The name of the security group.
  -h, --help                       help for CreateSecurityGroup
```

### Options inherited from parent commands

```
  -c, --columns string              columns to display - [+]<title>:<jq query for content>||<title>:<jq query for content>
      --config string               Path of profile file (by default, ~/.osc/config.json)
      --filter strings              comma separated list of filters for results - name:value,name:value, alias for jq filter 'select(.name | test("value"))'
      --jq string                   jq filter
      --no-upgrade                  do not check for new versions
  -O, --out-file string             redirect output to file
  -o, --output string               output format (raw, json, yaml, table, csv, none, base64) (default "raw")
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

