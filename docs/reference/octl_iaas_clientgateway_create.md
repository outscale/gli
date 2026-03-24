## octl iaas clientgateway create

alias for api CreateClientGateway

### Synopsis

> *alias for api CreateClientGateway*

Provides information about your client gateway.

This action registers information to identify the client gateway that you deployed in your network.

To open a tunnel to the client gateway, you must provide the communication protocol type, the fixed public IP of the gateway, and an Autonomous System Number (ASN).

For more information, see [About Client Gateways](https://docs.outscale.com/en/userguide/About-Client-Gateways.html).

```
octl iaas clientgateway create [flags]
```

### Options

```
      --bgp-asn int              The Autonomous System Number (ASN) used by the Border Gateway Protocol (BGP) to find the path to your client gateway through the Internet.
      --connection-type string   The communication protocol used to establish tunnel with your client gateway (always ipsec.1).
  -h, --help                     help for create
      --public-ip string         The public fixed IPv4 address of your client gateway.
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

* [octl iaas clientgateway](octl_iaas_clientgateway.md)	 - clientgateway commands

