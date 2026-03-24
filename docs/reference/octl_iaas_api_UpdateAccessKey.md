## octl iaas api UpdateAccessKey

Modifies the attributes of the specified access key of either the root user or an EIM user.

### Synopsis

Modifies the attributes of the specified access key of either the root user or an EIM user.

The parameter `ExpirationDate` is not required when updating the state of your access key. However, if you do not specify the expiration date of an access key when updating its state, it is then set to not expire.

```
octl iaas api UpdateAccessKey [flags]
```

### Options

```
      --AccessKeyId string       The ID of the access key.
      --ClearExpirationDate      If true, the current expiration date is deleted and the access key is set to not expire.
      --ClearTag                 If true, the current tag of the access key is deleted.
      --DryRun                   If true, checks whether you have the required permissions to perform the action.
      --ExpirationDate osctime   The date and time, or the date, at which you want the access key to expire, in ISO 8601 format (for example, 2020-06-14T00:00:00.000Z or 2020-06-14).
      --State string             The new state for the access key (ACTIVE | INACTIVE).
      --Tag string               A new tag to add to the access key.
      --UserName string          The name of the EIM user that the access key you want to modify is associated with.
  -h, --help                     help for UpdateAccessKey
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

