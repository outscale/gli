## octl iaas policyversion create

alias for api CreatePolicyVersion --PolicyOrn policy_orn

### Synopsis

> *alias for api CreatePolicyVersion --PolicyOrn policy_orn*

Creates a version of a specified managed policy.

A managed policy can have up to five versions.

**[IMPORTANT]**

A delay of up to 15 seconds can occur when attaching, detaching, or updating a managed policy.

```
octl iaas policyversion create policy_orn [flags]
```

### Options

```
      --default         If set to true, the new policy version is set as the default version and becomes the operative one.
      --document File   The file storing the policy document, corresponding to a JSON string that contains the policy.
  -h, --help            help for create
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

* [octl iaas policyversion](octl_iaas_policyversion.md)	 - policyversion commands

