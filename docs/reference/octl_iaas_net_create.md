## octl iaas net create

alias for api CreateNet

### Synopsis

> alias for api CreateNet

Creates a Net with a specified IP range.

The IP range (network range) of your Net must be between a /28 netmask (16 IPs) and a /16 netmask (65536 IPs).


For more information, see [About Nets](https://docs.outscale.com/en/userguide/About-Nets.html).

```
octl iaas net create [flags]
```

### Options

```
  -h, --help              help for create
      --ip-range string   [REQUIRED] The IP range for the Net, in CIDR notation (for example, 10.0.0.0/16).
      --tenancy string    The tenancy options for the VMs: - default if a VM created in a Net can be launched with any tenancy.
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

* [octl iaas net](octl_iaas_net.md)	 - net commands

