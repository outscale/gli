## octl iaas image update

alias for api UpdateImage --ImageId image_id

### Synopsis

> *alias for api UpdateImage --ImageId image_id*

Modifies the access permissions for an OUTSCALE machine image (OMI).

You must specify either the `Additions` or the `Removals` parameter.

After sharing an OMI with an OUTSCALE account, the other account can create a copy of it that they own. For more information about copying OMIs, see [CreateImage](#createimage).

```
octl iaas image update image_id [image_id]... [flags]
```

### Options

```
      --description string                                 A new description for the image.
  -h, --help                                               help for update
      --permission-to-launch-addition-account-id strings   One or more OUTSCALE account IDs that the permission is associated with.
      --permission-to-launch-addition-global-permission    A global permission for all accounts.
      --permission-to-launch-removal-account-id strings    One or more OUTSCALE account IDs that the permission is associated with.
      --permission-to-launch-removal-global-permission     A global permission for all accounts.
      --product-code strings                               The product codes associated with the OMI.
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
      --no-upgrade                 do not check for new versions
  -O, --out-file string            redirect output to file
  -o, --output string              output format (raw, json, yaml, table, csv, none, base64, text)
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

* [octl iaas image](octl_iaas_image.md)	 - image commands

