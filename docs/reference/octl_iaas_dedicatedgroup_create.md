## octl iaas dedicatedgroup create

alias for api CreateDedicatedGroup

### Synopsis

> *alias for api CreateDedicatedGroup*

Creates a dedicated group for virtual machines (VMs).

For more information, see [About Dedicated Groups](https://docs.outscale.com/en/userguide/About-Dedicated-Groups.html).

```
octl iaas dedicatedgroup create [flags]
```

### Options

```
      --cpu-generation int   The processor generation for the VMs in the dedicated group (for example, 4).
  -h, --help                 help for create
      --name string          A name for the dedicated group.
      --subregion string     The Subregion in which you want to create the dedicated group.
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

* [octl iaas dedicatedgroup](octl_iaas_dedicatedgroup.md)	 - dedicatedgroup commands

