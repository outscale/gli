## octl iaas publicip delete

alias for api DeletePublicIp --PublicIpId public_ip_id

### Synopsis

> *alias for api DeletePublicIp --PublicIpId public_ip_id*

Releases a public IP.

You can release a public IP associated with your OUTSCALE account. This address is released in the public IP pool and can be used by someone else. Before releasing a public IP, ensure you updated all your resources communicating with this address.

```
octl iaas publicip delete public_ip_id [public_ip_id]... [flags]
```

### Options

```
  -h, --help               help for delete
      --public-ip string   The public IP.
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

* [octl iaas publicip](octl_iaas_publicip.md)	 - publicip commands

