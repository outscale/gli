## octl iaas api UpdateRoute

Replaces an existing route within a route table in a Net.

### Synopsis

Replaces an existing route within a route table in a Net.

You must specify one of the following elements as the target:

* Net peering

* NAT virtual machine (VM)

* Internet service

* Virtual gateway

* NAT service

* Network interface card (NIC)

The routing algorithm is based on the most specific match.

```
octl iaas api UpdateRoute [flags]
```

### Options

```
      --DestinationIpRange string   The IP range used for the destination match, in CIDR notation (for example, 10.0.0.0/24).
      --DryRun                      If true, checks whether you have the required permissions to perform the action.
      --GatewayId string            The ID of an internet service or virtual gateway attached to your Net.
      --NatServiceId string         The ID of a NAT service.
      --NetPeeringId string         The ID of a Net peering.
      --NicId string                The ID of a network interface card (NIC).
      --RouteTableId string         The ID of the route table.
      --VmId string                 The ID of a NAT VM in your Net.
  -h, --help                        help for UpdateRoute
```

### Options inherited from parent commands

```
  -c, --columns string              columns to display - [+]<title>:<jq query for content>||<title>:<jq query for content>
      --config string               Path of profile file (by default, ~/.osc/config.json)
      --filter strings              comma separated list of filters for results - name:value,name:value, alias for jq filter 'select(.name | test("value"))'
      --jq string                   jq filter
      --no-upgrade                  do not check for new versions
  -O, --out-file string             redirect output to file
  -o, --output string               output format (raw, json, yaml, table, csv, none, base64, text) (default "raw")
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

