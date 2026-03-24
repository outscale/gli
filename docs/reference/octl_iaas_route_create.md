## octl iaas route create

alias for api CreateRoute

### Synopsis

> *alias for api CreateRoute*

Creates a route in a specified route table within a specified Net.

You must specify one of the following elements as the target:

* Net peering

* NAT VM

* Internet service

* Virtual gateway

* NAT service

* Network interface card (NIC)

The routing algorithm is based on the most specific match.

For more information, see [About Route Tables](https://docs.outscale.com/en/userguide/About-Route-Tables.html).

```
octl iaas route create [flags]
```

### Options

```
      --destination-ip-range string   The IP range used for the destination match, in CIDR notation (for example, 10.0.0.0/24).
      --gateway-id string             The ID of an internet service or virtual gateway attached to your Net.
  -h, --help                          help for create
      --id string                     The ID of the route table for which you want to create a route.
      --nat-service-id string         The ID of a NAT service.
      --net-peering-id string         The ID of a Net peering.
      --nic-id string                 The ID of a NIC.
      --vm-id string                  The ID of a NAT VM in your Net (attached to exactly one NIC).
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

* [octl iaas route](octl_iaas_route.md)	 - route commands

