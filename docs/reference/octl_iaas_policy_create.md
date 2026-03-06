## octl iaas policy create

alias for api CreatePolicy

### Synopsis

> *alias for api CreatePolicy*

Creates a managed policy to apply to a user.

This action creates a policy version and sets v1 as the default one.

```
octl iaas policy create [flags]
```

### Options

```
      --description string   A description for the policy.
      --document File        The file storing the policy document, corresponding to a JSON string that contains the policy.
  -h, --help                 help for create
      --name string          The name of the policy.
      --path string          The path of the policy.
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

* [octl iaas policy](octl_iaas_policy.md)	 - policy commands

