## octl storage api ListObjectVersions



```
octl storage api ListObjectVersions [flags]
```

### Options

```
      --Bucket string                      The bucket name that contains the objects.
      --Delimiter string                   A delimiter is a character that you specify to group keys.
      --EncodingType string                Encoding type used by Amazon S3 to encode the [object keys] in the response.
      --ExpectedBucketOwner string         The account ID of the expected bucket owner.
      --KeyMarker string                   Specifies the key to start with when listing objects in a bucket.
      --MaxKeys int32                      Sets the maximum number of keys returned in the response.
      --OptionalObjectAttributes strings   Specifies the optional fields that you want returned in the response.
      --Prefix string                      Use this parameter to select only those keys that begin with the specified prefix.
      --RequestPayer string                Confirms that the requester knows that they will be charged for the request.
      --VersionIdMarker string             Specifies the object version you want to start listing from.
  -h, --help                               help for ListObjectVersions
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

* [octl storage api](octl_storage_api.md)	 - storage api calls

