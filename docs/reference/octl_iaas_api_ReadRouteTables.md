## octl iaas api ReadRouteTables

Lists one or more of your route tables.

### Synopsis

Lists one or more of your route tables.

In your Net, each Subnet must be associated with a route table. If a Subnet is not explicitly associated with a route table, it is implicitly associated with the main route table of the Net.

```
octl iaas api ReadRouteTables [flags]
```

### Options

```
      --DryRun                                            If true, checks whether you have the required permissions to perform the action.
      --Filters.LinkRouteTableIds strings                 The IDs of the route tables involved in the associations.
      --Filters.LinkRouteTableLinkRouteTableIds strings   The IDs of the associations between the route tables and the Subnets.
      --Filters.LinkRouteTableMain                        If true, the route tables are the main ones for their Nets.
      --Filters.LinkSubnetIds strings                     The IDs of the Subnets involved in the associations.
      --Filters.NetIds strings                            The IDs of the Nets for the route tables.
      --Filters.RouteCreationMethods strings              The methods used to create a route.
      --Filters.RouteDestinationIpRanges strings          The IP ranges specified in routes in the tables.
      --Filters.RouteDestinationServiceIds strings        The service IDs specified in routes in the tables.
      --Filters.RouteGatewayIds strings                   The IDs of the gateways specified in routes in the tables.
      --Filters.RouteNatServiceIds strings                The IDs of the NAT services specified in routes in the tables.
      --Filters.RouteNetPeeringIds strings                The IDs of the Net peerings specified in routes in the tables.
      --Filters.RouteStates strings                       The states of routes in the route tables (always active).
      --Filters.RouteTableIds strings                     The IDs of the route tables.
      --Filters.RouteVmIds strings                        The IDs of the VMs specified in routes in the tables.
      --Filters.TagKeys strings                           The keys of the tags associated with the route tables.
      --Filters.TagValues strings                         The values of the tags associated with the route tables.
      --Filters.Tags strings                              The key/value combination of the tags associated with the route tables, in the following format: "Filters":{"Tags":["TAGKEY=TAGVALUE"]}.
      --NextPageToken string                              The token to request the next page of results.
      --ResultsPerPage int                                The maximum number of logs returned in a single response (between 1 and 1000, both included).
  -h, --help                                              help for ReadRouteTables
```

### Options inherited from parent commands

```
  -c, --columns string              columns to display - [+]<title>:<jq query for content>||<title>:<jq query for content>
      --config string               Path of profile file (by default, ~/.osc/config.json)
      --filter strings              comma separated list of filters for results - name:value,name:value, alias for jq filter 'select(.name | test("value"))'
      --jq string                   jq filter
      --no-upgrade                  do not check for new versions
  -O, --out-file string             redirect output to file
  -o, --output string               output format (raw, json, yaml, table, csv, none, base64) (default "raw")
      --payload string              JSON content for query body
      --profile string              Profile to use in profile file (by default, "default")
      --single                      convert single entry lists to a single object
      --template string             JSON template file for query body
  -v, --verbose                     Verbose output
      --waitfor string              jq expression to wait for - octl will query every waitfor-interval until the expression returns 1/true or a non empty result
      --waitfor-interval duration   interval between two waitfor iterations (default 5s)
      --waitfor-timeout duration    maximum duration of a wait (default 10m0s)
  -y, --yes                         answer yes to all prompts
```

### SEE ALSO

* [octl iaas api](octl_iaas_api.md)	 - iaas api calls

