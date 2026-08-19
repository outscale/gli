## octl iaas api CreateNetAccessPoint

Creates a Net access point to access an OUTSCALE service from this Net without using the Internet or public IPs.

### Synopsis

Creates a Net access point to access an OUTSCALE service from this Net without using the Internet or public IPs.

You specify the service using its name. For more information about the available services, see [ReadNetAccessPointServices](#readnetaccesspointservices).
 

To control the routing of traffic between the Net and the specified service, you can specify one or more route tables. Virtual machines placed in Subnets associated with the specified route table thus use the Net access point to access the service. When you specify a route table, a route is automatically added to it with the destination set to the prefix list ID of the service, and the target set to the ID of the access point.

When a Net access point is created, a public IP is automatically allocated to your OUTSCALE account and used for the Net access point. This public IP is not connected to the Internet. It is counted in your quota, but it is not billed.
 

For more information, see [About Net Access Points](https://docs.outscale.com/en/userguide/About-Net-Access-Points.html).

```
octl iaas api CreateNetAccessPoint [flags]
```

### Options

```
      --DryRun                  If true, checks whether you have the required permissions to perform the action.
      --NetId string            The ID of the Net.
      --RouteTableIds strings   One or more IDs of route tables to use for the connection.
      --ServiceName string      The name of the service (in the format com.outscale.region.service).
  -h, --help                    help for CreateNetAccessPoint
```

### Options inherited from parent commands

```
  -c, --columns string             columns to display - [+]<title>:<jq query for content>||<title>:<jq query for content>
      --config string              Path of profile file (by default, ~/.osc/config.json)
      --dry-run                    Display the request payload that would be sent to the API without sending it
      --elapsed                    add elapsed time column when using --watch (default true)
      --filter strings             comma separated list of filters for results - name:value,name:value, alias for jq filter 'select(.name | tostring | test("value"))'
      --interval duration          interval between two watch/waitfor iterations (default 5s)
      --jq string                  jq filter
      --max-pages int              maximum number of pages a command can fetch (default 20)
      --no-upgrade                 do not check for new versions
  -O, --out-file string            redirect output to file
  -o, --output string              output format (raw, json, yaml, table, csv, none, base64, text)
      --payload string             JSON content for query body
      --profile string             Profile to use in profile file (by default, "default")
  -s, --silent                     Hides all information messages
      --single                     convert single entry lists to a single object
      --template string            JSON template file for query body
  -v, --verbose                    Verbose output
      --waitfor string             repeatedly call the API until the specified jq expression returns 1/true or a non empty result
      --waitfor-timeout duration   maximum duration of a wait (default 10m0s)
      --watch                      repeatedly call the API and display changes
  -y, --yes                        answer yes to all prompts
```

### SEE ALSO

* [octl iaas api](octl_iaas_api.md)	 - iaas api calls

