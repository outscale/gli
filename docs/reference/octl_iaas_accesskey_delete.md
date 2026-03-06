## octl iaas accesskey delete

alias for api DeleteAccessKey --AccessKeyId access_key_id

### Synopsis

> *alias for api DeleteAccessKey --AccessKeyId access_key_id*

Deletes the specified access key of either your root account or an EIM user.

The access key of an EIM user must be in the `INACTIVE` state to be deleted.

```
octl iaas accesskey delete access_key_id [access_key_id]... [flags]
```

### Options

```
  -h, --help               help for delete
      --user-name string   The name of the EIM user the access key you want to delete is associated with.
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

