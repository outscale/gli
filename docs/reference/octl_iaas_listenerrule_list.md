## octl iaas listenerrule list

alias for api ReadListenerRules

### Synopsis

> *alias for api ReadListenerRules*

Lists one or more listener rules. By default, this action returns the full list of listener rules for the OUTSCALE account.

```
octl iaas listenerrule list [flags]
```

### Options

```
  -h, --help           help for list
      --name strings   The names of the listener rules.
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
      --payload string    JSON content for query body
      --profile string    Profile to use in profile file (by default, "default")
      --single            convert single entry lists to a single object
      --template string   JSON template file for query body
  -v, --verbose           Verbose output
  -y, --yes               answer yes to all prompts
```

### SEE ALSO

* [octl iaas listenerrule](octl_iaas_listenerrule.md)	 - listenerrule commands

