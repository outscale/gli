## octl iaas vpnconnection delete

alias for api DeleteVpnConnection --VpnConnectionId vpn_connection_id

### Synopsis

> *alias for api DeleteVpnConnection --VpnConnectionId vpn_connection_id*

Deletes a specified VPN connection.

If you want to delete a Net and all its dependencies, we recommend to detach the virtual gateway from the Net and delete the Net before deleting the VPN connection. This enables you to delete the Net without waiting for the VPN connection to be deleted.

```
octl iaas vpnconnection delete vpn_connection_id [vpn_connection_id]... [flags]
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

* [octl iaas vpnconnection](octl_iaas_vpnconnection.md)	 - vpnconnection commands

