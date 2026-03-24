
resource "outscale_net" "net" {
  ip_range = "10.0.0.0/16"
  tags {
    key   = "Name"
    value = "test"
  }
}

resource "outscale_subnet" "subnet1" {
    net_id     = outscale_net.net.net_id
    ip_range = "10.0.0.0/24"
}

resource "outscale_subnet" "subnet2" {
    net_id     = outscale_net.net.net_id
    ip_range = "10.0.1.0/24"
}

resource "outscale_internet_service" "is" {
  tags {
    key   = "Name"
    value = "test"
  }
}

resource "outscale_internet_service_link" "outscale_internet_service_link" {
    internet_service_id = outscale_internet_service.is.internet_service_id
    net_id              = outscale_net.net.net_id
}

resource "outscale_public_ip" "natip" {
  tags {
    key   = "Name"
    value = "test"
  }
}

resource "outscale_nat_service" "nat" {
    subnet_id    = outscale_subnet.subnet1.subnet_id
    public_ip_id = outscale_public_ip.natip.public_ip_id
    tags {
      key   = "Name"
      value = "test"
    }
}

resource "outscale_route_table" "rtbl1" {
    net_id = outscale_net.net.net_id
}

resource "outscale_route" "route1" {
    destination_ip_range = "0.0.0.0/0"
    gateway_id           = outscale_internet_service.is.internet_service_id
    route_table_id       = outscale_route_table.rtbl1.route_table_id
}

resource "outscale_route_table_link" "rtbl1" {
    subnet_id      = outscale_subnet.subnet1.subnet_id
    route_table_id = outscale_route_table.rtbl1.id
}

resource "outscale_route_table" "rtbl2" {
    net_id = outscale_net.net.net_id
}

resource "outscale_route" "route2" {
    destination_ip_range = "0.0.0.0/0"
    nat_service_id       = outscale_nat_service.nat.nat_service_id
    route_table_id       = outscale_route_table.rtbl2.route_table_id
}

resource "outscale_route_table_link" "rtbl2" {
    subnet_id      = outscale_subnet.subnet2.subnet_id
    route_table_id = outscale_route_table.rtbl2.id
}

