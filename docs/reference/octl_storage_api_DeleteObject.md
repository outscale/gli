## octl storage api DeleteObject



```
octl storage api DeleteObject [flags]
```

### Options

```
      --Bucket string                     The bucket name of the bucket containing the object.
      --BypassGovernanceRetention         Indicates whether S3 Object Lock should bypass Governance-mode restrictions to process this operation.
      --ExpectedBucketOwner string        The account ID of the expected bucket owner.
      --IfMatch string                    The If-Match header field makes the request method conditional on ETags.
      --IfMatchLastModifiedTime osctime   If present, the object is deleted only if its modification times matches the provided Timestamp .
      --IfMatchSize int                   If present, the object is deleted only if its size matches the provided size in bytes.
      --Key string                        name of the object to delete.
      --MFA string                        The concatenation of the authentication device's serial number, a space, and the value that is displayed on your authentication device.
      --RequestPayer string               Confirms that the requester knows that they will be charged for the request.
      --VersionId string                  Version ID used to reference a specific version of the object.
  -h, --help                              help for DeleteObject
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

