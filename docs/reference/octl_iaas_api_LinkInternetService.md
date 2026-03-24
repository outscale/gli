## octl iaas api LinkInternetService

Attaches an internet service to a Net.

### Synopsis

Attaches an internet service to a Net.

To enable the connection between the Internet and a Net, you must attach an internet service to this Net.

```
octl iaas api LinkInternetService [flags]
```

### Options

```
      --DryRun                     If true, checks whether you have the required permissions to perform the action.
      --InternetServiceId string   The ID of the internet service you want to attach.
      --NetId string               The ID of the Net to which you want to attach the internet service.
  -h, --help                       help for LinkInternetService
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

* [octl iaas api](octl_iaas_api.md)	 - iaas api calls

