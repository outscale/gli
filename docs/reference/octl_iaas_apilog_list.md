## octl iaas apilog list

alias for api ReadApiLogs

### Synopsis

> *alias for api ReadApiLogs*

Lists the logs of the API calls you have performed with this account.

**[IMPORTANT]**

Past logs are accessible for up to 32 days.

By default, the retrieved interval is 48 hours. If neither of the `QueryDateBefore` nor `QueryDateAfter` parameters are specified, logs from the past 48 hours are retrieved. If you only specify one of two, logs are retrieved from a 2-day interval based on the date you provided. To retrieve logs beyond a 2-day interval, specify both parameters.

For more information, see [About OMS](https://docs.outscale.com/en/userguide/About-OMS.html).

```
octl iaas apilog list [flags]
```

### Options

```
  -h, --help                        help for list
      --query-access-key strings    The access keys used for the logged calls.
      --query-api-name strings      The names of the APIs of the logged calls (always oapi for the OUTSCALE API).
      --query-call-name strings     The names of the logged calls.
      --query-date-after osctime    The date and time, or the date, after which you want to retrieve logged calls, in ISO 8601 format (for example, 2020-06-14T00:00:00.000Z or 2020-06-14).
      --query-date-before osctime   The date and time, or the date, before which you want to retrieve logged calls, in ISO 8601 format (for example, 2020-06-30T00:00:00.000Z or 2020-06-14).
      --query-ip-address strings    The IPs used for the logged calls.
      --query-user-agent strings    The user agents of the HTTP requests of the logged calls.
      --request-id strings          The request IDs provided in the responses of the logged calls.
      --response-status-code ints   The HTTP status codes of the logged calls.
      --with-account-id             If true, the account ID is displayed.
      --with-call-duration          If true, the duration of the call is displayed.
      --with-query-access-key       If true, the access key is displayed.
      --with-query-api-name         If true, the name of the API is displayed.
      --with-query-api-version      If true, the version of the API is displayed.
      --with-query-call-name        If true, the name of the call is displayed.
      --with-query-date             If true, the date of the call is displayed.
      --with-query-header-raw       If true, the raw header of the HTTP request is displayed.
      --with-query-header-size      If true, the size of the raw header of the HTTP request is displayed.
      --with-query-ip-address       If true, the IP is displayed.
      --with-query-payload-raw      If true, the raw payload of the HTTP request is displayed.
      --with-query-payload-size     If true, the size of the raw payload of the HTTP request is displayed.
      --with-query-user-agent       If true, the user agent of the HTTP request is displayed.
      --with-request-id             If true, the request ID is displayed.
      --with-response-size          If true, the size of the response is displayed.
      --with-response-status-code   If true, the HTTP status code of the response is displayed.
```

### Options inherited from parent commands

```
  -c, --columns string    columns to display - [+]<title>:<jq query for content>||<title>:<jq query for content>
      --config string     Path of profile file (by default, ~/.osc/config.json)
      --filter strings    comma separated list of filters for results - name:value,name:value, alias for jq filter 'select(.name | test("value"))'
      --jq string         jq filter
      --no-upgrade        do not check for new versions
  -O, --out-file string   redirect output to file
  -o, --output string     output format (raw, json, yaml, table, csv, none, base64) (default "raw")
      --profile string    Profile to use in profile file (by default, "default")
      --single            convert single entry lists to a single object
      --template string   JSON template for query body
  -v, --verbose           Verbose output
  -y, --yes               answer yes to all prompts
```

### SEE ALSO

* [octl iaas apilog](octl_iaas_apilog.md)	 - apilog commands

