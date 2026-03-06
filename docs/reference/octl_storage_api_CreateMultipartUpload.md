## octl storage api CreateMultipartUpload



```
octl storage api CreateMultipartUpload [flags]
```

### Options

```
      --ACL string                          The canned ACL to apply to the object.
      --Bucket string                       The name of the bucket where the multipart upload is initiated and where the object is uploaded.
      --BucketKeyEnabled                    Specifies whether Amazon S3 should use an S3 Bucket Key for object encryption with server-side encryption using Key Management Service (KMS) keys (SSE-KMS).
      --CacheControl string                 Specifies caching behavior along the request/reply chain.
      --ChecksumAlgorithm string            Indicates the algorithm that you want Amazon S3 to use to create the checksum for the object.
      --ChecksumType string                 Indicates the checksum type that you want Amazon S3 to use to calculate the object’s checksum value.
      --ContentDisposition string           Specifies presentational information for the object.
      --ContentEncoding string              Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field.
      --ContentLanguage string              The language that the content is in.
      --ContentType string                  A standard MIME type describing the format of the object data.
      --ExpectedBucketOwner string          The account ID of the expected bucket owner.
      --Expires osctime                     The date and time at which the object is no longer cacheable.
      --GrantFullControl string             Specify access permissions explicitly to give the grantee READ, READ_ACP, and WRITE_ACP permissions on the object.
      --GrantRead string                    Specify access permissions explicitly to allow grantee to read the object data and its metadata.
      --GrantReadACP string                 Specify access permissions explicitly to allows grantee to read the object ACL.
      --GrantWriteACP string                Specify access permissions explicitly to allows grantee to allow grantee to write the ACL for the applicable object.
      --Key string                          Object key for which the multipart upload is to be initiated.
      --ObjectLockLegalHoldStatus string    Specifies whether you want to apply a legal hold to the uploaded object.
      --ObjectLockMode string               Specifies the Object Lock mode that you want to apply to the uploaded object.
      --ObjectLockRetainUntilDate osctime   Specifies the date and time when you want the Object Lock to expire.
      --RequestPayer string                 Confirms that the requester knows that they will be charged for the request.
      --SSECustomerAlgorithm string         Specifies the algorithm to use when encrypting the object (for example, AES256).
      --SSECustomerKey string               Specifies the customer-provided encryption key for Amazon S3 to use in encrypting data.
      --SSECustomerKeyMD5 string            Specifies the 128-bit MD5 digest of the customer-provided encryption key according to RFC 1321.
      --SSEKMSEncryptionContext string      Specifies the Amazon Web Services KMS Encryption Context to use for object encryption.
      --SSEKMSKeyId string                  Specifies the KMS key ID (Key ID, Key ARN, or Key Alias) to use for object encryption.
      --ServerSideEncryption string         The server-side encryption algorithm used when you store this object in Amazon S3 or Amazon FSx.
      --StorageClass string                 By default, Amazon S3 uses the STANDARD Storage Class to store newly created objects.
      --Tagging string                      The tag-set for the object.
      --WebsiteRedirectLocation string      If the bucket is configured as a website, redirects requests for this object to another object in the same bucket or to an external URL.
  -h, --help                                help for CreateMultipartUpload
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

