## octl iaas image create

alias for api CreateImage

### Synopsis

> *alias for api CreateImage*

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
octl iaas image create [flags]
```

### Options

```
      --architecture string                 **(when registering from a snapshot)** The architecture of the OMI (i386 or x86_64).
      --boot-mode strings                   The boot modes compatible with the OMI.
      --description string                  A description for the new OMI.
      --file-location string                **(required when registering from a bucket by using a manifest file)** The pre-signed URL of the manifest file for the OMI you want to register.
  -h, --help                                help for create
      --name string                         A unique name for the new OMI.
      --no-reboot                           **(when creating from a VM)** If false, the VM shuts down before creating the OMI and then reboots.
      --product-code strings                The product codes associated with the OMI.
      --root-device-name string             **(required when registering from a snapshot)** The name of the root device for the new OMI.
      --source-image-id string              **(required when copying an OMI)** The ID of the OMI you want to copy.
      --source-region-name string           **(required when copying an OMI)** The name of the source Region (always the same as the Region of your account).
      --tpm-mandatory                       By default or if set to false, a virtual Trusted Platform Module (vTPM) is not mandatory on VMs created from this OMI.
      --vm-id string                        **(required when creating from a VM)** The ID of the VM from which you want to create the OMI.
      --volume-delete-on-vm-deletion        If set to true, the volume is deleted when terminating the VM.
      --volume-device-name string           The device name for the volume.
      --volume-iops int                     The number of I/O operations per second (IOPS).
      --volume-size int                     The size of the volume, in gibibytes (GiB).
      --volume-snapshot-id string           The ID of the snapshot used to create the volume.
      --volume-type string                  The type of the volume (standard | io1 | gp2).
      --volume-virtual-device-name string   The name of the virtual device (ephemeralN).
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

* [octl iaas image](octl_iaas_image.md)	 - image commands

