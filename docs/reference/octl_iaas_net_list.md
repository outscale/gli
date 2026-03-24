## octl iaas net list

alias for api ReadNets

### Synopsis

> *alias for api ReadNets*

Lists one or more Nets.

```
octl iaas net list [flags]
```

### Options

```
      --dhcp-options-set-id strings   The IDs of the DHCP options sets.
  -h, --help                          help for list
      --id strings                    The IDs of the Nets.
      --ip-range strings              The IP ranges for the Nets, in CIDR notation (for example, 10.0.0.0/16).
      --is-default                    If true, the Net used is the default one.
      --state strings                 The states of the Nets (pending | available | deleting).
      --tag strings                   The key/value combination of the tags associated with the Nets, in the following format: "Filters":{"Tags":["TAGKEY=TAGVALUE"]}.
      --tag-key strings               The keys of the tags associated with the Nets.
      --tag-value strings             The values of the tags associated with the Nets.
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

* [octl iaas net](octl_iaas_net.md)	 - net commands

