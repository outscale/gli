## octl iaas linkedpolicy list

alias for api ReadLinkedPolicies

### Synopsis

> *alias for api ReadLinkedPolicies*

Lists the managed policies linked to a specified user.

```
octl iaas linkedpolicy list [flags]
```

### Options

```
      --first-item int       The item starting the list of policies requested.
  -h, --help                 help for list
      --path-prefix string   The path prefix of the policies.
      --user-name string     The name of the user the policies are linked to.
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

* [octl iaas linkedpolicy](octl_iaas_linkedpolicy.md)	 - linkedpolicy commands

