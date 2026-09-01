## octl iaas api UnlinkPublicIp

Disassociates a public IP from the virtual machine (VM) or network interface card (NIC) it is associated with.

### Synopsis

Disassociates a public IP from the virtual machine (VM) or network interface card (NIC) it is associated with.



**[IMPORTANT]**

To disassociate the public IP from a NAT service, you need to delete the NAT service. For more information, see the [DeleteNatService](#deletenatservice) method.

```
octl iaas api UnlinkPublicIp [flags]
```

### Options

```
      --DryRun                  If true, checks whether you have the required permissions to perform the action.
      --LinkPublicIpId string   The ID representing the association of the public IP with the VM or the NIC.
      --PublicIp string         The public IP.
  -h, --help                    help for UnlinkPublicIp
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

* [octl iaas api](octl_iaas_api.md)	 - Call iaas API

