## octl iaas vm update

Modifies the specified attributes of a virtual machine (VM).

### Synopsis

Modifies the specified attributes of a virtual machine (VM).

You must stop the VM before modifying the following attributes:

* `NestedVirtualization`

* `Performance`

* `UserData`

* `VmType`

To complete the update of secure boot, you need to do a stop/start of the VM. A simple restart is not sufficient, as the update is done when the VM goes through the stopped state. For the difference between stop/start and restart, see [About VM Lifecycle](https://docs.outscale.com/en/userguide/About-VM-Lifecycle.html).

> alias for UpdateVm --VmId vm_id

```
octl iaas vm update vm_id [vm_id]... [flags]
```

### Options

```
      --action-on-next-boot-secure-boot string                One action to perform on the next boot of the VM (enable | disable | setup-mode | none).
      --bsu-optimized                                         This parameter is not available.
      --deletion-protection                                   If true, you cannot delete the VM unless you change this parameter back to false.
  -h, --help                                                  help for update
      --initiated-shutdown-behavior string                    The VM behavior when you stop it.
      --is-source-dest-checked                                (Net only) If true, the source/destination check is enabled.
      --keypair-name string                                   The name of a keypair you want to associate with the VM.
      --nested-virtualization                                 (dedicated tenancy only) If true, nested virtualization is enabled.
      --performance string                                    The performance of the VM.
      --security-group-id strings                             One or more IDs of security groups for the VM.
      --shutdown-behavior-configuration-guest-action string   The action performed by the orchestrator when the VM is shut down from the guest operating system.
      --shutdown-behavior-configuration-host-action string    The action performed by the orchestrator when the VM is shut down due to a host infrastructure failure.
      --type string                                           The type of VM.
      --user-data string                                      The file storing the data or script used to add a specific configuration to the VM (max size 500 KiB).
      --volume-delete-on-vm-deletion                          If set to true, the volume is deleted when terminating the VM.
      --volume-device-name string                             The device name for the volume.
      --volume-id string                                      The ID of the volume.
      --volume-no-device string                               Removes the device which is included in the block device mapping of the OMI.
      --volume-virtual-device-name string                     The name of the virtual device (ephemeralN).
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

* [octl iaas vm](octl_iaas_vm.md)	 - Manage Vm resources

