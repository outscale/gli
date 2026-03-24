## octl iaas api CreateApiAccessRule

Creates a rule to allow access to the API from your OUTSCALE account.

### Synopsis

Creates a rule to allow access to the API from your OUTSCALE account.

You need to specify at least the `CaIds` or the `IpRanges` parameter.

**[NOTE]**

By default, your account has a set of rules allowing global access, that you can delete.

For more information, see [About API Access Rules](https://docs.outscale.com/en/userguide/About-API-Access-Rules.html).

```
octl iaas api CreateApiAccessRule [flags]
```

### Options

```
      --CaIds strings        One or more IDs of Client Certificate Authorities (CAs).
      --Cns strings          One or more Client Certificate Common Names (CNs).
      --Description string   A description for the API access rule.
      --DryRun               If true, checks whether you have the required permissions to perform the action.
      --IpRanges strings     One or more IPs or CIDR blocks (for example, 192.0.2.0/16).
  -h, --help                 help for CreateApiAccessRule
```

### Options inherited from parent commands

```
  -c, --columns string              columns to display - [+]<title>:<jq query for content>||<title>:<jq query for content>
      --config string               Path of profile file (by default, ~/.osc/config.json)
      --filter strings              comma separated list of filters for results - name:value,name:value, alias for jq filter 'select(.name | test("value"))'
      --jq string                   jq filter
      --no-upgrade                  do not check for new versions
  -O, --out-file string             redirect output to file
  -o, --output string               output format (raw, json, yaml, table, csv, none, base64, text) (default "raw")
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

