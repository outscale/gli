## octl iaas api ReadApiLogs

Lists the logs of the API calls you have performed with this OUTSCALE account.

### Synopsis

Lists the logs of the API calls you have performed with this OUTSCALE account.

**[IMPORTANT]**

Past logs are accessible for up to 32 days.

By default, the retrieved interval is 48 hours. If neither of the `QueryDateBefore` nor `QueryDateAfter` parameters are specified, logs from the past 48 hours are retrieved. If you only specify one of two, logs are retrieved from a 2-day interval based on the date you provided. To retrieve logs beyond a 2-day interval, specify both parameters.

For more information, see [About OMS](https://docs.outscale.com/en/userguide/About-OMS.html).

```
octl iaas api ReadApiLogs [flags]
```

### Options

```
      --DryRun                             If true, checks whether you have the required permissions to perform the action.
      --Filters.QueryAccessKeys strings    The access keys used for the logged calls.
      --Filters.QueryApiNames strings      The names of the APIs of the logged calls (always oapi for the OUTSCALE API).
      --Filters.QueryCallNames strings     The names of the logged calls.
      --Filters.QueryDateAfter osctime     The date and time, or the date, after which you want to retrieve logged calls, in ISO 8601 format (for example, 2020-06-14T00:00:00.000Z or 2020-06-14).
      --Filters.QueryDateBefore osctime    The date and time, or the date, before which you want to retrieve logged calls, in ISO 8601 format (for example, 2020-06-30T00:00:00.000Z or 2020-06-14).
      --Filters.QueryIpAddresses strings   The IPs used for the logged calls.
      --Filters.QueryUserAgents strings    The user agents of the HTTP requests of the logged calls.
      --Filters.RequestIds strings         The request IDs provided in the responses of the logged calls.
      --Filters.ResponseStatusCodes ints   The HTTP status codes of the logged calls.
      --NextPageToken string               The token to request the next page of results.
      --ResultsPerPage int                 The maximum number of logs returned in a single response (between 1 and 1000, both included).
      --With.AccountId                     If true, the OUTSCALE account ID is displayed.
      --With.CallDuration                  If true, the duration of the call is displayed.
      --With.QueryAccessKey                If true, the access key is displayed.
      --With.QueryApiName                  If true, the name of the API is displayed.
      --With.QueryApiVersion               If true, the version of the API is displayed.
      --With.QueryCallName                 If true, the name of the call is displayed.
      --With.QueryDate                     If true, the date of the call is displayed.
      --With.QueryHeaderRaw                If true, the raw header of the HTTP request is displayed.
      --With.QueryHeaderSize               If true, the size of the raw header of the HTTP request is displayed.
      --With.QueryIpAddress                If true, the IP is displayed.
      --With.QueryPayloadRaw               If true, the raw payload of the HTTP request is displayed.
      --With.QueryPayloadSize              If true, the size of the raw payload of the HTTP request is displayed.
      --With.QueryUserAgent                If true, the user agent of the HTTP request is displayed.
      --With.RequestId                     If true, the request ID is displayed.
      --With.ResponseSize                  If true, the size of the response is displayed.
      --With.ResponseStatusCode            If true, the HTTP status code of the response is displayed.
  -h, --help                               help for ReadApiLogs
```

### Options inherited from parent commands

```
  -c, --columns string              columns to display - [+]<title>:<jq query for content>||<title>:<jq query for content>
      --config string               Path of profile file (by default, ~/.osc/config.json)
      --filter strings              comma separated list of filters for results - name:value,name:value, alias for jq filter 'select(.name | test("value"))'
      --jq string                   jq filter
      --no-upgrade                  do not check for new versions
  -O, --out-file string             redirect output to file
  -o, --output string               output format (raw, json, yaml, table, csv, none, base64) (default "raw")
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

* [octl iaas api](octl_iaas_api.md)	 - iaas api calls

