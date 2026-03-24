## octl storage object put

alias for api PutObject --Key key

### Synopsis

> *alias for api PutObject --Key key*



```
octl storage object put key [key]... [flags]
```

### Options

```
      --acl string                         The canned ACL to apply to the object.
      --body streamedFile                  Object data.
      --bucket string                      The bucket name to which the PUT action was initiated.
      --bucket-key-enabled                 Specifies whether Amazon S3 should use an S3 Bucket Key for object encryption with server-side encryption using Key Management Service (KMS) keys (SSE-KMS).
      --cache-control string               Can be used to specify caching behavior along the request/reply chain.
      --checksum-algorithm string          Indicates the algorithm used to create the checksum for the object when you use the SDK.
      --checksum-crc-32 string             This header can be used as a data integrity check to verify that the data received is the same data that was originally sent.
      --checksum-crc-32-c string           This header can be used as a data integrity check to verify that the data received is the same data that was originally sent.
      --checksum-crc-64-nvme string        This header can be used as a data integrity check to verify that the data received is the same data that was originally sent.
      --checksum-sha-1 string              This header can be used as a data integrity check to verify that the data received is the same data that was originally sent.
      --checksum-sha-256 string            This header can be used as a data integrity check to verify that the data received is the same data that was originally sent.
      --content-disposition string         Specifies presentational information for the object.
      --content-encoding string            Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field.
      --content-language string            The language the content is in.
      --content-length int                 Size of the body in bytes.
      --content-md-5 string                The Base64 encoded 128-bit MD5 digest of the message (without the headers) according to RFC 1864.
      --content-type string                A standard MIME type describing the format of the contents.
      --expected-bucket-owner string       The account ID of the expected bucket owner.
      --expire osctime                     The date and time at which the object is no longer cacheable.
      --grant-full-control string          Gives the grantee READ, READ_ACP, and WRITE_ACP permissions on the object.
      --grant-read string                  Allows grantee to read the object data and its metadata.
      --grant-read-acp string              Allows grantee to read the object ACL.
      --grant-write-acp string             Allows grantee to write the ACL for the applicable object.
  -h, --help                               help for put
      --if-match string                    Uploads the object only if the ETag (entity tag) value provided during the WRITE operation matches the ETag of the object in S3.
      --if-none-match string               Uploads the object only if the object key name does not already exist in the bucket specified.
      --lock-legal-hold-status string      Specifies whether a legal hold will be applied to this object.
      --lock-mode string                   The Object Lock mode that you want to apply to this object.
      --lock-retain-until-date osctime     The date and time when you want this object's Object Lock to expire.
      --server-side-encryption string      The server-side encryption algorithm that was used when you store this object in Amazon S3 or Amazon FSx.
      --sse-customer-algorithm string      Specifies the algorithm to use when encrypting the object (for example, AES256 ).
      --sse-customer-key string            Specifies the customer-provided encryption key for Amazon S3 to use in encrypting data.
      --sse-customer-key-md-5 string       Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321.
      --ssekm-encryption-context string    Specifies the Amazon Web Services KMS Encryption Context as an additional encryption context to use for object encryption.
      --ssekm-key-id string                Specifies the KMS key ID (Key ID, Key ARN, or Key Alias) to use for object encryption.
      --storage-class string               By default, Amazon S3 uses the STANDARD Storage Class to store newly created objects.
      --tagging string                     The tag-set for the object.
      --website-redirect-location string   If the bucket is configured as a website, redirects requests for this object to another object in the same bucket or to an external URL.
      --write-offset-byte int              Specifies the offset for appending data to existing objects in bytes.
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

* [octl storage object](octl_storage_object.md)	 - object commands

