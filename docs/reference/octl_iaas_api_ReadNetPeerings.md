## octl iaas api ReadNetPeerings

Lists one or more peering connections between two Nets.

### Synopsis

Lists one or more peering connections between two Nets.

```
octl iaas api ReadNetPeerings [flags]
```

### Options

```
      --DryRun                                  If true, checks whether you have the required permissions to perform the action.
      --Filters.AccepterNetAccountIds strings   The OUTSCALE account IDs of the owners of the peer Nets.
      --Filters.AccepterNetIpRanges strings     The IP ranges of the peer Nets, in CIDR notation (for example, 10.0.0.0/24).
      --Filters.AccepterNetNetIds strings       The IDs of the peer Nets.
      --Filters.ExpirationDates osctime         The dates and times at which the Net peerings expire, in ISO 8601 date-time format (for example, 2020-06-14T00:00:00.000Z).
      --Filters.NetPeeringIds strings           The IDs of the Net peerings.
      --Filters.SourceNetAccountIds strings     The OUTSCALE account IDs of the owners of the peer Nets.
      --Filters.SourceNetIpRanges strings       The IP ranges of the peer Nets.
      --Filters.SourceNetNetIds strings         The IDs of the peer Nets.
      --Filters.StateMessages strings           Additional information about the states of the Net peerings.
      --Filters.StateNames strings              The states of the Net peerings (pending-acceptance | active | rejected | failed | expired | deleted).
      --Filters.TagKeys strings                 The keys of the tags associated with the Net peerings.
      --Filters.TagValues strings               The values of the tags associated with the Net peerings.
      --Filters.Tags strings                    The key/value combination of the tags associated with the Net peerings, in the following format: "Filters":{"Tags":["TAGKEY=TAGVALUE"]}.
      --NextPageToken string                    The token to request the next page of results.
      --ResultsPerPage int                      The maximum number of logs returned in a single response (between 1 and 1000, both included).
  -h, --help                                    help for ReadNetPeerings
```

### Options inherited from parent commands

```
  -c, --columns string              columns to display - [+]<title>:<jq query for content>||<title>:<jq query for content>
      --config string               Path of profile file (by default, ~/.osc/config.json)
      --filter strings              comma separated list of filters for results - name:value,name:value, alias for jq filter 'select(.name | test("value"))'
      --jq string                   jq filter
      --no-upgrade                  do not check for new versions
  -O, --out-file string             redirect output to file
  -o, --output string               output format (raw, json, yaml, table, csv, none, base64) (default "raw")
      --payload string              JSON content for query body
      --profile string              Profile to use in profile file (by default, "default")
      --single                      convert single entry lists to a single object
      --template string             JSON template file for query body
  -v, --verbose                     Verbose output
      --waitfor string              jq expression to wait for - octl will query every waitfor-interval until the expression returns 1/true or a non empty result
      --waitfor-interval duration   interval between two waitfor iterations (default 5s)
      --waitfor-timeout duration    maximum duration of a wait (default 10m0s)
  -y, --yes                         answer yes to all prompts
```

### SEE ALSO

* [octl iaas api](octl_iaas_api.md)	 - iaas api calls

