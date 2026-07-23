## octl storage object delete

alias for api DeleteObject --Key key

### Synopsis

> *alias for api DeleteObject --Key key*



```
octl storage object delete key [key]... [flags]
```

### Options

```
      --bucket string                         [REQUIRED] The bucket name of the bucket containing the object.
      --bypass-governance-retention           Indicates whether S3 Object Lock should bypass Governance-mode restrictions to process this operation.
      --expected-bucket-owner string          The account ID of the expected bucket owner.
  -h, --help                                  help for delete
      --if-match string                       The If-Match header field makes the request method conditional on ETags.
      --if-match-last-modified-time osctime   If present, the object is deleted only if its modification times matches the provided Timestamp .
      --if-match-size int                     If present, the object is deleted only if its size matches the provided size in bytes.
      --mfa string                            The concatenation of the authentication device's serial number, a space, and the value that is displayed on your authentication device.
      --version-id string                     Version ID used to reference a specific version of the object.
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

* [octl storage object](octl_storage_object.md)	 - object commands

