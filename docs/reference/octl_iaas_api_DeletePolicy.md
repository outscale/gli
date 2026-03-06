## octl iaas api DeletePolicy

Deletes a managed policy.

### Synopsis

Deletes a managed policy.

Before deleting a managed policy, you must unlink all users linked to it and delete all the versions of the policy, except the default one, using the `DeletePolicyVersion` method.

```
octl iaas api DeletePolicy [flags]
```

### Options

```
      --DryRun             If true, checks whether you have the required permissions to perform the action.
      --PolicyOrn string   The OUTSCALE Resource Name (ORN) of the policy you want to delete.
  -h, --help               help for DeletePolicy
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

