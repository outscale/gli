## octl iaas netpeering delete

alias for api DeleteNetPeering --NetPeeringId net_peering_id

### Synopsis

> *alias for api DeleteNetPeering --NetPeeringId net_peering_id*

Deletes a Net peering.

If the Net peering is in the `active` state, it can be deleted either by the owner of the requester Net or the owner of the peer Net.

If it is in the `pending-acceptance` state, it can be deleted only by the owner of the requester Net.

If it is in the `rejected`, `failed`, or `expired` states, it cannot be deleted.

```
octl iaas netpeering delete net_peering_id [net_peering_id]... [flags]
```

### Options

```
  -h, --help   help for delete
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

* [octl iaas netpeering](octl_iaas_netpeering.md)	 - netpeering commands

