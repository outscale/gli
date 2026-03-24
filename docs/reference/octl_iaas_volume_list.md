## octl iaas volume list

alias for api ReadVolumes

### Synopsis

> *alias for api ReadVolumes*

Lists one or more specified Block Storage Unit (BSU) volumes.

```
octl iaas volume list [flags]
```

### Options

```
      --client-token strings                The idempotency tokens provided when creating the volumes.
      --creation-date osctime               The dates and times at which the volumes were created, in ISO 8601 date-time format (for example, 2020-06-30T00:00:00.000Z).
  -h, --help                                help for list
      --id strings                          The IDs of the volumes.
      --link-volume-delete-on-vm-deletion   Whether the volumes are deleted or not when terminating the VMs.
      --link-volume-device-name strings     The VM device names.
      --link-volume-link-date osctime       The dates and times at which the volumes were attached, in ISO 8601 date-time format (for example, 2020-06-30T00:00:00.000Z).
      --link-volume-link-state strings      The attachment states of the volumes (attaching | detaching | attached | detached).
      --link-volume-vm-id strings           One or more IDs of VMs.
      --size ints                           The sizes of the volumes, in gibibytes (GiB).
      --snapshot-id strings                 The snapshots from which the volumes were created.
      --state strings                       The states of the volumes (creating | available | in-use | deleting | error).
      --subregion strings                   The names of the Subregions in which the volumes were created.
      --tag strings                         The key/value combination of the tags associated with the volumes, in the following format: "Filters":{"Tags":["TAGKEY=TAGVALUE"]}.
      --tag-key strings                     The keys of the tags associated with the volumes.
      --tag-value strings                   The values of the tags associated with the volumes.
      --type strings                        The types of the volumes (standard | gp2 | io1).
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

* [octl iaas volume](octl_iaas_volume.md)	 - volume commands

