## octl iaas api CreateSecurityGroupRule

Adds one or more rules to a security group.

### Synopsis

Adds one or more rules to a security group.

Use the `SecurityGroupId` parameter to specify the security group for which you want to create a rule.

Use the `Flow` parameter to specify if you want an inbound rule or an outbound rule.

An inbound rule allows the security group to receive traffic:
* Either from a specific IP range (`IpRange` parameter) on a specific port range (`FromPortRange` and `ToPortRange` parameters) and specific protocol (`IpProtocol` parameter).
* Or from another specific security group (`SecurityGroupAccountIdToLink` and `SecurityGroupNameToLink` parameters).

(Net only) An outbound rule works similarly but allows the security group to send traffic rather than receive traffic.

Alternatively, you can use the `Rules` parameter to add several rules at the same time. Note that the `SecurityGroupName` subparameter can only be used for security groups in the public Cloud.

**[NOTE]**

* The modifications are effective as quickly as possible, but a small delay may occur.

* By default, traffic between two security groups is allowed through both public and private IPs. To restrict traffic to private IPs only, contact our Support team at support@outscale.com.

For more information, see [About Security Group Rules](https://docs.outscale.com/en/userguide/About-Security-Group-Rules.html).

```
octl iaas api CreateSecurityGroupRule [flags]
```

### Options

```
      --DryRun                                                     If true, checks whether you have the required permissions to perform the action.
      --Flow string                                                The direction of the flow: Inbound or Outbound.
      --FromPortRange int                                          The beginning of the port range for the TCP and UDP protocols, or an ICMP type number.
      --IpProtocol string                                          The IP protocol name (tcp, udp, icmp, or -1 for all protocols).
      --IpRange string                                             The IP range for the security group rule, in CIDR notation (for example, 10.0.0.0/16).
      --Rules.0.FromPortRange int                                  The beginning of the port range for the TCP and UDP protocols, or an ICMP type number.
      --Rules.0.IpProtocol string                                  The IP protocol name (tcp, udp, icmp, or -1 for all protocols).
      --Rules.0.IpRanges strings                                   One or more IP ranges for the security group rules, in CIDR notation (for example, ["10.0.0.0/24" , "10.0.1.0/24"]).
      --Rules.0.SecurityGroupRuleId string                         The ID of the security group rule.
      --Rules.0.SecurityGroupsMembers.0.AccountId string           The account ID that owns the source or destination security group.
      --Rules.0.SecurityGroupsMembers.0.SecurityGroupId string     The ID of a source or destination security group that you want to link to the security group of the rule.
      --Rules.0.SecurityGroupsMembers.0.SecurityGroupName string   The name of a source or destination security group that you want to link to the security group of the rule.
      --Rules.0.ServiceIds strings                                 One or more service IDs to allow traffic from a Net to access the corresponding OUTSCALE services.
      --Rules.0.ToPortRange int                                    The end of the port range for the TCP and UDP protocols, or an ICMP code number.
      --SecurityGroupAccountIdToLink string                        The account ID that owns the source or destination security group specified in the SecurityGroupNameToLink parameter.
      --SecurityGroupId string                                     The ID of the security group for which you want to create a rule.
      --SecurityGroupNameToLink string                             The ID of a source or destination security group that you want to link to the security group of the rule.
      --ToPortRange int                                            The end of the port range for the TCP and UDP protocols, or an ICMP code number.
  -h, --help                                                       help for CreateSecurityGroupRule
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

###### Auto generated by spf13/cobra on 6-Mar-2026
