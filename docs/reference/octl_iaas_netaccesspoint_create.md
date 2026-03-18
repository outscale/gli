## octl iaas netaccesspoint create

alias for api CreateNetAccessPoint

### Synopsis

> *alias for api CreateNetAccessPoint*

Creates a Net access point to access an OUTSCALE service from this Net without using the Internet and public IPs.

You specify the service using its name. For more information about the available services, see [ReadNetAccessPointServices](#readnetaccesspointservices).
 

To control the routing of traffic between the Net and the specified service, you can specify one or more route tables. Virtual machines placed in Subnets associated with the specified route table thus use the Net access point to access the service. When you specify a route table, a route is automatically added to it with the destination set to the prefix list ID of the service, and the target set to the ID of the access point.

When a Net access point is created, a public IP is automatically allocated to your OUTSCALE account and used for the Net access point. This public IP is not connected to the Internet. It is counted in your quota, but it is not billed.
 

For more information, see [About Net Access Points](https://docs.outscale.com/en/userguide/About-Net-Access-Points.html).

```
octl iaas netaccesspoint create [flags]
```

### Options

```
  -h, --help                     help for create
      --net-id string            The ID of the Net.
      --route-table-id strings   One or more IDs of route tables to use for the connection.
      --service-name string      The name of the service (in the format com.outscale.region.service).
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

* [octl iaas netaccesspoint](octl_iaas_netaccesspoint.md)	 - netaccesspoint commands

