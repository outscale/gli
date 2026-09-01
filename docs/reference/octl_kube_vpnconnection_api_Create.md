## octl kube vpnconnection api Create



```
octl kube vpnconnection api Create [flags]
```

### Options

```
      --DryRun strings                                                      
      --FieldManager string                                                 
      --FieldValidation string                                              
      --ObjectMeta.Annotations stringToString                                (default [])
      --ObjectMeta.DeletionGracePeriodSeconds int                           
      --ObjectMeta.DeletionTimestamp fileOrJson                             
      --ObjectMeta.Finalizers strings                                       
      --ObjectMeta.GenerateName string                                      
      --ObjectMeta.Labels stringToString                                     (default [])
      --ObjectMeta.Name string                                              
      --ObjectMeta.Namespace string                                         
      --ObjectMeta.ResourceVersion string                                   
      --ObjectMeta.SelfLink string                                          
      --ObjectMeta.UID string                                               
      --Spec.BgpASN int                                                     
      --Spec.ClientGatewayConfigurationStorage.ConfigMapName string         
      --Spec.ClientGatewayConfigurationStorage.Namespace string             
      --Spec.ClientGatewayConfigurationStorage.SecretName string            
      --Spec.PublicIP string                                                
      --Spec.StaticRoutesOnly                                               
      --Spec.VpnOptions.Phase2Options.PreSharedKey.SecretKey string         
      --Spec.VpnOptions.Phase2Options.PreSharedKey.SecretName string        
      --Spec.VpnOptions.Phase2Options.PreSharedKey.SecretNamespace string   
      --Spec.VpnOptions.TunnelInsideIpRange string                          
      --Spec.VpnRoutes strings                                              
  -h, --help                                                                help for Create
```

### Options inherited from parent commands

```
      --cluster string             [REQUIRED] ID of cluster
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

* [octl kube vpnconnection api](octl_kube_vpnconnection_api.md)	 - Call vpnconnection API

