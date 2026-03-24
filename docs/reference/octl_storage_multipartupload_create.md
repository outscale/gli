## octl storage multipartupload create

alias for api CreateMultipartUpload

### Synopsis

> *alias for api CreateMultipartUpload*



```
octl storage multipartupload create [flags]
```

### Options

```
      --acl string                              The canned ACL to apply to the object.
      --bucket string                           The name of the bucket where the multipart upload is initiated and where the object is uploaded.
      --bucket-key-enabled                      Specifies whether Amazon S3 should use an S3 Bucket Key for object encryption with server-side encryption using Key Management Service (KMS) keys (SSE-KMS).
      --cache-control string                    Specifies caching behavior along the request/reply chain.
      --checksum-algorithm string               Indicates the algorithm that you want Amazon S3 to use to create the checksum for the object.
      --checksum-type string                    Indicates the checksum type that you want Amazon S3 to use to calculate the object’s checksum value.
      --content-disposition string              Specifies presentational information for the object.
      --content-encoding string                 Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field.
      --content-language string                 The language that the content is in.
      --content-type string                     A standard MIME type describing the format of the object data.
      --expected-bucket-owner string            The account ID of the expected bucket owner.
      --expire osctime                          The date and time at which the object is no longer cacheable.
      --grant-full-control string               Specify access permissions explicitly to give the grantee READ, READ_ACP, and WRITE_ACP permissions on the object.
      --grant-read string                       Specify access permissions explicitly to allow grantee to read the object data and its metadata.
      --grant-read-acp string                   Specify access permissions explicitly to allows grantee to read the object ACL.
      --grant-write-acp string                  Specify access permissions explicitly to allows grantee to allow grantee to write the ACL for the applicable object.
  -h, --help                                    help for create
      --key string                              Object key for which the multipart upload is to be initiated.
      --object-lock-legal-hold-status string    Specifies whether you want to apply a legal hold to the uploaded object.
      --object-lock-mode string                 Specifies the Object Lock mode that you want to apply to the uploaded object.
      --object-lock-retain-until-date osctime   Specifies the date and time when you want the Object Lock to expire.
      --server-side-encryption string           The server-side encryption algorithm used when you store this object in Amazon S3 or Amazon FSx.
      --sse-customer-algorithm string           Specifies the algorithm to use when encrypting the object (for example, AES256).
      --sse-customer-key string                 Specifies the customer-provided encryption key for Amazon S3 to use in encrypting data.
      --sse-customer-key-md-5 string            Specifies the 128-bit MD5 digest of the customer-provided encryption key according to RFC 1321.
      --ssekm-encryption-context string         Specifies the Amazon Web Services KMS Encryption Context to use for object encryption.
      --ssekm-key-id string                     Specifies the KMS key ID (Key ID, Key ARN, or Key Alias) to use for object encryption.
      --storage-class string                    By default, Amazon S3 uses the STANDARD Storage Class to store newly created objects.
      --tagging string                          The tag-set for the object.
      --website-redirect-location string        If the bucket is configured as a website, redirects requests for this object to another object in the same bucket or to an external URL.
```

### Options inherited from parent commands

```
  -c, --columns string              columns to display - [+]<title>:<jq query for content>||<title>:<jq query for content>
      --config string               Path of profile file (by default, ~/.osc/config.json)
      --filter strings              comma separated list of filters for results - name:value,name:value, alias for jq filter 'select(.name | test("value"))'
      --jq string                   jq filter
      --no-upgrade                  do not check for new versions
  -O, --out-file string             redirect output to file
  -o, --output string               output format (raw, json, yaml, table, csv, none, base64) (default "raw")
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

* [octl storage multipartupload](octl_storage_multipartupload.md)	 - multipartupload commands

