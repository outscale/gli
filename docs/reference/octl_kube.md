## octl kube

OUTSCALE Kubernetes as a Service (OKS) management

### Options

```
  -h, --help   help for kube
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

* [octl](octl.md)	 - A modern CLI for Outscale services
* [octl kube api](octl_kube_api.md)	 - Call kube API
* [octl kube cluster](octl_kube_cluster.md)	 - Manage Cluster resources
* [octl kube ippool](octl_kube_ippool.md)	 - Manage IP pool resources
* [octl kube kubectl](octl_kube_kubectl.md)	 - Launch kubectl commands on a cluster
* [octl kube netpeering](octl_kube_netpeering.md)	 - Manage netpeering resources
* [octl kube nodepool](octl_kube_nodepool.md)	 - Manage nodepool resources
* [octl kube oosaccess](octl_kube_oosaccess.md)	 - Manage OOS Access resources
* [octl kube project](octl_kube_project.md)	 - Manage Project resources
* [octl kube publicip](octl_kube_publicip.md)	 - Manage ProjectPublicIpsIPS resources
* [octl kube quota](octl_kube_quota.md)	 - Manage ProjectQuota resources
* [octl kube secret](octl_kube_secret.md)	 - Create secret for CCM or CSI driver deployment
* [octl kube vpnconnection](octl_kube_vpnconnection.md)	 - vpnconnection commands

