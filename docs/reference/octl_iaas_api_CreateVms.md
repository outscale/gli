## octl iaas api CreateVms

Creates virtual machines (VMs), and then launches them.

### Synopsis

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
octl iaas api CreateVms [flags]
```

### Options

```
      --ActionsOnNextBoot.SecureBoot string              One action to perform on the next boot of the VM (enable | disable | setup-mode | none).
      --BlockDeviceMappings.0.Bsu.DeleteOnVmDeletion     If set to true, the volume is deleted when terminating the VM.
      --BlockDeviceMappings.0.Bsu.Iops int               The number of I/O operations per second (IOPS).
      --BlockDeviceMappings.0.Bsu.SnapshotId string      The ID of the snapshot used to create the volume.
      --BlockDeviceMappings.0.Bsu.VolumeSize int         The size of the volume, in gibibytes (GiB).
      --BlockDeviceMappings.0.Bsu.VolumeType string      The type of the volume (standard | io1 | gp2).
      --BlockDeviceMappings.0.DeviceName string          The device name for the volume.
      --BlockDeviceMappings.0.NoDevice string            Removes the device which is included in the block device mapping of the OMI.
      --BlockDeviceMappings.0.VirtualDeviceName string   The name of the virtual device (ephemeralN).
      --BootMode string                                  Information about the boot mode of the VM.
      --BootOnCreation                                   If true, the VM is started on creation.
      --BsuOptimized                                     This parameter is not available.
      --ClientToken string                               A unique identifier which enables you to manage the idempotency.
      --DeletionProtection                               If true, you cannot delete the VM unless you change this parameter back to false.
      --DryRun                                           If true, checks whether you have the required permissions to perform the action.
      --ImageId string                                   The ID of the OMI used to create the VM.
      --KeypairName string                               The name of the keypair.
      --MaxVmsCount int                                  The maximum number of VMs you want to create.
      --MinVmsCount int                                  The minimum number of VMs you want to create.
      --NestedVirtualization                             (dedicated tenancy only) If true, nested virtualization is enabled.
      --Nics.0.DeleteOnVmDeletion                        If true, the NIC is deleted when the VM is terminated.
      --Nics.0.Description string                        The description of the NIC, if you are creating a NIC when creating the VM.
      --Nics.0.DeviceNumber int                          The index of the VM device for the NIC attachment (between 0 and 7, both included).
      --Nics.0.NicId string                              The ID of the NIC, if you are attaching an existing NIC when creating a VM.
      --Nics.0.PrivateIps.0.IsPrimary                    If true, the IP is the primary private IP of the NIC.
      --Nics.0.PrivateIps.0.PrivateIp string             A private IP for the NIC.
      --Nics.0.SecondaryPrivateIpCount int               The number of secondary private IPs, if you create a NIC when creating a VM.
      --Nics.0.SecurityGroupIds strings                  One or more IDs of security groups for the NIC, if you create a NIC when creating a VM.
      --Nics.0.SubnetId string                           The ID of the Subnet for the NIC, if you create a NIC when creating a VM.
      --Performance string                               The performance of the VM.
      --Placement.SubregionName string                   The name of the Subregion.
      --Placement.Tenancy string                         The tenancy of the VM (default, dedicated, or a dedicated group ID).
      --PrivateIps strings                               One or more private IPs of the VM.
      --SecurityGroupIds strings                         One or more IDs of security group for the VMs.
      --SecurityGroups strings                           One or more names of security groups for the VMs.
      --SubnetId string                                  The ID of the Subnet in which you want to create the VM.
      --TpmEnabled                                       If true, a virtual Trusted Platform Module (vTPM) is enabled on the VM.
      --UserData string                                  Data or script used to add a specific configuration to the VM.
      --VmInitiatedShutdownBehavior string               The VM behavior when you stop it.
      --VmType string                                    The type of VM.
  -h, --help                                             help for CreateVms
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

* [octl iaas api](octl_iaas_api.md)	 - iaas api calls

