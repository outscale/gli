## octl storage part list

alias for api ListParts

### Synopsis

> *alias for api ListParts*



```
octl storage part list [flags]
```

### Options

```
      --bucket string                  The name of the bucket to which the parts are being uploaded.
      --expected-bucket-owner string   The account ID of the expected bucket owner.
  -h, --help                           help for list
      --key string                     Object key for which the multipart upload was initiated.
      --max-part int32                 Sets the maximum number of parts to return.
      --number-marker string           Specifies the part after which listing should begin.
      --upload-id string               Upload ID identifying the multipart upload whose parts are being listed.
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

* [octl storage part](octl_storage_part.md)	 - part commands

