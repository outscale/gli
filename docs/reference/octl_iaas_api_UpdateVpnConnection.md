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
      --VpnConnectionId string                                        The ID of the VPN connection you want to modify.
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

