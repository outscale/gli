## octl storage objectversion list

alias for api ListObjectVersions

### Synopsis

> *alias for api ListObjectVersions*



```
octl storage objectversion list [flags]
```

### Options

```
      --bucket string                       The bucket name that contains the objects.
      --delimiter string                    A delimiter is a character that you specify to group keys.
      --encoding-type string                Encoding type used by Amazon S3 to encode the [object keys] in the response.
      --expected-bucket-owner string        The account ID of the expected bucket owner.
  -h, --help                                help for list
      --id-marker string                    Specifies the object version you want to start listing from.
      --key-marker string                   Specifies the key to start with when listing objects in a bucket.
      --optional-object-attribute strings   Specifies the optional fields that you want returned in the response.
      --prefix string                       Use this parameter to select only those keys that begin with the specified prefix.
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

* [octl storage objectversion](octl_storage_objectversion.md)	 - objectversion commands

