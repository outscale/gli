## octl iaas api CreateImage

Creates an OUTSCALE machine image (OMI).

### Synopsis

Creates an OUTSCALE machine image (OMI).

You can use this method for different use cases:
* **Creating from a VM**: You create an OMI from one of your virtual machines (VMs).<br>
* **Copying an OMI**: You copy an existing OMI. The source OMI can be one of your own OMIs, or an OMI owned by another OUTSCALE account that has granted you permission via the [UpdateImage](#updateimage) method.<br>
* **Registering from a snapshot**: You register an OMI from an existing snapshot. The source snapshot can be one of your own snapshots, or a snapshot owned by another OUTSCALE account that has granted you permission via the [UpdateSnapshot](#updatesnapshot) method.<br>
* **Registering from a bucket by using a manifest file**: You register an OMI from the manifest file of an OMI that was exported to an OUTSCALE Object Storage (OOS) bucket. First, the owner of the source OMI must export it to the bucket by using the [CreateImageExportTask](#createimageexporttask) method. Then, they must grant you permission to read the manifest file via a pre-signed URL. For more information, see [Creating a Pre-Signed URL](https://docs.outscale.com/en/userguide/Creating-a-Pre-Signed-URL.html).

**[TIP]**

Registering from a bucket enables you to copy an OMI across Regions.

For more information, see [About OMIs](https://docs.outscale.com/en/userguide/About-OMIs.html).

```
octl iaas api CreateImage [flags]
```

### Options

```
      --Architecture string                              **(when registering from a snapshot)** The architecture of the OMI (i386 or x86_64).
      --BlockDeviceMappings.0.Bsu.DeleteOnVmDeletion     If set to true, the volume is deleted when terminating the VM.
      --BlockDeviceMappings.0.Bsu.Iops int               The number of I/O operations per second (IOPS).
      --BlockDeviceMappings.0.Bsu.SnapshotId string      The ID of the snapshot used to create the volume.
      --BlockDeviceMappings.0.Bsu.VolumeSize int         The size of the volume, in gibibytes (GiB).
      --BlockDeviceMappings.0.Bsu.VolumeType string      The type of the volume (standard | io1 | gp2).
      --BlockDeviceMappings.0.DeviceName string          The device name for the volume.
      --BlockDeviceMappings.0.VirtualDeviceName string   The name of the virtual device (ephemeralN).
      --BootModes strings                                The boot modes compatible with the OMI.
      --Description string                               A description for the new OMI.
      --DryRun                                           If true, checks whether you have the required permissions to perform the action.
      --FileLocation string                              **(required when registering from a bucket by using a manifest file)** The pre-signed URL of the manifest file for the OMI you want to register.
      --ImageName string                                 A unique name for the new OMI.
      --NoReboot                                         **(when creating from a VM)** If false, the VM shuts down before creating the OMI and then reboots.
      --ProductCodes strings                             The product codes associated with the OMI.
      --RootDeviceName string                            **(required when registering from a snapshot)** The name of the root device for the new OMI.
      --SourceImageId string                             **(required when copying an OMI)** The ID of the OMI you want to copy.
      --SourceRegionName string                          **(required when copying an OMI)** The name of the source Region (always the same as the Region of your account).
      --TpmMandatory                                     By default or if set to false, a virtual Trusted Platform Module (vTPM) is not mandatory on VMs created from this OMI.
      --VmId string                                      **(required when creating from a VM)** The ID of the VM from which you want to create the OMI.
  -h, --help                                             help for CreateImage
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

