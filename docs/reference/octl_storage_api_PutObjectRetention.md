## octl storage api PutObjectRetention



```
octl storage api PutObjectRetention [flags]
```

### Options

```
      --Bucket string                       The bucket name that contains the object you want to apply this Object Retention configuration to.
      --BypassGovernanceRetention           Indicates whether this action should bypass Governance-mode restrictions.
      --ChecksumAlgorithm string            Indicates the algorithm used to create the checksum for the object when you use the SDK.
      --ContentMD5 string                   The MD5 hash for the request body.
      --ExpectedBucketOwner string          The account ID of the expected bucket owner.
      --Key string                          The key name for the object that you want to apply this Object Retention configuration to.
      --RequestPayer string                 Confirms that the requester knows that they will be charged for the request.
      --Retention.Mode string               
      --Retention.RetainUntilDate osctime   
      --VersionId string                    The version ID for the object that you want to apply this Object Retention configuration to.
  -h, --help                                help for PutObjectRetention
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

