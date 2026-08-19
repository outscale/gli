## octl kube project

project commands

### Options

```
  -h, --help   help for project
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

* [octl kube](octl_kube.md)	 - OUTSCALE Kubernetes as a Service (OKS) management
* [octl kube project clusters](octl_kube_project_clusters.md)	 - alias for api ListClustersByProjectID
* [octl kube project create](octl_kube_project_create.md)	 - alias for api CreateProject
* [octl kube project delete](octl_kube_project_delete.md)	 - alias for api DeleteProject  id_or_name
* [octl kube project describe](octl_kube_project_describe.md)	 - alias for api GetProject  id_or_name
* [octl kube project list](octl_kube_project_list.md)	 - alias for api ListProjects
* [octl kube project nets](octl_kube_project_nets.md)	 - alias for api GetProjectNets id_or_name
* [octl kube project public-ips](octl_kube_project_public-ips.md)	 - alias for api GetProjectPublicIps id_or_name
* [octl kube project quotas](octl_kube_project_quotas.md)	 - alias for api GetProjectQuotas id_or_name
* [octl kube project snapshots](octl_kube_project_snapshots.md)	 - alias for api GetProjectSnapshots id_or_name
* [octl kube project update](octl_kube_project_update.md)	 - alias for api UpdateProject  id_or_name
* [octl kube project use](octl_kube_project_use.md)	 - Set a default project for cluster commands, reset it without args

