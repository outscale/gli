## octl iaas vmsstophistory list

alias for api ReadVmsStopHistory

### Synopsis

> alias for api ReadVmsStopHistory

Lists the stop history of one or more VMs.

```
octl iaas vmsstophistory list [flags]
```

### Options

```
  -h, --help                       help for list
      --state-reason strings       The reason explaining why the VM stopped.
      --stop-date-after osctime    The date and time (UTC), or the date, after which you want to retrieve VM stops, in ISO 8601 format (for example, 2026-06-14T00:00:00.000Z or 2026-06-14).
      --stop-date-before osctime   The date and time (UTC), or the date, before which you want to retrieve VM stops, in ISO 8601 format (for example, 2026-06-14T00:00:00.000Z or 2026-06-14).
      --vm-id strings              The IDs of the stopped VM(s).
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

* [octl iaas vmsstophistory](octl_iaas_vmsstophistory.md)	 - vmsstophistory commands

