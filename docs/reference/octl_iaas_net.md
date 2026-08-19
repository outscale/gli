## octl iaas net

net commands

### Options

```
  -h, --help   help for net
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

* [octl iaas](octl_iaas.md)	 - OUTSCALE IaaS management
* [octl iaas net create](octl_iaas_net_create.md)	 - alias for api CreateNet
* [octl iaas net delete](octl_iaas_net_delete.md)	 - alias for api DeleteNet --NetId net_id
* [octl iaas net dependencies](octl_iaas_net_dependencies.md)	 - Shows all dependencies of a net
* [octl iaas net describe](octl_iaas_net_describe.md)	 - alias for api ReadNets --Filters.NetIds net_id
* [octl iaas net list](octl_iaas_net_list.md)	 - alias for api ReadNets
* [octl iaas net teardown](octl_iaas_net_teardown.md)	 - Tears down a net and its subresources
* [octl iaas net update](octl_iaas_net_update.md)	 - alias for api UpdateNet --NetId net_id

