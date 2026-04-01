## octl storage api PutBucketAcl



```
octl storage api PutBucketAcl [flags]
```

### Options

```
      --ACL string                                                 The canned ACL to apply to the bucket.
      --AccessControlPolicy.Grants.0.Grantee.DisplayName string    Screen name of the grantee.
      --AccessControlPolicy.Grants.0.Grantee.EmailAddress string   Email address of the grantee.
      --AccessControlPolicy.Grants.0.Grantee.ID string             The canonical user ID of the grantee.
      --AccessControlPolicy.Grants.0.Grantee.Type string           of grantee This member is required.
      --AccessControlPolicy.Grants.0.Grantee.URI string            of the grantee group.
      --AccessControlPolicy.Grants.0.Permission string             Specifies the permission given to the grantee.
      --AccessControlPolicy.Owner.DisplayName string               Container for the display name of the owner.
      --AccessControlPolicy.Owner.ID string                        Container for the ID of the owner.
      --Bucket string                                              The bucket to which to apply the ACL.
      --ChecksumAlgorithm string                                   Indicates the algorithm used to create the checksum for the object when you use the SDK.
      --ContentMD5 string                                          The base64-encoded 128-bit MD5 digest of the data.
      --ExpectedBucketOwner string                                 The account ID of the expected bucket owner.
      --GrantFullControl string                                    Allows grantee the read, write, read ACP, and write ACP permissions on the bucket.
      --GrantRead string                                           Allows grantee to list the objects in the bucket.
      --GrantReadACP string                                        Allows grantee to read the bucket ACL.
      --GrantWrite string                                          Allows grantee to create new objects in the bucket.
      --GrantWriteACP string                                       Allows grantee to write the ACL for the applicable bucket.
  -h, --help                                                       help for PutBucketAcl
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

