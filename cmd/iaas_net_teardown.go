/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>

SPDX-License-Identifier: BSD-3-Clause
*/
package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/outscale/octl/pkg/debug"
	"github.com/outscale/octl/pkg/messages"
	"github.com/outscale/octl/pkg/spinner"
	"github.com/outscale/osc-sdk-go/v3/pkg/osc"
	"github.com/spf13/cobra"
)

var teardownCmd = &cobra.Command{
	Use:   "teardown net_id",
	Short: "Tears down a net and its subresources",
	Args:  cobra.ExactArgs(1),
	Run:   teardownNet,
}

func teardownNet(cmd *cobra.Command, args []string) {
	debug.Println(cmd.Name() + " called")

	displayNet(cmd, args, true)
	if !messages.Prompt("Are you sure you want to teardown this net ?") {
		return
	}

	netID := args[0]
	p := loadProfile(cmd)
	cl, err := osc.NewClient(p, sdkOptions(cmd)...)
	if err != nil {
		messages.ExitErr(err)
	}
	r, _, err := listResources(cmd.Context(), cl, netID)
	if err != nil {
		messages.ExitErr(err)
	}
	tmout, _ := cmd.Flags().GetDuration("timeout")
	doTeardown(cmd.Context(), r, tmout)
}

var deletionOrder = []ResourceType{
	LoadBalancer,
	VM,
	RouteTable,
	NetAccessPoint,
	NATService,
	VPNConnection,
	VirtualGateway,
	NetPeering,
	NIC,
	PublicIP,
	InternetService,
	SecurityGroup,
	Subnet,
	Net,
}

func doTeardown(ctx context.Context, r *Resource, tmout time.Duration) {
	for _, typ := range deletionOrder {
		doTeardownType(ctx, r, typ, tmout)
	}
}

func doTeardownType(ctx context.Context, r *Resource, typ ResourceType, tmout time.Duration) {
	for _, c := range r.children {
		doTeardownType(ctx, c, typ, tmout)
	}
	if r.delete == nil || r.Type != typ {
		return
	}
	start := time.Now()
	ctx, cancel := context.WithTimeout(ctx, tmout)
	defer cancel()
	spinCancel := spinner.Run(ctx, "Deleting "+r.String()+" ...", false)
	err := r.delete(ctx)
	if err != nil {
		t := time.NewTicker(20 * time.Second)
		defer t.Stop()
	LOOPRETRY:
		for {
			select {
			case <-ctx.Done():
				break LOOPRETRY
			case <-t.C:
				err = r.delete(ctx)
				if err == nil {
					break LOOPRETRY
				}
			}
		}
	}
	spinCancel()
	if err != nil {
		messages.ExitErr(fmt.Errorf("unable to delete %s in %s: %w", r, tmout, err))
	}
	messages.Success("%s was deleted in %s.", r, time.Since(start))
}

func onlyError[T any](_ T, err error) error {
	if oerr := osc.AsErrorResponse(err); oerr != nil {
		if osc.IsNotFound(err) {
			return nil
		}
	}
	return err
}

func deleteNet(cl *osc.Client, id string) deleteFunc {
	return func(ctx context.Context) error {
		return onlyError(cl.DeleteNet(ctx, osc.DeleteNetRequest{NetId: id}))
	}
}

func deleteSubnet(cl *osc.Client, id string) deleteFunc {
	return func(ctx context.Context) error {
		return onlyError(cl.DeleteSubnet(ctx, osc.DeleteSubnetRequest{SubnetId: id}))
	}
}

func deleteInternetService(cl *osc.Client, id, netID string) deleteFunc {
	return func(ctx context.Context) error {
		err := onlyError(cl.UnlinkInternetService(ctx, osc.UnlinkInternetServiceRequest{InternetServiceId: id, NetId: netID}))
		if err != nil {
			return err
		}
		return onlyError(cl.DeleteInternetService(ctx, osc.DeleteInternetServiceRequest{InternetServiceId: id}))
	}
}

func deleteNetPeering(cl *osc.Client, id string) deleteFunc {
	return func(ctx context.Context) error {
		return onlyError(cl.DeleteNetPeering(ctx, osc.DeleteNetPeeringRequest{NetPeeringId: id}))
	}
}

func deleteNetAccessPoint(cl *osc.Client, id string) deleteFunc {
	return func(ctx context.Context) error {
		return onlyError(cl.DeleteNetAccessPoint(ctx, osc.DeleteNetAccessPointRequest{NetAccessPointId: id}))
	}
}

func deleteNATService(cl *osc.Client, id string) deleteFunc {
	return func(ctx context.Context) error {
		return onlyError(cl.DeleteNatService(ctx, osc.DeleteNatServiceRequest{NatServiceId: id}))
	}
}

func deleteRouteTable(cl *osc.Client, id string, links []string) deleteFunc {
	return func(ctx context.Context) error {
		for _, linkid := range links {
			err := onlyError(cl.UnlinkRouteTable(ctx, osc.UnlinkRouteTableRequest{LinkRouteTableId: linkid}))
			if err != nil {
				return err
			}
		}
		return onlyError(cl.DeleteRouteTable(ctx, osc.DeleteRouteTableRequest{RouteTableId: id}))
	}
}

func deleteLoadBalancer(cl *osc.Client, name string) deleteFunc {
	return func(ctx context.Context) error {
		return onlyError(cl.DeleteLoadBalancer(ctx, osc.DeleteLoadBalancerRequest{LoadBalancerName: name}))
	}
}

func deleteSecurityGroup(cl *osc.Client, id string) deleteFunc {
	return func(ctx context.Context) error {
		err := onlyError(cl.DeleteSecurityGroup(ctx, osc.DeleteSecurityGroupRequest{SecurityGroupId: &id}))
		if oerr := osc.AsErrorResponse(err); oerr != nil && oerr.GetCode() == "9051" {
			return nil
		}
		return err
	}
}

func deleteVirtualGateway(cl *osc.Client, id, netID string) deleteFunc {
	return func(ctx context.Context) error {
		err := onlyError(cl.UnlinkVirtualGateway(ctx, osc.UnlinkVirtualGatewayRequest{VirtualGatewayId: id, NetId: netID}))
		if err != nil {
			return err
		}
		return onlyError(cl.DeleteVirtualGateway(ctx, osc.DeleteVirtualGatewayRequest{VirtualGatewayId: id}))
	}
}

func deleteVpnConnection(cl *osc.Client, id string) deleteFunc {
	return func(ctx context.Context) error {
		return onlyError(cl.DeleteVpnConnection(ctx, osc.DeleteVpnConnectionRequest{VpnConnectionId: id}))
	}
}

func deletePublicIP(cl *osc.Client, id string) deleteFunc {
	return func(ctx context.Context) error {
		return onlyError(cl.DeletePublicIp(ctx, osc.DeletePublicIpRequest{PublicIpId: &id}))
	}
}

func deleteVm(cl *osc.Client, id string) deleteFunc {
	return func(ctx context.Context) error {
		return onlyError(cl.DeleteVms(ctx, osc.DeleteVmsRequest{VmIds: []string{id}}))
	}
}

func deleteNIC(cl *osc.Client, id, link string) deleteFunc {
	return func(ctx context.Context) error {
		if link != "" {
			err := onlyError(cl.UnlinkNic(ctx, osc.UnlinkNicRequest{LinkNicId: link}))
			if err != nil {
				return err
			}
		}
		return onlyError(cl.DeleteNic(ctx, osc.DeleteNicRequest{NicId: id}))
	}
}
