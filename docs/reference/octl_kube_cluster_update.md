## octl kube cluster update

Updates the configuration of an existing cluster by its ID.

### Synopsis

Updates the configuration of an existing cluster by its ID. The request must include the updated cluster details in the request body. Returns the updated cluster information

> alias for UpdateCluster  id_or_name

```
octl kube cluster update id_or_name [id_or_name]... [flags]
```

### Options

```
      --admin-whitelist strings                                        
      --admission-flag-disable-admission-plugin strings                List of Kubernetes admission plugins to disable
      --admission-flag-enable-admission-plugin strings                 List of Kubernetes admission plugins to enable
      --auth-oidc-client-id string                                     The client id that all tokens must be issued for.
      --auth-oidc-group-claim strings                                  
      --auth-oidc-group-prefix string                                  
      --auth-oidc-issuer-url string                                    The URL of the provider that allows the API server to discover public signing keys.
      --auth-oidc-required-claim stringToString                         (default [])
      --auth-oidc-username-claim string                                
      --auth-oidc-username-prefix string                               
      --auto-maintenance-minor-upgrade-maintenance-duration-hour int   Duration of the maintenance window in hours
      --auto-maintenance-minor-upgrade-maintenance-enabled             Flag to enable or disable the maintenance window
      --auto-maintenance-minor-upgrade-maintenance-start-hour int      Hour of the day when maintenance window starts (0-23)
      --auto-maintenance-minor-upgrade-maintenance-tz string           Timezone for the maintenance window
      --auto-maintenance-minor-upgrade-maintenance-week-day string     Day of the week for the maintenance window
      --auto-maintenance-patch-upgrade-maintenance-duration-hour int   Duration of the maintenance window in hours
      --auto-maintenance-patch-upgrade-maintenance-enabled             Flag to enable or disable the maintenance window
      --auto-maintenance-patch-upgrade-maintenance-start-hour int      Hour of the day when maintenance window starts (0-23)
      --auto-maintenance-patch-upgrade-maintenance-tz string           Timezone for the maintenance window
      --auto-maintenance-patch-upgrade-maintenance-week-day string     Day of the week for the maintenance window
      --control-plane string                                           Size of control plane deployment for the cluster
      --description string                                             
      --disable-api-termination                                        
  -h, --help                                                           help for update
      --maintenance-window-duration-hour int                           Duration of the maintenance window in hours
      --maintenance-window-start-hour int                              Hour of the day when maintenance window starts (0-23)
      --maintenance-window-tz string                                   Timezone for the maintenance window
      --maintenance-window-week-day string                             Day of the week for the maintenance window
      --quirk strings                                                  
      --tags stringToString                                            Tags (key=value,key=value) (default [])
      --version string                                                 
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
      --project string             Name or ID of project
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

* [octl kube cluster](octl_kube_cluster.md)	 - Manage Cluster resources

