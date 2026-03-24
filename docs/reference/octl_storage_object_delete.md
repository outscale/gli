## octl storage object delete

alias for api DeleteObject --Key key

### Synopsis

> *alias for api DeleteObject --Key key*



```
octl storage object delete key [key]... [flags]
```

### Options

```
      --bucket string                         The bucket name of the bucket containing the object.
      --bypass-governance-retention           Indicates whether S3 Object Lock should bypass Governance-mode restrictions to process this operation.
      --expected-bucket-owner string          The account ID of the expected bucket owner.
  -h, --help                                  help for delete
      --if-match string                       Deletes the object if the ETag (entity tag) value provided during the delete operation matches the ETag of the object in S3.
      --if-match-last-modified-time osctime   If present, the object is deleted only if its modification times matches the provided Timestamp .
      --if-match-size int                     If present, the object is deleted only if its size matches the provided size in bytes.
      --mfa string                            The concatenation of the authentication device's serial number, a space, and the value that is displayed on your authentication device.
      --version-id string                     Version ID used to reference a specific version of the object.
```

### Options inherited from parent commands

```
  -c, --columns string              columns to display - [+]<title>:<jq query for content>||<title>:<jq query for content>
      --config string               Path of profile file (by default, ~/.osc/config.json)
      --filter strings              comma separated list of filters for results - name:value,name:value, alias for jq filter 'select(.name | test("value"))'
      --jq string                   jq filter
      --no-upgrade                  do not check for new versions
  -O, --out-file string             redirect output to file
  -o, --output string               output format (raw, json, yaml, table, csv, none, base64) (default "raw")
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

* [octl storage object](octl_storage_object.md)	 - object commands

