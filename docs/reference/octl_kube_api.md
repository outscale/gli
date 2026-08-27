## octl kube api

kube api calls

### Options

```
  -h, --help   help for api
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
* [octl kube api CreateCluster](octl_kube_api_CreateCluster.md)	 - Creates a new cluster with the provided configuration.
* [octl kube api CreateProject](octl_kube_api_CreateProject.md)	 - Creates a new project with the provided details.
* [octl kube api DeleteCluster](octl_kube_api_DeleteCluster.md)	 - Deletes a specific cluster by its ID.
* [octl kube api DeleteProject](octl_kube_api_DeleteProject.md)	 - Deletes a specific project by its ID.
* [octl kube api GetCPSubregions](octl_kube_api_GetCPSubregions.md)	 - Retrieves the available subregions for cluster deployment.
* [octl kube api GetCluster](octl_kube_api_GetCluster.md)	 - Retrieves detailed information about a specific cluster by its ID.
* [octl kube api GetClusterTemplate](octl_kube_api_GetClusterTemplate.md)	 - Retrieves the default cluster template, including predefined control plane configurations, networking settings, and maintenance schedules.
* [octl kube api GetControlPlanePlans](octl_kube_api_GetControlPlanePlans.md)	 - Retrieves the available control plane plans for cluster deployment.
* [octl kube api GetKubeconfig](octl_kube_api_GetKubeconfig.md)	 - 
* [octl kube api GetKubeconfigWithPubkeyNACL](octl_kube_api_GetKubeconfigWithPubkeyNACL.md)	 - Retrieves the kubeconfig for a specific cluster by its ID with optional NaCl public key encryption.
* [octl kube api GetKubernetesVersions](octl_kube_api_GetKubernetesVersions.md)	 - Retrieves the available Kubernetes versions for cluster creation or upgrades.
* [octl kube api GetNetPeeringAcceptanceTemplate](octl_kube_api_GetNetPeeringAcceptanceTemplate.md)	 - NetPeering Acceptance manifest
* [octl kube api GetNetPeeringRequestTemplate](octl_kube_api_GetNetPeeringRequestTemplate.md)	 - NetPeering Request manifest
* [octl kube api GetNodepoolTemplate](octl_kube_api_GetNodepoolTemplate.md)	 - Retrieves the default node pool template, including predefined configurations for node scaling, storage, and upgrade strategies.
* [octl kube api GetProject](octl_kube_api_GetProject.md)	 - Retrieves detailed information about a specific project by its ID.
* [octl kube api GetProjectNets](octl_kube_api_GetProjectNets.md)	 - 
* [octl kube api GetProjectPublicIps](octl_kube_api_GetProjectPublicIps.md)	 - 
* [octl kube api GetProjectQuotas](octl_kube_api_GetProjectQuotas.md)	 - 
* [octl kube api GetProjectSnapshots](octl_kube_api_GetProjectSnapshots.md)	 - 
* [octl kube api GetProjectTemplate](octl_kube_api_GetProjectTemplate.md)	 - Retrieves the default project template, including predefined network configurations, region, and metadata.
* [octl kube api GetQuotas](octl_kube_api_GetQuotas.md)	 - 
* [octl kube api ListAllClusters](octl_kube_api_ListAllClusters.md)	 - 
* [octl kube api ListClustersByProjectID](octl_kube_api_ListClustersByProjectID.md)	 - 
* [octl kube api ListProjects](octl_kube_api_ListProjects.md)	 - Retrieves a list of all projects with optional filters for name, status, CIDR, and deletion status.
* [octl kube api UpdateCluster](octl_kube_api_UpdateCluster.md)	 - Updates the configuration of an existing cluster by its ID.
* [octl kube api UpdateProject](octl_kube_api_UpdateProject.md)	 - Updates the details of an existing project by its ID.
* [octl kube api UpgradeCluster](octl_kube_api_UpgradeCluster.md)	 - Upgrades a specific cluster by its ID to the latest available version.

