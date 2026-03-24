## octl storage bucket versioning

versioning commands

### Options

```
  -h, --help   help for versioning
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

* [octl storage bucket](octl_storage_bucket.md)	 - bucket commands
* [octl storage bucket versioning describe](octl_storage_bucket_versioning_describe.md)	 - Display versioning configuration, alias for api GetBucketVersioning --Bucket bucket
* [octl storage bucket versioning disable](octl_storage_bucket_versioning_disable.md)	 - Disable versioning, alias for api PutBucketVersioning --Bucket bucket --VersioningConfiguration.Status Suspended
* [octl storage bucket versioning enable](octl_storage_bucket_versioning_enable.md)	 - Enable versioning, alias for api PutBucketVersioning --Bucket bucket --VersioningConfiguration.Status Enabled

