## octl iaas api CreateNetPeering

Requests a Net peering between a Net you own and a peer Net that belongs to you or another account.

### Synopsis

Requests a Net peering between a Net you own and a peer Net that belongs to you or another account.

This action creates a Net peering that remains in the `pending-acceptance` state until it is accepted by the owner of the peer Net. If the owner of the peer Net does not accept the request within 7 days, the state of the Net peering becomes `expired`. For more information, see [AcceptNetPeering](#acceptnetpeering).

**[IMPORTANT]**

* The two Nets must not have overlapping IP ranges. Otherwise, the Net peering is in the `failed` state.

* A peering connection between two Nets works both ways. If an A-to-B connection is already created and accepted, creating a B-to-A connection is not necessary and would be automatically rejected.

For more information, see [About Net Peerings](https://docs.outscale.com/en/userguide/About-Net-Peerings.html).

```
octl iaas api CreateNetPeering [flags]
```

### Options

```
      --AccepterNetId string     The ID of the Net you want to connect with.
      --AccepterOwnerId string   The account ID of the owner of the Net you want to connect with.
      --DryRun                   If true, checks whether you have the required permissions to perform the action.
      --SourceNetId string       The ID of the Net you send the peering request from.
  -h, --help                     help for CreateNetPeering
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

