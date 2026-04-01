## octl storage api PutBucketLifecycleConfiguration



```
octl storage api PutBucketLifecycleConfiguration [flags]
```

### Options

```
      --Bucket string                                                                                 The name of the bucket for which to set the configuration.
      --ChecksumAlgorithm string                                                                      Indicates the algorithm used to create the checksum for the object when you use the SDK.
      --ExpectedBucketOwner string                                                                    The account ID of the expected bucket owner.
      --LifecycleConfiguration.Rules.0.AbortIncompleteMultipartUpload.DaysAfterInitiation int32       Specifies the number of days after which Amazon S3 aborts an incomplete multipart upload.
      --LifecycleConfiguration.Rules.0.Expiration.Date osctime                                        Indicates at what date the object is to be moved or deleted.
      --LifecycleConfiguration.Rules.0.Expiration.Days int32                                          Indicates the lifetime, in days, of the objects that are subject to the rule.
      --LifecycleConfiguration.Rules.0.Expiration.ExpiredObjectDeleteMarker                           Indicates whether Amazon S3 will remove a delete marker with no noncurrent versions.
      --LifecycleConfiguration.Rules.0.Filter.And.ObjectSizeGreaterThan int                           Minimum object size to which the rule applies.
      --LifecycleConfiguration.Rules.0.Filter.And.ObjectSizeLessThan int                              Maximum object size to which the rule applies.
      --LifecycleConfiguration.Rules.0.Filter.And.Prefix string                                       identifying one or more objects to which the rule applies.
      --LifecycleConfiguration.Rules.0.Filter.And.Tags.0.Key string                                   Name of the object key.
      --LifecycleConfiguration.Rules.0.Filter.And.Tags.0.Value string                                 of the tag.
      --LifecycleConfiguration.Rules.0.Filter.ObjectSizeGreaterThan int                               Minimum object size to which the rule applies.
      --LifecycleConfiguration.Rules.0.Filter.ObjectSizeLessThan int                                  Maximum object size to which the rule applies.
      --LifecycleConfiguration.Rules.0.Filter.Prefix string                                           identifying one or more objects to which the rule applies.
      --LifecycleConfiguration.Rules.0.Filter.Tag.Key string                                          Name of the object key.
      --LifecycleConfiguration.Rules.0.Filter.Tag.Value string                                        of the tag.
      --LifecycleConfiguration.Rules.0.ID string                                                      Unique identifier for the rule.
      --LifecycleConfiguration.Rules.0.NoncurrentVersionExpiration.NewerNoncurrentVersions int32      Specifies how many noncurrent versions Amazon S3 will retain.
      --LifecycleConfiguration.Rules.0.NoncurrentVersionExpiration.NoncurrentDays int32               Specifies the number of days an object is noncurrent before Amazon S3 can perform the associated action.
      --LifecycleConfiguration.Rules.0.NoncurrentVersionTransitions.0.NewerNoncurrentVersions int32   Specifies how many noncurrent versions Amazon S3 will retain in the same storage class before transitioning objects.
      --LifecycleConfiguration.Rules.0.NoncurrentVersionTransitions.0.NoncurrentDays int32            Specifies the number of days an object is noncurrent before Amazon S3 can perform the associated action.
      --LifecycleConfiguration.Rules.0.NoncurrentVersionTransitions.0.StorageClass string             The class of storage used to store the object.
      --LifecycleConfiguration.Rules.0.Prefix string                                                  identifying one or more objects to which the rule applies.
      --LifecycleConfiguration.Rules.0.Status string                                                  If 'Enabled', the rule is currently being applied.
      --LifecycleConfiguration.Rules.0.Transitions.0.Date osctime                                     Indicates when objects are transitioned to the specified storage class.
      --LifecycleConfiguration.Rules.0.Transitions.0.Days int32                                       Indicates the number of days after creation when objects are transitioned to the specified storage class.
      --LifecycleConfiguration.Rules.0.Transitions.0.StorageClass string                              The storage class to which you want the object to transition.
      --TransitionDefaultMinimumObjectSize string                                                     Indicates which default minimum object size behavior is applied to the lifecycle configuration.
  -h, --help                                                                                          help for PutBucketLifecycleConfiguration
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

