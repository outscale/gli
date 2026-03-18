## octl iaas api ReadQuotas

Lists one or more of your quotas.

### Synopsis

Lists one or more of your quotas.

For more information, see [About Your Account](https://docs.outscale.com/en/userguide/About-Your-OUTSCALE-Account.html).

```
octl iaas api ReadQuotas [flags]
```

### Options

```
      --DryRun                              If true, checks whether you have the required permissions to perform the action.
      --Filters.Collections strings         The group names of the quotas.
      --Filters.QuotaNames strings          The names of the quotas.
      --Filters.QuotaTypes strings          The resource IDs if these are resource-specific quotas, global if they are not.
      --Filters.ShortDescriptions strings   The description of the quotas.
      --NextPageToken string                The token to request the next page of results.
      --ResultsPerPage int                  The maximum number of logs returned in a single response (between 1 and 1000, both included).
  -h, --help                                help for ReadQuotas
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

* [octl iaas api](octl_iaas_api.md)	 - iaas api calls

