## octl iaas snapshot list

alias for api ReadSnapshots

### Synopsis

> *alias for api ReadSnapshots*

Lists one or more snapshots that are available to you and the permissions to create volumes from them.

```
octl iaas snapshot list [flags]
```

### Options

```
      --account-alias strings                            The account aliases of the owners of the snapshots.
      --account-id strings                               The OUTSCALE account IDs of the owners of the snapshots.
      --client-token strings                             The idempotency tokens provided when creating the snapshots.
      --description strings                              The descriptions of the snapshots.
      --from-creation-date osctime                       The beginning of the time period, in ISO 8601 date-time format (for example, 2020-06-14T00:00:00.000Z).
  -h, --help                                             help for list
      --id strings                                       The IDs of the snapshots.
      --permission-to-create-volume-account-id strings   The OUTSCALE account IDs which have permissions to create volumes.
      --permission-to-create-volume-global-permission    If true, lists all public volumes.
      --progress ints                                    The progresses of the snapshots, as a percentage.
      --state strings                                    The states of the snapshots (in-queue | pending | completed | error | deleting).
      --tag strings                                      The key/value combination of the tags associated with the snapshots, in the following format: "Filters":{"Tags":["TAGKEY=TAGVALUE"]}.
      --tag-key strings                                  The keys of the tags associated with the snapshots.
      --tag-value strings                                The values of the tags associated with the snapshots.
      --to-creation-date osctime                         The end of the time period, in ISO 8601 date-time format (for example, 2020-06-30T00:00:00.000Z).
      --volume-id strings                                The IDs of the volumes used to create the snapshots.
      --volume-size ints                                 The sizes of the volumes used to create the snapshots, in gibibytes (GiB).
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

* [octl iaas snapshot](octl_iaas_snapshot.md)	 - snapshot commands

