## octl storage bucket

bucket commands

### Options

```
  -h, --help   help for bucket
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

* [octl storage](octl_storage.md)	 - OUTSCALE Object Storage (OOS) management
* [octl storage bucket acl](octl_storage_bucket_acl.md)	 - acl commands
* [octl storage bucket cors](octl_storage_bucket_cors.md)	 - cors commands
* [octl storage bucket create](octl_storage_bucket_create.md)	 - alias for api CreateBucket
* [octl storage bucket delete](octl_storage_bucket_delete.md)	 - alias for api DeleteBucket --Bucket bucket
* [octl storage bucket describe](octl_storage_bucket_describe.md)	 - Display a bucket, alias for api HeadBucket --Bucket bucket
* [octl storage bucket encryption](octl_storage_bucket_encryption.md)	 - encryption commands
* [octl storage bucket lifecycle](octl_storage_bucket_lifecycle.md)	 - lifecycle commands
* [octl storage bucket list](octl_storage_bucket_list.md)	 - alias for api ListBuckets
* [octl storage bucket objectlock](octl_storage_bucket_objectlock.md)	 - objectlock commands
* [octl storage bucket policy](octl_storage_bucket_policy.md)	 - policy commands
* [octl storage bucket versioning](octl_storage_bucket_versioning.md)	 - versioning commands
* [octl storage bucket website](octl_storage_bucket_website.md)	 - website commands

