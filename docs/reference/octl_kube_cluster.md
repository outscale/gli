## octl kube cluster

cluster commands

### Options

```
  -h, --help   help for cluster
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

* [octl kube](octl_kube.md)	 - OUTSCALE Kubernetes as a Service (OKS) management
* [octl kube cluster create](octl_kube_cluster_create.md)	 - alias for api CreateCluster
* [octl kube cluster delete](octl_kube_cluster_delete.md)	 - alias for api DeleteCluster  id
* [octl kube cluster describe](octl_kube_cluster_describe.md)	 - alias for api GetCluster  id
* [octl kube cluster kubeconfig](octl_kube_cluster_kubeconfig.md)	 - alias for api GetKubeconfig cluster_name_or_id
* [octl kube cluster list](octl_kube_cluster_list.md)	 - alias for api ListAllClusters
* [octl kube cluster update](octl_kube_cluster_update.md)	 - alias for api UpdateCluster  id

