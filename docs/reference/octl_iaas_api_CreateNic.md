## octl iaas api CreateNic

Creates a network interface card (NIC) in the specified Subnet.

### Synopsis

Creates a network interface card (NIC) in the specified Subnet.

For more information, see [About NICs](https://docs.outscale.com/en/userguide/About-NICs.html).

```
octl iaas api CreateNic [flags]
```

### Options

```
      --Description string              A description for the NIC.
      --DryRun                          If true, checks whether you have the required permissions to perform the action.
      --PrivateIps.0.IsPrimary          If true, the IP is the primary private IP of the NIC.
      --PrivateIps.0.PrivateIp string   A private IP for the NIC.
      --SecurityGroupIds strings        One or more IDs of security groups for the NIC.
      --SubnetId string                 The ID of the Subnet in which you want to create the NIC.
  -h, --help                            help for CreateNic
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

