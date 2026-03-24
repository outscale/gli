## octl iaas netpeering

netpeering commands

### Options

```
  -h, --help   help for netpeering
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

* [octl iaas](octl_iaas.md)	 - OUTSCALE IaaS management
* [octl iaas netpeering accept](octl_iaas_netpeering_accept.md)	 - alias for api AcceptNetPeering --NetPeeringId netpeering_id
* [octl iaas netpeering create](octl_iaas_netpeering_create.md)	 - alias for api CreateNetPeering
* [octl iaas netpeering delete](octl_iaas_netpeering_delete.md)	 - alias for api DeleteNetPeering --NetPeeringId net_peering_id
* [octl iaas netpeering describe](octl_iaas_netpeering_describe.md)	 - alias for api ReadNetPeerings --Filters.NetPeeringIds net_peering_id
* [octl iaas netpeering list](octl_iaas_netpeering_list.md)	 - alias for api ReadNetPeerings
* [octl iaas netpeering reject](octl_iaas_netpeering_reject.md)	 - alias for api RejectNetPeering --NetPeeringId netpeering_id

