## octl iaas api LinkNic

Attaches a network interface card (NIC) to a virtual machine (VM).

### Synopsis

Attaches a network interface card (NIC) to a virtual machine (VM).

The interface and the VM must be in the same Subregion. The VM can be either `running` or `stopped`. The NIC must be in the `available` state. For more information, see [Attaching a NIC to a VM](https://docs.outscale.com/en/userguide/Attaching-a-NIC-to-a-VM.html).

```
octl iaas api LinkNic [flags]
```

### Options

```
      --DeviceNumber int   The index of the VM device for the NIC attachment (between 1 and 7, both included).
      --DryRun             If true, checks whether you have the required permissions to perform the action.
      --NicId string       The ID of the NIC you want to attach.
      --VmId string        The ID of the VM to which you want to attach the NIC.
  -h, --help               help for LinkNic
```

### Options inherited from parent commands

```
  -c, --columns string    columns to display - [+]<title>:<jq query for content>||<title>:<jq query for content>
      --config string     Path of profile file (by default, ~/.osc/config.json)
      --filter strings    comma separated list of filters for results - name:value,name:value, alias for jq filter 'select(.name | test("value"))'
      --jq string         jq filter
      --no-upgrade        do not check for new versions
  -O, --out-file string   redirect output to file
  -o, --output string     output format (raw, json, yaml, table, csv, none, base64) (default "raw")
      --profile string    Profile to use in profile file (by default, "default")
      --single            convert single entry lists to a single object
      --template string   JSON template for query body
  -v, --verbose           Verbose output
  -y, --yes               answer yes to all prompts
```

### SEE ALSO

* [octl iaas api](octl_iaas_api.md)	 - iaas api calls

