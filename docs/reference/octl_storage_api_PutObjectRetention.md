## octl storage api PutObjectRetention



```
octl storage api PutObjectRetention [flags]
```

### Options

```
      --Bucket string                       [REQUIRED] The bucket name that contains the object you want to apply this Object Retention configuration to.
      --BypassGovernanceRetention           Indicates whether this action should bypass Governance-mode restrictions.
      --ChecksumAlgorithm string            Indicates the algorithm used to create the checksum for the object when you use the SDK.
      --ContentMD5 string                   The MD5 hash for the request body.
      --ExpectedBucketOwner string          The account ID of the expected bucket owner.
      --Key string                          [REQUIRED] The key name for the object that you want to apply this Object Retention configuration to.
      --RequestPayer string                 Confirms that the requester knows that they will be charged for the request.
      --Retention.Mode string               Indicates the Retention mode for the specified object.
      --Retention.RetainUntilDate osctime   The date on which this Object Lock Retention will expire.
      --VersionId string                    The version ID for the object that you want to apply this Object Retention configuration to.
  -h, --help                                help for PutObjectRetention
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

