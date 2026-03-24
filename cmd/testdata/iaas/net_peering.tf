resource "outscale_net" "net_remote" {
  ip_range = "10.1.0.0/16"
  tags {
    key   = "Name"
    value = "test-netpeering"
  }
}

resource "outscale_net_peering" "np" {
    accepter_net_id = outscale_net.net_remote.net_id
    source_net_id   = outscale_net.net.net_id
}

resource "outscale_net_peering_acceptation" "np" {
    net_peering_id = outscale_net_peering.np.net_peering_id
}

resource "outscale_net_access_point" "nap" {
    net_id          = outscale_net.net.net_id
    route_table_ids = [outscale_route_table.rtbl1.route_table_id]
    service_name    = "com.outscale.eu-west-2.api"
}
