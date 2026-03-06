## octl storage api CreateBucket



```
octl storage api CreateBucket [flags]
```

### Options

```
      --ACL string                                               The canned ACL to apply to the bucket.
      --Bucket string                                            The name of the bucket to create.
      --CreateBucketConfiguration.Bucket.DataRedundancy string   
      --CreateBucketConfiguration.Bucket.Type string             
      --CreateBucketConfiguration.Location.Name string           
      --CreateBucketConfiguration.Location.Type string           
      --CreateBucketConfiguration.LocationConstraint string      
      --CreateBucketConfiguration.Tags.0.Key string              
      --CreateBucketConfiguration.Tags.0.Value string            
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

