## octl iaas usergroup delete

alias for api DeleteUserGroup --UserGroupName user_group_name

### Synopsis

> *alias for api DeleteUserGroup --UserGroupName user_group_name*

Deletes a specified user group.

**[WARNING]**

The user group must be empty of any user and must not have any linked policy. Otherwise, you need to force the deletion.

If you force the deletion, all inline policies will be deleted with the user group.

```
octl iaas usergroup delete user_group_name [user_group_name]... [flags]
```

### Options

```
      --force         If true, forces the deletion of the user group even if it is not empty.
  -h, --help          help for delete
      --path string   The path to the group.
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

* [octl iaas usergroup](octl_iaas_usergroup.md)	 - usergroup commands

