## octl iaas dhcpoption list

alias for api ReadDhcpOptions

### Synopsis

> *alias for api ReadDhcpOptions*

Gets information about the content of one or more DHCP options sets.

```
octl iaas dhcpoption list [flags]
```

### Options

```
      --default                      If true, lists all default DHCP options set.
      --domain-name strings          The domain names used for the DHCP options sets.
      --domain-name-server strings   The IPs of the domain name servers used for the DHCP options sets.
  -h, --help                         help for list
      --id strings                   The IDs of the DHCP options sets.
      --log-server strings           The IPs of the log servers used for the DHCP options sets.
      --ntp-server strings           The IPs of the Network Time Protocol (NTP) servers used for the DHCP options sets.
      --tag strings                  The key/value combination of the tags associated with the DHCP options sets, in the following format: "Filters":{"Tags":["TAGKEY=TAGVALUE"]}.
      --tag-key strings              The keys of the tags associated with the DHCP options sets.
      --tag-value strings            The values of the tags associated with the DHCP options sets.
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

* [octl iaas dhcpoption](octl_iaas_dhcpoption.md)	 - dhcpoption commands

