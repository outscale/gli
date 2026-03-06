## octl iaas quota list

alias for api ReadQuotas

### Synopsis

> *alias for api ReadQuotas*

Lists one or more of your quotas.

For more information, see [About Your Account](https://docs.outscale.com/en/userguide/About-Your-Account.html).

```
octl iaas quota list [flags]
```

### Options

```
      --collection strings          The group names of the quotas.
  -h, --help                        help for list
      --name strings                The names of the quotas.
      --short-description strings   The description of the quotas.
      --type strings                The resource IDs if these are resource-specific quotas, global if they are not.
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

* [octl iaas quota](octl_iaas_quota.md)	 - quota commands

