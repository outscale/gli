## octl iaas loadbalancer update

alias for api UpdateLoadBalancer --LoadBalancerName load_balancer_name

### Synopsis

> *alias for api UpdateLoadBalancer --LoadBalancerName load_balancer_name*

Modifies the specified attribute of a load balancer. You can specify only one attribute at a time.

You can set a new SSL certificate to an SSL or HTTPS listener of a load balancer.

This certificate replaces any certificate used on the same load balancer and port.

You can also replace the currently enabled policy for the load balancer with another one.

If the `PolicyNames` parameter is empty, the currently enabled policy is disabled.

```
octl iaas loadbalancer update load_balancer_name [load_balancer_name]... [flags]
```

### Options

```
      --access-log-is-enabled                  If true, access logs are enabled for your load balancer.
      --access-log-osu-bucket-name string      The name of the OOS bucket for the access logs.
      --access-log-osu-bucket-prefix string    The path to the folder of the access logs in your OOS bucket (by default, the root level of your bucket).
      --access-log-publication-interval int    The time interval for the publication of access logs in the OOS bucket, in minutes.
      --health-check-check-interval int        The number of seconds between two requests (between 5 and 600 both included).
      --health-check-healthy-threshold int     The number of consecutive successful requests before considering the VM as healthy (between 2 and 10 both included).
      --health-check-path string               If you use the HTTP or HTTPS protocols, the request URL path.
      --health-check-port int                  The port number (between 1 and 65535, both included).
      --health-check-protocol string           The protocol for the URL of the VM (HTTP | HTTPS | TCP | SSL).
      --health-check-timeout int               The maximum waiting time for a response before considering the VM as unhealthy, in seconds (between 2 and 60 both included).
      --health-check-unhealthy-threshold int   The number of consecutive failed requests before considering the VM as unhealthy (between 2 and 10 both included).
  -h, --help                                   help for update
      --policy-name strings                    The name of the policy you want to enable for the listener.
      --port int                               The port on which the load balancer is listening (between 1 and 65535, both included).
      --public-ip string                       (internet-facing only) The public IP you want to associate with the load balancer.
      --secured-cookie                         If true, secure cookies are enabled for the load balancer.
      --security-group strings                 (Net only) One or more IDs of security groups you want to assign to the load balancer.
      --server-certificate-id string           The OUTSCALE Resource Name (ORN) of the server certificate.
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

