## octl iaas keypair delete

alias for api DeleteKeypair --KeypairId keypair_id

### Synopsis

> *alias for api DeleteKeypair --KeypairId keypair_id*

Deletes the specified keypair.

This action deletes the public key stored by 3DS OUTSCALE, thus deleting the keypair.

```
octl iaas keypair delete keypair_id [keypair_id]... [flags]
```

### Options

```
  -h, --help          help for delete
      --name string   The name of the keypair you want to delete.
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

* [octl iaas keypair](octl_iaas_keypair.md)	 - keypair commands

