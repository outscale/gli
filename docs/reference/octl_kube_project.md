## octl kube project

Manage Project resources

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

* [octl kube](octl_kube.md)	 - OUTSCALE Kubernetes as a Service (OKS) management
* [octl kube project clusters](octl_kube_project_clusters.md)	 - 
* [octl kube project create](octl_kube_project_create.md)	 - Creates a new project with the provided details.
* [octl kube project delete](octl_kube_project_delete.md)	 - Deletes a specific project by its ID.
* [octl kube project describe](octl_kube_project_describe.md)	 - Retrieves detailed information about a specific project by its ID.
* [octl kube project list](octl_kube_project_list.md)	 - Retrieves a list of all projects with optional filters for name, status, CIDR, and deletion status.
* [octl kube project nets](octl_kube_project_nets.md)	 - 
* [octl kube project public-ips](octl_kube_project_public-ips.md)	 - 
* [octl kube project quotas](octl_kube_project_quotas.md)	 - 
* [octl kube project snapshots](octl_kube_project_snapshots.md)	 - 
* [octl kube project update](octl_kube_project_update.md)	 - Updates the details of an existing project by its ID.
* [octl kube project use](octl_kube_project_use.md)	 - Set a default project for cluster commands, reset it without args

