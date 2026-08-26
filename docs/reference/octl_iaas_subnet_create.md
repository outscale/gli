## octl iaas subnet create

alias for api CreateSubnet

### Synopsis

> alias for api CreateSubnet

Creates a Subnet in an existing Net.

To create a Subnet in a Net, you have to provide the ID of the Net and the IP range for the Subnet (its network range). Once the Subnet is created, you cannot modify its IP range.


For more information, see [About Nets](https://docs.outscale.com/en/userguide/About-Nets.html).

```
octl iaas subnet create [flags]
```

### Options

```
  -h, --help               help for create
      --ip-range string    [REQUIRED] The IP range in the Subnet, in CIDR notation (for example, 10.0.0.0/16).
      --net-id string      [REQUIRED] The ID of the Net for which you want to create a Subnet.
      --subregion string   The name of the Subregion in which you want to create the Subnet.
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
  -o, --output string              output format (raw, json, yaml, table, csv, none, base64, text)
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

* [octl iaas subnet](octl_iaas_subnet.md)	 - subnet commands

