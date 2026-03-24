## octl iaas api ReadEntitiesLinkedToPolicy

Lists all entities (account, users, or user groups) linked to a specific managed policy.

### Synopsis

Lists all entities (account, users, or user groups) linked to a specific managed policy.

```
octl iaas api ReadEntitiesLinkedToPolicy [flags]
```

### Options

```
      --EntitiesType strings   The type of entity linked to the policy you want to get information about.
      --FirstItem int          The item starting the list of entities requested.
      --PolicyOrn string       The OUTSCALE Resource Name (ORN) of the policy.
      --ResultsPerPage int     The maximum number of items that can be returned in a single response (by default, 100).
  -h, --help                   help for ReadEntitiesLinkedToPolicy
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

