## octl storage api PutBucketEncryption



```
octl storage api PutBucketEncryption [flags]
```

### Options

```
      --Bucket string                                                                                        Specifies default encryption for a bucket using server-side encryption with different key options.
      --ChecksumAlgorithm string                                                                             Indicates the algorithm used to create the checksum for the request when you use the SDK.
      --ContentMD5 string                                                                                    The Base64 encoded 128-bit MD5 digest of the server-side encryption configuration.
      --ExpectedBucketOwner string                                                                           The account ID of the expected bucket owner.
      --ServerSideEncryptionConfiguration.Rules.0.ApplyServerSideEncryptionByDefault.KMSMasterKeyID string   
      --ServerSideEncryptionConfiguration.Rules.0.ApplyServerSideEncryptionByDefault.SSEAlgorithm string     
      --ServerSideEncryptionConfiguration.Rules.0.BlockedEncryptionTypes.EncryptionType strings              
      --ServerSideEncryptionConfiguration.Rules.0.BucketKeyEnabled                                           
  -h, --help                                                                                                 help for PutBucketEncryption
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

