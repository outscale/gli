## octl iaas api LinkFlexibleGpu

Attaches one of your allocated flexible GPUs (fGPUs) to one of your virtual machines (VMs).

### Synopsis

Attaches one of your allocated flexible GPUs (fGPUs) to one of your virtual machines (VMs).

To complete the linking of the fGPU, you need to do a stop/start of the VM. A simple restart is not sufficient, as the linking of the fGPU is done when the VM goes through the `stopped` state. For the difference between stop/start and restart, see [About VM Lifecycle](https://docs.outscale.com/en/userguide/About-VM-Lifecycle.html).



**[NOTE]**

You can attach fGPUs only to VMs with the `highest` (1) performance flag. For more information see [About Flexible GPUs](https://docs.outscale.com/en/userguide/About-Flexible-GPUs.html) and [VM Types](https://docs.outscale.com/en/userguide/VM-Types.html).

```
octl iaas api LinkFlexibleGpu [flags]
```

### Options

```
      --DryRun                 If true, checks whether you have the required permissions to perform the action.
      --FlexibleGpuId string   [REQUIRED] The ID of the fGPU you want to attach.
      --VmId string            [REQUIRED] The ID of the VM you want to attach the fGPU to.
  -h, --help                   help for LinkFlexibleGpu
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

* [octl iaas api](octl_iaas_api.md)	 - Call iaas API

