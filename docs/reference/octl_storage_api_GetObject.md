## octl storage api GetObject



```
octl storage api GetObject [flags]
```

### Options

```
      --Bucket string                       The bucket name containing the object.
      --ChecksumMode string                 To retrieve the checksum, this mode must be enabled.
      --ExpectedBucketOwner string          The account ID of the expected bucket owner.
      --IfMatch string                      Return the object only if its entity tag (ETag) is the same as the one specified in this header; otherwise, return a 412 Precondition Failed error.
      --IfModifiedSince osctime             Return the object only if it has been modified since the specified time; otherwise, return a 304 Not Modified error.
      --IfNoneMatch string                  Return the object only if its entity tag (ETag) is different from the one specified in this header; otherwise, return a 304 Not Modified error.
      --IfUnmodifiedSince osctime           Return the object only if it has not been modified since the specified time; otherwise, return a 412 Precondition Failed error.
      --Key string                          of the object to get.
      --PartNumber int32                    Part number of the object being read.
      --Range string                        Downloads the specified byte range of an object.
      --RequestPayer string                 Confirms that the requester knows that they will be charged for the request.
      --ResponseCacheControl string         Sets the Cache-Control header of the response.
      --ResponseContentDisposition string   Sets the Content-Disposition header of the response.
      --ResponseContentEncoding string      Sets the Content-Encoding header of the response.
      --ResponseContentLanguage string      Sets the Content-Language header of the response.
      --ResponseContentType string          Sets the Content-Type header of the response.
      --ResponseExpires osctime             Sets the Expires header of the response.
      --SSECustomerAlgorithm string         Specifies the algorithm to use when decrypting the object (for example, AES256 ).
      --SSECustomerKey string               Specifies the customer-provided encryption key that you originally provided for Amazon S3 to encrypt the data before storing it.
      --SSECustomerKeyMD5 string            Specifies the 128-bit MD5 digest of the customer-provided encryption key according to RFC 1321.
      --VersionId string                    Version ID used to reference a specific version of the object.
  -h, --help                                help for GetObject
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

