## octl iaas api DeleteVirtualGateway

Deletes a specified virtual gateway.

### Synopsis

Deletes a specified virtual gateway.

**[IMPORTANT]**

Before deleting a virtual gateway, we recommend detaching it from any associated Net, DirectLink, and DirectLink interface, and deleting the VPN connection.

```
octl iaas api DeleteVirtualGateway [flags]
```

### Options

```
      --DryRun                    If true, checks whether you have the required permissions to perform the action.
      --VirtualGatewayId string   The ID of the virtual gateway you want to delete.
  -h, --help                      help for DeleteVirtualGateway
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
      --payload string    JSON content for query body
      --profile string    Profile to use in profile file (by default, "default")
      --single            convert single entry lists to a single object
      --template string   JSON template file for query body
  -v, --verbose           Verbose output
  -y, --yes               answer yes to all prompts
```

### SEE ALSO

* [octl iaas api](octl_iaas_api.md)	 - iaas api calls

