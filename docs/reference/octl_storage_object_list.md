## octl storage object list

alias for api ListObjectsV2

### Synopsis

> *alias for api ListObjectsV2*



```
octl storage object list [flags]
```

### Options

```
      --bucket string                       Directory buckets - When you use this operation with a directory bucket, you must use virtual-hosted-style requests in the format Bucket-name.s3express-zone-id.region-code.amazonaws.com .
      --delimiter string                    A delimiter is a character that you use to group keys.
      --encoding-type string                Encoding type used by Amazon S3 to encode the [object keys] in the response.
      --expected-bucket-owner string        The account ID of the expected bucket owner.
      --fetch-owner                         The owner field is not present in ListObjectsV2 by default.
  -h, --help                                help for list
      --optional-object-attribute strings   Specifies the optional fields that you want returned in the response.
      --prefix string                       Limits the response to keys that begin with the specified prefix.
      --start-after string                  is where you want Amazon S3 to start listing from.
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

* [octl storage object](octl_storage_object.md)	 - object commands

