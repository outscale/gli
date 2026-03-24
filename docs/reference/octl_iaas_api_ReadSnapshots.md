## octl iaas api ReadSnapshots

Lists one or more snapshots that are available to you and the permissions to create volumes from them.

### Synopsis

Lists one or more snapshots that are available to you and the permissions to create volumes from them.

```
octl iaas api ReadSnapshots [flags]
```

### Options

```
      --DryRun                                                If true, checks whether you have the required permissions to perform the action.
      --Filters.AccountAliases strings                        The account aliases of the owners of the snapshots.
      --Filters.AccountIds strings                            The OUTSCALE account IDs of the owners of the snapshots.
      --Filters.ClientTokens strings                          The idempotency tokens provided when creating the snapshots.
      --Filters.Descriptions strings                          The descriptions of the snapshots.
      --Filters.FromCreationDate osctime                      The beginning of the time period, in ISO 8601 date-time format (for example, 2020-06-14T00:00:00.000Z).
      --Filters.PermissionsToCreateVolumeAccountIds strings   The OUTSCALE account IDs which have permissions to create volumes.
      --Filters.PermissionsToCreateVolumeGlobalPermission     If true, lists all public volumes.
      --Filters.Progresses ints                               The progresses of the snapshots, as a percentage.
      --Filters.SnapshotIds strings                           The IDs of the snapshots.
      --Filters.States strings                                The states of the snapshots (in-queue | pending | completed | error | deleting).
      --Filters.TagKeys strings                               The keys of the tags associated with the snapshots.
      --Filters.TagValues strings                             The values of the tags associated with the snapshots.
      --Filters.Tags strings                                  The key/value combination of the tags associated with the snapshots, in the following format: "Filters":{"Tags":["TAGKEY=TAGVALUE"]}.
      --Filters.ToCreationDate osctime                        The end of the time period, in ISO 8601 date-time format (for example, 2020-06-30T00:00:00.000Z).
      --Filters.VolumeIds strings                             The IDs of the volumes used to create the snapshots.
      --Filters.VolumeSizes ints                              The sizes of the volumes used to create the snapshots, in gibibytes (GiB).
      --NextPageToken string                                  The token to request the next page of results.
      --ResultsPerPage int                                    The maximum number of logs returned in a single response (between 1 and 1000, both included).
  -h, --help                                                  help for ReadSnapshots
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

* [octl iaas api](octl_iaas_api.md)	 - iaas api calls

