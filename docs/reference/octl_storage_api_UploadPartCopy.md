## octl storage api UploadPartCopy



```
octl storage api UploadPartCopy [flags]
```

### Options

```
      --Bucket string                           [REQUIRED] The bucket name.
      --CopySource string                       [REQUIRED] Specifies the source object for the copy operation.
      --CopySourceIfMatch string                Copies the object if its entity tag (ETag) matches the specified tag.
      --CopySourceIfModifiedSince osctime       Copies the object if it has been modified since the specified time.
      --CopySourceIfNoneMatch string            Copies the object if its entity tag (ETag) is different than the specified ETag.
      --CopySourceIfUnmodifiedSince osctime     Copies the object if it hasn't been modified since the specified time.
      --CopySourceRange string                  The range of bytes to copy from the source object.
      --CopySourceSSECustomerAlgorithm string   Specifies the algorithm to use when decrypting the source object (for example, AES256 ).
      --CopySourceSSECustomerKey string         Specifies the customer-provided encryption key for Amazon S3 to use to decrypt the source object.
      --CopySourceSSECustomerKeyMD5 string      Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321.
      --ExpectedBucketOwner string              The account ID of the expected destination bucket owner.
      --ExpectedSourceBucketOwner string        The account ID of the expected source bucket owner.
      --Key string                              [REQUIRED] Object key for which the multipart upload was initiated.
      --PartNumber int32                        [REQUIRED] Part number of part being copied.
      --RequestPayer string                     Confirms that the requester knows that they will be charged for the request.
      --SSECustomerAlgorithm string             Specifies the algorithm to use when encrypting the object (for example, AES256).
      --SSECustomerKey string                   Specifies the customer-provided encryption key for Amazon S3 to use in encrypting data.
      --SSECustomerKeyMD5 string                Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321.
      --UploadId string                         [REQUIRED] Upload ID identifying the multipart upload whose part is being copied.
  -h, --help                                    help for UploadPartCopy
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

