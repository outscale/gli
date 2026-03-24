## octl iaas nic link

alias for api LinkNic --NicId nic_id

### Synopsis

> *alias for api LinkNic --NicId nic_id*

Attaches a network interface card (NIC) to a virtual machine (VM).

The interface and the VM must be in the same Subregion. The VM can be either `running` or `stopped`. The NIC must be in the `available` state. For more information, see [Attaching a NIC to a VM](https://docs.outscale.com/en/userguide/Attaching-a-NIC-to-a-VM.html).

```
octl iaas nic link nic_id [flags]
```

### Options

```
      --device-number int   The index of the VM device for the NIC attachment (between 1 and 7, both included).
  -h, --help                help for link
      --vm-id string        The ID of the VM to which you want to attach the NIC.
```

### Options inherited from parent commands

```
  -c, --columns string              columns to display - [+]<title>:<jq query for content>||<title>:<jq query for content>
      --config string               Path of profile file (by default, ~/.osc/config.json)
      --filter strings              comma separated list of filters for results - name:value,name:value, alias for jq filter 'select(.name | test("value"))'
      --jq string                   jq filter
      --no-upgrade                  do not check for new versions
  -O, --out-file string             redirect output to file
  -o, --output string               output format (raw, json, yaml, table, csv, none, base64) (default "raw")
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

* [octl iaas nic](octl_iaas_nic.md)	 - nic commands

