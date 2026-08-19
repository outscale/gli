## octl iaas clientgateway list

alias for api ReadClientGateways

### Synopsis

> *alias for api ReadClientGateways*

Lists one or more of your client gateways.

```
octl iaas clientgateway list [flags]
```

### Options

```
      --bgp-asn ints              The Border Gateway Protocol (BGP) Autonomous System Numbers (ASNs) of the connections.
      --connection-type strings   The types of communication tunnels used by the client gateways (always ipsec.1).
  -h, --help                      help for list
      --id strings                The IDs of the client gateways.
      --public-ip strings         The public IPv4 addresses of the client gateways.
      --state strings             The states of the client gateways (pending | available | deleting | deleted).
      --tag strings               The key/value combination of the tags associated with the client gateways, in the following format: "Filters":{"Tags":["TAGKEY=TAGVALUE"]}.
      --tag-key strings           The keys of the tags associated with the client gateways.
      --tag-value strings         The values of the tags associated with the client gateways.
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

* [octl iaas clientgateway](octl_iaas_clientgateway.md)	 - clientgateway commands

