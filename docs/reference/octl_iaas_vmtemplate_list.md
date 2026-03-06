## octl iaas vmtemplate list

alias for api ReadVmTemplates

### Synopsis

> *alias for api ReadVmTemplates*

> [WARNING]

> This feature is currently under development and may not function properly.

Lists one or more virtual machine (VM) templates.

```
octl iaas vmtemplate list [flags]
```

### Options

```
      --cpu-core ints             The number of vCores.
      --cpu-generation strings    The processor generations (for example, v4).
      --cpu-performance strings   The performances of the VMs.
      --description strings       The descriptions of the VM templates.
  -h, --help                      help for list
      --id strings                The IDs of the VM templates.
      --image-id strings          The IDs of the OMIs.
      --keypair-name strings      The names of the keypairs.
      --name strings              The names of the VM templates.
      --ram ints                  The amount of RAM.
      --tag strings               The key/value combination of the tags associated with the VM templates, in the following format: "Filters":{"Tags":["TAGKEY=TAGVALUE"]}.
      --tag-key strings           The keys of the tags associated with the VM templates.
      --tag-value strings         The values of the tags associated with the VM templates.
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

* [octl iaas vmtemplate](octl_iaas_vmtemplate.md)	 - vmtemplate commands

