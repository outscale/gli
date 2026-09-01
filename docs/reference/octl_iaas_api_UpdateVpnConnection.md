## octl iaas api UpdateVpnConnection

Modifies the specified attributes of a VPN connection.

### Synopsis

Modifies the specified attributes of a VPN connection.

```
octl iaas api UpdateVpnConnection [flags]
```

### Options

```
      --ClientGatewayId string                                        The ID of the client gateway.
      --DryRun                                                        If true, checks whether you have the required permissions to perform the action.
      --VirtualGatewayId string                                       The ID of the virtual gateway.
      --VpnConnectionId string                                        [REQUIRED] The ID of the VPN connection you want to modify.
      --VpnOptions.Phase1Options.DpdTimeoutAction string              This parameter is not available.
      --VpnOptions.Phase1Options.DpdTimeoutSeconds int                This parameter is not available.
      --VpnOptions.Phase1Options.IkeVersions strings                  This parameter is not available.
      --VpnOptions.Phase1Options.Phase1DhGroupNumbers ints            This parameter is not available.
      --VpnOptions.Phase1Options.Phase1EncryptionAlgorithms strings   This parameter is not available.
      --VpnOptions.Phase1Options.Phase1IntegrityAlgorithms strings    This parameter is not available.
      --VpnOptions.Phase1Options.Phase1LifetimeSeconds int            This parameter is not available.
      --VpnOptions.Phase1Options.ReplayWindowSize int                 This parameter is not available.
      --VpnOptions.Phase1Options.StartupAction string                 This parameter is not available.
      --VpnOptions.Phase2Options.Phase2DhGroupNumbers ints            This parameter is not available.
      --VpnOptions.Phase2Options.Phase2EncryptionAlgorithms strings   This parameter is not available.
      --VpnOptions.Phase2Options.Phase2IntegrityAlgorithms strings    This parameter is not available.
      --VpnOptions.Phase2Options.Phase2LifetimeSeconds int            This parameter is not available.
      --VpnOptions.Phase2Options.PreSharedKey string                  The pre-shared key to establish the initial authentication between the client gateway and the virtual gateway.
      --VpnOptions.TunnelInsideIpRange string                         The range of inside IPs for the tunnel.
  -h, --help                                                          help for UpdateVpnConnection
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
  -o, --output string              output format (raw, json, yaml, table, csv, none, text)
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

* [octl iaas api](octl_iaas_api.md)	 - Call iaas API

