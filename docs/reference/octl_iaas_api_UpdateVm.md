## octl iaas api UpdateVm

Modifies the specified attributes of a virtual machine (VM).

### Synopsis

Modifies the specified attributes of a virtual machine (VM).

You must stop the VM before modifying the following attributes:

* `NestedVirtualization`

* `Performance`

* `UserData`

* `VmType`

To complete the update of secure boot, you need to do a stop/start of the VM. A simple restart is not sufficient, as the update is done when the VM goes through the stopped state. For the difference between stop/start and restart, see [About VM Lifecycle](https://docs.outscale.com/en/userguide/About-VM-Lifecycle.html).

```
octl iaas api UpdateVm [flags]
```

### Options

```
      --ActionsOnNextBoot.SecureBoot string              One action to perform on the next boot of the VM (enable | disable | setup-mode | none).
      --BlockDeviceMappings.0.Bsu.DeleteOnVmDeletion     If set to true, the volume is deleted when terminating the VM.
      --BlockDeviceMappings.0.Bsu.VolumeId string        The ID of the volume.
      --BlockDeviceMappings.0.DeviceName string          The device name for the volume.
      --BlockDeviceMappings.0.NoDevice string            Removes the device which is included in the block device mapping of the OMI.
      --BlockDeviceMappings.0.VirtualDeviceName string   The name of the virtual device (ephemeralN).
      --BsuOptimized                                     This parameter is not available.
      --DeletionProtection                               If true, you cannot delete the VM unless you change this parameter back to false.
      --DryRun                                           If true, checks whether you have the required permissions to perform the action.
      --IsSourceDestChecked                              (Net only) If true, the source/destination check is enabled.
      --KeypairName string                               The name of a keypair you want to associate with the VM.
      --NestedVirtualization                             (dedicated tenancy only) If true, nested virtualization is enabled.
      --Performance string                               The performance of the VM.
      --SecurityGroupIds strings                         One or more IDs of security groups for the VM.
      --UserData string                                  The Base64-encoded MIME user data, limited to 500 kibibytes (KiB).
      --VmId string                                      The ID of the VM.
      --VmInitiatedShutdownBehavior string               The VM behavior when you stop it.
      --VmType string                                    The type of VM.
  -h, --help                                             help for UpdateVm
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

* [octl iaas api](octl_iaas_api.md)	 - iaas api calls

