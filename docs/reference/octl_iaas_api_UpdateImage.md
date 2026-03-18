## octl iaas api UpdateImage

Modifies the access permissions for an OUTSCALE machine image (OMI).

### Synopsis

Modifies the access permissions for an OUTSCALE machine image (OMI).

You must specify either the `Additions` or the `Removals` parameter.

After sharing an OMI with an OUTSCALE account, the other account can create a copy of it that they own. For more information about copying OMIs, see [CreateImage](#createimage).

```
octl iaas api UpdateImage [flags]
```

### Options

```
      --Description string                                 A new description for the image.
      --DryRun                                             If true, checks whether you have the required permissions to perform the action.
      --ImageId string                                     The ID of the OMI you want to modify.
      --PermissionsToLaunch.Additions.AccountIds strings   One or more OUTSCALE account IDs that the permission is associated with.
      --PermissionsToLaunch.Additions.GlobalPermission     A global permission for all accounts.
      --PermissionsToLaunch.Removals.AccountIds strings    One or more OUTSCALE account IDs that the permission is associated with.
      --PermissionsToLaunch.Removals.GlobalPermission      A global permission for all accounts.
      --ProductCodes strings                               The product codes associated with the OMI.
  -h, --help                                               help for UpdateImage
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
      --payload string    JSON content for query body
      --profile string    Profile to use in profile file (by default, "default")
      --single            convert single entry lists to a single object
      --template string   JSON template file for query body
  -v, --verbose           Verbose output
  -y, --yes               answer yes to all prompts
```

### SEE ALSO

* [octl iaas api](octl_iaas_api.md)	 - iaas api calls

