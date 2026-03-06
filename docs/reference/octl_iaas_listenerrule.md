## octl iaas listenerrule

listenerrule commands

### Options

```
  -h, --help   help for listenerrule
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

* [octl iaas](octl_iaas.md)	 - OUTSCALE IaaS management
* [octl iaas listenerrule create](octl_iaas_listenerrule_create.md)	 - alias for api CreateListenerRule
* [octl iaas listenerrule delete](octl_iaas_listenerrule_delete.md)	 - alias for api DeleteListenerRule --ListenerRuleName listener_rule_name
* [octl iaas listenerrule describe](octl_iaas_listenerrule_describe.md)	 - alias for api ReadListenerRules --Filters.ListenerRuleNames listener_rule_name
* [octl iaas listenerrule list](octl_iaas_listenerrule_list.md)	 - alias for api ReadListenerRules
* [octl iaas listenerrule update](octl_iaas_listenerrule_update.md)	 - alias for api UpdateListenerRule --ListenerRuleName listener_rule_name

