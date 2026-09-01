## octl storage object

Manage Object resources

### Options

```
  -h, --help   help for object
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

* [octl storage](octl_storage.md)	 - OUTSCALE Object Storage (OOS) management
* [octl storage object acl](octl_storage_object_acl.md)	 - acl commands
* [octl storage object copy](octl_storage_object_copy.md)	 - Copy an object.
* [octl storage object delete](octl_storage_object_delete.md)	 - 
* [octl storage object describe](octl_storage_object_describe.md)	 - Display an object metadata.
* [octl storage object download](octl_storage_object_download.md)	 - Download an object.
* [octl storage object list](octl_storage_object_list.md)	 - 
* [octl storage object presign](octl_storage_object_presign.md)	 - Create a pre-signed URL
* [octl storage object put](octl_storage_object_put.md)	 - 
* [octl storage object retention](octl_storage_object_retention.md)	 - retention commands
* [octl storage object tagging](octl_storage_object_tagging.md)	 - tagging commands
* [octl storage object versions](octl_storage_object_versions.md)	 - Display an object metadata.

