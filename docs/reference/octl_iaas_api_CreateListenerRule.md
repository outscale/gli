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
      --Listener.LoadBalancerName string       [REQUIRED] The name of the load balancer to which the listener is attached.
      --Listener.LoadBalancerPort int          [REQUIRED] The port of load balancer on which the load balancer is listening (between 1 and 65535 both included).
      --ListenerRule.Action string             The type of action for the rule (always forward).
      --ListenerRule.HostNamePattern string    A host-name pattern for the rule, with a maximum length of 128 characters.
      --ListenerRule.ListenerRuleName string   [REQUIRED] A human-readable name for the listener rule.
      --ListenerRule.PathPattern string        A path pattern for the rule, with a maximum length of 128 characters.
      --ListenerRule.Priority int              [REQUIRED] The priority level of the listener rule, between 1 and 19999 both included.
      --VmIds strings                          The IDs of the backend VMs.
  -h, --help                                   help for CreateListenerRule
```

### Options inherited from parent commands

```
  -c, --columns string             columns to display - [+]<title>:<jq query for content>||<title>:<jq query for content>
      --config string              Path of profile file (by default, ~/.osc/config.json)
      --dry-run                    Display the request payload that would be sent to the API without sending it
      --elapsed                    add elapsed time column when using --watch (default true)
      --filter strings             comma separated list of filters for results - name:value,name:value, alias for jq filter 'select(.name | tostring | test("value"))'
      --interval duration          interval between two watch/waitfor iterations (default 5s)
      --jq string                  jq filter
      --max-pages int              maximum number of pages a command can fetch (default 20)
      --no-upgrade                 do not check for new versions
  -O, --out-file string            redirect output to file
  -o, --output string              output format (raw, json, yaml, table, csv, none, text)
      --payload string             JSON content for query body
      --profile string             Profile to use in profile file (by default, "default")
  -s, --silent                     Hides all information messages
      --single                     convert single entry lists to a single object
      --template string            JSON template file for query body
  -v, --verbose                    Verbose output
      --waitfor string             repeatedly call the API until the specified jq expression returns 1/true or a non empty result
      --waitfor-timeout duration   maximum duration of a wait (default 10m0s)
      --watch                      repeatedly call the API and display changes
  -y, --yes                        answer yes to all prompts
```

### SEE ALSO

* [octl iaas api](octl_iaas_api.md)	 - Call iaas API

