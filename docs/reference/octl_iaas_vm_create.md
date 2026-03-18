## octl iaas vm create

alias for api CreateVms

### Synopsis

> *alias for api CreateVms*

Creates virtual machines (VMs), and then launches them.

This action enables you to create a specified number of VMs using an OUTSCALE machine image (OMI) that you are allowed to use, and then to automatically launch them.

The VMs remain in the `pending` state until they are created and ready to be used. Once automatically launched, they are in the `running` state.

To check the state of your VMs, call the [ReadVms](#readvms) method.

The metadata server enables you to get the public key provided when the VM is launched. Official OMIs contain a script to get this public key and put it inside the VM to provide secure access without password.

If not specified, the security group used by the service is the default one.

**[NOTE]**

When you attach a security group to a VM, it is actually attached to the primary network interface of the VM.

For more information, see [About VMs](https://docs.outscale.com/en/userguide/About-VMs.html).

```
octl iaas vm create [flags]
```

### Options

```
      --action-on-next-boot-secure-boot string   One action to perform on the next boot of the VM (enable | disable | setup-mode | none).
      --boot-mode string                         Information about the boot mode of the VM.
      --boot-on-creation                         If true, the VM is started on creation.
      --bsu-optimized                            This parameter is not available.
      --client-token string                      A unique identifier which enables you to manage the idempotency.
      --deletion-protection                      If true, you cannot delete the VM unless you change this parameter back to false.
  -h, --help                                     help for create
      --image-id string                          The ID of the OMI used to create the VM.
      --initiated-shutdown-behavior string       The VM behavior when you stop it.
      --keypair-name string                      The name of the keypair.
      --max-vm-count int                         The maximum number of VMs you want to create.
      --min-vm-count int                         The minimum number of VMs you want to create.
      --nested-virtualization                    (dedicated tenancy only) If true, nested virtualization is enabled.
      --nic-delete-on-vm-deletion                If true, the NIC is deleted when the VM is terminated.
      --nic-description string                   The description of the NIC, if you are creating a NIC when creating the VM.
      --nic-device-number int                    The index of the VM device for the NIC attachment (between 0 and 7, both included).
      --nic-id string                            The ID of the NIC, if you are attaching an existing NIC when creating a VM.
      --nic-private-ip string                    A private IP for the NIC.
      --nic-private-ip-is-primary                If true, the IP is the primary private IP of the NIC.
      --nic-secondary-private-ip-count int       The number of secondary private IPs, if you create a NIC when creating a VM.
      --nic-security-group-id strings            One or more IDs of security groups for the NIC, if you create a NIC when creating a VM.
      --nic-subnet-id string                     The ID of the Subnet for the NIC, if you create a NIC when creating a VM.
      --performance string                       The performance of the VM.
      --placement-subregion string               The name of the Subregion.
      --placement-tenancy string                 The tenancy of the VM (default, dedicated, or a dedicated group ID).
      --private-ip strings                       One or more private IPs of the VM.
      --security-group strings                   One or more names of security groups for the VMs.
      --security-group-id strings                One or more IDs of security group for the VMs.
      --subnet-id string                         The ID of the Subnet in which you want to create the VM.
      --tpm-enabled                              If true, a virtual Trusted Platform Module (vTPM) is enabled on the VM.
      --type string                              The type of VM.
      --user-data base64File                     The file storing the data or script used to add a specific configuration to the VM (max size 500 KiB).
      --volume-delete-on-vm-deletion             If set to true, the volume is deleted when terminating the VM.
      --volume-device-name string                The device name for the volume.
      --volume-iops int                          The number of I/O operations per second (IOPS).
      --volume-no-device string                  Removes the device which is included in the block device mapping of the OMI.
      --volume-size int                          The size of the volume, in gibibytes (GiB).
      --volume-snapshot-id string                The ID of the snapshot used to create the volume.
      --volume-type string                       The type of the volume (standard | io1 | gp2).
      --volume-virtual-device-name string        The name of the virtual device (ephemeralN).
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
      --payload string    JSON content for query body
      --profile string    Profile to use in profile file (by default, "default")
      --single            convert single entry lists to a single object
      --template string   JSON template file for query body
  -v, --verbose           Verbose output
  -y, --yes               answer yes to all prompts
```

### SEE ALSO

* [octl iaas vm](octl_iaas_vm.md)	 - vm commands

