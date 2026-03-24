## octl iaas loadbalancer

loadbalancer commands

### Options

```
  -h, --help   help for loadbalancer
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

* [octl iaas](octl_iaas.md)	 - OUTSCALE IaaS management
* [octl iaas loadbalancer backends](octl_iaas_loadbalancer_backends.md)	 - alias for api ReadVmsHealth --LoadBalancerName load_balancer_name
* [octl iaas loadbalancer create](octl_iaas_loadbalancer_create.md)	 - alias for api CreateLoadBalancer
* [octl iaas loadbalancer delete](octl_iaas_loadbalancer_delete.md)	 - alias for api DeleteLoadBalancer --LoadBalancerName load_balancer_name
* [octl iaas loadbalancer describe](octl_iaas_loadbalancer_describe.md)	 - alias for api ReadLoadBalancers --Filters.LoadBalancerNames load_balancer_name
* [octl iaas loadbalancer list](octl_iaas_loadbalancer_list.md)	 - alias for api ReadLoadBalancers
* [octl iaas loadbalancer update](octl_iaas_loadbalancer_update.md)	 - alias for api UpdateLoadBalancer --LoadBalancerName load_balancer_name

