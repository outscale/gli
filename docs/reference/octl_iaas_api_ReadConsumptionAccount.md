## octl iaas api ReadConsumptionAccount

Gets information about the consumption of your OUTSCALE account for each billable resource within the specified time period.

### Synopsis

Gets information about the consumption of your OUTSCALE account for each billable resource within the specified time period.

```
octl iaas api ReadConsumptionAccount [flags]
```

### Options

```
      --DryRun                If true, checks whether you have the required permissions to perform the action.
      --FromDate osctime      The beginning of the time period, in ISO 8601 date format (for example, 2020-06-14).
      --Overall               If false, returns only the consumption of the specific account that sends this request.
      --ShowPrice             If true, the response also includes the unit price of the consumed resource (UnitPrice) and the total price of the consumed resource during the specified time period (Price), in the currency of the Region's catalog.
      --ShowResourceDetails   By default or if false, returns the consumption aggregated by resource type.
      --ToDate osctime        The end of the time period, in ISO 8601 date format (for example, 2020-06-30).
  -h, --help                  help for ReadConsumptionAccount
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

