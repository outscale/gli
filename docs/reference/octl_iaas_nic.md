## octl iaas nic

nic commands

### Options

```
  -h, --help   help for nic
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
      --profile string    Profile to use in profile file (by default, "default")
      --single            convert single entry lists to a single object
      --template string   JSON template for query body
  -v, --verbose           Verbose output
  -y, --yes               answer yes to all prompts
```

### SEE ALSO

* [octl iaas](octl_iaas.md)	 - OUTSCALE IaaS management
* [octl iaas nic create](octl_iaas_nic_create.md)	 - alias for api CreateNic
* [octl iaas nic delete](octl_iaas_nic_delete.md)	 - alias for api DeleteNic --NicId nic_id
* [octl iaas nic describe](octl_iaas_nic_describe.md)	 - alias for api ReadNics --Filters.NicIds nic_id
* [octl iaas nic link](octl_iaas_nic_link.md)	 - alias for api LinkNic --NicId nic_id
* [octl iaas nic list](octl_iaas_nic_list.md)	 - alias for api ReadNics
* [octl iaas nic update](octl_iaas_nic_update.md)	 - alias for api UpdateNic --NicId nic_id

