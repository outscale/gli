## octl iaas api UpdateLoadBalancer

Modifies the specified attribute of a load balancer.

### Synopsis

Modifies the specified attribute of a load balancer. You can specify only one attribute at a time.

You can set a new SSL certificate to an SSL or HTTPS listener of a load balancer.

This certificate replaces any certificate used on the same load balancer and port.

You can also replace the currently enabled policy for the load balancer with another one.

If the `PolicyNames` parameter is empty, the currently enabled policy is disabled.

```
octl iaas api UpdateLoadBalancer [flags]
```

### Options

```
      --AccessLog.IsEnabled                  If true, access logs are enabled for your load balancer.
      --AccessLog.OsuBucketName string       The name of the OOS bucket for the access logs.
      --AccessLog.OsuBucketPrefix string     The path to the folder of the access logs in your OOS bucket (by default, the root level of your bucket).
      --AccessLog.PublicationInterval int    The time interval for the publication of access logs in the OOS bucket, in minutes.
      --DryRun                               If true, checks whether you have the required permissions to perform the action.
      --HealthCheck.CheckInterval int        The number of seconds between two requests (between 5 and 600 both included).
      --HealthCheck.HealthyThreshold int     The number of consecutive successful requests before considering the VM as healthy (between 2 and 10 both included).
      --HealthCheck.Path string              If you use the HTTP or HTTPS protocols, the request URL path.
      --HealthCheck.Port int                 The port number (between 1 and 65535, both included).
      --HealthCheck.Protocol string          The protocol for the URL of the VM (HTTP | HTTPS | TCP | SSL).
      --HealthCheck.Timeout int              The maximum waiting time for a response before considering the VM as unhealthy, in seconds (between 2 and 60 both included).
      --HealthCheck.UnhealthyThreshold int   The number of consecutive failed requests before considering the VM as unhealthy (between 2 and 10 both included).
      --LoadBalancerName string              The name of the load balancer.
      --LoadBalancerPort int                 The port on which the load balancer is listening (between 1 and 65535, both included).
      --PolicyNames strings                  The name of the policy you want to enable for the listener.
      --PublicIp string                      (internet-facing only) The public IP you want to associate with the load balancer.
      --SecuredCookies                       If true, secure cookies are enabled for the load balancer.
      --SecurityGroups strings               (Net only) One or more IDs of security groups you want to assign to the load balancer.
      --ServerCertificateId string           The OUTSCALE Resource Name (ORN) of the server certificate.
  -h, --help                                 help for UpdateLoadBalancer
```

### Options inherited from parent commands

```
  -c, --columns string              columns to display - [+]<title>:<jq query for content>||<title>:<jq query for content>
      --config string               Path of profile file (by default, ~/.osc/config.json)
      --filter strings              comma separated list of filters for results - name:value,name:value, alias for jq filter 'select(.name | test("value"))'
      --jq string                   jq filter
      --no-upgrade                  do not check for new versions
  -O, --out-file string             redirect output to file
  -o, --output string               output format (raw, json, yaml, table, csv, none, base64) (default "raw")
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

