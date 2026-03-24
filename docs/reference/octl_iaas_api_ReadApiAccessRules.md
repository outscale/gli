## octl iaas api ReadApiAccessRules

Lists one or more API access rules.

### Synopsis

Lists one or more API access rules.

```
octl iaas api ReadApiAccessRules [flags]
```

### Options

```
      --DryRun                             If true, checks whether you have the required permissions to perform the action.
      --Filters.ApiAccessRuleIds strings   One or more IDs of API access rules.
      --Filters.CaIds strings              One or more IDs of Client Certificate Authorities (CAs).
      --Filters.Cns strings                One or more Client Certificate Common Names (CNs).
      --Filters.Descriptions strings       One or more descriptions of API access rules.
      --Filters.IpRanges strings           One or more IPs or CIDR blocks (for example, 192.0.2.0/16).
  -h, --help                               help for ReadApiAccessRules
```

### Options inherited from parent commands

```
  -c, --columns string              columns to display - [+]<title>:<jq query for content>||<title>:<jq query for content>
      --config string               Path of profile file (by default, ~/.osc/config.json)
      --filter strings              comma separated list of filters for results - name:value,name:value, alias for jq filter 'select(.name | test("value"))'
      --jq string                   jq filter
      --no-upgrade                  do not check for new versions
  -O, --out-file string             redirect output to file
  -o, --output string               output format (raw, json, yaml, table, csv, none, base64) (default "raw")
      --payload string              JSON content for query body
      --profile string              Profile to use in profile file (by default, "default")
      --single                      convert single entry lists to a single object
      --template string             JSON template file for query body
  -v, --verbose                     Verbose output
      --waitfor string              jq expression to wait for - octl will query every waitfor-interval until the expression returns 1/true or a non empty result
      --waitfor-interval duration   interval between two waitfor iterations (default 5s)
      --waitfor-timeout duration    maximum duration of a wait (default 10m0s)
  -y, --yes                         answer yes to all prompts
```

### SEE ALSO

* [octl iaas api](octl_iaas_api.md)	 - iaas api calls

