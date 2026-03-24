## octl kube nodepool api Create



```
octl kube nodepool api Create [flags]
```

### Options

```
      --ObjectMeta.CreationTimestamp string                             
      --ObjectMeta.DeletionGracePeriodSeconds int                       
      --ObjectMeta.DeletionTimestamp string                             
      --ObjectMeta.Finalizers strings                                   
      --ObjectMeta.GenerateName string                                  
      --ObjectMeta.Generation int                                       
      --ObjectMeta.ManagedFields.0.APIVersion string                    
      --ObjectMeta.ManagedFields.0.FieldsType string                    
      --ObjectMeta.ManagedFields.0.FieldsV1 string                      
      --ObjectMeta.ManagedFields.0.Manager string                       
      --ObjectMeta.ManagedFields.0.Operation string                     
      --ObjectMeta.ManagedFields.0.Subresource string                   
      --ObjectMeta.ManagedFields.0.Time string                          
      --ObjectMeta.Name string                                          
      --ObjectMeta.Namespace string                                     
      --ObjectMeta.OwnerReferences.0.APIVersion string                  
      --ObjectMeta.OwnerReferences.0.BlockOwnerDeletion                 
      --ObjectMeta.OwnerReferences.0.Controller                         
      --ObjectMeta.OwnerReferences.0.Kind string                        
      --ObjectMeta.OwnerReferences.0.Name string                        
      --ObjectMeta.OwnerReferences.0.UID string                         
      --ObjectMeta.ResourceVersion string                               
      --ObjectMeta.SelfLink string                                      
      --ObjectMeta.UID string                                           
      --Spec.AutoHealing                                                
      --Spec.Autoscaling                                                
      --Spec.DesiredNodes int                                           
      --Spec.Fgpu.K8sOperator                                           
      --Spec.Fgpu.Model string                                          
      --Spec.Flavour string                                             
      --Spec.Helms strings                                              
      --Spec.IpPoolName string                                          
      --Spec.MaxNodes int                                               
      --Spec.MinNodes int                                               
      --Spec.NodeType string                                            
      --Spec.PhysicalPlacement.NodeAttractCluster string                
      --Spec.PhysicalPlacement.NodeAttractServer string                 
      --Spec.PhysicalPlacement.NodeRepulseCluster string                
      --Spec.PhysicalPlacement.NodeRepulseServer string                 
      --Spec.Taint                                                      
      --Spec.UpgradeStrategy.AutoUpgradeEnabled                         
      --Spec.UpgradeStrategy.AutoUpgradeMaintenance.DurationHours int   
      --Spec.UpgradeStrategy.AutoUpgradeMaintenance.Rrule string        
      --Spec.UpgradeStrategy.AutoUpgradeMaintenance.StartHour int       
      --Spec.UpgradeStrategy.AutoUpgradeMaintenance.WeekDay string      
      --Spec.UpgradeStrategy.MaxSurge int                               
      --Spec.UpgradeStrategy.MaxUnavailable int                         
      --Spec.Volumes.0.Device string                                    
      --Spec.Volumes.0.Dir string                                       
      --Spec.Volumes.0.Iops int                                         
      --Spec.Volumes.0.Name string                                      
      --Spec.Volumes.0.Size int                                         
      --Spec.Volumes.0.Snapshot string                                  
      --Spec.Volumes.0.Type string                                      
      --Spec.Zones strings                                              
      --Status.LastError.Message string                                 
      --Status.LastError.Posted string                                  
      --Status.Progress.Attached int                                    
      --Status.Progress.Pending int                                     
      --Status.Progress.Ready int                                       
      --Status.Progress.Running int                                     
      --Status.Progress.ShuttingDown int                                
      --Status.Progress.Stopped int                                     
      --Status.Progress.Stopping int                                    
      --Status.State.Expires string                                     
      --Status.State.Name string                                        
      --TypeMeta.APIVersion string                                      
      --TypeMeta.Kind string                                            
  -h, --help                                                            help for Create
```

### Options inherited from parent commands

```
      --cluster string              Name or ID of cluster
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

* [octl kube nodepool api](octl_kube_nodepool_api.md)	 - nodepool api calls

