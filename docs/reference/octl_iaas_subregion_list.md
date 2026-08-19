## octl iaas subregion list

alias for api ReadSubregions

### Synopsis

> *alias for api ReadSubregions*

Lists one or more of the enabled Subregions that you can access in the current Region.

For more information, see [About Regions and Subregions](https://docs.outscale.com/en/userguide/About-Regions-and-Subregions.html).

```
octl iaas subregion list [flags]
```

### Options

```
  -h, --help                  help for list
      --region-name strings   The names of the Regions containing the Subregions.
      --state strings         The states of the Subregions.
      --subregion strings     The names of the Subregions.
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

* [octl iaas subregion](octl_iaas_subregion.md)	 - subregion commands

