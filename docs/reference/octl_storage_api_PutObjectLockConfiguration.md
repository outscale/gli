## octl storage api PutObjectLockConfiguration



```
octl storage api PutObjectLockConfiguration [flags]
```

### Options

```
      --Bucket string                                               The bucket whose Object Lock configuration you want to create or replace.
      --ChecksumAlgorithm string                                    Indicates the algorithm used to create the checksum for the object when you use the SDK.
      --ContentMD5 string                                           The MD5 hash for the request body.
      --ExpectedBucketOwner string                                  The account ID of the expected bucket owner.
      --ObjectLockConfiguration.ObjectLockEnabled string            Indicates whether this bucket has an Object Lock configuration enabled.
      --ObjectLockConfiguration.Rule.DefaultRetention.Days int32    The number of days that you want to specify for the default retention period.
      --ObjectLockConfiguration.Rule.DefaultRetention.Mode string   The default Object Lock retention mode you want to apply to new objects placed in the specified bucket.
      --ObjectLockConfiguration.Rule.DefaultRetention.Years int32   The number of years that you want to specify for the default retention period.
      --RequestPayer string                                         Confirms that the requester knows that they will be charged for the request.
      --Token string                                                A token to allow Object Lock to be enabled for an existing bucket.
  -h, --help                                                        help for PutObjectLockConfiguration
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

