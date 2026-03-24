## octl iaas api LinkLoadBalancerBackendMachines

Attaches one or more virtual machines (VMs) to a specified load balancer.

### Synopsis

Attaches one or more virtual machines (VMs) to a specified load balancer. You need to specify at least the `BackendIps` or the `BackendVmIds` parameter.

The VMs can be in different Subnets and different Subregions than the load balancer, as long as the VMs and load balancers are all in the public Cloud or all in the same Net. It may take a little time for a VM to be registered with the load balancer. Once the VM is registered with a load balancer, it receives traffic and requests from this load balancer and is called a backend VM.

```
octl iaas api LinkLoadBalancerBackendMachines [flags]
```

### Options

```
      --BackendIps strings        One or more public IPs of backend VMs.
      --BackendVmIds strings      One or more IDs of backend VMs.
      --DryRun                    If true, checks whether you have the required permissions to perform the action.
      --LoadBalancerName string   The name of the load balancer.
  -h, --help                      help for LinkLoadBalancerBackendMachines
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

* [octl iaas api](octl_iaas_api.md)	 - iaas api calls

