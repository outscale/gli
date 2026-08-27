## octl iaas api ReadVpnConnections

Lists one or more VPN connections.

### Synopsis

Lists one or more VPN connections.

```
octl iaas api ReadVpnConnections [flags]
```

### Options

```
      --DryRun                                     If true, checks whether you have the required permissions to perform the action.
      --Filters.BgpAsns ints                       The Border Gateway Protocol (BGP) Autonomous System Numbers (ASNs) of the connections.
      --Filters.ClientGatewayIds strings           The IDs of the client gateways.
      --Filters.ConnectionTypes strings            The types of the VPN connections (always ipsec.1).
      --Filters.RouteDestinationIpRanges strings   The destination IP ranges.
      --Filters.States strings                     The states of the VPN connections (pending | available | deleting | deleted).
      --Filters.StaticRoutesOnly                   If false, the VPN connection uses dynamic routing with Border Gateway Protocol (BGP).
      --Filters.TagKeys strings                    The keys of the tags associated with the VPN connections.
      --Filters.TagValues strings                  The values of the tags associated with the VPN connections.
      --Filters.Tags strings                       The key/value combination of the tags associated with the VPN connections, in the following format: "Filters":{"Tags":["TAGKEY=TAGVALUE"]}.
      --Filters.VirtualGatewayIds strings          The IDs of the virtual gateways.
      --Filters.VpnConnectionIds strings           The IDs of the VPN connections.
      --NextPageToken string                       The token to request the next page of results.
      --ResultsPerPage int                         The maximum number of logs returned in a single response (between 1 and 1000, both included).
  -h, --help                                       help for ReadVpnConnections
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

* [octl iaas api](octl_iaas_api.md)	 - iaas api calls

