## octl storage api UploadPartCopy



```
octl storage api UploadPartCopy [flags]
```

### Options

```
      --Bucket string                           The bucket name.
      --CopySource string                       Specifies the source object for the copy operation.
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
      --Key string                              Object key for which the multipart upload was initiated.
      --PartNumber int32                        Part number of part being copied.
      --RequestPayer string                     Confirms that the requester knows that they will be charged for the request.
      --SSECustomerAlgorithm string             Specifies the algorithm to use when encrypting the object (for example, AES256).
      --SSECustomerKey string                   Specifies the customer-provided encryption key for Amazon S3 to use in encrypting data.
      --SSECustomerKeyMD5 string                Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321.
      --UploadId string                         Upload ID identifying the multipart upload whose part is being copied.
  -h, --help                                    help for UploadPartCopy
```

### Options inherited from parent commands

```
  -c, --columns string    columns to display - [+]<title>:<jq query for content>||<title>:<jq query for content>
      --config string     Path of profile file (by default, ~/.osc/config.json)
      --filter strings    comma separated list of filters for results - name:value,name:value, alias for jq filter 'select(.name | test("value"))'
      --jq string         jq filter
      --no-upgrade        do not check for new versions
  -O, --out-file string   redirect output to file
  -o, --output string     output format (raw, json, yaml, table, csv, none, base64) (default "raw")
      --profile string    Profile to use in profile file (by default, "default")
      --single            convert single entry lists to a single object
      --template string   JSON template for query body
  -v, --verbose           Verbose output
  -y, --yes               answer yes to all prompts
```

### SEE ALSO

* [octl storage api](octl_storage_api.md)	 - storage api calls

