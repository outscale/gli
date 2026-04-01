## octl storage object tagging

tagging commands

### Options

```
  -h, --help   help for tagging
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

* [octl storage object](octl_storage_object.md)	 - object commands
* [octl storage object tagging configure](octl_storage_object_tagging_configure.md)	 - Add tags to an object, alias for api PutObjectTagging --Key key
* [octl storage object tagging delete](octl_storage_object_tagging_delete.md)	 - Delete all tags from an object, alias for api DeleteObjectTagging --Key key
* [octl storage object tagging describe](octl_storage_object_tagging_describe.md)	 - List object tags, alias for api GetObjectTagging --Key key

