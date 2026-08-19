## octl iaas apilog tail

alias for api ReadApiLogs

### Synopsis

> *alias for api ReadApiLogs*

Lists the logs of the API calls you have performed with this OUTSCALE account.

**[IMPORTANT]**

Past logs are accessible for up to 32 days.

By default, the retrieved interval is 48 hours. If neither of the `QueryDateBefore` nor `QueryDateAfter` parameters are specified, logs from the past 48 hours are retrieved. If you only specify one of two, logs are retrieved from a 2-day interval based on the date you provided. To retrieve logs beyond a 2-day interval, specify both parameters.

For more information, see [About OMS](https://docs.outscale.com/en/userguide/About-OMS.html).

```
octl iaas apilog tail [flags]
```

### Options

```
      --access-key strings   The access keys used for the logged calls.
      --api strings          The names of the APIs of the logged calls (always oapi for the OUTSCALE API).
      --call strings         The names of the logged calls.
  -h, --help                 help for tail
      --ip strings           The IPs used for the logged calls.
      --status ints          The HTTP status codes of the logged calls.
      --user-agent strings   The user agents of the HTTP requests of the logged calls.
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

* [octl iaas apilog](octl_iaas_apilog.md)	 - apilog commands

