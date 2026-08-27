## octl storage api DeleteObjects



```
octl storage api DeleteObjects [flags]
```

### Options

```
      --Bucket string                               [REQUIRED] The bucket name containing the objects to delete.
      --BypassGovernanceRetention                   Specifies whether you want to delete this object even if it has a Governance-type Object Lock in place.
      --ChecksumAlgorithm string                    Indicates the algorithm used to create the checksum for the object when you use the SDK.
      --Delete.Objects.0.ETag string                An entity tag (ETag) is an identifier assigned by a web server to a specific version of a resource found at a URL.
      --Delete.Objects.0.Key string                 [REQUIRED] name of the object.
      --Delete.Objects.0.LastModifiedTime osctime   If present, the objects are deleted only if its modification times matches the provided Timestamp .
      --Delete.Objects.0.Size int                   If present, the objects are deleted only if its size matches the provided size in bytes.
      --Delete.Objects.0.VersionId string           Version ID for the specific version of the object to delete.
      --Delete.Quiet                                Element to enable quiet mode for the request.
      --ExpectedBucketOwner string                  The account ID of the expected bucket owner.
      --MFA string                                  The concatenation of the authentication device's serial number, a space, and the value that is displayed on your authentication device.
      --RequestPayer string                         Confirms that the requester knows that they will be charged for the request.
  -h, --help                                        help for DeleteObjects
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
  -o, --output string              output format (raw, json, yaml, table, csv, none, text)
      --payload string             JSON content for query body
      --profile string             Profile to use in profile file (by default, "default")
  -s, --silent                     Hides all information messages
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

