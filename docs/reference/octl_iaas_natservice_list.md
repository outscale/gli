## octl iaas natservice list

alias for api ReadNatServices

### Synopsis

> *alias for api ReadNatServices*

Lists one or more network address translation (NAT) services.

```
octl iaas natservice list [flags]
```

### Options

```
      --client-token strings   The idempotency tokens provided when creating the NAT services.
  -h, --help                   help for list
      --id strings             The IDs of the NAT services.
      --net-id strings         The IDs of the Nets in which the NAT services are.
      --state strings          The states of the NAT services (pending | available | deleting | deleted).
      --subnet-id strings      The IDs of the Subnets in which the NAT services are.
      --tag strings            The key/value combination of the tags associated with the NAT services, in the following format: "Filters":{"Tags":["TAGKEY=TAGVALUE"]}.
      --tag-key strings        The keys of the tags associated with the NAT services.
      --tag-value strings      The values of the tags associated with the NAT services.
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

* [octl iaas natservice](octl_iaas_natservice.md)	 - natservice commands

