## octl iaas consoleoutput list

alias for api ReadConsoleOutput

### Synopsis

> *alias for api ReadConsoleOutput*

Gets the console output for a virtual machine (VM). This console is not in real-time. It is refreshed every two seconds and provides the most recent 64 KiB output.

**[IMPORTANT]**

On Windows VMs, the console is handled only on the first boot. It returns no output after the first boot.

```
octl iaas consoleoutput list [flags]
```

### Options

```
  -h, --help           help for list
      --vm-id string   The ID of the VM.
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

* [octl iaas consoleoutput](octl_iaas_consoleoutput.md)	 - consoleoutput commands

