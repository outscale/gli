## octl kube nodepool create

alias for api Create

### Synopsis

> alias for api Create



```
octl kube nodepool create [flags]
```

### Options

```
      --auto-healing                       
      --auto-upgrade-enabled               
      --autoscaling                        
      --cluster string                     [REQUIRED] Name or ID of cluster
      --desired-nodes int                  [REQUIRED] 
      --fgpu-model string                  
      --fgpu-operator                      
      --flavour string                     
      --helms strings                      
  -h, --help                               help for create
      --ip-pool string                     
      --max-nodes int                      
      --min-nodes int                      
      --name string                        [REQUIRED] 
      --node-annotations stringToString     (default [])
      --node-labels stringToString          (default [])
      --node-type string                   [REQUIRED] 
      --placement-attract-cluster string   
      --placement-attract-server string    
      --placement-repulse-cluster string   
      --placement-repulse-server string    
      --project string                     Name or ID of project
      --taint                              
      --upgrade-duration-hour int          
      --upgrade-max-surge int              
      --upgrade-max-unavailable int        
      --upgrade-rrule string               
      --upgrade-start-hour int             
      --upgrade-week-day string            
      --volume-iops int                    
      --volume-size int                    
      --volume-snapshot string             
      --volume-type string                 
      --zones strings                      [REQUIRED] 
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

* [octl kube nodepool](octl_kube_nodepool.md)	 - nodepool commands

