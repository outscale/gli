## octl iaas api

iaas api calls

### Options

```
  -h, --help   help for api
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

* [octl iaas](octl_iaas.md)	 - OUTSCALE IaaS management
* [octl iaas api AcceptNetPeering](octl_iaas_api_AcceptNetPeering.md)	 - Accepts a Net peering request.
* [octl iaas api AddUserToUserGroup](octl_iaas_api_AddUserToUserGroup.md)	 - Adds a user to a specified group.
* [octl iaas api CheckAuthentication](octl_iaas_api_CheckAuthentication.md)	 - Validates the authenticity of the OUTSCALE account.
* [octl iaas api CreateAccessKey](octl_iaas_api_CreateAccessKey.md)	 - Creates an access key for either the root user or an EIM user.
* [octl iaas api CreateAccount](octl_iaas_api_CreateAccount.md)	 - Creates an OUTSCALE account.
* [octl iaas api CreateApiAccessRule](octl_iaas_api_CreateApiAccessRule.md)	 - Creates a rule to allow access to the API from your OUTSCALE account.
* [octl iaas api CreateCa](octl_iaas_api_CreateCa.md)	 - Creates a Client Certificate Authority (CA).
* [octl iaas api CreateClientGateway](octl_iaas_api_CreateClientGateway.md)	 - Provides information about your client gateway.
* [octl iaas api CreateDedicatedGroup](octl_iaas_api_CreateDedicatedGroup.md)	 - Creates a dedicated group for virtual machines (VMs).
* [octl iaas api CreateDhcpOptions](octl_iaas_api_CreateDhcpOptions.md)	 - Creates a set of DHCP options, that you can then associate with a Net using the [UpdateNet](#updatenet) method.
* [octl iaas api CreateDirectLink](octl_iaas_api_CreateDirectLink.md)	 - Creates a DirectLink between a customer network and a specified DirectLink location.
* [octl iaas api CreateDirectLinkInterface](octl_iaas_api_CreateDirectLinkInterface.md)	 - Creates a DirectLink interface.
* [octl iaas api CreateFlexibleGpu](octl_iaas_api_CreateFlexibleGpu.md)	 - Allocates a flexible GPU (fGPU) to your OUTSCALE account.
* [octl iaas api CreateImage](octl_iaas_api_CreateImage.md)	 - Creates an OUTSCALE machine image (OMI).
* [octl iaas api CreateImageExportTask](octl_iaas_api_CreateImageExportTask.md)	 - Exports an OUTSCALE machine image (OMI) to an OUTSCALE Object Storage (OOS) bucket.
* [octl iaas api CreateInternetService](octl_iaas_api_CreateInternetService.md)	 - Creates an internet service you can use with a Net.
* [octl iaas api CreateKeypair](octl_iaas_api_CreateKeypair.md)	 - Creates a keypair to use with your virtual machines (VMs).
* [octl iaas api CreateListenerRule](octl_iaas_api_CreateListenerRule.md)	 - Creates a rule for traffic redirection for the specified listener.
* [octl iaas api CreateLoadBalancer](octl_iaas_api_CreateLoadBalancer.md)	 - Creates a load balancer.
* [octl iaas api CreateLoadBalancerListeners](octl_iaas_api_CreateLoadBalancerListeners.md)	 - Creates one or more listeners for a specified load balancer.
* [octl iaas api CreateLoadBalancerPolicy](octl_iaas_api_CreateLoadBalancerPolicy.md)	 - Creates a stickiness policy with sticky session lifetimes defined by the browser lifetime.
* [octl iaas api CreateLoadBalancerTags](octl_iaas_api_CreateLoadBalancerTags.md)	 - Adds one or more tags to the specified load balancers.
* [octl iaas api CreateNatService](octl_iaas_api_CreateNatService.md)	 - Creates a network address translation (NAT) service in the specified public Subnet of a Net.
* [octl iaas api CreateNet](octl_iaas_api_CreateNet.md)	 - Creates a Net with a specified IP range.
* [octl iaas api CreateNetAccessPoint](octl_iaas_api_CreateNetAccessPoint.md)	 - Creates a Net access point to access an OUTSCALE service from this Net without using the Internet and public IPs.
* [octl iaas api CreateNetPeering](octl_iaas_api_CreateNetPeering.md)	 - Requests a Net peering between a Net you own and a peer Net that belongs to you or another OUTSCALE account.
* [octl iaas api CreateNic](octl_iaas_api_CreateNic.md)	 - Creates a network interface card (NIC) in the specified Subnet.
* [octl iaas api CreatePolicy](octl_iaas_api_CreatePolicy.md)	 - Creates a managed policy to apply to a user.
* [octl iaas api CreatePolicyVersion](octl_iaas_api_CreatePolicyVersion.md)	 - Creates a version of a specified managed policy.
* [octl iaas api CreateProductType](octl_iaas_api_CreateProductType.md)	 - Creates a product type you can associate with an OMI for consumption monitoring and billing purposes.
* [octl iaas api CreatePublicIp](octl_iaas_api_CreatePublicIp.md)	 - Acquires a public IP for your account.
* [octl iaas api CreateRoute](octl_iaas_api_CreateRoute.md)	 - Creates a route in a specified route table within a specified Net.
* [octl iaas api CreateRouteTable](octl_iaas_api_CreateRouteTable.md)	 - Creates a route table for a specified Net.
* [octl iaas api CreateSecurityGroup](octl_iaas_api_CreateSecurityGroup.md)	 - Creates a security group.
* [octl iaas api CreateSecurityGroupRule](octl_iaas_api_CreateSecurityGroupRule.md)	 - Adds one or more rules to a security group.
* [octl iaas api CreateServerCertificate](octl_iaas_api_CreateServerCertificate.md)	 - Creates a server certificate and its matching private key.
* [octl iaas api CreateSnapshot](octl_iaas_api_CreateSnapshot.md)	 - Creates a snapshot.
* [octl iaas api CreateSnapshotExportTask](octl_iaas_api_CreateSnapshotExportTask.md)	 - Exports a snapshot to an OUTSCALE Object Storage (OOS) bucket that belongs to you.
* [octl iaas api CreateSubnet](octl_iaas_api_CreateSubnet.md)	 - Creates a Subnet in an existing Net.
* [octl iaas api CreateTags](octl_iaas_api_CreateTags.md)	 - Adds one or more tags to the specified resources.
* [octl iaas api CreateUser](octl_iaas_api_CreateUser.md)	 - Creates an EIM user for your OUTSCALE account.
* [octl iaas api CreateUserGroup](octl_iaas_api_CreateUserGroup.md)	 - Creates a group to which you can add users.
* [octl iaas api CreateVirtualGateway](octl_iaas_api_CreateVirtualGateway.md)	 - Creates a virtual gateway.
* [octl iaas api CreateVmGroup](octl_iaas_api_CreateVmGroup.md)	 - > [WARNING] > This feature is currently under development and may not function properly.
* [octl iaas api CreateVmTemplate](octl_iaas_api_CreateVmTemplate.md)	 - > [WARNING] > This feature is currently under development and may not function properly.
* [octl iaas api CreateVms](octl_iaas_api_CreateVms.md)	 - Creates virtual machines (VMs), and then launches them.
* [octl iaas api CreateVolume](octl_iaas_api_CreateVolume.md)	 - Creates a Block Storage Unit (BSU) volume in a specified Region.
* [octl iaas api CreateVpnConnection](octl_iaas_api_CreateVpnConnection.md)	 - Creates a VPN connection between a specified virtual gateway and a specified client gateway.
* [octl iaas api CreateVpnConnectionRoute](octl_iaas_api_CreateVpnConnectionRoute.md)	 - Creates a static route to a VPN connection.
* [octl iaas api DeleteAccessKey](octl_iaas_api_DeleteAccessKey.md)	 - Deletes the specified access key of either the root user or an EIM user.
* [octl iaas api DeleteApiAccessRule](octl_iaas_api_DeleteApiAccessRule.md)	 - Deletes a specified API access rule.
* [octl iaas api DeleteCa](octl_iaas_api_DeleteCa.md)	 - Deletes a specified Client Certificate Authority (CA).
* [octl iaas api DeleteClientGateway](octl_iaas_api_DeleteClientGateway.md)	 - Deletes a client gateway.
* [octl iaas api DeleteDedicatedGroup](octl_iaas_api_DeleteDedicatedGroup.md)	 - Deletes a specified dedicated group of virtual machines (VMs).
* [octl iaas api DeleteDhcpOptions](octl_iaas_api_DeleteDhcpOptions.md)	 - Deletes a specified DHCP options set.
* [octl iaas api DeleteDirectLink](octl_iaas_api_DeleteDirectLink.md)	 - Deletes a specified DirectLink.
* [octl iaas api DeleteDirectLinkInterface](octl_iaas_api_DeleteDirectLinkInterface.md)	 - Deletes a specified DirectLink interface.
* [octl iaas api DeleteExportTask](octl_iaas_api_DeleteExportTask.md)	 - Deletes an export task.
* [octl iaas api DeleteFlexibleGpu](octl_iaas_api_DeleteFlexibleGpu.md)	 - Releases a flexible GPU (fGPU) from your OUTSCALE account.
* [octl iaas api DeleteImage](octl_iaas_api_DeleteImage.md)	 - Deletes an OUTSCALE machine image (OMI) so that you cannot use it anymore to launch virtual machines (VMs).
* [octl iaas api DeleteInternetService](octl_iaas_api_DeleteInternetService.md)	 - Deletes an internet service.
* [octl iaas api DeleteKeypair](octl_iaas_api_DeleteKeypair.md)	 - Deletes the specified keypair.
* [octl iaas api DeleteListenerRule](octl_iaas_api_DeleteListenerRule.md)	 - Deletes a listener rule.
* [octl iaas api DeleteLoadBalancer](octl_iaas_api_DeleteLoadBalancer.md)	 - Deletes a specified load balancer.
* [octl iaas api DeleteLoadBalancerListeners](octl_iaas_api_DeleteLoadBalancerListeners.md)	 - Deletes listeners of a specified load balancer.
* [octl iaas api DeleteLoadBalancerPolicy](octl_iaas_api_DeleteLoadBalancerPolicy.md)	 - Deletes a specified policy from a load balancer.
* [octl iaas api DeleteLoadBalancerTags](octl_iaas_api_DeleteLoadBalancerTags.md)	 - Deletes one or more tags from the specified load balancers.
* [octl iaas api DeleteNatService](octl_iaas_api_DeleteNatService.md)	 - Deletes a specified network address translation (NAT) service.
* [octl iaas api DeleteNet](octl_iaas_api_DeleteNet.md)	 - Deletes a specified Net.
* [octl iaas api DeleteNetAccessPoint](octl_iaas_api_DeleteNetAccessPoint.md)	 - Deletes a specified Net access point.
* [octl iaas api DeleteNetPeering](octl_iaas_api_DeleteNetPeering.md)	 - Deletes a Net peering.
* [octl iaas api DeleteNic](octl_iaas_api_DeleteNic.md)	 - Deletes the specified network interface card (NIC).
* [octl iaas api DeletePolicy](octl_iaas_api_DeletePolicy.md)	 - Deletes a managed policy.
* [octl iaas api DeletePolicyVersion](octl_iaas_api_DeletePolicyVersion.md)	 - Deletes a specified version of a managed policy, if it is not set as the default one.
* [octl iaas api DeleteProductType](octl_iaas_api_DeleteProductType.md)	 - Deletes a specified product type that belongs to you.
* [octl iaas api DeletePublicIp](octl_iaas_api_DeletePublicIp.md)	 - Releases a public IP.
* [octl iaas api DeleteRoute](octl_iaas_api_DeleteRoute.md)	 - Deletes a route from a specified route table.
* [octl iaas api DeleteRouteTable](octl_iaas_api_DeleteRouteTable.md)	 - Deletes a specified route table.
* [octl iaas api DeleteSecurityGroup](octl_iaas_api_DeleteSecurityGroup.md)	 - Deletes a specified security group.
* [octl iaas api DeleteSecurityGroupRule](octl_iaas_api_DeleteSecurityGroupRule.md)	 - Deletes one or more inbound or outbound rules from a security group.
* [octl iaas api DeleteServerCertificate](octl_iaas_api_DeleteServerCertificate.md)	 - Deletes a specified server certificate.
* [octl iaas api DeleteSnapshot](octl_iaas_api_DeleteSnapshot.md)	 - Deletes a specified snapshot.
* [octl iaas api DeleteSubnet](octl_iaas_api_DeleteSubnet.md)	 - Deletes a specified Subnet.
* [octl iaas api DeleteTags](octl_iaas_api_DeleteTags.md)	 - Deletes one or more tags from the specified resources.
* [octl iaas api DeleteUser](octl_iaas_api_DeleteUser.md)	 - Deletes a specified EIM user.
* [octl iaas api DeleteUserGroup](octl_iaas_api_DeleteUserGroup.md)	 - Deletes a specified user group.
* [octl iaas api DeleteUserGroupPolicy](octl_iaas_api_DeleteUserGroupPolicy.md)	 - Deletes a specified inline policy from a specific group.
* [octl iaas api DeleteUserPolicy](octl_iaas_api_DeleteUserPolicy.md)	 - Deletes a specified inline policy from a specific user.
* [octl iaas api DeleteVirtualGateway](octl_iaas_api_DeleteVirtualGateway.md)	 - Deletes a specified virtual gateway.
* [octl iaas api DeleteVmGroup](octl_iaas_api_DeleteVmGroup.md)	 - > [WARNING] > This feature is currently under development and may not function properly.
* [octl iaas api DeleteVmTemplate](octl_iaas_api_DeleteVmTemplate.md)	 - > [WARNING] > This feature is currently under development and may not function properly.
* [octl iaas api DeleteVms](octl_iaas_api_DeleteVms.md)	 - Terminates one or more virtual machines (VMs).
* [octl iaas api DeleteVolume](octl_iaas_api_DeleteVolume.md)	 - Deletes a specified Block Storage Unit (BSU) volume.
* [octl iaas api DeleteVpnConnection](octl_iaas_api_DeleteVpnConnection.md)	 - Deletes a specified VPN connection.
* [octl iaas api DeleteVpnConnectionRoute](octl_iaas_api_DeleteVpnConnectionRoute.md)	 - Deletes a static route to a VPN connection previously created using the CreateVpnConnectionRoute method.
* [octl iaas api DeregisterVmsInLoadBalancer](octl_iaas_api_DeregisterVmsInLoadBalancer.md)	 - > [WARNING] > Deprecated: This call is deprecated and will be removed.
* [octl iaas api DisableOutscaleLogin](octl_iaas_api_DisableOutscaleLogin.md)	 - Disables the possibility of logging in using the Outscale credentials of your root user when identity federation is activated.
* [octl iaas api DisableOutscaleLoginForUsers](octl_iaas_api_DisableOutscaleLoginForUsers.md)	 - Disables the possibility of logging in using the Outscale credentials of your EIM users when identity federation is activated.
* [octl iaas api DisableOutscaleLoginPerUsers](octl_iaas_api_DisableOutscaleLoginPerUsers.md)	 - Disables the possibility for one or more specific users to log in using their Outscale credentials when identity federation is activated.
* [octl iaas api EnableOutscaleLogin](octl_iaas_api_EnableOutscaleLogin.md)	 - Enables the possibility of logging in using the Outscale credentials of your root user when identity federation is activated.
* [octl iaas api EnableOutscaleLoginForUsers](octl_iaas_api_EnableOutscaleLoginForUsers.md)	 - Enables the possibility for all your EIM users to log in using their Outscale credentials when identity federation is activated.
* [octl iaas api EnableOutscaleLoginPerUsers](octl_iaas_api_EnableOutscaleLoginPerUsers.md)	 - Enables the possibility for one or more specific users to log in using their Outscale credentials when identity federation is activated.
* [octl iaas api LinkFlexibleGpu](octl_iaas_api_LinkFlexibleGpu.md)	 - Attaches one of your allocated flexible GPUs (fGPUs) to one of your virtual machines (VMs).
* [octl iaas api LinkInternetService](octl_iaas_api_LinkInternetService.md)	 - Attaches an internet service to a Net.
* [octl iaas api LinkLoadBalancerBackendMachines](octl_iaas_api_LinkLoadBalancerBackendMachines.md)	 - Attaches one or more virtual machines (VMs) to a specified load balancer.
* [octl iaas api LinkManagedPolicyToUserGroup](octl_iaas_api_LinkManagedPolicyToUserGroup.md)	 - Links a managed policy to a specific group.
* [octl iaas api LinkNic](octl_iaas_api_LinkNic.md)	 - Attaches a network interface card (NIC) to a virtual machine (VM).
* [octl iaas api LinkPolicy](octl_iaas_api_LinkPolicy.md)	 - Links a managed policy to a specific user.
* [octl iaas api LinkPrivateIps](octl_iaas_api_LinkPrivateIps.md)	 - Assigns one or more secondary private IPs to a specified network interface card (NIC).
* [octl iaas api LinkPublicIp](octl_iaas_api_LinkPublicIp.md)	 - Associates a public IP with a virtual machine (VM) or a network interface card (NIC), in the public Cloud or in a Net.
* [octl iaas api LinkRouteTable](octl_iaas_api_LinkRouteTable.md)	 - Associates a Subnet with a route table.
* [octl iaas api LinkVirtualGateway](octl_iaas_api_LinkVirtualGateway.md)	 - Attaches a virtual gateway to a Net.
* [octl iaas api LinkVolume](octl_iaas_api_LinkVolume.md)	 - Attaches a Block Storage Unit (BSU) volume to a virtual machine (VM).
* [octl iaas api PutUserGroupPolicy](octl_iaas_api_PutUserGroupPolicy.md)	 - Creates or updates an inline policy included in a specified group.
* [octl iaas api PutUserPolicy](octl_iaas_api_PutUserPolicy.md)	 - Creates or updates an inline policy included in a specified user.
* [octl iaas api ReadAccessKeys](octl_iaas_api_ReadAccessKeys.md)	 - Lists the access key IDs of either your root user or an EIM user.
* [octl iaas api ReadAccounts](octl_iaas_api_ReadAccounts.md)	 - Gets information about the account that sent the request.
* [octl iaas api ReadAdminPassword](octl_iaas_api_ReadAdminPassword.md)	 - Gets the administrator password for a Windows running virtual machine (VM).
* [octl iaas api ReadApiAccessPolicy](octl_iaas_api_ReadApiAccessPolicy.md)	 - Gets information about the API access policy of your OUTSCALE account.
* [octl iaas api ReadApiAccessRules](octl_iaas_api_ReadApiAccessRules.md)	 - Lists one or more API access rules.
* [octl iaas api ReadApiLogs](octl_iaas_api_ReadApiLogs.md)	 - Lists the logs of the API calls you have performed with this OUTSCALE account.
* [octl iaas api ReadCO2EmissionAccount](octl_iaas_api_ReadCO2EmissionAccount.md)	 - Gets information about the estimated carbon footprint of your account for the current Region within the specified time period.
* [octl iaas api ReadCas](octl_iaas_api_ReadCas.md)	 - Gets information about one or more of your Client Certificate Authorities (CAs).
* [octl iaas api ReadCatalog](octl_iaas_api_ReadCatalog.md)	 - Returns the price list of OUTSCALE services for the current Region.
* [octl iaas api ReadCatalogs](octl_iaas_api_ReadCatalogs.md)	 - Returns the price list of OUTSCALE services for the current Region within a specific time period.
* [octl iaas api ReadClientGateways](octl_iaas_api_ReadClientGateways.md)	 - Lists one or more of your client gateways.
* [octl iaas api ReadConsoleOutput](octl_iaas_api_ReadConsoleOutput.md)	 - Gets the console output for a virtual machine (VM).
* [octl iaas api ReadConsumptionAccount](octl_iaas_api_ReadConsumptionAccount.md)	 - Gets information about the consumption of your OUTSCALE account for each billable resource within the specified time period.
* [octl iaas api ReadDedicatedGroups](octl_iaas_api_ReadDedicatedGroups.md)	 - List one or more dedicated groups of virtual machines (VMs).
* [octl iaas api ReadDhcpOptions](octl_iaas_api_ReadDhcpOptions.md)	 - Gets information about the content of one or more DHCP options sets.
* [octl iaas api ReadDirectLinkInterfaces](octl_iaas_api_ReadDirectLinkInterfaces.md)	 - Lists one or more of your DirectLink interfaces.
* [octl iaas api ReadDirectLinks](octl_iaas_api_ReadDirectLinks.md)	 - Lists all DirectLinks in the Region.
* [octl iaas api ReadEntitiesLinkedToPolicy](octl_iaas_api_ReadEntitiesLinkedToPolicy.md)	 - Lists all entities (account, users, or user groups) linked to a specific managed policy.
* [octl iaas api ReadFlexibleGpuCatalog](octl_iaas_api_ReadFlexibleGpuCatalog.md)	 - Lists all flexible GPUs available in the public catalog.
* [octl iaas api ReadFlexibleGpus](octl_iaas_api_ReadFlexibleGpus.md)	 - Lists one or more flexible GPUs (fGPUs) allocated to your OUTSCALE account.
* [octl iaas api ReadImageExportTasks](octl_iaas_api_ReadImageExportTasks.md)	 - Lists one or more image export tasks.
* [octl iaas api ReadImages](octl_iaas_api_ReadImages.md)	 - Lists one or more OUTSCALE machine images (OMIs) you can use.
* [octl iaas api ReadInternetServices](octl_iaas_api_ReadInternetServices.md)	 - Lists one or more of your internet services.
* [octl iaas api ReadKeypairs](octl_iaas_api_ReadKeypairs.md)	 - Lists one or more of your keypairs.
* [octl iaas api ReadLinkedPolicies](octl_iaas_api_ReadLinkedPolicies.md)	 - Lists the managed policies linked to a specified user.
* [octl iaas api ReadListenerRules](octl_iaas_api_ReadListenerRules.md)	 - Lists one or more listener rules.
* [octl iaas api ReadLoadBalancerTags](octl_iaas_api_ReadLoadBalancerTags.md)	 - Lists the tags associated with one or more specified load balancers.
* [octl iaas api ReadLoadBalancers](octl_iaas_api_ReadLoadBalancers.md)	 - Lists one or more load balancers and their attributes.
* [octl iaas api ReadLocations](octl_iaas_api_ReadLocations.md)	 - Lists the locations, corresponding to datacenters, where you can set up a DirectLink.
* [octl iaas api ReadManagedPoliciesLinkedToUserGroup](octl_iaas_api_ReadManagedPoliciesLinkedToUserGroup.md)	 - Lists the managed policies linked to a specified group.
* [octl iaas api ReadNatServices](octl_iaas_api_ReadNatServices.md)	 - Lists one or more network address translation (NAT) services.
* [octl iaas api ReadNetAccessPointServices](octl_iaas_api_ReadNetAccessPointServices.md)	 - Lists OUTSCALE services available to create Net access points.
* [octl iaas api ReadNetAccessPoints](octl_iaas_api_ReadNetAccessPoints.md)	 - Lists one or more Net access points.
* [octl iaas api ReadNetPeerings](octl_iaas_api_ReadNetPeerings.md)	 - Lists one or more peering connections between two Nets.
* [octl iaas api ReadNets](octl_iaas_api_ReadNets.md)	 - Lists one or more Nets.
* [octl iaas api ReadNics](octl_iaas_api_ReadNics.md)	 - Lists one or more network interface cards (NICs).
* [octl iaas api ReadPolicies](octl_iaas_api_ReadPolicies.md)	 - Lists all the managed policies available for your OUTSCALE account.
* [octl iaas api ReadPolicy](octl_iaas_api_ReadPolicy.md)	 - Lists information about a specified managed policy.
* [octl iaas api ReadPolicyVersion](octl_iaas_api_ReadPolicyVersion.md)	 - Lists information about a specified version of a managed policy.
* [octl iaas api ReadPolicyVersions](octl_iaas_api_ReadPolicyVersions.md)	 - Lists information about all the policy versions of a specified managed policy.
* [octl iaas api ReadProductTypes](octl_iaas_api_ReadProductTypes.md)	 - Lists one or more product types.
* [octl iaas api ReadPublicCatalog](octl_iaas_api_ReadPublicCatalog.md)	 - Returns the price list of OUTSCALE products and services for the Region specified in the endpoint of the request.
* [octl iaas api ReadPublicIpRanges](octl_iaas_api_ReadPublicIpRanges.md)	 - Gets the public IPv4 addresses in CIDR notation for the Region specified in the endpoint of the request.
* [octl iaas api ReadPublicIps](octl_iaas_api_ReadPublicIps.md)	 - Lists one or more public IPs allocated to your OUTSCALE account.
* [octl iaas api ReadQuotas](octl_iaas_api_ReadQuotas.md)	 - Lists one or more of your quotas.
* [octl iaas api ReadRegions](octl_iaas_api_ReadRegions.md)	 - Lists one or more Regions of the OUTSCALE Cloud.
* [octl iaas api ReadRouteTables](octl_iaas_api_ReadRouteTables.md)	 - Lists one or more of your route tables.
* [octl iaas api ReadSecurityGroups](octl_iaas_api_ReadSecurityGroups.md)	 - Lists one or more security groups.
* [octl iaas api ReadServerCertificates](octl_iaas_api_ReadServerCertificates.md)	 - Lists your server certificates.
* [octl iaas api ReadSnapshotExportTasks](octl_iaas_api_ReadSnapshotExportTasks.md)	 - Lists one or more snapshot export tasks.
* [octl iaas api ReadSnapshots](octl_iaas_api_ReadSnapshots.md)	 - Lists one or more snapshots that are available to you and the permissions to create volumes from them.
* [octl iaas api ReadSubnets](octl_iaas_api_ReadSubnets.md)	 - Lists one or more of your Subnets.
* [octl iaas api ReadSubregions](octl_iaas_api_ReadSubregions.md)	 - Lists one or more of the enabled Subregions that you can access in the current Region.
* [octl iaas api ReadTags](octl_iaas_api_ReadTags.md)	 - Lists one or more tags for your resources.
* [octl iaas api ReadUnitPrice](octl_iaas_api_ReadUnitPrice.md)	 - Gets unit price information for the specified parameters.
* [octl iaas api ReadUserGroup](octl_iaas_api_ReadUserGroup.md)	 - Lists information about a specified user group, including its users.
* [octl iaas api ReadUserGroupPolicies](octl_iaas_api_ReadUserGroupPolicies.md)	 - Lists the names of the inline policies embedded in a specific group.
* [octl iaas api ReadUserGroupPolicy](octl_iaas_api_ReadUserGroupPolicy.md)	 - Returns information about an inline policy included in a specified group.
* [octl iaas api ReadUserGroups](octl_iaas_api_ReadUserGroups.md)	 - Lists all the user groups of the OUTSCALE account.
* [octl iaas api ReadUserGroupsPerUser](octl_iaas_api_ReadUserGroupsPerUser.md)	 - Lists the groups a specified user belongs to.
* [octl iaas api ReadUserPolicies](octl_iaas_api_ReadUserPolicies.md)	 - Lists the names of inline policies included in a specified user.
* [octl iaas api ReadUserPolicy](octl_iaas_api_ReadUserPolicy.md)	 - Returns information about an inline policy included in a specified user.
* [octl iaas api ReadUsers](octl_iaas_api_ReadUsers.md)	 - Lists all EIM users in the OUTSCALE account.
* [octl iaas api ReadVirtualGateways](octl_iaas_api_ReadVirtualGateways.md)	 - Lists one or more virtual gateways.
* [octl iaas api ReadVmGroups](octl_iaas_api_ReadVmGroups.md)	 - > [WARNING] > This feature is currently under development and may not function properly.
* [octl iaas api ReadVmTemplates](octl_iaas_api_ReadVmTemplates.md)	 - > [WARNING] > This feature is currently under development and may not function properly.
* [octl iaas api ReadVmTypes](octl_iaas_api_ReadVmTypes.md)	 - Lists one or more predefined VM types.
* [octl iaas api ReadVms](octl_iaas_api_ReadVms.md)	 - Lists one or more of your virtual machines (VMs).
* [octl iaas api ReadVmsHealth](octl_iaas_api_ReadVmsHealth.md)	 - Lists the state of one or more backend virtual machines (VMs) registered with a specified load balancer.
* [octl iaas api ReadVmsState](octl_iaas_api_ReadVmsState.md)	 - Lists the status of one or more virtual machines (VMs).
* [octl iaas api ReadVolumeUpdateTasks](octl_iaas_api_ReadVolumeUpdateTasks.md)	 - Lists one or more specified tasks of volume update.
* [octl iaas api ReadVolumes](octl_iaas_api_ReadVolumes.md)	 - Lists one or more specified Block Storage Unit (BSU) volumes.
* [octl iaas api ReadVpnConnections](octl_iaas_api_ReadVpnConnections.md)	 - Lists one or more VPN connections.
* [octl iaas api RebootVms](octl_iaas_api_RebootVms.md)	 - Reboots one or more virtual machines (VMs).
* [octl iaas api RegisterVmsInLoadBalancer](octl_iaas_api_RegisterVmsInLoadBalancer.md)	 - > [WARNING] > Deprecated: This call is deprecated and will be removed.
* [octl iaas api RejectNetPeering](octl_iaas_api_RejectNetPeering.md)	 - Rejects a Net peering request.
* [octl iaas api RemoveUserFromUserGroup](octl_iaas_api_RemoveUserFromUserGroup.md)	 - Removes a specified user from a specified group.
* [octl iaas api ScaleDownVmGroup](octl_iaas_api_ScaleDownVmGroup.md)	 - > [WARNING] > This feature is currently under development and may not function properly.
* [octl iaas api ScaleUpVmGroup](octl_iaas_api_ScaleUpVmGroup.md)	 - > [WARNING] > This feature is currently under development and may not function properly.
* [octl iaas api SetDefaultPolicyVersion](octl_iaas_api_SetDefaultPolicyVersion.md)	 - Sets a specified version of a managed policy as the default (operative) one.
* [octl iaas api StartVms](octl_iaas_api_StartVms.md)	 - Start one or more virtual machines (VMs).
* [octl iaas api StopVms](octl_iaas_api_StopVms.md)	 - Stops one or more running virtual machines (VMs).
* [octl iaas api UnlinkFlexibleGpu](octl_iaas_api_UnlinkFlexibleGpu.md)	 - Detaches a flexible GPU (fGPU) from a virtual machine (VM).
* [octl iaas api UnlinkInternetService](octl_iaas_api_UnlinkInternetService.md)	 - Detaches an internet service from a Net.
* [octl iaas api UnlinkLoadBalancerBackendMachines](octl_iaas_api_UnlinkLoadBalancerBackendMachines.md)	 - Detaches one or more backend virtual machines (VMs) from a load balancer.
* [octl iaas api UnlinkManagedPolicyFromUserGroup](octl_iaas_api_UnlinkManagedPolicyFromUserGroup.md)	 - Unlinks a managed policy from a specific group.
* [octl iaas api UnlinkNic](octl_iaas_api_UnlinkNic.md)	 - Detaches a network interface card (NIC) from a virtual machine (VM).
* [octl iaas api UnlinkPolicy](octl_iaas_api_UnlinkPolicy.md)	 - Removes a managed policy from a specific user.
* [octl iaas api UnlinkPrivateIps](octl_iaas_api_UnlinkPrivateIps.md)	 - Unassigns one or more secondary private IPs from a network interface card (NIC).
* [octl iaas api UnlinkPublicIp](octl_iaas_api_UnlinkPublicIp.md)	 - Disassociates a public IP from the virtual machine (VM) or network interface card (NIC) it is associated with.
* [octl iaas api UnlinkRouteTable](octl_iaas_api_UnlinkRouteTable.md)	 - Disassociates a Subnet from a route table.
* [octl iaas api UnlinkVirtualGateway](octl_iaas_api_UnlinkVirtualGateway.md)	 - Detaches a virtual gateway from a Net.
* [octl iaas api UnlinkVolume](octl_iaas_api_UnlinkVolume.md)	 - Detaches a Block Storage Unit (BSU) volume from a virtual machine (VM).
* [octl iaas api UpdateAccessKey](octl_iaas_api_UpdateAccessKey.md)	 - Modifies the attributes of the specified access key of either the root user or an EIM user.
* [octl iaas api UpdateAccount](octl_iaas_api_UpdateAccount.md)	 - Updates the OUTSCALE account information for the account that sends the request.
* [octl iaas api UpdateApiAccessPolicy](octl_iaas_api_UpdateApiAccessPolicy.md)	 - Updates the API access policy of your OUTSCALE account.
* [octl iaas api UpdateApiAccessRule](octl_iaas_api_UpdateApiAccessRule.md)	 - Modifies a specified API access rule.
* [octl iaas api UpdateCa](octl_iaas_api_UpdateCa.md)	 - Modifies the specified attribute of a Client Certificate Authority (CA).
* [octl iaas api UpdateDedicatedGroup](octl_iaas_api_UpdateDedicatedGroup.md)	 - Modifies the name of a specified dedicated group.
* [octl iaas api UpdateDirectLinkInterface](octl_iaas_api_UpdateDirectLinkInterface.md)	 - Modifies the maximum transmission unit (MTU) of a DirectLink interface.
* [octl iaas api UpdateFlexibleGpu](octl_iaas_api_UpdateFlexibleGpu.md)	 - Modifies a flexible GPU (fGPU) behavior.
* [octl iaas api UpdateImage](octl_iaas_api_UpdateImage.md)	 - Modifies the access permissions for an OUTSCALE machine image (OMI).
* [octl iaas api UpdateListenerRule](octl_iaas_api_UpdateListenerRule.md)	 - Updates the pattern of the listener rule.
* [octl iaas api UpdateLoadBalancer](octl_iaas_api_UpdateLoadBalancer.md)	 - Modifies the specified attribute of a load balancer.
* [octl iaas api UpdateNet](octl_iaas_api_UpdateNet.md)	 - Associates a DHCP options set with a specified Net.
* [octl iaas api UpdateNetAccessPoint](octl_iaas_api_UpdateNetAccessPoint.md)	 - Modifies the attributes of a Net access point.
* [octl iaas api UpdateNic](octl_iaas_api_UpdateNic.md)	 - Modifies the specified network interface card (NIC).
* [octl iaas api UpdateRoute](octl_iaas_api_UpdateRoute.md)	 - Replaces an existing route within a route table in a Net.
* [octl iaas api UpdateRoutePropagation](octl_iaas_api_UpdateRoutePropagation.md)	 - Configures the propagation of routes to a specified route table of a Net by a virtual gateway.
* [octl iaas api UpdateRouteTableLink](octl_iaas_api_UpdateRouteTableLink.md)	 - Replaces the route table associated with a specific Subnet in a Net with another one.
* [octl iaas api UpdateServerCertificate](octl_iaas_api_UpdateServerCertificate.md)	 - Modifies the name and/or the path of a specified server certificate.
* [octl iaas api UpdateSnapshot](octl_iaas_api_UpdateSnapshot.md)	 - Modifies the permissions for a specified snapshot.
* [octl iaas api UpdateSubnet](octl_iaas_api_UpdateSubnet.md)	 - Modifies the specified attribute of a Subnet.
* [octl iaas api UpdateUser](octl_iaas_api_UpdateUser.md)	 - Modifies the name and/or the path of a specified EIM user.
* [octl iaas api UpdateUserGroup](octl_iaas_api_UpdateUserGroup.md)	 - Modifies the name and/or the path of a specified group.
* [octl iaas api UpdateVm](octl_iaas_api_UpdateVm.md)	 - Modifies the specified attributes of a virtual machine (VM).
* [octl iaas api UpdateVmGroup](octl_iaas_api_UpdateVmGroup.md)	 - > [WARNING] > This feature is currently under development and may not function properly.
* [octl iaas api UpdateVmTemplate](octl_iaas_api_UpdateVmTemplate.md)	 - > [WARNING] > This feature is currently under development and may not function properly.
* [octl iaas api UpdateVolume](octl_iaas_api_UpdateVolume.md)	 - Modifies the specified attributes of a volume.
* [octl iaas api UpdateVpnConnection](octl_iaas_api_UpdateVpnConnection.md)	 - Modifies the specified attributes of a VPN connection.

