package cmd

import (
	"context"
	"errors"
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/outscale/goutils/sdk/tags"
	"github.com/outscale/octl/pkg/debug"
	"github.com/outscale/octl/pkg/messages"
	"github.com/outscale/octl/pkg/spinner"
	"github.com/outscale/octl/pkg/tree"
	"github.com/outscale/osc-sdk-go/v3/pkg/osc"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

var depsCmd = &cobra.Command{
	Use:     "dependencies net_id",
	Aliases: []string{"deps"},
	Short:   "Shows all dependencies of a net",
	Run:     displayNet(false),
}

func displayNet(failOnVms bool) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			messages.ExitErr(fmt.Errorf("not enough arguments for %s", cmd.Name()))
		}
		netID := args[0]

		p := loadProfile(cmd)
		cl, err := osc.NewClient(p, sdkOptions(cmd)...)
		if err != nil {
			messages.ExitErr(err)
		}
		r, hasVM, err := listResources(cmd.Context(), cl, netID)
		if err == nil {
			err = tree.WriteTo(r, os.Stdout)
		}
		if err != nil {
			messages.ExitErr(err)
		}

		teardownVMs, _ := cmd.Flags().GetBool("teardown-vms")
		if failOnVms && hasVM && !teardownVMs {
			messages.ExitErr(errors.New("cannot teardown a net having VMs unless --teardown-vms is set"))
		}
	}
}

type ResourceType string

const (
	Net             ResourceType = "net"
	Subnet          ResourceType = "subnet"
	InternetService ResourceType = "internet service"
	NetPeering      ResourceType = "net peering"
	VirtualGateway  ResourceType = "virtual gateway"
	VPNConnection   ResourceType = "vpn connection"
	NATService      ResourceType = "nat service"
	RouteTable      ResourceType = "route table"
	SecurityGroup   ResourceType = "security group"
	LoadBalancer    ResourceType = "load balancer"
	NetAccessPoint  ResourceType = "net access point"
	PublicIP        ResourceType = "public ip"
	VM              ResourceType = "vm"
	NIC             ResourceType = "nic"
)

type (
	deleteFunc func(ctx context.Context) error

	Resource struct {
		Type ResourceType
		ID   string
		Name string

		children []*Resource

		delete deleteFunc
	}
)

func (r *Resource) String() string {
	b := strings.Builder{}
	_, _ = b.WriteString(string(r.Type))
	_, _ = b.WriteString("/")
	_, _ = b.WriteString(r.ID)
	if r.Name != "" {
		_, _ = b.WriteString(" (")
		_, _ = b.WriteString(r.Name)
		_, _ = b.WriteString(")")
	}
	return b.String()
}

func (r *Resource) Children() []tree.Tree {
	return lo.Map(r.children, func(c *Resource, _ int) tree.Tree {
		return tree.Tree(c)
	})
}

func (r *Resource) addToChild(childID string, childType ResourceType, subr *Resource) bool {
	child, found := lo.Find(r.children, func(r *Resource) bool { return r.ID == childID && r.Type == childType })
	if found {
		child.children = append(child.children, subr)
		return true
	}
	for _, c := range r.children {
		if c.addToChild(childID, childType, subr) {
			return true
		}
	}
	debug.Println("parent", childType, childID, "not found for", subr.Type, subr.ID)
	return false
}

func listResources(ctx context.Context, cl *osc.Client, netID string) (*Resource, bool, error) {
	cancel := spinner.Run(ctx, "Building dependency tree...")
	defer cancel()
	var securityGroups []string
	nets, err := cl.ReadNets(ctx, osc.ReadNetsRequest{Filters: &osc.FiltersNet{NetIds: &[]string{netID}}})
	if err != nil {
		return nil, false, fmt.Errorf("list nets: %w", err)
	}
	if len(*nets.Nets) == 0 {
		return nil, false, errors.New("no net found")
	}
	root := &Resource{
		Type:   Net,
		ID:     netID,
		Name:   tags.Must(tags.GetName((*nets.Nets)[0].Tags)),
		delete: deleteNet(cl, netID),
	}

	subnets, err := cl.ReadSubnets(ctx, osc.ReadSubnetsRequest{Filters: &osc.FiltersSubnet{NetIds: &[]string{netID}}})
	if err != nil {
		return nil, false, fmt.Errorf("list subnets: %w", err)
	}
	for _, subnet := range *subnets.Subnets {
		sr := &Resource{
			Type:   Subnet,
			ID:     subnet.SubnetId,
			Name:   tags.Must(tags.GetName(subnet.Tags)),
			delete: deleteSubnet(cl, subnet.SubnetId),
		}
		root.children = append(root.children, sr)
	}

	iss, err := cl.ReadInternetServices(ctx, osc.ReadInternetServicesRequest{Filters: &osc.FiltersInternetService{LinkNetIds: &[]string{netID}}})
	if err != nil {
		return nil, false, fmt.Errorf("list internet services: %w", err)
	}
	for _, is := range *iss.InternetServices {
		root.children = append(root.children, &Resource{
			Type:   InternetService,
			ID:     is.InternetServiceId,
			Name:   tags.Must(tags.GetName(is.Tags)),
			delete: deleteInternetService(cl, is.InternetServiceId, netID),
		})
	}

	nps, err := cl.ReadNetPeerings(ctx, osc.ReadNetPeeringsRequest{Filters: &osc.FiltersNetPeering{SourceNetNetIds: &[]string{netID}}})
	if err != nil {
		return nil, false, fmt.Errorf("list net peerings: %w", err)
	}
	for _, np := range *nps.NetPeerings {
		root.children = append(root.children, &Resource{
			Type:   NetPeering,
			ID:     np.NetPeeringId,
			Name:   tags.Must(tags.GetName(np.Tags)),
			delete: deleteNetPeering(cl, np.NetPeeringId),
		})
	}
	nps, err = cl.ReadNetPeerings(ctx, osc.ReadNetPeeringsRequest{Filters: &osc.FiltersNetPeering{AccepterNetNetIds: &[]string{netID}}})
	if err != nil {
		return nil, false, fmt.Errorf("list net peerings: %w", err)
	}
	for _, np := range *nps.NetPeerings {
		root.children = append(root.children, &Resource{
			Type:   NetPeering,
			ID:     np.NetPeeringId,
			Name:   tags.Must(tags.GetName(np.Tags)),
			delete: deleteNetPeering(cl, np.NetPeeringId),
		})
	}

	naps, err := cl.ReadNetAccessPoints(ctx, osc.ReadNetAccessPointsRequest{Filters: &osc.FiltersNetAccessPoint{NetIds: &[]string{netID}}})
	if err != nil {
		return nil, false, fmt.Errorf("list net access points: %w", err)
	}
	for _, nap := range *naps.NetAccessPoints {
		root.children = append(root.children, &Resource{
			Type:   NetAccessPoint,
			ID:     nap.NetAccessPointId,
			Name:   tags.Must(tags.GetName(nap.Tags)),
			delete: deleteNetAccessPoint(cl, nap.NetAccessPointId),
		})
	}

	vgws, err := cl.ReadVirtualGateways(ctx, osc.ReadVirtualGatewaysRequest{Filters: &osc.FiltersVirtualGateway{LinkNetIds: &[]string{netID}}})
	if err != nil {
		return nil, false, fmt.Errorf("list virtual gateways: %w", err)
	}
	for _, vgw := range *vgws.VirtualGateways {
		root.children = append(root.children, &Resource{
			Type:   VirtualGateway,
			ID:     vgw.VirtualGatewayId,
			Name:   tags.Must(tags.GetName(vgw.Tags)),
			delete: deleteVirtualGateway(cl, vgw.VirtualGatewayId, netID),
		})
	}
	vgwids := lo.Map(*vgws.VirtualGateways, func(vgw osc.VirtualGateway, _ int) string {
		return vgw.VirtualGatewayId
	})
	vpns, err := cl.ReadVpnConnections(ctx, osc.ReadVpnConnectionsRequest{Filters: &osc.FiltersVpnConnection{VirtualGatewayIds: &vgwids}})
	if err != nil {
		return nil, false, fmt.Errorf("list net peerings: %w", err)
	}
	for _, vpn := range *vpns.VpnConnections {
		root.addToChild(vpn.VirtualGatewayId, VirtualGateway, &Resource{
			Type:   VPNConnection,
			ID:     vpn.VpnConnectionId,
			Name:   tags.Must(tags.GetName(vpn.Tags)),
			delete: deleteVpnConnection(cl, vpn.VpnConnectionId),
		})
	}

	nats, err := cl.ReadNatServices(ctx, osc.ReadNatServicesRequest{Filters: &osc.FiltersNatService{NetIds: &[]string{netID}}})
	if err != nil {
		return nil, false, fmt.Errorf("list NAT services: %w", err)
	}
	for _, nat := range *nats.NatServices {
		sr := &Resource{
			Type:   NATService,
			ID:     nat.NatServiceId,
			Name:   tags.Must(tags.GetName(nat.Tags)),
			delete: deleteNATService(cl, nat.NatServiceId),
		}
		root.addToChild(nat.SubnetId, Subnet, sr)
		for _, ip := range nat.PublicIps {
			pip := &Resource{
				Type:   PublicIP,
				ID:     ip.PublicIpId,
				Name:   ip.PublicIp,
				delete: deletePublicIP(cl, ip.PublicIpId),
			}
			sr.children = append(sr.children, pip)
		}
	}

	rtbls, err := cl.ReadRouteTables(ctx, osc.ReadRouteTablesRequest{Filters: &osc.FiltersRouteTable{NetIds: &[]string{netID}}})
	if err != nil {
		return nil, false, fmt.Errorf("list route tables: %w", err)
	}
	for _, rtbl := range *rtbls.RouteTables {
		links := lo.Map(rtbl.LinkRouteTables, func(l osc.LinkRouteTable, _ int) string { return l.LinkRouteTableId })
		sr := &Resource{
			Type:   RouteTable,
			ID:     rtbl.RouteTableId,
			Name:   tags.Must(tags.GetName(rtbl.Tags)),
			delete: deleteRouteTable(cl, rtbl.RouteTableId, links),
		}
		for _, link := range rtbl.LinkRouteTables {
			root.addToChild(link.SubnetId, Subnet, sr)
		}
	}

	lbus, err := cl.ReadLoadBalancers(ctx, osc.ReadLoadBalancersRequest{})
	if err != nil {
		return nil, false, fmt.Errorf("list load balancers: %w", err)
	}
	for _, lbu := range *lbus.LoadBalancers {
		if !slices.ContainsFunc(*subnets.Subnets, func(sn osc.Subnet) bool { return slices.Contains(lbu.Subnets, sn.SubnetId) }) {
			continue
		}

		sr := &Resource{
			Type:   LoadBalancer,
			ID:     lbu.LoadBalancerName,
			Name:   tags.Must(tags.GetName(lbu.Tags)),
			delete: deleteLoadBalancer(cl, lbu.LoadBalancerName),
		}
		for _, subnet := range lbu.Subnets {
			root.addToChild(subnet, Subnet, sr)
		}
		for _, sg := range lbu.SecurityGroups {
			securityGroups = append(securityGroups, sg)
			sgr := &Resource{
				Type:   SecurityGroup,
				ID:     sg,
				delete: deleteSecurityGroup(cl, sg),
			}
			sr.children = append(sr.children, sgr)
		}
		if lbu.PublicIp != nil {
			pips, err := cl.ReadPublicIps(ctx, osc.ReadPublicIpsRequest{
				Filters: &osc.FiltersPublicIp{PublicIps: &[]string{*lbu.PublicIp}},
			})
			if err != nil {
				return nil, false, fmt.Errorf("list public ip: %w", err)
			}
			for _, pip := range *pips.PublicIps {
				pipr := &Resource{
					Type:   PublicIP,
					ID:     pip.PublicIpId,
					Name:   pip.PublicIp,
					delete: deletePublicIP(cl, pip.PublicIpId),
				}
				sr.children = append(sr.children, pipr)
			}
		}
	}

	vms, err := cl.ReadVms(ctx, osc.ReadVmsRequest{Filters: &osc.FiltersVm{NetIds: &[]string{netID}}})
	if err != nil {
		return nil, false, fmt.Errorf("list vms: %w", err)
	}
	hasVM := len(*vms.Vms) > 0
	for _, vm := range *vms.Vms {
		sr := &Resource{
			Type:   VM,
			ID:     vm.VmId,
			Name:   tags.Must(tags.GetName(vm.Tags)),
			delete: deleteVm(cl, vm.VmId),
		}
		root.addToChild(*vm.SubnetId, Subnet, sr)
		for _, sg := range vm.SecurityGroups {
			securityGroups = append(securityGroups, sg.SecurityGroupId)
			sgr := &Resource{
				Type:   SecurityGroup,
				ID:     sg.SecurityGroupId,
				Name:   sg.SecurityGroupName,
				delete: deleteSecurityGroup(cl, sg.SecurityGroupId),
			}
			sr.children = append(sr.children, sgr)
		}
	}
	nics, err := cl.ReadNics(ctx, osc.ReadNicsRequest{Filters: &osc.FiltersNic{NetIds: &[]string{netID}}})
	if err != nil {
		return nil, false, fmt.Errorf("list nics: %w", err)
	}
	for _, nic := range *nics.Nics {
		nicr := &Resource{
			Type: NIC,
			ID:   nic.NicId,
		}
		if nic.LinkNic != nil {
			root.addToChild(nic.LinkNic.VmId, VM, nicr)
			nicr.delete = deleteNIC(cl, nic.NicId, nic.LinkNic.LinkNicId)
		} else {
			root.addToChild(nic.SubnetId, Subnet, nicr)
			nicr.delete = deleteNIC(cl, nic.NicId, "")
			// ths sg is already reported at vm level if linked
			for _, sg := range nic.SecurityGroups {
				securityGroups = append(securityGroups, sg.SecurityGroupId)
				sgr := &Resource{
					Type:   SecurityGroup,
					ID:     sg.SecurityGroupId,
					Name:   sg.SecurityGroupName,
					delete: deleteSecurityGroup(cl, sg.SecurityGroupId),
				}
				nicr.children = append(nicr.children, sgr)
			}
		}
		if nic.LinkPublicIp != nil {
			pips, err := cl.ReadPublicIps(ctx, osc.ReadPublicIpsRequest{
				Filters: &osc.FiltersPublicIp{PublicIps: &[]string{nic.LinkPublicIp.PublicIp}},
			})
			if err != nil {
				return nil, false, fmt.Errorf("list public ip: %w", err)
			}
			for _, pip := range *pips.PublicIps {
				pipr := &Resource{
					Type:   PublicIP,
					ID:     pip.PublicIpId,
					Name:   pip.PublicIp,
					delete: deletePublicIP(cl, pip.PublicIpId),
				}
				nicr.children = append(nicr.children, pipr)
			}
		}
	}

	sgs, err := cl.ReadSecurityGroups(ctx, osc.ReadSecurityGroupsRequest{Filters: &osc.FiltersSecurityGroup{NetIds: &[]string{netID}}})
	if err != nil {
		return nil, false, fmt.Errorf("list security groups: %w", err)
	}
	for _, sg := range *sgs.SecurityGroups {
		if slices.Contains(securityGroups, sg.SecurityGroupId) || sg.SecurityGroupName == "default" {
			continue
		}
		sgr := &Resource{
			Type:   SecurityGroup,
			ID:     sg.SecurityGroupId,
			Name:   sg.SecurityGroupName,
			delete: deleteSecurityGroup(cl, sg.SecurityGroupId),
		}
		root.children = append(root.children, sgr)
	}

	return root, hasVM, nil
}
