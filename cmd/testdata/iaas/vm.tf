resource "outscale_keypair" "kp" {
    keypair_name = "octl"
}

resource "outscale_security_group" "vmsg" {
  description         = "vm security group"
  security_group_name = "vm_security_group1"
  net_id              = outscale_net.net.net_id
}

resource "outscale_vm" "vm1" {
    image_id                 = "ami-0016b8a0"
    vm_type                  = "tinav7.c1r1p3"
    keypair_name             = outscale_keypair.kp.keypair_name
    security_group_ids       = [outscale_security_group.vmsg.security_group_id]
    subnet_id                = outscale_subnet.subnet2.subnet_id
    tags {
        key   = "name"
        value = "test"
    }
}

resource "outscale_vm" "vm2" {
    image_id                 = "ami-0016b8a0"
    vm_type                  = "tinav7.c1r1p3"
    keypair_name             = outscale_keypair.kp.keypair_name
    tags {
        key   = "name"
        value = "test"
    }
    nics {
        nic_id        = outscale_nic.vmnic.nic_id
        device_number = "0"
    }
}

resource "outscale_public_ip" "vmip" {
  tags {
    key   = "Name"
    value = "test"
  }
}

resource "outscale_public_ip_link" "vmip" {
    public_ip = outscale_public_ip.vmip.public_ip
    nic_id    = outscale_nic.vmnic.nic_id
}

resource "outscale_nic" "vmnic" {
    subnet_id          = outscale_subnet.subnet2.subnet_id
    security_group_ids = [outscale_security_group.vmsg.security_group_id]
}

resource "outscale_nic" "nicnovm" {
    subnet_id          = outscale_subnet.subnet2.subnet_id
    security_group_ids = [outscale_security_group.vmsg.security_group_id]
}

resource "outscale_security_group" "sgnovm" {
  description         = "sg without vm"
  security_group_name = "novmsg"
  net_id              = outscale_net.net.net_id
}
