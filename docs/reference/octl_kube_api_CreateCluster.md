## octl kube api CreateCluster

Creates a new cluster with the provided configuration.

### Synopsis

Creates a new cluster with the provided configuration. The request must include the cluster details in the request body. all clusters are associated to a project

```
octl kube api CreateCluster [flags]
```

### Options

```
      --AdminLbu                                                     Flag to enable admin load balancer for cluster management
      --AdminWhitelist strings                                       List of CIDR blocks or IP addresses allowed to access the Kubernetes API
      --AdmissionFlags.DisableAdmissionPlugins strings               List of Kubernetes admission plugins to disable
      --AdmissionFlags.EnableAdmissionPlugins strings                List of Kubernetes admission plugins to enable
      --Auth.Oidc.ClientId string                                    The client id that all tokens must be issued for.
      --Auth.Oidc.GroupsClaim strings                                
      --Auth.Oidc.GroupsPrefix string                                
      --Auth.Oidc.IssuerUrl string                                   The URL of the provider that allows the API server to discover public signing keys.
      --Auth.Oidc.UsernameClaim string                               
      --Auth.Oidc.UsernamePrefix string                              
      --AutoMaintenances.MinorUpgradeMaintenance.DurationHours int   Duration of the maintenance window in hours
      --AutoMaintenances.MinorUpgradeMaintenance.Enabled             Flag to enable or disable the maintenance window
      --AutoMaintenances.MinorUpgradeMaintenance.StartHour int       Hour of the day when maintenance window starts (0-23)
      --AutoMaintenances.MinorUpgradeMaintenance.Tz string           Timezone for the maintenance window
      --AutoMaintenances.MinorUpgradeMaintenance.WeekDay string      Day of the week for the maintenance window
      --AutoMaintenances.PatchUpgradeMaintenance.DurationHours int   Duration of the maintenance window in hours
      --AutoMaintenances.PatchUpgradeMaintenance.Enabled             Flag to enable or disable the maintenance window
      --AutoMaintenances.PatchUpgradeMaintenance.StartHour int       Hour of the day when maintenance window starts (0-23)
      --AutoMaintenances.PatchUpgradeMaintenance.Tz string           Timezone for the maintenance window
      --AutoMaintenances.PatchUpgradeMaintenance.WeekDay string      Day of the week for the maintenance window
      --CidrPods string                                              CIDR block for Kubernetes pods network
      --CidrService string                                           CIDR block for Kubernetes services network
      --ClusterDns string                                            IP address for cluster DNS service
      --ControlPlanes string                                         Size of control plane deployment for the cluster
      --CpMultiAz                                                    Flag to enable multi-availability zone for the control plane
      --CpSubregions strings                                         List of subregions where control plane components are deployed
      --Description string                                           
      --DisableApiTermination                                        Flag to prevent accidental cluster deletion
      --MaintenanceWindow.DurationHours int                          Duration of the maintenance window in hours
      --MaintenanceWindow.StartHour int                              Hour of the day when maintenance window starts (0-23)
      --MaintenanceWindow.Tz string                                  Timezone for the maintenance window
      --MaintenanceWindow.WeekDay string                             Day of the week for the maintenance window
      --Name string                                                  Unique cluster name per project, must start with a letter and contain only lowercase letters, numbers, or hyphens
      --ProjectId string                                             Unique identifier of the project this cluster belongs to
      --Quirks strings                                               
      --Version string                                               Version of Kubernetes to be deployed
  -h, --help                                                         help for CreateCluster
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

* [octl kube api](octl_kube_api.md)	 - kube api calls

