## octl iaas loadbalancerlistener create

alias for api CreateLoadBalancerListeners

### Synopsis

> *alias for api CreateLoadBalancerListeners*

Creates one or more listeners for a specified load balancer.

For more information, see [About Load Balancers](https://docs.outscale.com/en/userguide/About-Load-Balancers.html).

```
octl iaas loadbalancerlistener create [flags]
```

### Options

```
  -h, --help                                     help for create
      --listener-backend-port int                The port on which the backend VM is listening (between 1 and 65535, both included).
      --listener-backend-protocol string         The protocol for routing traffic to backend VMs (HTTP | HTTPS | TCP | SSL).
      --listener-load-balancer-port int          The port on which the load balancer is listening (between 1 and 65535, both included).
      --listener-load-balancer-protocol string   The routing protocol (HTTP | HTTPS | TCP | SSL).
      --listener-server-certificate-id string    The OUTSCALE Resource Name (ORN) of the server certificate.
      --name string                              The name of the load balancer for which you want to create listeners.
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

* [octl iaas loadbalancerlistener](octl_iaas_loadbalancerlistener.md)	 - loadbalancerlistener commands

