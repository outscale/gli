## octl storage bucket create

alias for api CreateBucket

### Synopsis

> *alias for api CreateBucket*



```
octl storage bucket create [flags]
```

### Options

```
      --acl string                       The canned ACL to apply to the bucket.
      --bucket string                    The name of the bucket to create.
      --grant-full-control string        Allows grantee the read, write, read ACP, and write ACP permissions on the bucket.
      --grant-read string                Allows grantee to list the objects in the bucket.
      --grant-read-acp string            Allows grantee to read the bucket ACL.
      --grant-write string               Allows grantee to create new objects in the bucket.
      --grant-write-acp string           Allows grantee to write the ACL for the applicable bucket.
  -h, --help                             help for create
      --namespace string                 Specifies the namespace where you want to create your general purpose bucket.
      --object-lock-enabled-for-bucket   Specifies whether you want S3 Object Lock to be enabled for the new bucket.
      --object-ownership string          The container element for object ownership for a bucket's ownership controls.
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
      --payload string    JSON content for query body
      --profile string    Profile to use in profile file (by default, "default")
      --single            convert single entry lists to a single object
      --template string   JSON template file for query body
  -v, --verbose           Verbose output
  -y, --yes               answer yes to all prompts
```

### SEE ALSO

* [octl storage bucket](octl_storage_bucket.md)	 - bucket commands

