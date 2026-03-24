## octl iaas tag create

alias for api CreateTags

### Synopsis

> *alias for api CreateTags*

Adds one or more tags to the specified resources.

If a tag with the same key already exists for the resource, the tag value is replaced.

You can tag the following resources using their IDs:

* Client gateways (cgw-xxxxxxxx)

* DHCP options (dopt-xxxxxxxx)

* Flexible GPU (fgpu-xxxxxxxx)

* Images (ami-xxxxxxxx)

* Internet services (igw-xxxxxxxx)

* Keypairs (key-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx)

* NAT services (nat-xxxxxxxx)

* Net endpoints (vpce-xxxxxxxx)

* Net peerings (vpcx-xxxxxxxx)

* Nets (vpc-xxxxxxxx)

* Network interface cards (NIC) (eni-xxxxxxxx)

* OMI export tasks (image-export-xxxxxxxx)

* Public IPs (eipalloc-xxxxxxxx)

* Route tables (rtb-xxxxxxxx)

* Security groups (sg-xxxxxxxx)

* Snapshot export tasks (snap-export-xxxxxxxx)
* Snapshots (snap-xxxxxxxx)

* Subnets (subnet-xxxxxxxx)

* Virtual gateways (vgw-xxxxxxxx)

* Virtual machines (VMs) (i-xxxxxxxx)

* Volumes (vol-xxxxxxxx)

* VPN connections (vpn-xxxxxxxx)

For more information, see [About Tags](https://docs.outscale.com/en/userguide/About-Tags.html).

```
octl iaas tag create [flags]
```

### Options

```
  -h, --help                  help for create
      --key string            The key of the tag, between 1 and 255 characters.
      --resource-id strings   One or more resource IDs.
      --value string          The value of the tag, between 0 and 255 characters.
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

* [octl iaas tag](octl_iaas_tag.md)	 - tag commands

