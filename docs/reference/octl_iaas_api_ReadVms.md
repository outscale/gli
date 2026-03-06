## octl iaas api ReadVms

Lists one or more of your virtual machines (VMs).

### Synopsis

Lists one or more of your virtual machines (VMs).

If you provide one or more VM IDs, this action returns a description for all of these VMs. If you do not provide any VM ID, this action returns a description for all of the VMs that belong to you. If you provide an invalid VM ID, an error is returned. If you provide the ID of a VM that does not belong to you, the description of this VM is not included in the response. The refresh interval for data returned by this action is one hour, meaning that a terminated VM may appear in the response.

```
octl iaas api ReadVms [flags]
```

### Options

```
      --DryRun                                                If true, checks whether you have the required permissions to perform the action.
      --Filters.Architectures strings                         The architectures of the VMs (i386 | x86_64).
      --Filters.BlockDeviceMappingDeleteOnVmDeletion          Whether the BSU volumes are deleted when terminating the VMs.
      --Filters.BlockDeviceMappingDeviceNames strings         The device names for the BSU volumes (in the format /dev/sdX, /dev/sdXX, /dev/xvdX, or /dev/xvdXX).
      --Filters.BlockDeviceMappingLinkDates strings           The link dates for the BSU volumes mapped to the VMs (for example, 2016-01-23T18:45:30.000Z).
      --Filters.BlockDeviceMappingStates strings              The states for the BSU volumes (attaching | attached | detaching | detached).
      --Filters.BlockDeviceMappingVolumeIds strings           The volume IDs of the BSU volumes.
      --Filters.BootModes strings                             The boot modes of the VMs.
      --Filters.ClientTokens strings                          The idempotency tokens provided when launching the VMs.
      --Filters.CreationDates strings                         The dates when the VMs were launched.
      --Filters.ImageIds strings                              The IDs of the OMIs used to launch the VMs.
      --Filters.IsSourceDestChecked                           Whether the source/destination checking is enabled (true) or disabled (false).
      --Filters.KeypairNames strings                          The names of the keypairs used when launching the VMs.
      --Filters.LaunchNumbers ints                            The numbers for the VMs when launching a group of several VMs (for example, 0, 1, 2, and so on).
      --Filters.Lifecycles strings                            Whether the VMs are Spot Instances (spot).
      --Filters.NetIds strings                                The IDs of the Nets in which the VMs are running.
      --Filters.NicAccountIds strings                         The IDs of the NICs.
      --Filters.NicDescriptions strings                       The descriptions of the NICs.
      --Filters.NicIsSourceDestChecked                        Whether the source/destination checking is enabled (true) or disabled (false).
      --Filters.NicLinkNicDeleteOnVmDeletion                  Whether the NICs are deleted when the VMs they are attached to are deleted.
      --Filters.NicLinkNicDeviceNumbers ints                  The device numbers the NICs are attached to.
      --Filters.NicLinkNicLinkNicDates strings                The dates and times (UTC) when the NICs were attached to the VMs.
      --Filters.NicLinkNicLinkNicIds strings                  The IDs of the NIC attachments.
      --Filters.NicLinkNicStates strings                      The states of the attachments.
      --Filters.NicLinkNicVmAccountIds strings                The account IDs of the owners of the VMs the NICs are attached to.
      --Filters.NicLinkNicVmIds strings                       The IDs of the VMs the NICs are attached to.
      --Filters.NicLinkPublicIpAccountIds strings             The account IDs of the owners of the public IPs associated with the NICs.
      --Filters.NicLinkPublicIpLinkPublicIpIds strings        The association IDs returned when the public IPs were associated with the NICs.
      --Filters.NicLinkPublicIpPublicIpIds strings            The allocation IDs returned when the public IPs were allocated to their accounts.
      --Filters.NicLinkPublicIpPublicIps strings              The public IPs associated with the NICs.
      --Filters.NicMacAddresses strings                       The Media Access Control (MAC) addresses of the NICs.
      --Filters.NicNetIds strings                             The IDs of the Nets where the NICs are located.
      --Filters.NicNicIds strings                             The IDs of the NICs.
      --Filters.NicPrivateIpsLinkPublicIpAccountIds strings   The account IDs of the owner of the public IPs associated with the private IPs.
      --Filters.NicPrivateIpsLinkPublicIpIds strings          The public IPs associated with the private IPs.
      --Filters.NicPrivateIpsPrimaryIp                        Whether the private IPs are the primary IPs associated with the NICs.
      --Filters.NicPrivateIpsPrivateIps strings               The private IPs of the NICs.
      --Filters.NicSecurityGroupIds strings                   The IDs of the security groups associated with the NICs.
      --Filters.NicSecurityGroupNames strings                 The names of the security groups associated with the NICs.
      --Filters.NicStates strings                             The states of the NICs (available | in-use).
      --Filters.NicSubnetIds strings                          The IDs of the Subnets for the NICs.
      --Filters.NicSubregionNames strings                     The Subregions where the NICs are located.
      --Filters.Platforms strings                             The platforms.
      --Filters.PrivateIps strings                            The private IPs of the VMs.
      --Filters.ProductCodes strings                          The product codes associated with the OMI used to create the VMs.
      --Filters.PublicIps strings                             The public IPs of the VMs.
      --Filters.ReservationIds strings                        The IDs of the reservation of the VMs, created every time you launch VMs.
      --Filters.RootDeviceNames strings                       The names of the root devices for the VMs (for example, /dev/sda1)
      --Filters.RootDeviceTypes strings                       The root devices types used by the VMs (always ebs)
      --Filters.SecurityGroupIds strings                      The IDs of the security groups for the VMs (only in the public Cloud).
      --Filters.SecurityGroupNames strings                    The names of the security groups for the VMs (only in the public Cloud).
      --Filters.StateReasonCodes ints                         The reason codes for the state changes.
      --Filters.StateReasonMessages strings                   The messages describing the state changes.
      --Filters.StateReasons strings                          The reasons explaining the current states of the VMs.
      --Filters.SubnetIds strings                             The IDs of the Subnets for the VMs.
      --Filters.SubregionNames strings                        The names of the Subregions of the VMs.
      --Filters.TagKeys strings                               The keys of the tags associated with the VMs.
      --Filters.TagValues strings                             The values of the tags associated with the VMs.
      --Filters.Tags strings                                  The key/value combination of the tags associated with the VMs, in the following format: "Filters":{"Tags":["TAGKEY=TAGVALUE"]}.
      --Filters.Tenancies strings                             The tenancies of the VMs (dedicated | default | host).
      --Filters.TpmEnabled                                    Whether a virtual Trusted Platform Module (vTPM) is enabled (true) or disabled (false) on the VM.
      --Filters.VmIds strings                                 One or more IDs of VMs.
      --Filters.VmSecurityGroupIds strings                    The IDs of the security groups for the VMs.
      --Filters.VmSecurityGroupNames strings                  The names of the security group for the VMs.
      --Filters.VmStateCodes ints                             The state codes of the VMs: -1 (quarantine), 0 (pending), 16 (running), 32 (shutting-down), 48 (terminated), 64 (stopping), and 80 (stopped).
      --Filters.VmStateNames strings                          The state names of the VMs (pending | running | stopping | stopped | shutting-down | terminated | quarantine).
      --Filters.VmTypes strings                               The VM types (for example, t2.micro).
      --NextPageToken string                                  The token to request the next page of results.
      --ResultsPerPage int                                    The maximum number of logs returned in a single response (between 1 and 1000, both included).
  -h, --help                                                  help for ReadVms
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

###### Auto generated by spf13/cobra on 6-Mar-2026
