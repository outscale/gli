## octl iaas api LinkPublicIp

Associates a public IP with a virtual machine (VM) or a network interface card (NIC), in the public Cloud or in a Net.

### Synopsis

Associates a public IP with a virtual machine (VM) or a network interface card (NIC), in the public Cloud or in a Net. You can associate a public IP with only one VM or network interface at a time.

To associate a public IP in a Net, ensure that the Net has an internet service attached. For more information, see the [LinkInternetService](#linkinternetservice) method.

By default, the public IP is disassociated every time you stop and start the VM. For a persistent association, you can add the `osc.fcu.eip.auto-attach` tag to the VM with the public IP as value. For more information, see the [CreateTags](#createtags) method.

**[IMPORTANT]**

You can associate a public IP with a network address translation (NAT) service only when creating the NAT service. To modify its public IP, you need to delete the NAT service and re-create it with the new public IP. For more information, see the [CreateNatService](#createnatservice) method.

```
octl iaas api LinkPublicIp [flags]
```

### Options

```
      --AllowRelink         If true, allows the public IP to be associated with the VM or NIC that you specify even if it is already associated with another VM or NIC.
      --DryRun              If true, checks whether you have the required permissions to perform the action.
      --NicId string        (Net only) The ID of the NIC.
      --PrivateIp string    (Net only) The primary or secondary private IP of the specified NIC.
      --PublicIp string     The public IP.
      --PublicIpId string   The allocation ID of the public IP.
      --VmId string         The ID of the VM.
  -h, --help                help for LinkPublicIp
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

* [octl iaas api](octl_iaas_api.md)	 - iaas api calls

