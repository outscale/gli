## octl iaas unitprice list

alias for api ReadUnitPrice

### Synopsis

> *alias for api ReadUnitPrice*

Gets unit price information for the specified parameters.

```
octl iaas unitprice list [flags]
```

### Options

```
  -h, --help               help for list
      --operation string   The operation associated with the catalog entry (for example, RunInstances-OD or CreateVolume).
      --service string     The service associated with the catalog entry (for example, TinaOS-FCU or TinaOS-OOS).
      --type string        The type associated with the catalog entry (for example, BSU:VolumeIOPS:io1 or BoxUsage:tinav6.c6r16p3).
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

* [octl iaas unitprice](octl_iaas_unitprice.md)	 - unitprice commands

