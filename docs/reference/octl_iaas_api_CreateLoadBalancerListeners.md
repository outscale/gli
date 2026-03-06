## octl iaas api CreateLoadBalancerListeners

Creates one or more listeners for a specified load balancer.

### Synopsis

Creates one or more listeners for a specified load balancer.

For more information, see [About Load Balancers](https://docs.outscale.com/en/userguide/About-Load-Balancers.html).

```
octl iaas api CreateLoadBalancerListeners [flags]
```

### Options

```
      --DryRun                                    If true, checks whether you have the required permissions to perform the action.
      --Listeners.0.BackendPort int               The port on which the backend VM is listening (between 1 and 65535, both included).
      --Listeners.0.BackendProtocol string        The protocol for routing traffic to backend VMs (HTTP | HTTPS | TCP | SSL).
      --Listeners.0.LoadBalancerPort int          The port on which the load balancer is listening (between 1 and 65535, both included).
      --Listeners.0.LoadBalancerProtocol string   The routing protocol (HTTP | HTTPS | TCP | SSL).
      --Listeners.0.ServerCertificateId string    The OUTSCALE Resource Name (ORN) of the server certificate.
      --LoadBalancerName string                   The name of the load balancer for which you want to create listeners.
  -h, --help                                      help for CreateLoadBalancerListeners
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

