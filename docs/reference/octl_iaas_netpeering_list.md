## octl iaas netpeering list

alias for api ReadNetPeerings

### Synopsis

> alias for api ReadNetPeerings

Lists one or more peering connections between two Nets.

```
octl iaas netpeering list [flags]
```

### Options

```
      --accepter-net-account-id strings   The OUTSCALE account IDs of the owners of the peer Nets.
      --accepter-net-ip-range strings     The IP ranges of the peer Nets, in CIDR notation (for example, 10.0.0.0/24).
      --accepter-net-net-id strings       The IDs of the peer Nets.
      --expiration-date osctime           The dates and times at which the Net peerings expire, in ISO 8601 date-time format (for example, 2020-06-14T00:00:00.000Z).
  -h, --help                              help for list
      --id strings                        The IDs of the Net peerings.
      --source-net-account-id strings     The OUTSCALE account IDs of the owners of the peer Nets.
      --source-net-ip-range strings       The IP ranges of the peer Nets.
      --source-net-net-id strings         The IDs of the peer Nets.
      --state-message strings             Additional information about the states of the Net peerings.
      --state-name strings                The states of the Net peerings (pending-acceptance | active | rejected | failed | expired | deleted).
      --tag strings                       The key/value combination of the tags associated with the Net peerings, in the following format: "Filters":{"Tags":["TAGKEY=TAGVALUE"]}.
      --tag-key strings                   The keys of the tags associated with the Net peerings.
      --tag-value strings                 The values of the tags associated with the Net peerings.
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

* [octl iaas netpeering](octl_iaas_netpeering.md)	 - netpeering commands

