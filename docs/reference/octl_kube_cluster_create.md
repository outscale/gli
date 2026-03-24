## octl kube cluster create

alias for api CreateCluster

### Synopsis

> *alias for api CreateCluster*

Creates a new cluster with the provided configuration. The request must include the cluster details in the request body. all clusters are associated to a project

```
octl kube cluster create [flags]
```

### Options

```
      --allow-from strings                  List of CIDR blocks or IP addresses allowed to access the Kubernetes API
      --cidr-pods string                    CIDR block for Kubernetes pods network
      --cidr-services string                CIDR block for Kubernetes services network
      --control-plane string                Size of control plane deployment for the cluster
      --description string                  
      --disable-admission-plugins strings   List of Kubernetes admission plugins to disable
      --disable-api-termination             Flag to prevent accidental cluster deletion
      --enable-admission-plugins strings    List of Kubernetes admission plugins to enable
  -h, --help                                help for create
      --maintenance-duration-hour int       Duration of the maintenance window in hours
      --maintenance-start-hour int          Hour of the day when maintenance window starts (0-23)
      --maintenance-tz string               Timezone for the maintenance window
      --maintenance-week-day string         Day of the week for the maintenance window
      --multi-az                            Flag to enable multi-availability zone for the control plane
      --name string                         Unique cluster name per project, must start with a letter and contain only lowercase letters, numbers, or hyphens
      --oidc-client-id string               The client id that all tokens must be issued for.
      --oidc-group-claim strings            
      --oidc-group-prefix string            
      --oidc-issuer-url string              The URL of the provider that allows the API server to discover public signing keys.
      --oidc-username-claim string          
      --oidc-username-prefix string         
      --project string                      Unique identifier of the project this cluster belongs to
      --quirk strings                       
      --subregions strings                  List of subregions where control plane components are deployed
      --upgrade-minor-duration-hour int     Duration of the maintenance window in hours
      --upgrade-minor-enabled               Flag to enable or disable the maintenance window
      --upgrade-minor-start-hour int        Hour of the day when maintenance window starts (0-23)
      --upgrade-minor-tz string             Timezone for the maintenance window
      --upgrade-minor-week-day string       Day of the week for the maintenance window
      --upgrade-patch-duration-hour int     Duration of the maintenance window in hours
      --upgrade-patch-enabled               Flag to enable or disable the maintenance window
      --upgrade-patch-start-hour int        Hour of the day when maintenance window starts (0-23)
      --upgrade-patch-tz string             Timezone for the maintenance window
      --upgrade-patch-week-day string       Day of the week for the maintenance window
      --version string                      Version of Kubernetes to be deployed
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

* [octl kube cluster](octl_kube_cluster.md)	 - cluster commands

