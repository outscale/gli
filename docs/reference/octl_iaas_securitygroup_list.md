## octl iaas securitygroup list

alias for api ReadSecurityGroups

### Synopsis

> *alias for api ReadSecurityGroups*

Lists one or more security groups.

You can specify either the name of the security groups or their IDs.

```
octl iaas securitygroup list [flags]
```

### Options

```
      --description strings                         The descriptions of the security groups.
  -h, --help                                        help for list
      --id strings                                  The IDs of the security groups.
      --inbound-rule-account-id strings             The OUTSCALE account IDs that have been granted permissions.
      --inbound-rule-from-port-range ints           The beginnings of the port ranges for the TCP and UDP protocols, or the ICMP type numbers.
      --inbound-rule-ip-range strings               The IP ranges that have been granted permissions, in CIDR notation (for example, 10.0.0.0/24).
      --inbound-rule-protocol strings               The IP protocols for the permissions (tcp | udp | icmp, or a protocol number, or -1 for all protocols).
      --inbound-rule-security-group-id strings      The IDs of the security groups that have been granted permissions.
      --inbound-rule-security-group-name strings    The names of the security groups that have been granted permissions.
      --inbound-rule-to-port-range ints             The ends of the port ranges for the TCP and UDP protocols, or the ICMP code numbers.
      --name strings                                The names of the security groups.
      --net-id strings                              The IDs of the Nets specified when the security groups were created.
      --outbound-rule-account-id strings            The OUTSCALE account IDs that have been granted permissions.
      --outbound-rule-from-port-range ints          The beginnings of the port ranges for the TCP and UDP protocols, or the ICMP type numbers.
      --outbound-rule-ip-range strings              The IP ranges that have been granted permissions, in CIDR notation (for example, 10.0.0.0/24).
      --outbound-rule-protocol strings              The IP protocols for the permissions (tcp | udp | icmp, or a protocol number, or -1 for all protocols).
      --outbound-rule-security-group-id strings     The IDs of the security groups that have been granted permissions.
      --outbound-rule-security-group-name strings   The names of the security groups that have been granted permissions.
      --outbound-rule-to-port-range ints            The ends of the port ranges for the TCP and UDP protocols, or the ICMP code numbers.
      --tag strings                                 The key/value combination of the tags associated with the security groups, in the following format: "Filters":{"Tags":["TAGKEY=TAGVALUE"]}.
      --tag-key strings                             The keys of the tags associated with the security groups.
      --tag-value strings                           The values of the tags associated with the security groups.
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

* [octl iaas securitygroup](octl_iaas_securitygroup.md)	 - securitygroup commands

