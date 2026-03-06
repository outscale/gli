## octl iaas api CreateListenerRule

Creates a rule for traffic redirection for the specified listener.

### Synopsis

Creates a rule for traffic redirection for the specified listener. Each rule must have either the `HostNamePattern` or `PathPattern` parameter specified. Rules are treated in priority order, from the highest value to the lowest value.

Once the rule is created, you need to register backend VMs with it. For more information, see the [RegisterVmsInLoadBalancer](#registervmsinloadbalancer) method.

For more information, see [About Load Balancers](https://docs.outscale.com/en/userguide/About-Load-Balancers.html).

```
octl iaas api CreateListenerRule [flags]
```

### Options

```
      --DryRun                                 If true, checks whether you have the required permissions to perform the action.
      --Listener.LoadBalancerName string       The name of the load balancer to which the listener is attached.
      --Listener.LoadBalancerPort int          The port of load balancer on which the load balancer is listening (between 1 and 65535 both included).
      --ListenerRule.Action string             The type of action for the rule (always forward).
      --ListenerRule.HostNamePattern string    A host-name pattern for the rule, with a maximum length of 128 characters.
      --ListenerRule.ListenerRuleName string   A human-readable name for the listener rule.
      --ListenerRule.PathPattern string        A path pattern for the rule, with a maximum length of 128 characters.
      --ListenerRule.Priority int              The priority level of the listener rule, between 1 and 19999 both included.
      --VmIds strings                          The IDs of the backend VMs.
  -h, --help                                   help for CreateListenerRule
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

