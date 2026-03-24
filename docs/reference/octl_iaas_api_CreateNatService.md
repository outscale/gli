## octl iaas api CreateNatService

Creates a network address translation (NAT) service in the specified public Subnet of a Net.

### Synopsis

Creates a network address translation (NAT) service in the specified public Subnet of a Net.

A NAT service enables virtual machines (VMs) placed in the private Subnet of this Net to connect to the Internet, without being accessible from the Internet.

When creating a NAT service, you specify the allocation ID of the public IP you want to use as public IP for the NAT service. Once the NAT service is created, you need to create a route in the route table of the private Subnet, with 0.0.0.0/0 as destination and the ID of the NAT service as target. For more information, see [LinkPublicIP](#linkpublicip) and [CreateRoute](#createroute).

This action also enables you to create multiple NAT services in the same Net (one per public Subnet).

**[IMPORTANT]**

You cannot modify the public IP associated with a NAT service after its creation. To do so, you need to delete the NAT service and create a new one with another public IP.

For more information, see [About NAT Services](https://docs.outscale.com/en/userguide/About-NAT-Services.html).

```
octl iaas api CreateNatService [flags]
```

### Options

```
      --ClientToken string   A unique identifier which enables you to manage the idempotency.
      --DryRun               If true, checks whether you have the required permissions to perform the action.
      --PublicIpId string    The allocation ID of the public IP to associate with the NAT service.
      --SubnetId string      The ID of the Subnet in which you want to create the NAT service.
  -h, --help                 help for CreateNatService
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

