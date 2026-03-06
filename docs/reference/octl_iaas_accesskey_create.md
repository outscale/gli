## octl iaas accesskey create

alias for api CreateAccessKey

### Synopsis

> *alias for api CreateAccessKey*

Creates an access key for either your root account or an EIM user. The new key is automatically set to `ACTIVE`.

For more information, see [About Access Keys](https://docs.outscale.com/en/userguide/About-Access-Keys.html).

```
octl iaas accesskey create [flags]
```

### Options

```
      --expiration-date osctime   The date and time, or the date, at which you want the access key to expire, in ISO 8601 format (for example, 2020-06-14T00:00:00.000Z, or 2020-06-14).
  -h, --help                      help for create
      --tag string                A tag to add to the access key.
      --user-name string          The name of the EIM user that owns the key to be created.
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

