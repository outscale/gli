# Chaining commands

Commands may be chained, and attributes returned by a command can be reinjected in a subsequent command, using JQ queries within `{{}}` placeholders:

Create a Nic and link it to a VM:
```shell
octl iaas nic create --subnet-id subnet-foo -o json | octl iaas nic link {{.NicId}} --vm-id i-foo --device-number 7
```

Delete all subnets from a net:
```shell
octl iaas subnet list --net-id vpc-foo -o json | octl iaas subnet del '{{.[].SubnetId}}'
```