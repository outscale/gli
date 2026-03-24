## octl iaas policy list

alias for api ReadPolicies

### Synopsis

> *alias for api ReadPolicies*

Lists all the managed policies available for your OUTSCALE account.

```
octl iaas policy list [flags]
```

### Options

```
      --first-item int       The item starting the list of policies requested.
  -h, --help                 help for list
      --only-linked          If set to true, lists only the policies attached to a user.
      --path-prefix string   The path prefix you can use to filter the results.
      --scope string         The scope of the policies.
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

* [octl iaas policy](octl_iaas_policy.md)	 - policy commands

