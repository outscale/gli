## octl iaas api UnlinkRouteTable

Disassociates a Subnet from a route table.

### Synopsis

Disassociates a Subnet from a route table.

After disassociation, the Subnet can no longer use the routes in this route table, but uses the routes in the main route table of the Net instead.

```
octl iaas api UnlinkRouteTable [flags]
```

### Options

```
      --DryRun                    If true, checks whether you have the required permissions to perform the action.
      --LinkRouteTableId string   The ID of the association between the route table and the Subnet.
  -h, --help                      help for UnlinkRouteTable
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

