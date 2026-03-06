## octl iaas clientgateway

clientgateway commands

### Options

```
  -h, --help   help for clientgateway
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

* [octl iaas](octl_iaas.md)	 - OUTSCALE IaaS management
* [octl iaas clientgateway create](octl_iaas_clientgateway_create.md)	 - alias for api CreateClientGateway
* [octl iaas clientgateway delete](octl_iaas_clientgateway_delete.md)	 - alias for api DeleteClientGateway --ClientGatewayId client_gateway_id
* [octl iaas clientgateway describe](octl_iaas_clientgateway_describe.md)	 - alias for api ReadClientGateways --Filters.ClientGatewayIds client_gateway_id
* [octl iaas clientgateway list](octl_iaas_clientgateway_list.md)	 - alias for api ReadClientGateways

