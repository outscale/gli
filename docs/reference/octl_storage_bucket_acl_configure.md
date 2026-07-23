## octl storage bucket acl configure

Update ACL configuration, alias for api PutBucketAcl --Bucket bucket

### Synopsis

> *Update ACL configuration, alias for api PutBucketAcl --Bucket bucket*



```
octl storage bucket acl configure bucket [flags]
```

### Options

```
      --acl string                  The canned ACL to apply to the bucket.
      --from-file string            The file storing the ACL config (JSON format, i.e. {"Grants":[...]})
      --from-string string          The ACL config (JSON format, i.e. {"Grants":[...]})
      --grant-full-control string   Allows grantee the read, write, read ACP, and write ACP permissions on the bucket.
      --grant-read string           Allows grantee to list the objects in the bucket.
      --grant-write string          Allows grantee to create new objects in the bucket.
      --grant-write-acp string      Allows grantee to write the ACL for the applicable bucket.
  -h, --help                        help for configure
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

* [octl storage bucket acl](octl_storage_bucket_acl.md)	 - acl commands

