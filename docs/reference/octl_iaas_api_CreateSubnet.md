## octl iaas api CreateSubnet

Creates a Subnet in an existing Net.

### Synopsis

Creates a Subnet in an existing Net.

To create a Subnet in a Net, you have to provide the ID of the Net and the IP range for the Subnet (its network range). Once the Subnet is created, you cannot modify its IP range.

For more information, see [About Nets](https://docs.outscale.com/en/userguide/About-Nets.html).

```
octl iaas api CreateSubnet [flags]
```

### Options

```
      --DryRun                 If true, checks whether you have the required permissions to perform the action.
      --IpRange string         The IP range in the Subnet, in CIDR notation (for example, 10.0.0.0/16).
      --NetId string           The ID of the Net for which you want to create a Subnet.
      --SubregionName string   The name of the Subregion in which you want to create the Subnet.
  -h, --help                   help for CreateSubnet
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

