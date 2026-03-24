## octl iaas vm

vm commands

### Options

```
  -h, --help   help for vm
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

* [octl iaas](octl_iaas.md)	 - OUTSCALE IaaS management
* [octl iaas vm create](octl_iaas_vm_create.md)	 - alias for api CreateVms
* [octl iaas vm delete](octl_iaas_vm_delete.md)	 - alias for api DeleteVms --VmIds vm_id
* [octl iaas vm describe](octl_iaas_vm_describe.md)	 - alias for api ReadVms --Filters.VmIds vm_id
* [octl iaas vm list](octl_iaas_vm_list.md)	 - alias for api ReadVms
* [octl iaas vm readconsole](octl_iaas_vm_readconsole.md)	 - alias for api ReadConsoleOutput --VmId vm_id
* [octl iaas vm start](octl_iaas_vm_start.md)	 - alias for api StartVms --VmIds vm_id
* [octl iaas vm states](octl_iaas_vm_states.md)	 - alias for api ReadVmsState
* [octl iaas vm stop](octl_iaas_vm_stop.md)	 - alias for api StopVms --VmIds vm_id
* [octl iaas vm update](octl_iaas_vm_update.md)	 - alias for api UpdateVm --VmId vm_id

