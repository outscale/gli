## octl storage

OUTSCALE Object Storage (OOS) management

### Options

```
  -h, --help   help for storage
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

* [octl](octl.md)	 - A modern CLI for Outscale services
* [octl storage api](octl_storage_api.md)	 - storage api calls
* [octl storage bucket](octl_storage_bucket.md)	 - bucket commands
* [octl storage multipartupload](octl_storage_multipartupload.md)	 - multipartupload commands
* [octl storage object](octl_storage_object.md)	 - object commands
* [octl storage objectversion](octl_storage_objectversion.md)	 - objectversion commands
* [octl storage part](octl_storage_part.md)	 - part commands

