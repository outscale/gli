## octl iaas api ReadNetAccessPoints

Lists one or more Net access points.

### Synopsis

Lists one or more Net access points.

```
octl iaas api ReadNetAccessPoints [flags]
```

### Options

```
      --DryRun                              If true, checks whether you have the required permissions to perform the action.
      --Filters.NetAccessPointIds strings   The IDs of the Net access points.
      --Filters.NetIds strings              The IDs of the Nets.
      --Filters.ServiceNames strings        The names of the services.
      --Filters.States strings              The states of the Net access points (pending | available | deleting | deleted).
      --Filters.TagKeys strings             The keys of the tags associated with the Net access points.
      --Filters.TagValues strings           The values of the tags associated with the Net access points.
      --Filters.Tags strings                The key/value combination of the tags associated with the Net access points, in the following format: "Filters":{"Tags":["TAGKEY=TAGVALUE"]}.
      --NextPageToken string                The token to request the next page of results.
      --ResultsPerPage int                  The maximum number of logs returned in a single response (between 1 and 1000, both included).
  -h, --help                                help for ReadNetAccessPoints
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

