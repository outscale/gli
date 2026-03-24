## octl iaas api CreateVirtualGateway

Creates a virtual gateway.

### Synopsis

Creates a virtual gateway.

A virtual gateway is the access point on the Net side of a VPN connection.

For more information, see [About Virtual Gateways](https://docs.outscale.com/en/userguide/About-Virtual-Gateways.html).

```
octl iaas api CreateVirtualGateway [flags]
```

### Options

```
      --ConnectionType string   The type of VPN connection supported by the virtual gateway (always ipsec.1).
      --DryRun                  If true, checks whether you have the required permissions to perform the action.
  -h, --help                    help for CreateVirtualGateway
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

