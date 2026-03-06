## octl iaas virtualgateway list

alias for api ReadVirtualGateways

### Synopsis

> *alias for api ReadVirtualGateways*

Lists one or more virtual gateways.

```
octl iaas virtualgateway list [flags]
```

### Options

```
      --connection-type strings   The types of the virtual gateways (always ipsec.1).
  -h, --help                      help for list
      --id strings                The IDs of the virtual gateways.
      --link-net-id strings       The IDs of the Nets the virtual gateways are attached to.
      --link-state strings        The current states of the attachments between the virtual gateways and the Nets (attaching | attached | detaching | detached).
      --state strings             The states of the virtual gateways (pending | available | deleting | deleted).
      --tag strings               The key/value combination of the tags associated with the virtual gateways, in the following format: "Filters":{"Tags":["TAGKEY=TAGVALUE"]}.
      --tag-key strings           The keys of the tags associated with the virtual gateways.
      --tag-value strings         The values of the tags associated with the virtual gateways.
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

* [octl iaas virtualgateway](octl_iaas_virtualgateway.md)	 - virtualgateway commands

