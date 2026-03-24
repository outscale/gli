## octl storage api CopyObject



```
octl storage api CopyObject [flags]
```

### Options

```
      --ACL string                              The canned access control list (ACL) to apply to the object.
      --Bucket string                           The name of the destination bucket.
      --BucketKeyEnabled                        Specifies whether Amazon S3 should use an S3 Bucket Key for object encryption with server-side encryption using Key Management Service (KMS) keys (SSE-KMS).
      --CacheControl string                     Specifies the caching behavior along the request/reply chain.
      --ChecksumAlgorithm string                Indicates the algorithm that you want Amazon S3 to use to create the checksum for the object.
      --ContentDisposition string               Specifies presentational information for the object.
      --ContentEncoding string                  Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field.
      --ContentLanguage string                  The language the content is in.
      --ContentType string                      A standard MIME type that describes the format of the object data.
      --CopySource string                       Specifies the source object for the copy operation.
      --CopySourceIfMatch string                Copies the object if its entity tag (ETag) matches the specified tag.
      --CopySourceIfModifiedSince osctime       Copies the object if it has been modified since the specified time.
      --CopySourceIfNoneMatch string            Copies the object if its entity tag (ETag) is different than the specified ETag.
      --CopySourceIfUnmodifiedSince osctime     Copies the object if it hasn't been modified since the specified time.
      --CopySourceSSECustomerAlgorithm string   Specifies the algorithm to use when decrypting the source object (for example, AES256 ).
      --CopySourceSSECustomerKey string         Specifies the customer-provided encryption key for Amazon S3 to use to decrypt the source object.
      --CopySourceSSECustomerKeyMD5 string      Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321.
      --ExpectedBucketOwner string              The account ID of the expected destination bucket owner.
      --ExpectedSourceBucketOwner string        The account ID of the expected source bucket owner.
      --Expires osctime                         The date and time at which the object is no longer cacheable.
      --GrantFullControl string                 Gives the grantee READ, READ_ACP, and WRITE_ACP permissions on the object.
      --GrantRead string                        Allows grantee to read the object data and its metadata.
      --GrantReadACP string                     Allows grantee to read the object ACL.
      --GrantWriteACP string                    Allows grantee to write the ACL for the applicable object.
      --Key string                              The key of the destination object.
      --MetadataDirective string                Specifies whether the metadata is copied from the source object or replaced with metadata that's provided in the request.
      --ObjectLockLegalHoldStatus string        Specifies whether you want to apply a legal hold to the object copy.
      --ObjectLockMode string                   The Object Lock mode that you want to apply to the object copy.
      --ObjectLockRetainUntilDate osctime       The date and time when you want the Object Lock of the object copy to expire.
      --RequestPayer string                     Confirms that the requester knows that they will be charged for the request.
      --SSECustomerAlgorithm string             Specifies the algorithm to use when encrypting the object (for example, AES256 ).
      --SSECustomerKey string                   Specifies the customer-provided encryption key for Amazon S3 to use in encrypting data.
      --SSECustomerKeyMD5 string                Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321.
      --SSEKMSEncryptionContext string          Specifies the Amazon Web Services KMS Encryption Context as an additional encryption context to use for the destination object encryption.
      --SSEKMSKeyId string                      Specifies the KMS key ID (Key ID, Key ARN, or Key Alias) to use for object encryption.
      --ServerSideEncryption string             The server-side encryption algorithm used when storing this object in Amazon S3.
      --StorageClass string                     If the x-amz-storage-class header is not used, the copied object will be stored in the STANDARD Storage Class by default.
      --Tagging string                          The tag-set for the object copy in the destination bucket.
      --TaggingDirective string                 Specifies whether the object tag-set is copied from the source object or replaced with the tag-set that's provided in the request.
      --WebsiteRedirectLocation string          If the destination bucket is configured as a website, redirects requests for this object copy to another object in the same bucket or to an external URL.
  -h, --help                                    help for CopyObject
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

