## octl iaas policyversion

Manage PolicyVersion resources

### Options

```
  -h, --help   help for policyversion
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

* [octl iaas](octl_iaas.md)	 - OUTSCALE IaaS management
* [octl iaas policyversion create](octl_iaas_policyversion_create.md)	 - Creates a version of a specified managed policy.
* [octl iaas policyversion delete](octl_iaas_policyversion_delete.md)	 - Deletes a specified version of a managed policy, if it is not set as the default one.
* [octl iaas policyversion describe](octl_iaas_policyversion_describe.md)	 - Lists information about a specified version of a managed policy.
* [octl iaas policyversion list](octl_iaas_policyversion_list.md)	 - Lists information about all the policy versions of a specified managed policy.
* [octl iaas policyversion set-default](octl_iaas_policyversion_set-default.md)	 - Sets a specified version of a managed policy as the default (operative) one.

