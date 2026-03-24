## octl iaas api LinkPrivateIps

Assigns one or more secondary private IPs to a specified network interface card (NIC).

### Synopsis

Assigns one or more secondary private IPs to a specified network interface card (NIC). This action is only available in a Net. The private IPs to be assigned can be added individually using the `PrivateIps` parameter, or you can specify the number of private IPs to be automatically chosen within the Subnet range using the `SecondaryPrivateIpCount` parameter. You can specify only one of these two parameters. If none of these parameters are specified, a private IP is chosen within the Subnet range.

```
octl iaas api LinkPrivateIps [flags]
```

### Options

```
      --AllowRelink                   If true, allows an IP that is already assigned to another NIC in the same Subnet to be assigned to the NIC you specified.
      --DryRun                        If true, checks whether you have the required permissions to perform the action.
      --NicId string                  The ID of the NIC.
      --PrivateIps strings            The secondary private IP or IPs you want to assign to the NIC within the IP range of the Subnet.
      --SecondaryPrivateIpCount int   The number of secondary private IPs to assign to the NIC.
  -h, --help                          help for LinkPrivateIps
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

