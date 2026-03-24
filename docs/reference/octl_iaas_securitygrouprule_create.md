## octl iaas securitygrouprule create

alias for api CreateSecurityGroupRule

### Synopsis

> *alias for api CreateSecurityGroupRule*

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
octl iaas securitygrouprule create [flags]
```

### Options

```
      --account-id-to-link string                               The OUTSCALE account ID that owns the source or destination security group specified in the SecurityGroupNameToLink parameter.
      --flow string                                             The direction of the flow: Inbound or Outbound.
      --from-port-range int                                     The beginning of the port range for the TCP and UDP protocols, or an ICMP type number.
  -h, --help                                                    help for create
      --id string                                               The ID of the security group for which you want to create a rule.
      --ip-protocol string                                      The IP protocol name (tcp, udp, icmp, or -1 for all protocols).
      --ip-range string                                         The IP range for the security group rule, in CIDR notation (for example, 10.0.0.0/16).
      --name-to-link string                                     The ID of a source or destination security group that you want to link to the security group of the rule.
      --rule-from-port-range int                                The beginning of the port range for the TCP and UDP protocols, or an ICMP type number.
      --rule-ip-protocol string                                 The IP protocol name (tcp, udp, icmp, or -1 for all protocols).
      --rule-ip-range strings                                   One or more IP ranges for the security group rules, in CIDR notation (for example, ["10.0.0.0/24" , "10.0.1.0/24"]).
      --rule-security-group-member-account-id string            The OUTSCALE account ID that owns the source or destination security group.
      --rule-security-group-member-security-group-id string     The ID of a source or destination security group that you want to link to the security group of the rule.
      --rule-security-group-member-security-group-name string   The name of a source or destination security group that you want to link to the security group of the rule.
      --rule-security-group-rule-id string                      The ID of the security group rule.
      --rule-service-id strings                                 One or more service IDs to allow traffic from a Net to access the corresponding OUTSCALE services.
      --rule-to-port-range int                                  The end of the port range for the TCP and UDP protocols, or an ICMP code number.
      --to-port-range int                                       The end of the port range for the TCP and UDP protocols, or an ICMP code number.
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

* [octl iaas securitygrouprule](octl_iaas_securitygrouprule.md)	 - securitygrouprule commands

