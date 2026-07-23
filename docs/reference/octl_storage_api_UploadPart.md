## octl storage api UploadPart



```
octl storage api UploadPart [flags]
```

### Options

```
      --Body streamedFile             Object data.
      --Bucket string                 [REQUIRED] The name of the bucket to which the multipart upload was initiated.
      --ChecksumAlgorithm string      Indicates the algorithm used to create the checksum for the object when you use the SDK.
      --ChecksumCRC32 string          This header can be used as a data integrity check to verify that the data received is the same data that was originally sent.
      --ChecksumCRC32C string         This header can be used as a data integrity check to verify that the data received is the same data that was originally sent.
      --ChecksumSHA1 string           This header can be used as a data integrity check to verify that the data received is the same data that was originally sent.
      --ChecksumSHA256 string         This header can be used as a data integrity check to verify that the data received is the same data that was originally sent.
      --ContentLength int             Size of the body in bytes.
      --ContentMD5 string             The base64-encoded 128-bit MD5 digest of the part data.
      --ExpectedBucketOwner string    The account ID of the expected bucket owner.
      --Key string                    [REQUIRED] Object key for which the multipart upload was initiated.
      --PartNumber int32              [REQUIRED] Part number of part being uploaded.
      --RequestPayer string           Confirms that the requester knows that they will be charged for the request.
      --SSECustomerAlgorithm string   Specifies the algorithm to use when encrypting the object (for example, AES256).
      --SSECustomerKey string         Specifies the customer-provided encryption key for Amazon S3 to use in encrypting data.
      --SSECustomerKeyMD5 string      Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321.
      --UploadId string               [REQUIRED] Upload ID identifying the multipart upload whose part is being uploaded.
  -h, --help                          help for UploadPart
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

