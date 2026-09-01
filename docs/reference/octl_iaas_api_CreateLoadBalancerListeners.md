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
      --LoadBalancerName string                   [REQUIRED] The name of the load balancer for which you want to create listeners.
  -h, --help                                      help for CreateLoadBalancerListeners
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

