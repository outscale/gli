## octl iaas apiaccessrule update

alias for api UpdateApiAccessRule --ApiAccessRuleId api_access_rule_id

### Synopsis

> *alias for api UpdateApiAccessRule --ApiAccessRuleId api_access_rule_id*

Modifies a specified API access rule.

**[WARNING]**

- The new rule you specify fully replaces the old rule. Therefore, for a parameter that is not specified, any previously set value is deleted.

- If, as result of your modification, you no longer have access to the APIs, you will need to contact the Support team to regain access. For more information, see [Technical Support](https://docs.outscale.com/en/userguide/Technical-Support.html).

```
octl iaas apiaccessrule update api_access_rule_id [api_access_rule_id]... [flags]
```

### Options

```
      --ca-id strings        One or more IDs of Client Certificate Authorities (CAs).
      --cn strings           One or more Client Certificate Common Names (CNs).
      --description string   A new description for the API access rule.
  -h, --help                 help for update
      --ip-range strings     One or more IPs or CIDR blocks (for example, 192.0.2.0/16).
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

* [octl iaas apiaccessrule](octl_iaas_apiaccessrule.md)	 - apiaccessrule commands

