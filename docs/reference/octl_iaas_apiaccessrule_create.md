## octl iaas apiaccessrule create

alias for api CreateApiAccessRule

### Synopsis

> *alias for api CreateApiAccessRule*

Creates a rule to allow access to the API from your OUTSCALE account.

You need to specify at least the `CaIds` or the `IpRanges` parameter.

**[NOTE]**

By default, your account has a set of rules allowing global access, that you can delete.

For more information, see [About API Access Rules](https://docs.outscale.com/en/userguide/About-API-Access-Rules.html).

```
octl iaas apiaccessrule create [flags]
```

### Options

```
      --ca-id strings        One or more IDs of Client Certificate Authorities (CAs).
      --cn strings           One or more Client Certificate Common Names (CNs).
      --description string   A description for the API access rule.
  -h, --help                 help for create
      --ip-range strings     One or more IPs or CIDR blocks (for example, 192.0.2.0/16).
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

* [octl iaas apiaccessrule](octl_iaas_apiaccessrule.md)	 - apiaccessrule commands

