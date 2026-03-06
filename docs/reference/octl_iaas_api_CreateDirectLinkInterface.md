## octl iaas api CreateDirectLinkInterface

Creates a DirectLink interface.

### Synopsis

Creates a DirectLink interface.

DirectLink interfaces enable you to reach one of your Nets through a virtual gateway.

For more information, see [About DirectLink](https://docs.outscale.com/en/userguide/About-DirectLink.html).

```
octl iaas api CreateDirectLinkInterface [flags]
```

### Options

```
      --DirectLinkId string                                  The ID of the existing DirectLink for which you want to create the DirectLink interface.
      --DirectLinkInterface.BgpAsn int                       
      --DirectLinkInterface.BgpKey string                    
      --DirectLinkInterface.ClientPrivateIp string           
      --DirectLinkInterface.DirectLinkInterfaceName string   
      --DirectLinkInterface.OutscalePrivateIp string         
      --DirectLinkInterface.VirtualGatewayId string          
      --DirectLinkInterface.Vlan int                         
      --DryRun                                               If true, checks whether you have the required permissions to perform the action.
  -h, --help                                                 help for CreateDirectLinkInterface
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

