## octl storage api CreateBucket



```
octl storage api CreateBucket [flags]
```

### Options

```
      --ACL string                                               The canned ACL to apply to the bucket.
      --Bucket string                                            [REQUIRED] The name of the bucket to create.
      --CreateBucketConfiguration.Bucket.DataRedundancy string   The number of Zone (Availability Zone or Local Zone) that's used for redundancy for the bucket.
      --CreateBucketConfiguration.Bucket.Type string             The type of bucket.
      --CreateBucketConfiguration.Location.Name string           The name of the location where the bucket will be created.
      --CreateBucketConfiguration.Location.Type string           The type of location where the bucket will be created.
      --CreateBucketConfiguration.LocationConstraint string      Specifies the Region where the bucket will be created.
      --GrantFullControl string                                  Allows grantee the read, write, read ACP, and write ACP permissions on the bucket.
      --GrantRead string                                         Allows grantee to list the objects in the bucket.
      --GrantReadACP string                                      Allows grantee to read the bucket ACL.
      --GrantWrite string                                        Allows grantee to create new objects in the bucket.
      --GrantWriteACP string                                     Allows grantee to write the ACL for the applicable bucket.
      --ObjectLockEnabledForBucket                               Specifies whether you want S3 Object Lock to be enabled for the new bucket.
      --ObjectOwnership string                                   The container element for object ownership for a bucket's ownership controls.
  -h, --help                                                     help for CreateBucket
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

