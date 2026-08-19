## octl iaas apiaccessrule delete

alias for api DeleteApiAccessRule --ApiAccessRuleId api_access_rule_id

### Synopsis

> *alias for api DeleteApiAccessRule --ApiAccessRuleId api_access_rule_id*

Deletes a specified API access rule.

**[IMPORTANT]**

You cannot delete the last remaining API access rule. However, if you delete all the API access rules that allow you to access the APIs, you need to contact the Support team to regain access. For more information, see [Technical Support](https://docs.outscale.com/en/userguide/Technical-Support.html).

```
octl iaas apiaccessrule delete api_access_rule_id [api_access_rule_id]... [flags]
```

### Options

```
  -h, --help   help for delete
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

* [octl iaas apiaccessrule](octl_iaas_apiaccessrule.md)	 - apiaccessrule commands

