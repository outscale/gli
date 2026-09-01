## octl iaas volume

Manage Volume resources

### Options

```
  -h, --help   help for volume
```

### Options inherited from parent commands

```
  -c, --columns string             columns to display - [+]<title>:<jq query for content>||<title>:<jq query for content>
      --config string              Path of profile file (by default, ~/.osc/config.json)
      --dry-run                    Display the request payload that would be sent to the API without sending it
      --elapsed                    add elapsed time column when using --watch (default true)
      --filter strings             comma separated list of filters for results - name:value,name:value, alias for jq filter 'select(.name | tostring | test("value"))'
      --interval duration          interval between two watch/waitfor iterations (default 5s)
      --jq string                  jq filter
      --max-pages int              maximum number of pages a command can fetch (default 20)
      --no-upgrade                 do not check for new versions
  -O, --out-file string            redirect output to file
  -o, --output string              output format (raw, json, yaml, table, csv, none, text)
      --payload string             JSON content for query body
      --profile string             Profile to use in profile file (by default, "default")
  -s, --silent                     Hides all information messages
      --single                     convert single entry lists to a single object
      --template string            JSON template file for query body
  -v, --verbose                    Verbose output
      --waitfor string             repeatedly call the API until the specified jq expression returns 1/true or a non empty result
      --waitfor-timeout duration   maximum duration of a wait (default 10m0s)
      --watch                      repeatedly call the API and display changes
  -y, --yes                        answer yes to all prompts
```

### SEE ALSO

* [octl iaas](octl_iaas.md)	 - OUTSCALE IaaS management
* [octl iaas volume create](octl_iaas_volume_create.md)	 - Creates a Block Storage Unit (BSU) volume in a specified Region.
* [octl iaas volume delete](octl_iaas_volume_delete.md)	 - Deletes a specified Block Storage Unit (BSU) volume.
* [octl iaas volume describe](octl_iaas_volume_describe.md)	 - Lists one or more specified Block Storage Unit (BSU) volumes.
* [octl iaas volume link](octl_iaas_volume_link.md)	 - Attaches a Block Storage Unit (BSU) volume to a virtual machine (VM).
* [octl iaas volume list](octl_iaas_volume_list.md)	 - Lists one or more specified Block Storage Unit (BSU) volumes.
* [octl iaas volume unlink](octl_iaas_volume_unlink.md)	 - Detaches a Block Storage Unit (BSU) volume from a virtual machine (VM).
* [octl iaas volume update](octl_iaas_volume_update.md)	 - Modifies the specified attributes of a volume.

