## octl iaas vpnconnection update

alias for api UpdateVpnConnection --VpnConnectionId vpn_connection_id

### Synopsis

> alias for api UpdateVpnConnection --VpnConnectionId vpn_connection_id

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
  -c, --columns string             columns to display - [+]<title>:<jq query for content>||<title>:<jq query for content>
      --config string              Path of profile file (by default, ~/.osc/config.json)
      --dry-run                    Display the request payload that would be sent to the API without sending it
      --elapsed                    add elapsed time column when using --watch (default true)
      --filter strings             comma separated list of filters for results - name:value,name:value, alias for jq filter 'select(.name | tostring | test("value"))'
      --interval duration          interval between two watch/waitfor iterations (default 5s)
      --jq string                  jq filter
      --max-pages int              maximum number of pages a command can fetch (default 20)
      --no-upgrade                 do not check for new versions
  -O, --out-file string            redirect output to file
  -o, --output string              output format (raw, json, yaml, table, csv, none, base64, text)
      --payload string             JSON content for query body
      --profile string             Profile to use in profile file (by default, "default")
  -s, --silent                     Hides all information messages
      --single                     convert single entry lists to a single object
      --template string            JSON template file for query body
  -v, --verbose                    Verbose output
      --waitfor string             repeatedly call the API until the specified jq expression returns 1/true or a non empty result
      --waitfor-timeout duration   maximum duration of a wait (default 10m0s)
      --watch                      repeatedly call the API and display changes
  -y, --yes                        answer yes to all prompts
```

### SEE ALSO

* [octl iaas vpnconnection](octl_iaas_vpnconnection.md)	 - vpnconnection commands

