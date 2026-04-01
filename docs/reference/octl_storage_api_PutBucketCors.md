## octl storage api PutBucketCors



```
octl storage api PutBucketCors [flags]
```

### Options

```
      --Bucket string                                          Specifies the bucket impacted by the cors configuration.
      --CORSConfiguration.CORSRules.0.AllowedHeaders strings   Headers that are specified in the Access-Control-Request-Headers header.
      --CORSConfiguration.CORSRules.0.AllowedMethods strings   An HTTP method that you allow the origin to execute.
      --CORSConfiguration.CORSRules.0.AllowedOrigins strings   One or more origins you want customers to be able to access the bucket from.
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

