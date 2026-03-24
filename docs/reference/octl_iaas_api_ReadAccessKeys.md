## octl iaas api ReadAccessKeys

Lists the access key IDs of either your root user or an EIM user.

### Synopsis

Lists the access key IDs of either your root user or an EIM user.

```
octl iaas api ReadAccessKeys [flags]
```

### Options

```
      --DryRun                         If true, checks whether you have the required permissions to perform the action.
      --Filters.AccessKeyIds strings   The IDs of the access keys.
      --Filters.States strings         The states of the access keys (ACTIVE | INACTIVE).
      --Tag string                     The tag added to the access key.
      --UserName string                The name of the EIM user.
  -h, --help                           help for ReadAccessKeys
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

