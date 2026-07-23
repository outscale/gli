## octl storage api PutBucketEncryption



```
octl storage api PutBucketEncryption [flags]
```

### Options

```
      --Bucket string                                                                                        [REQUIRED] Specifies default encryption for a bucket using server-side encryption with different key options.
      --ChecksumAlgorithm string                                                                             Indicates the algorithm used to create the checksum for the object when you use the SDK.
      --ContentMD5 string                                                                                    The base64-encoded 128-bit MD5 digest of the server-side encryption configuration.
      --ExpectedBucketOwner string                                                                           The account ID of the expected bucket owner.
      --ServerSideEncryptionConfiguration.Rules.0.ApplyServerSideEncryptionByDefault.KMSMasterKeyID string   Amazon Web Services Key Management Service (KMS) customer managed key ID to use for the default encryption.
      --ServerSideEncryptionConfiguration.Rules.0.ApplyServerSideEncryptionByDefault.SSEAlgorithm string     [REQUIRED] Server-side encryption algorithm to use for the default encryption.
      --ServerSideEncryptionConfiguration.Rules.0.BucketKeyEnabled                                           Specifies whether Amazon S3 should use an S3 Bucket Key with server-side encryption using KMS (SSE-KMS) for new objects in the bucket.
  -h, --help                                                                                                 help for PutBucketEncryption
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

