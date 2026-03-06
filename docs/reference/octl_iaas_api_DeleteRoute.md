## octl iaas api DeleteRoute

Deletes a route from a specified route table.

### Synopsis

Deletes a route from a specified route table.

```
octl iaas api DeleteRoute [flags]
```

### Options

```
      --DestinationIpRange string   The exact IP range for the route.
      --DryRun                      If true, checks whether you have the required permissions to perform the action.
      --RouteTableId string         The ID of the route table from which you want to delete a route.
  -h, --help                        help for DeleteRoute
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

* [octl iaas api](octl_iaas_api.md)	 - iaas api calls

