## octl iaas natservice

natservice commands

### Options

```
  -h, --help   help for natservice
```

### Options inherited from parent commands

```
  -c, --columns string              columns to display - [+]<title>:<jq query for content>||<title>:<jq query for content>
      --config string               Path of profile file (by default, ~/.osc/config.json)
      --filter strings              comma separated list of filters for results - name:value,name:value, alias for jq filter 'select(.name | test("value"))'
      --jq string                   jq filter
      --no-upgrade                  do not check for new versions
  -O, --out-file string             redirect output to file
  -o, --output string               output format (raw, json, yaml, table, csv, none, base64, text) (default "raw")
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

* [octl iaas](octl_iaas.md)	 - OUTSCALE IaaS management
* [octl iaas natservice create](octl_iaas_natservice_create.md)	 - alias for api CreateNatService
* [octl iaas natservice delete](octl_iaas_natservice_delete.md)	 - alias for api DeleteNatService --NatServiceId nat_service_id
* [octl iaas natservice describe](octl_iaas_natservice_describe.md)	 - alias for api ReadNatServices --Filters.NatServiceIds nat_service_id
* [octl iaas natservice list](octl_iaas_natservice_list.md)	 - alias for api ReadNatServices

