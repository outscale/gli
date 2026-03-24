## octl iaas vmtemplate update

alias for api UpdateVmTemplate --VmTemplateId vm_template_id

### Synopsis

> *alias for api UpdateVmTemplate --VmTemplateId vm_template_id*

> [WARNING]

> This feature is currently under development and may not function properly.

Modifies the specified attributes of a template of virtual machines (VMs).

```
octl iaas vmtemplate update vm_template_id [vm_template_id]... [flags]
```

### Options

```
      --description string   A new description for the VM template.
  -h, --help                 help for update
      --name string          A new name for your VM template.
      --tag-key string       The key of the tag, between 1 and 255 characters.
      --tag-value string     The value of the tag, between 0 and 255 characters.
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

* [octl iaas vmtemplate](octl_iaas_vmtemplate.md)	 - vmtemplate commands

