## octl kube secret

Create secret for CCM or CSI driver deployment

### Synopsis

Create secret for CCM/CSI driver/Cluster-API provider deployments using the selected AK/SK.

CCM:
```shell
octl kube secret --name osc-secret | kubectl apply -f -
```
CSI driver:
```shell
octl kube secret --name osc-csi-bsu | kubectl apply -f -
```
Cluster-API:
```shell
octl kube secret --name cluster-api-provider-outscale --namespace cluster-api-provider-outscale-system | kubectl apply -f -
```


```
octl kube secret [flags]
```

### Options

```
  -h, --help               help for secret
      --name string        [REQUIRED] name of secret
      --namespace string   namespace of secret (default "kube-system")
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

* [octl kube](octl_kube.md)	 - OUTSCALE Kubernetes as a Service (OKS) management

