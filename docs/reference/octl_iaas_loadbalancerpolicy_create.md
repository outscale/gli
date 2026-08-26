## octl iaas loadbalancerpolicy create

alias for api CreateLoadBalancerPolicy

### Synopsis

> alias for api CreateLoadBalancerPolicy

Creates a stickiness policy with sticky session lifetimes defined by the browser lifetime.

The created policy can be used with HTTP or HTTPS listeners only.

If this policy is implemented by a load balancer, this load balancer uses this cookie in all incoming requests to direct them to the specified backend server virtual machine (VM). If this cookie is not present, the load balancer sends the request to any other server according to its load-balancing algorithm.



You can also create a stickiness policy with sticky session lifetimes following the lifetime of an application-generated cookie.

Unlike the other type of stickiness policy, the lifetime of the special Load Balancer Unit (LBU) cookie follows the lifetime of the application-generated cookie specified in the policy configuration. The load balancer inserts a new stickiness cookie only when the application response includes a new application cookie.

The session stops being sticky if the application cookie is removed or expires, until a new application cookie is issued.


For more information, see [About Load Balancers](https://docs.outscale.com/en/userguide/About-Load-Balancers.html).

```
octl iaas loadbalancerpolicy create [flags]
```

### Options

```
      --cookie-expiration-period int   The lifetime of the cookie, in seconds.
      --cookie-name string             The name of the application cookie used for stickiness, between 1 and 255 characters.
  -h, --help                           help for create
      --name string                    [REQUIRED] The name of the load balancer for which you want to create a policy.
      --policy-name string             [REQUIRED] The unique name of the policy, with a maximum length of 32 alphanumeric characters and dashes (-).
      --policy-type string             [REQUIRED] The type of stickiness policy you want to create: app or load_balancer.
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
  -o, --output string              output format (raw, json, yaml, table, csv, none, base64, text)
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

* [octl iaas loadbalancerpolicy](octl_iaas_loadbalancerpolicy.md)	 - loadbalancerpolicy commands

