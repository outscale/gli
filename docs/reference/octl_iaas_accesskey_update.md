## octl iaas accesskey update

Modifies the attributes of the specified access key of either the root user or an EIM user.

### Synopsis

Modifies the attributes of the specified access key of either the root user or an EIM user.


The parameter `ExpirationDate` is not required when updating the state of your access key. However, if you do not specify the expiration date of an access key when updating its state, it is then set to not expire.

> alias for UpdateAccessKey --AccessKeyId access_key_id

```
octl iaas accesskey update access_key_id [access_key_id]... [flags]
```

### Options

```
      --clear-expiration-date    If true, the current expiration date is deleted and the access key is set to not expire.
      --clear-tag                If true, the current tag of the access key is deleted.
      --expiration-date string   The date and time, or the date, at which you want the access key to expire, in ISO 8601 format (for example, 2020-06-14T00:00:00.000Z or 2020-06-14).
  -h, --help                     help for update
      --state string             The new state for the access key (ACTIVE | INACTIVE).
      --tag string               A new tag to add to the access key.
      --user-name string         The name of the EIM user that the access key you want to modify is associated with.
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

* [octl iaas accesskey](octl_iaas_accesskey.md)	 - Manage AccessKey resources

