## octl iaas api UpdateApiAccessRule

Modifies a specified API access rule.

### Synopsis

Modifies a specified API access rule.

**[WARNING]**

- The new rule you specify fully replaces the old rule. Therefore, for a parameter that is not specified, any previously set value is deleted.

- If, as result of your modification, you no longer have access to the APIs, you will need to contact the Support team to regain access. For more information, see [Technical Support](https://docs.outscale.com/en/userguide/Technical-Support.html).

```
octl iaas api UpdateApiAccessRule [flags]
```

### Options

```
      --ApiAccessRuleId string   The ID of the API access rule you want to update.
      --CaIds strings            One or more IDs of Client Certificate Authorities (CAs).
      --Cns strings              One or more Client Certificate Common Names (CNs).
      --Description string       A new description for the API access rule.
      --DryRun                   If true, checks whether you have the required permissions to perform the action.
      --IpRanges strings         One or more IPs or CIDR blocks (for example, 192.0.2.0/16).
  -h, --help                     help for UpdateApiAccessRule
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

