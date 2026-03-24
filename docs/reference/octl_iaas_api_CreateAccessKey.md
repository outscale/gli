## octl iaas api CreateAccessKey

Creates an access key for either the root user or an EIM user.

### Synopsis

Creates an access key for either the root user or an EIM user. The new key is automatically set to `ACTIVE`.

For more information, see [About Access Keys](https://docs.outscale.com/en/userguide/About-Access-Keys.html).

```
octl iaas api CreateAccessKey [flags]
```

### Options

```
      --DryRun                   If true, checks whether you have the required permissions to perform the action.
      --ExpirationDate osctime   The date and time, or the date, at which you want the access key to expire, in ISO 8601 format (for example, 2020-06-14T00:00:00.000Z, or 2020-06-14).
      --Tag string               A tag to add to the access key.
      --UserName string          The name of the EIM user that owns the key to be created.
  -h, --help                     help for CreateAccessKey
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

