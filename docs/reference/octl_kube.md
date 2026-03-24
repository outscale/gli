## octl kube

OUTSCALE Kubernetes as a Service (OKS) management

### Options

```
  -h, --help   help for kube
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

* [octl](octl.md)	 - A modern CLI for Outscale services
* [octl kube api](octl_kube_api.md)	 - kube api calls
* [octl kube cluster](octl_kube_cluster.md)	 - cluster commands
* [octl kube kubeconfig](octl_kube_kubeconfig.md)	 - kubeconfig commands
* [octl kube kubectl](octl_kube_kubectl.md)	 - 
* [octl kube nodepool](octl_kube_nodepool.md)	 - nodepool commands
* [octl kube project](octl_kube_project.md)	 - project commands
* [octl kube publicip](octl_kube_publicip.md)	 - publicip commands
* [octl kube quota](octl_kube_quota.md)	 - quota commands

