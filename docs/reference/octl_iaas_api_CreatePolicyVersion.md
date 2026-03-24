## octl iaas api CreatePolicyVersion

Creates a version of a specified managed policy.

### Synopsis

Creates a version of a specified managed policy.

A managed policy can have up to five versions.

**[IMPORTANT]**

A delay of up to 15 seconds can occur when attaching, detaching, or updating a managed policy.

```
octl iaas api CreatePolicyVersion [flags]
```

### Options

```
      --Document string    The policy document, corresponding to a JSON string that contains the policy.
      --PolicyOrn string   The OUTSCALE Resource Name (ORN) of the policy.
      --SetAsDefault       If set to true, the new policy version is set as the default version and becomes the operative one.
  -h, --help               help for CreatePolicyVersion
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

* [octl iaas api](octl_iaas_api.md)	 - iaas api calls

