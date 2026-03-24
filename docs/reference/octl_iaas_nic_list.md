## octl iaas nic list

alias for api ReadNics

### Synopsis

> *alias for api ReadNics*

Lists one or more network interface cards (NICs).

A NIC is a virtual network interface that you can attach to a virtual machine (VM) in a Net.

```
octl iaas nic list [flags]
```

### Options

```
      --description strings                            The descriptions of the NICs.
  -h, --help                                           help for list
      --id strings                                     The IDs of the NICs.
      --is-source-dest-check                           Whether the source/destination checking is enabled (true) or disabled (false).
      --link-nic-delete-on-vm-deletion                 Whether the NICs are deleted when the VMs they are attached to are terminated.
      --link-nic-device-number ints                    The device numbers the NICs are attached to.
      --link-nic-link-nic-id strings                   The attachment IDs of the NICs.
      --link-nic-state strings                         The states of the attachments.
      --link-nic-vm-account-id strings                 The OUTSCALE account IDs of the owners of the VMs the NICs are attached to.
      --link-nic-vm-id strings                         The IDs of the VMs the NICs are attached to.
      --link-public-ip strings                         The public IPs associated with the NICs.
      --link-public-ip-account-id strings              The OUTSCALE account IDs of the owners of the public IPs associated with the NICs.
      --link-public-ip-link-public-ip-id strings       The association IDs returned when the public IPs were associated with the NICs.
      --link-public-ip-public-dns-name strings         The public DNS names associated with the public IPs.
      --link-public-ip-public-ip-id strings            The allocation IDs returned when the public IPs were allocated to their accounts.
      --mac-address strings                            The Media Access Control (MAC) addresses of the NICs.
      --net-id strings                                 The IDs of the Nets where the NICs are located.
      --private-dns-name strings                       The private DNS names associated with the primary private IPs.
      --private-ip-link-public-ip strings              The public IPs associated with the private IPs.
      --private-ip-link-public-ip-account-id strings   The OUTSCALE account IDs of the owner of the public IPs associated with the private IPs.
      --private-ip-primary-ip                          Whether the private IP is the primary IP associated with the NIC.
      --private-ip-private-ip strings                  The private IPs of the NICs.
      --security-group-id strings                      The IDs of the security groups associated with the NICs.
      --security-group-name strings                    The names of the security groups associated with the NICs.
      --state strings                                  The states of the NICs.
      --subnet-id strings                              The IDs of the Subnets for the NICs.
      --subregion strings                              The Subregions where the NICs are located.
      --tag strings                                    The key/value combination of the tags associated with the NICs, in the following format: "Filters":{"Tags":["TAGKEY=TAGVALUE"]}.
      --tag-key strings                                The keys of the tags associated with the NICs.
      --tag-value strings                              The values of the tags associated with the NICs.
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

* [octl iaas nic](octl_iaas_nic.md)	 - nic commands

