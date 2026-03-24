## octl iaas loadbalancer create

alias for api CreateLoadBalancer

### Synopsis

> *alias for api CreateLoadBalancer*

Creates a load balancer.

The load balancer is created with a unique Domain Name Service (DNS) name. It receives the incoming traffic and routes it to its registered virtual machines (VMs).

By default, this action creates an Internet-facing load balancer, resolving to public IPs. To create an internal load balancer in a Net, resolving to private IPs, use the `LoadBalancerType` parameter.

You must specify either the `Subnets` or the `SubregionNames` parameters.

For more information, see [About Load Balancers](https://docs.outscale.com/en/userguide/About-Load-Balancers.html).

```
octl iaas loadbalancer create [flags]
```

### Options

```
  -h, --help                                     help for create
      --listener-backend-port int                The port on which the backend VM is listening (between 1 and 65535, both included).
      --listener-backend-protocol string         The protocol for routing traffic to backend VMs (HTTP | HTTPS | TCP | SSL).
      --listener-load-balancer-port int          The port on which the load balancer is listening (between 1 and 65535, both included).
      --listener-load-balancer-protocol string   The routing protocol (HTTP | HTTPS | TCP | SSL).
      --listener-server-certificate-id string    The OUTSCALE Resource Name (ORN) of the server certificate.
      --name string                              The unique name of the load balancer, with a maximum length of 32 alphanumeric characters and dashes (-).
      --public-ip string                         (internet-facing only) The public IP you want to associate with the load balancer.
      --security-group strings                   (Net only) One or more IDs of security groups you want to assign to the load balancer.
      --subnet strings                           (Net only) The ID of the Subnet in which you want to create the load balancer.
      --subregion strings                        (public Cloud only) The Subregion in which you want to create the load balancer.
      --tag-key string                           The key of the tag, between 1 and 255 characters.
      --tag-value string                         The value of the tag, between 0 and 255 characters.
      --type string                              The type of load balancer: internet-facing or internal.
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

* [octl iaas loadbalancer](octl_iaas_loadbalancer.md)	 - loadbalancer commands

