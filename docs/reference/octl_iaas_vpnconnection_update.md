## octl iaas vpnconnection update

alias for api UpdateVpnConnection --VpnConnectionId vpn_connection_id

### Synopsis

> *alias for api UpdateVpnConnection --VpnConnectionId vpn_connection_id*

Modifies the specified attributes of a VPN connection.

```
octl iaas vpnconnection update vpn_connection_id [vpn_connection_id]... [flags]
```

### Options

```
      --client-gateway-id string                                           The ID of the client gateway.
  -h, --help                                                               help for update
      --virtual-gateway-id string                                          The ID of the virtual gateway.
      --vpn-options-phase-1-options-dpd-timeout-action string              This parameter is not available.
      --vpn-options-phase-1-options-dpd-timeout-second int                 This parameter is not available.
      --vpn-options-phase-1-options-ike-version strings                    This parameter is not available.
      --vpn-options-phase-1-options-phase-1-dh-group-number ints           This parameter is not available.
      --vpn-options-phase-1-options-phase-1-encryption-algorithm strings   This parameter is not available.
      --vpn-options-phase-1-options-phase-1-integrity-algorithm strings    This parameter is not available.
      --vpn-options-phase-1-options-phase-1-lifetime-second int            This parameter is not available.
      --vpn-options-phase-1-options-replay-window-size int                 This parameter is not available.
      --vpn-options-phase-1-options-startup-action string                  This parameter is not available.
      --vpn-options-phase-2-options-phase-2-dh-group-number ints           This parameter is not available.
      --vpn-options-phase-2-options-phase-2-encryption-algorithm strings   This parameter is not available.
      --vpn-options-phase-2-options-phase-2-integrity-algorithm strings    This parameter is not available.
      --vpn-options-phase-2-options-phase-2-lifetime-second int            This parameter is not available.
      --vpn-options-phase-2-options-pre-shared-key string                  The pre-shared key to establish the initial authentication between the client gateway and the virtual gateway.
      --vpn-options-tunnel-inside-ip-range string                          The range of inside IPs for the tunnel.
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

