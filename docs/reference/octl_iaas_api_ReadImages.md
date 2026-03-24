## octl iaas api ReadImages

Lists one or more OUTSCALE machine images (OMIs) you can use.

### Synopsis

Lists one or more OUTSCALE machine images (OMIs) you can use.

```
octl iaas api ReadImages [flags]
```

### Options

```
      --DryRun                                          If true, checks whether you have the required permissions to perform the action.
      --Filters.AccountAliases strings                  The account aliases of the owners of the OMIs.
      --Filters.AccountIds strings                      The OUTSCALE account IDs of the owners of the OMIs.
      --Filters.Architectures strings                   The architectures of the OMIs (i386 | x86_64).
      --Filters.BlockDeviceMappingDeleteOnVmDeletion    Whether the volumes are deleted or not when terminating the VM.
      --Filters.BlockDeviceMappingDeviceNames strings   The device names for the volumes.
      --Filters.BlockDeviceMappingSnapshotIds strings   The IDs of the snapshots used to create the volumes.
      --Filters.BlockDeviceMappingVolumeSizes ints      The sizes of the volumes, in gibibytes (GiB).
      --Filters.BlockDeviceMappingVolumeTypes strings   The types of volumes (standard | gp2 | io1).
      --Filters.BootModes strings                       The boot modes compatible with the OMIs.
      --Filters.Descriptions strings                    The descriptions of the OMIs, provided when they were created.
      --Filters.FileLocations strings                   The locations of the buckets where the OMI files are stored.
      --Filters.Hypervisors strings                     The hypervisor type of the OMI (always xen).
      --Filters.ImageIds strings                        The IDs of the OMIs.
      --Filters.ImageNames strings                      The names of the OMIs, provided when they were created.
      --Filters.PermissionsToLaunchAccountIds strings   The OUTSCALE account IDs which have launch permissions for the OMIs.
      --Filters.PermissionsToLaunchGlobalPermission     If true, lists all public OMIs.
      --Filters.ProductCodeNames strings                The names of the product codes associated with the OMI.
      --Filters.ProductCodes strings                    The product codes associated with the OMI.
      --Filters.RootDeviceNames strings                 The name of the root device.
      --Filters.RootDeviceTypes strings                 The types of root device used by the OMIs (bsu or ebs).
      --Filters.SecureBoot                              Whether secure boot is activated or not.
      --Filters.States strings                          The states of the OMIs (pending | available | failed).
      --Filters.TagKeys strings                         The keys of the tags associated with the OMIs.
      --Filters.TagValues strings                       The values of the tags associated with the OMIs.
      --Filters.Tags strings                            The key/value combination of the tags associated with the OMIs, in the following format: "Filters":{"Tags":["TAGKEY=TAGVALUE"]}.
      --Filters.TpmMandatory                            Whether a virtual Trusted Platform Module (vTPM) is mandatory for VMs created from this OMI (true) or not (false).
      --Filters.VirtualizationTypes strings             The virtualization types (always hvm).
      --NextPageToken string                            The token to request the next page of results.
      --ResultsPerPage int                              The maximum number of logs returned in a single response (between 1 and 1000, both included).
  -h, --help                                            help for ReadImages
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

* [octl iaas api](octl_iaas_api.md)	 - iaas api calls

