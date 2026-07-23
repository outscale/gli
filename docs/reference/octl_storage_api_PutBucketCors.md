## octl storage api PutBucketCors



```
octl storage api PutBucketCors [flags]
```

### Options

```
      --Bucket string                                          [REQUIRED] Specifies the bucket impacted by the cors configuration.
      --CORSConfiguration.CORSRules.0.AllowedHeaders strings   Headers that are specified in the Access-Control-Request-Headers header.
      --CORSConfiguration.CORSRules.0.AllowedMethods strings   [REQUIRED] An HTTP method that you allow the origin to execute.
      --CORSConfiguration.CORSRules.0.AllowedOrigins strings   [REQUIRED] One or more origins you want customers to be able to access the bucket from.
      --CORSConfiguration.CORSRules.0.ExposeHeaders strings    One or more headers in the response that you want customers to be able to access from their applications (for example, from a JavaScript XMLHttpRequest object).
      --CORSConfiguration.CORSRules.0.ID string                Unique identifier for the rule.
      --CORSConfiguration.CORSRules.0.MaxAgeSeconds int32      The time in seconds that your browser is to cache the preflight response for the specified resource.
      --ChecksumAlgorithm string                               Indicates the algorithm used to create the checksum for the object when you use the SDK.
      --ContentMD5 string                                      The base64-encoded 128-bit MD5 digest of the data.
      --ExpectedBucketOwner string                             The account ID of the expected bucket owner.
  -h, --help                                                   help for PutBucketCors
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

