## octl storage multipartupload list

alias for api ListMultipartUploads

### Synopsis

> *alias for api ListMultipartUploads*



```
octl storage multipartupload list [flags]
```

### Options

```
      --bucket string                  The name of the bucket to which the multipart upload was initiated.
      --delimiter string               Character you use to group keys.
      --encoding-type string           Encoding type used by Amazon S3 to encode the [object keys] in the response.
      --expected-bucket-owner string   The account ID of the expected bucket owner.
  -h, --help                           help for list
      --key-marker string              Specifies the multipart upload after which listing should begin.
      --max-upload int32               Sets the maximum number of multipart uploads, from 1 to 1,000, to return in the response body.
      --prefix string                  Lists in-progress uploads only for those keys that begin with the specified prefix.
      --upload-id-marker string        Together with key-marker, specifies the multipart upload after which listing should begin.
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

* [octl storage multipartupload](octl_storage_multipartupload.md)	 - multipartupload commands

