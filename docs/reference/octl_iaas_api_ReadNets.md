## octl iaas api ReadNets

Lists one or more Nets.

### Synopsis

Lists one or more Nets.

```
octl iaas api ReadNets [flags]
```

### Options

```
      --DryRun                              If true, checks whether you have the required permissions to perform the action.
      --Filters.DhcpOptionsSetIds strings   The IDs of the DHCP options sets.
      --Filters.IpRanges strings            The IP ranges for the Nets, in CIDR notation (for example, 10.0.0.0/16).
      --Filters.IsDefault                   If true, the Net used is the default one.
      --Filters.NetIds strings              The IDs of the Nets.
      --Filters.States strings              The states of the Nets (pending | available | deleting).
      --Filters.TagKeys strings             The keys of the tags associated with the Nets.
      --Filters.TagValues strings           The values of the tags associated with the Nets.
      --Filters.Tags strings                The key/value combination of the tags associated with the Nets, in the following format: "Filters":{"Tags":["TAGKEY=TAGVALUE"]}.
      --NextPageToken string                The token to request the next page of results.
      --ResultsPerPage int                  The maximum number of logs returned in a single response (between 1 and 1000, both included).
  -h, --help                                help for ReadNets
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

