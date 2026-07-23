## octl storage api ListObjectsV2



```
octl storage api ListObjectsV2 [flags]
```

### Options

```
      --Bucket string                      [REQUIRED] Directory buckets - When you use this operation with a directory bucket, you must use virtual-hosted-style requests in the format Bucket-name.s3express-zone-id.region-code.amazonaws.com .
      --ContinuationToken string           indicates to Amazon S3 that the list is being continued on this bucket with a token.
      --Delimiter string                   A delimiter is a character that you use to group keys.
      --EncodingType string                Encoding type used by Amazon S3 to encode the [object keys] in the response.
      --ExpectedBucketOwner string         The account ID of the expected bucket owner.
      --FetchOwner                         The owner field is not present in ListObjectsV2 by default.
      --MaxKeys int32                      Sets the maximum number of keys returned in the response.
      --OptionalObjectAttributes strings   Specifies the optional fields that you want returned in the response.
      --Prefix string                      Limits the response to keys that begin with the specified prefix.
      --RequestPayer string                Confirms that the requester knows that she or he will be charged for the list objects request in V2 style.
      --StartAfter string                  is where you want Amazon S3 to start listing from.
  -h, --help                               help for ListObjectsV2
```

### Options inherited from parent commands

```
  -c, --columns string             columns to display - [+]<title>:<jq query for content>||<title>:<jq query for content>
      --config string              Path of profile file (by default, ~/.osc/config.json)
      --dry-run                    Display the request payload that would be sent to the API without sending it
      --elapsed                    add elapsed time column when using --watch (default true)
      --filter strings             comma separated list of filters for results - name:value,name:value, alias for jq filter 'select(.name | tostring | test("value"))'
      --interval duration          interval between two watch/waitfor iterations (default 5s)
      --jq string                  jq filter
      --max-pages int              maximum number of pages a command can fetch (default 20)
      --no-auto-content-type       Disable automatic content-type detection
      --no-upgrade                 do not check for new versions
  -O, --out-file string            redirect output to file
  -o, --output string              output format (raw, json, yaml, table, csv, none, base64, text)
      --payload string             JSON content for query body
      --profile string             Profile to use in profile file (by default, "default")
      --single                     convert single entry lists to a single object
      --template string            JSON template file for query body
  -v, --verbose                    Verbose output
      --waitfor string             repeatedly call the API until the specified jq expression returns 1/true or a non empty result
      --waitfor-timeout duration   maximum duration of a wait (default 10m0s)
      --watch                      repeatedly call the API and display changes
  -y, --yes                        answer yes to all prompts
```

### SEE ALSO

* [octl storage api](octl_storage_api.md)	 - storage api calls

