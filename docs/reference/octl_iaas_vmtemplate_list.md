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

* [octl iaas vmtemplate](octl_iaas_vmtemplate.md)	 - vmtemplate commands

