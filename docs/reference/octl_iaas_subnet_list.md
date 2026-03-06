## octl iaas subnet list

alias for api ReadSubnets

### Synopsis

> *alias for api ReadSubnets*

Lists one or more of your Subnets.

If you do not specify any Subnet ID, this action describes all of your Subnets.

```
octl iaas subnet list [flags]
```

### Options

```
      --available-ip-count ints   The number of available IPs.
  -h, --help                      help for list
      --id strings                The IDs of the Subnets.
      --ip-range strings          The IP ranges in the Subnets, in CIDR notation (for example, 10.0.0.0/16).
      --net-id strings            The IDs of the Nets in which the Subnets are.
      --state strings             The states of the Subnets (pending | available | deleted).
      --subregion strings         The names of the Subregions in which the Subnets are located.
      --tag strings               The key/value combination of the tags associated with the Subnets, in the following format: "Filters":{"Tags":["TAGKEY=TAGVALUE"]}.
      --tag-key strings           The keys of the tags associated with the Subnets.
      --tag-value strings         The values of the tags associated with the Subnets.
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

* [octl iaas subnet](octl_iaas_subnet.md)	 - subnet commands

