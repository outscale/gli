## octl storage api PutObject



```
octl storage api PutObject [flags]
```

### Options

```
      --ACL string                          The canned ACL to apply to the object.
      --Body streamedFile                   Object data.
      --Bucket string                       The bucket name to which the PUT action was initiated.
      --BucketKeyEnabled                    Specifies whether Amazon S3 should use an S3 Bucket Key for object encryption with server-side encryption using Key Management Service (KMS) keys (SSE-KMS).
      --CacheControl string                 Can be used to specify caching behavior along the request/reply chain.
      --ChecksumAlgorithm string            Indicates the algorithm used to create the checksum for the object when you use the SDK.
      --ChecksumCRC32 string                This header can be used as a data integrity check to verify that the data received is the same data that was originally sent.
      --ChecksumCRC32C string               This header can be used as a data integrity check to verify that the data received is the same data that was originally sent.
      --ChecksumCRC64NVME string            This header can be used as a data integrity check to verify that the data received is the same data that was originally sent.
      --ChecksumSHA1 string                 This header can be used as a data integrity check to verify that the data received is the same data that was originally sent.
      --ChecksumSHA256 string               This header can be used as a data integrity check to verify that the data received is the same data that was originally sent.
      --ContentDisposition string           Specifies presentational information for the object.
      --ContentEncoding string              Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field.
      --ContentLanguage string              The language the content is in.
      --ContentLength int                   Size of the body in bytes.
      --ContentMD5 string                   The Base64 encoded 128-bit MD5 digest of the message (without the headers) according to RFC 1864.
      --ContentType string                  A standard MIME type describing the format of the contents.
      --ExpectedBucketOwner string          The account ID of the expected bucket owner.
      --Expires osctime                     The date and time at which the object is no longer cacheable.
      --GrantFullControl string             Gives the grantee READ, READ_ACP, and WRITE_ACP permissions on the object.
      --GrantRead string                    Allows grantee to read the object data and its metadata.
      --GrantReadACP string                 Allows grantee to read the object ACL.
      --GrantWriteACP string                Allows grantee to write the ACL for the applicable object.
      --IfMatch string                      Uploads the object only if the ETag (entity tag) value provided during the WRITE operation matches the ETag of the object in S3.
      --IfNoneMatch string                  Uploads the object only if the object key name does not already exist in the bucket specified.
      --Key string                          Object key for which the PUT action was initiated.
      --ObjectLockLegalHoldStatus string    Specifies whether a legal hold will be applied to this object.
      --ObjectLockMode string               The Object Lock mode that you want to apply to this object.
      --ObjectLockRetainUntilDate osctime   The date and time when you want this object's Object Lock to expire.
      --RequestPayer string                 Confirms that the requester knows that they will be charged for the request.
      --SSECustomerAlgorithm string         Specifies the algorithm to use when encrypting the object (for example, AES256 ).
      --SSECustomerKey string               Specifies the customer-provided encryption key for Amazon S3 to use in encrypting data.
      --SSECustomerKeyMD5 string            Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321.
      --SSEKMSEncryptionContext string      Specifies the Amazon Web Services KMS Encryption Context as an additional encryption context to use for object encryption.
      --SSEKMSKeyId string                  Specifies the KMS key ID (Key ID, Key ARN, or Key Alias) to use for object encryption.
      --ServerSideEncryption string         The server-side encryption algorithm that was used when you store this object in Amazon S3 or Amazon FSx.
      --StorageClass string                 By default, Amazon S3 uses the STANDARD Storage Class to store newly created objects.
      --Tagging string                      The tag-set for the object.
      --WebsiteRedirectLocation string      If the bucket is configured as a website, redirects requests for this object to another object in the same bucket or to an external URL.
      --WriteOffsetBytes int                Specifies the offset for appending data to existing objects in bytes.
  -h, --help                                help for PutObject
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
      --payload string    JSON content for query body
      --profile string    Profile to use in profile file (by default, "default")
      --single            convert single entry lists to a single object
      --template string   JSON template file for query body
  -v, --verbose           Verbose output
  -y, --yes               answer yes to all prompts
```

### SEE ALSO

* [octl storage api](octl_storage_api.md)	 - storage api calls

