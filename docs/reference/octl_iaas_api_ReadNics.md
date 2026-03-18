## octl iaas api ReadNics

Lists one or more network interface cards (NICs).

### Synopsis

Lists one or more network interface cards (NICs).

A NIC is a virtual network interface that you can attach to a virtual machine (VM) in a Net.

```
octl iaas api ReadNics [flags]
```

### Options

```
      --DryRun                                             If true, checks whether you have the required permissions to perform the action.
      --Filters.Descriptions strings                       The descriptions of the NICs.
      --Filters.IsSourceDestCheck                          Whether the source/destination checking is enabled (true) or disabled (false).
      --Filters.LinkNicDeleteOnVmDeletion                  Whether the NICs are deleted when the VMs they are attached to are terminated.
      --Filters.LinkNicDeviceNumbers ints                  The device numbers the NICs are attached to.
      --Filters.LinkNicLinkNicIds strings                  The attachment IDs of the NICs.
      --Filters.LinkNicStates strings                      The states of the attachments.
      --Filters.LinkNicVmAccountIds strings                The OUTSCALE account IDs of the owners of the VMs the NICs are attached to.
      --Filters.LinkNicVmIds strings                       The IDs of the VMs the NICs are attached to.
      --Filters.LinkPublicIpAccountIds strings             The OUTSCALE account IDs of the owners of the public IPs associated with the NICs.
      --Filters.LinkPublicIpLinkPublicIpIds strings        The association IDs returned when the public IPs were associated with the NICs.
      --Filters.LinkPublicIpPublicDnsNames strings         The public DNS names associated with the public IPs.
      --Filters.LinkPublicIpPublicIpIds strings            The allocation IDs returned when the public IPs were allocated to their accounts.
      --Filters.LinkPublicIpPublicIps strings              The public IPs associated with the NICs.
      --Filters.MacAddresses strings                       The Media Access Control (MAC) addresses of the NICs.
      --Filters.NetIds strings                             The IDs of the Nets where the NICs are located.
      --Filters.NicIds strings                             The IDs of the NICs.
      --Filters.PrivateDnsNames strings                    The private DNS names associated with the primary private IPs.
      --Filters.PrivateIpsLinkPublicIpAccountIds strings   The OUTSCALE account IDs of the owner of the public IPs associated with the private IPs.
      --Filters.PrivateIpsLinkPublicIpPublicIps strings    The public IPs associated with the private IPs.
      --Filters.PrivateIpsPrimaryIp                        Whether the private IP is the primary IP associated with the NIC.
      --Filters.PrivateIpsPrivateIps strings               The private IPs of the NICs.
      --Filters.SecurityGroupIds strings                   The IDs of the security groups associated with the NICs.
      --Filters.SecurityGroupNames strings                 The names of the security groups associated with the NICs.
      --Filters.States strings                             The states of the NICs.
      --Filters.SubnetIds strings                          The IDs of the Subnets for the NICs.
      --Filters.SubregionNames strings                     The Subregions where the NICs are located.
      --Filters.TagKeys strings                            The keys of the tags associated with the NICs.
      --Filters.TagValues strings                          The values of the tags associated with the NICs.
      --Filters.Tags strings                               The key/value combination of the tags associated with the NICs, in the following format: "Filters":{"Tags":["TAGKEY=TAGVALUE"]}.
      --NextPageToken string                               The token to request the next page of results.
      --ResultsPerPage int                                 The maximum number of logs returned in a single response (between 1 and 1000, both included).
  -h, --help                                               help for ReadNics
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

