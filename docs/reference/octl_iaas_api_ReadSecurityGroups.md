## octl iaas api ReadSecurityGroups

Lists one or more security groups.

### Synopsis

Lists one or more security groups.

You can specify either the name of the security groups or their IDs.

```
octl iaas api ReadSecurityGroups [flags]
```

### Options

```
      --DryRun                                           If true, checks whether you have the required permissions to perform the action.
      --Filters.Descriptions strings                     The descriptions of the security groups.
      --Filters.InboundRuleAccountIds strings            The OUTSCALE account IDs that have been granted permissions.
      --Filters.InboundRuleFromPortRanges ints           The beginnings of the port ranges for the TCP and UDP protocols, or the ICMP type numbers.
      --Filters.InboundRuleIpRanges strings              The IP ranges that have been granted permissions, in CIDR notation (for example, 10.0.0.0/24).
      --Filters.InboundRuleProtocols strings             The IP protocols for the permissions (tcp | udp | icmp, or a protocol number, or -1 for all protocols).
      --Filters.InboundRuleSecurityGroupIds strings      The IDs of the security groups that have been granted permissions.
      --Filters.InboundRuleSecurityGroupNames strings    The names of the security groups that have been granted permissions.
      --Filters.InboundRuleToPortRanges ints             The ends of the port ranges for the TCP and UDP protocols, or the ICMP code numbers.
      --Filters.NetIds strings                           The IDs of the Nets specified when the security groups were created.
      --Filters.OutboundRuleAccountIds strings           The OUTSCALE account IDs that have been granted permissions.
      --Filters.OutboundRuleFromPortRanges ints          The beginnings of the port ranges for the TCP and UDP protocols, or the ICMP type numbers.
      --Filters.OutboundRuleIpRanges strings             The IP ranges that have been granted permissions, in CIDR notation (for example, 10.0.0.0/24).
      --Filters.OutboundRuleProtocols strings            The IP protocols for the permissions (tcp | udp | icmp, or a protocol number, or -1 for all protocols).
      --Filters.OutboundRuleSecurityGroupIds strings     The IDs of the security groups that have been granted permissions.
      --Filters.OutboundRuleSecurityGroupNames strings   The names of the security groups that have been granted permissions.
      --Filters.OutboundRuleToPortRanges ints            The ends of the port ranges for the TCP and UDP protocols, or the ICMP code numbers.
      --Filters.SecurityGroupIds strings                 The IDs of the security groups.
      --Filters.SecurityGroupNames strings               The names of the security groups.
      --Filters.TagKeys strings                          The keys of the tags associated with the security groups.
      --Filters.TagValues strings                        The values of the tags associated with the security groups.
      --Filters.Tags strings                             The key/value combination of the tags associated with the security groups, in the following format: "Filters":{"Tags":["TAGKEY=TAGVALUE"]}.
      --NextPageToken string                             The token to request the next page of results.
      --ResultsPerPage int                               The maximum number of logs returned in a single response (between 1 and 1000, both included).
  -h, --help                                             help for ReadSecurityGroups
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

