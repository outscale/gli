## octl iaas publicip list

alias for api ReadPublicIps

### Synopsis

> *alias for api ReadPublicIps*

Lists one or more public IPs allocated to your OUTSCALE account.

By default, this action returns information about all your public IPs: available or associated with a virtual machine (VM), a network interface card (NIC) or a NAT service.

```
octl iaas publicip list [flags]
```

### Options

```
  -h, --help                        help for list
      --id strings                  The IDs of the public IPs.
      --link-public-ip-id strings   The IDs representing the associations of public IPs with VMs or NICs.
      --nic-account-id strings      The OUTSCALE account IDs of the owners of the NICs.
      --nic-id strings              The IDs of the NICs.
      --placement strings           Whether the public IPs are for use in the public Cloud or in a Net.
      --private-ip strings          The private IPs associated with the public IPs.
      --public-ip strings           The public IPs.
      --tag strings                 The key/value combination of the tags associated with the public IPs, in the following format: "Filters":{"Tags":["TAGKEY=TAGVALUE"]}.
      --tag-key strings             The keys of the tags associated with the public IPs.
      --tag-value strings           The values of the tags associated with the public IPs.
      --vm-id strings               The IDs of the VMs.
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

* [octl iaas publicip](octl_iaas_publicip.md)	 - publicip commands

