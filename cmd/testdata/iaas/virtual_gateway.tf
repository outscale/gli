resource "outscale_virtual_gateway" "vgw" {
  connection_type = "ipsec.1"
}

resource "outscale_virtual_gateway_link" "outscale_virtual_gateway_link" {
  virtual_gateway_id = outscale_virtual_gateway.vgw.virtual_gateway_id
  net_id             = outscale_net.net.net_id
}

resource "outscale_client_gateway" "cgw" {
  bgp_asn         = 64512
  public_ip       = "198.18.7.207"
  connection_type = "ipsec.1"
}


resource "outscale_vpn_connection" "vpn" {
  client_gateway_id  = outscale_client_gateway.cgw.client_gateway_id
  virtual_gateway_id = outscale_virtual_gateway.vgw.virtual_gateway_id
  connection_type    = "ipsec.1"
  static_routes_only = true
}

resource "outscale_vpn_connection_route" "route" {
  vpn_connection_id    = outscale_vpn_connection.vpn.vpn_connection_id
  destination_ip_range = "40.0.0.0/16"
}
