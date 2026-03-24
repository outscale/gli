resource "outscale_security_group" "lbusg1" {
  security_group_name = "lbusg1"
  description         = "lbusg1"
  net_id              = outscale_net.net.net_id
  tags {
      key   = "Name"
      value = "test"
  }
}

resource "outscale_security_group" "lbusg2" {
  security_group_name = "lbusg2"
  description         = "lbusg2"
  net_id              = outscale_net.net.net_id
  tags {
      key   = "Name"
      value = "test"
  }
}

resource "outscale_public_ip" "lbuip" {
  tags {
    key   = "Name"
    value = "test"
  }
}

resource "outscale_load_balancer" "lbu" {
    load_balancer_name = "test-octl"
    listeners {
        backend_port           = 80
        backend_protocol       = "TCP"
        load_balancer_protocol = "TCP"
        load_balancer_port     = 80
    }
    subnets         = [outscale_subnet.subnet1.subnet_id]
    security_groups = [outscale_security_group.lbusg1.security_group_id, outscale_security_group.lbusg2.security_group_id]
    public_ip       = outscale_public_ip.lbuip.public_ip
    tags {
        key   = "Name"
        value = "test"
    }
}

resource "outscale_load_balancer" "lbu-internal" {
    load_balancer_name = "test-octl-internal"
    listeners {
        backend_port           = 80
        backend_protocol       = "TCP"
        load_balancer_protocol = "TCP"
        load_balancer_port     = 80
    }
    load_balancer_type = "internal"
    subnets            = [outscale_subnet.subnet2.subnet_id]
    security_groups    = [outscale_security_group.lbusg1.security_group_id, outscale_security_group.lbusg2.security_group_id]
    tags {
        key   = "Name"
        value = "test"
    }
}
