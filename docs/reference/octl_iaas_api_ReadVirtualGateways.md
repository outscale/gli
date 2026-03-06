## octl iaas api ReadVirtualGateways

Lists one or more virtual gateways.

### Synopsis

Lists one or more virtual gateways.

```
octl iaas api ReadVirtualGateways [flags]
```

### Options

```
      --DryRun                              If true, checks whether you have the required permissions to perform the action.
      --Filters.ConnectionTypes strings     The types of the virtual gateways (always ipsec.1).
      --Filters.LinkNetIds strings          The IDs of the Nets the virtual gateways are attached to.
      --Filters.LinkStates strings          The current states of the attachments between the virtual gateways and the Nets (attaching | attached | detaching | detached).
      --Filters.States strings              The states of the virtual gateways (pending | available | deleting | deleted).
      --Filters.TagKeys strings             The keys of the tags associated with the virtual gateways.
      --Filters.TagValues strings           The values of the tags associated with the virtual gateways.
      --Filters.Tags strings                The key/value combination of the tags associated with the virtual gateways, in the following format: "Filters":{"Tags":["TAGKEY=TAGVALUE"]}.
      --Filters.VirtualGatewayIds strings   The IDs of the virtual gateways.
      --NextPageToken string                The token to request the next page of results.
      --ResultsPerPage int                  The maximum number of logs returned in a single response (between 1 and 1000, both included).
  -h, --help                                help for ReadVirtualGateways
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

