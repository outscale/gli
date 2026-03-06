## octl iaas api CreateVpnConnectionRoute

Creates a static route to a VPN connection.

### Synopsis

Creates a static route to a VPN connection.

This enables you to select the network flows sent by the virtual gateway to the target VPN connection.

For more information, see [About Routing Configuration for VPN Connections](https://docs.outscale.com/en/userguide/About-Routing-Configuration-for-VPN-Connections.html).

```
octl iaas api CreateVpnConnectionRoute [flags]
```

### Options

```
      --DestinationIpRange string   The network prefix of the route, in CIDR notation (for example, 10.12.0.0/16).
      --DryRun                      If true, checks whether you have the required permissions to perform the action.
      --VpnConnectionId string      The ID of the target VPN connection of the static route.
  -h, --help                        help for CreateVpnConnectionRoute
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

