## octl iaas api DeleteApiAccessRule

Deletes a specified API access rule.

### Synopsis

Deletes a specified API access rule.

**[IMPORTANT]**

You cannot delete the last remaining API access rule. However, if you delete all the API access rules that allow you to access the APIs, you need to contact the Support team to regain access. For more information, see [Technical Support](https://docs.outscale.com/en/userguide/Technical-Support.html).

```
octl iaas api DeleteApiAccessRule [flags]
```

### Options

```
      --ApiAccessRuleId string   The ID of the API access rule you want to delete.
      --DryRun                   If true, checks whether you have the required permissions to perform the action.
  -h, --help                     help for DeleteApiAccessRule
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

