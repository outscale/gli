## octl kube api

kube api calls

### Options

```
  -h, --help   help for api
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
* [octl kube api CreateCluster](octl_kube_api_CreateCluster.md)	 - Creates a new cluster with the provided configuration.
* [octl kube api CreateProject](octl_kube_api_CreateProject.md)	 - Creates a new project with the provided details.
* [octl kube api DeleteCluster](octl_kube_api_DeleteCluster.md)	 - request returning *DetailResponse
* [octl kube api DeleteProject](octl_kube_api_DeleteProject.md)	 - request returning *DetailResponse
* [octl kube api GetCPSubregions](octl_kube_api_GetCPSubregions.md)	 - request returning *CPSubregionsResponse
* [octl kube api GetCluster](octl_kube_api_GetCluster.md)	 - request returning *ClusterResponse
* [octl kube api GetClusterTemplate](octl_kube_api_GetClusterTemplate.md)	 - request returning *TemplateResponseClusterInputTemplate
* [octl kube api GetControlPlanePlans](octl_kube_api_GetControlPlanePlans.md)	 - request returning *ControlPlanesResponse
* [octl kube api GetKubeconfig](octl_kube_api_GetKubeconfig.md)	 - 
* [octl kube api GetKubeconfigWithPubkeyNACL](octl_kube_api_GetKubeconfigWithPubkeyNACL.md)	 - request returning *KubeconfigResponse
* [octl kube api GetKubernetesVersions](octl_kube_api_GetKubernetesVersions.md)	 - request returning *KubernetesVersionsResponse
* [octl kube api GetNetPeeringAcceptanceTemplate](octl_kube_api_GetNetPeeringAcceptanceTemplate.md)	 - request returning *TemplateResponseNetPeeringAcceptance
* [octl kube api GetNetPeeringRequestTemplate](octl_kube_api_GetNetPeeringRequestTemplate.md)	 - request returning *TemplateResponseNetPeeringRequest
* [octl kube api GetNodepoolTemplate](octl_kube_api_GetNodepoolTemplate.md)	 - request returning *TemplateResponseNodepool
* [octl kube api GetProject](octl_kube_api_GetProject.md)	 - request returning *ProjectResponse
* [octl kube api GetProjectNets](octl_kube_api_GetProjectNets.md)	 - 
* [octl kube api GetProjectPublicIps](octl_kube_api_GetProjectPublicIps.md)	 - 
* [octl kube api GetProjectQuotas](octl_kube_api_GetProjectQuotas.md)	 - 
* [octl kube api GetProjectSnapshots](octl_kube_api_GetProjectSnapshots.md)	 - 
* [octl kube api GetProjectTemplate](octl_kube_api_GetProjectTemplate.md)	 - request returning *TemplateResponseProjectInput
* [octl kube api GetQuotas](octl_kube_api_GetQuotas.md)	 - 
* [octl kube api ListAllClusters](octl_kube_api_ListAllClusters.md)	 - 
* [octl kube api ListClustersByProjectID](octl_kube_api_ListClustersByProjectID.md)	 - 
* [octl kube api ListProjects](octl_kube_api_ListProjects.md)	 - request returning *ProjectResponseList
* [octl kube api UpdateCluster](octl_kube_api_UpdateCluster.md)	 - Updates the configuration of an existing cluster by its ID.
* [octl kube api UpdateProject](octl_kube_api_UpdateProject.md)	 - Updates the details of an existing project by its ID.
* [octl kube api UpgradeCluster](octl_kube_api_UpgradeCluster.md)	 - request returning *ClusterResponse

