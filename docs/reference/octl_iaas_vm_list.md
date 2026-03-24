## octl iaas vm list

alias for api ReadVms

### Synopsis

> *alias for api ReadVms*

Lists one or more of your virtual machines (VMs).

If you provide one or more VM IDs, this action returns a description for all of these VMs. If you do not provide any VM ID, this action returns a description for all of the VMs that belong to you. If you provide an invalid VM ID, an error is returned. If you provide the ID of a VM that does not belong to you, the description of this VM is not included in the response. The refresh interval for data returned by this action is one hour, meaning that a terminated VM may appear in the response.

```
octl iaas vm list [flags]
```

### Options

```
      --architecture strings                               The architectures of the VMs (i386 | x86_64).
      --boot-mode strings                                  The boot modes of the VMs.
      --client-token strings                               The idempotency tokens provided when launching the VMs.
      --creation-date strings                              The dates when the VMs were launched.
  -h, --help                                               help for list
      --id strings                                         One or more IDs of VMs.
      --image-id strings                                   The IDs of the OMIs used to launch the VMs.
      --is-source-dest-checked                             Whether the source/destination checking is enabled (true) or disabled (false).
      --keypair-name strings                               The names of the keypairs used when launching the VMs.
      --launch-number ints                                 The numbers for the VMs when launching a group of several VMs (for example, 0, 1, 2, and so on).
      --lifecycle strings                                  Whether the VMs are Spot Instances (spot).
      --net-id strings                                     The IDs of the Nets in which the VMs are running.
      --nic-account-id strings                             The IDs of the NICs.
      --nic-description strings                            The descriptions of the NICs.
      --nic-id strings                                     The IDs of the NICs.
      --nic-is-source-dest-checked                         Whether the source/destination checking is enabled (true) or disabled (false).
      --nic-link-nic-delete-on-vm-deletion                 Whether the NICs are deleted when the VMs they are attached to are deleted.
      --nic-link-nic-device-number ints                    The device numbers the NICs are attached to.
      --nic-link-nic-link-nic-date strings                 The dates and times (UTC) when the NICs were attached to the VMs.
      --nic-link-nic-link-nic-id strings                   The IDs of the NIC attachments.
      --nic-link-nic-state strings                         The states of the attachments.
      --nic-link-nic-vm-account-id strings                 The OUTSCALE account IDs of the owners of the VMs the NICs are attached to.
      --nic-link-nic-vm-id strings                         The IDs of the VMs the NICs are attached to.
      --nic-link-public-ip strings                         The public IPs associated with the NICs.
      --nic-link-public-ip-account-id strings              The OUTSCALE account IDs of the owners of the public IPs associated with the NICs.
      --nic-link-public-ip-link-public-ip-id strings       The association IDs returned when the public IPs were associated with the NICs.
      --nic-link-public-ip-public-ip-id strings            The allocation IDs returned when the public IPs were allocated to their accounts.
      --nic-mac-address strings                            The Media Access Control (MAC) addresses of the NICs.
      --nic-net-id strings                                 The IDs of the Nets where the NICs are located.
      --nic-private-ip strings                             The private IPs of the NICs.
      --nic-private-ip-link-public-ip-account-id strings   The OUTSCALE account IDs of the owner of the public IPs associated with the private IPs.
      --nic-private-ip-link-public-ip-id strings           The public IPs associated with the private IPs.
      --nic-private-ip-primary-ip                          Whether the private IPs are the primary IPs associated with the NICs.
      --nic-security-group-id strings                      The IDs of the security groups associated with the NICs.
      --nic-security-group-name strings                    The names of the security groups associated with the NICs.
      --nic-state strings                                  The states of the NICs (available | in-use).
      --nic-subnet-id strings                              The IDs of the Subnets for the NICs.
      --nic-subregion strings                              The Subregions where the NICs are located.
      --platform strings                                   The platforms.
      --private-ip strings                                 The private IPs of the VMs.
      --product-code strings                               The product codes associated with the OMI used to create the VMs.
      --public-ip strings                                  The public IPs of the VMs.
      --reservation-id strings                             The IDs of the reservation of the VMs, created every time you launch VMs.
      --root-device-name strings                           The names of the root devices for the VMs (for example, /dev/sda1)
      --root-device-type strings                           The root devices types used by the VMs (always ebs)
      --security-group-id strings                          The IDs of the security groups for the VMs (only in the public Cloud).
      --security-group-name strings                        The names of the security groups for the VMs (only in the public Cloud).
      --state-code ints                                    The state codes of the VMs: -1 (quarantine), 0 (pending), 16 (running), 32 (shutting-down), 48 (terminated), 64 (stopping), and 80 (stopped).
      --state-name strings                                 The state names of the VMs (pending | running | stopping | stopped | shutting-down | terminated | quarantine).
      --state-reason strings                               The reasons explaining the current states of the VMs.
      --state-reason-code ints                             The reason codes for the state changes.
      --state-reason-message strings                       The messages describing the state changes.
      --subnet-id strings                                  The IDs of the Subnets for the VMs.
      --subregion strings                                  The names of the Subregions of the VMs.
      --tag strings                                        The key/value combination of the tags associated with the VMs, in the following format: "Filters":{"Tags":["TAGKEY=TAGVALUE"]}.
      --tag-key strings                                    The keys of the tags associated with the VMs.
      --tag-value strings                                  The values of the tags associated with the VMs.
      --tenancy strings                                    The tenancies of the VMs (dedicated | default | host).
      --tpm-enabled                                        Whether a virtual Trusted Platform Module (vTPM) is enabled (true) or disabled (false) on the VM.
      --type strings                                       The VM types (for example, t2.micro).
      --vm-security-group-id strings                       The IDs of the security groups for the VMs.
      --vm-security-group-name strings                     The names of the security group for the VMs.
      --volume-delete-on-vm-deletion                       Whether the BSU volumes are deleted when terminating the VMs.
      --volume-device-name strings                         The device names for the BSU volumes (in the format /dev/sdX, /dev/sdXX, /dev/xvdX, or /dev/xvdXX).
      --volume-id strings                                  The volume IDs of the BSU volumes.
      --volume-link-date strings                           The link dates for the BSU volumes mapped to the VMs (for example, 2016-01-23T18:45:30.000Z).
      --volume-state strings                               The states for the BSU volumes (attaching | attached | detaching | detached).
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

* [octl iaas vm](octl_iaas_vm.md)	 - vm commands

