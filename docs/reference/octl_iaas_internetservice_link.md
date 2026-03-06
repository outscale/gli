## octl iaas internetservice link

alias for api LinkInternetService --InternetServiceId service_id

### Synopsis

> *alias for api LinkInternetService --InternetServiceId service_id*

Attaches an internet service to a Net.

To enable the connection between the Internet and a Net, you must attach an internet service to this Net.

```
octl iaas internetservice link service_id [flags]
```

### Options

```
  -h, --help            help for link
      --net-id string   The ID of the Net to which you want to attach the internet service.
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

* [octl iaas internetservice](octl_iaas_internetservice.md)	 - internetservice commands

