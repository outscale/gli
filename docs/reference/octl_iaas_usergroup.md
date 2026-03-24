## octl iaas usergroup

usergroup commands

### Options

```
  -h, --help   help for usergroup
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

* [octl iaas](octl_iaas.md)	 - OUTSCALE IaaS management
* [octl iaas usergroup create](octl_iaas_usergroup_create.md)	 - alias for api CreateUserGroup
* [octl iaas usergroup delete](octl_iaas_usergroup_delete.md)	 - alias for api DeleteUserGroup --UserGroupName user_group_name
* [octl iaas usergroup list](octl_iaas_usergroup_list.md)	 - alias for api ReadUserGroup
* [octl iaas usergroup list](octl_iaas_usergroup_list.md)	 - alias for api ReadUserGroups
* [octl iaas usergroup update](octl_iaas_usergroup_update.md)	 - alias for api UpdateUserGroup --UserGroupName user_group_name

