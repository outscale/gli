## octl iaas api PutUserPolicy

Creates or updates an inline policy included in a specified user.

### Synopsis

Creates or updates an inline policy included in a specified user.

The policy is automatically applied to the user after its creation.

**[IMPORTANT]**

A delay of up to 15 seconds can occur when creating or deleting an inline policy.

```
octl iaas api PutUserPolicy [flags]
```

### Options

```
      --DryRun                  If true, checks whether you have the required permissions to perform the action.
      --PolicyDocument string   The policy document, corresponding to a JSON string that contains the policy.
      --PolicyName string       The name of the policy (between 1 and 128 characters).
      --UserName string         The name of the user.
  -h, --help                    help for PutUserPolicy
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

* [octl iaas api](octl_iaas_api.md)	 - iaas api calls

