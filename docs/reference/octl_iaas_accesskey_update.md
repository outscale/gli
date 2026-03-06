## octl iaas accesskey update

alias for api UpdateAccessKey --AccessKeyId access_key_id

### Synopsis

> *alias for api UpdateAccessKey --AccessKeyId access_key_id*

Modifies the attributes of the specified access key of either your root account or an EIM user.

The parameter `ExpirationDate` is not required when updating the state of your access key. However, if you do not specify the expiration date of an access key when updating its state, it is then set to not expire.

```
octl iaas accesskey update access_key_id [access_key_id]... [flags]
```

### Options

```
      --expiration-date osctime   The date and time, or the date, at which you want the access key to expire, in ISO 8601 format (for example, 2020-06-14T00:00:00.000Z or 2020-06-14).
  -h, --help                      help for update
      --state string              The new state for the access key (ACTIVE | INACTIVE).
      --user-name string          The name of the EIM user that the access key you want to modify is associated with.
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

* [octl iaas accesskey](octl_iaas_accesskey.md)	 - accesskey commands

