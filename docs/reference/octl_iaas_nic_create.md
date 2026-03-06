## octl iaas nic create

alias for api CreateNic

### Synopsis

> *alias for api CreateNic*

Creates a network interface card (NIC) in the specified Subnet.

For more information, see [About NICs](https://docs.outscale.com/en/userguide/About-NICs.html).

```
octl iaas nic create [flags]
```

### Options

```
      --description string             A description for the NIC.
  -h, --help                           help for create
      --private-ip-is-primary          If true, the IP is the primary private IP of the NIC.
      --private-ip-private-ip string   The private IP of the NIC.
      --security-group-id strings      One or more IDs of security groups for the NIC.
      --subnet-id string               The ID of the Subnet in which you want to create the NIC.
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

* [octl iaas nic](octl_iaas_nic.md)	 - nic commands

