## octl iaas keypair list

alias for api ReadKeypairs

### Synopsis

> *alias for api ReadKeypairs*

Lists one or more of your keypairs.

```
octl iaas keypair list [flags]
```

### Options

```
      --fingerprint strings   The fingerprints of the keypairs.
  -h, --help                  help for list
      --id strings            The IDs of the keypairs.
      --name strings          The names of the keypairs.
      --tag strings           The key/value combination of the tags associated with the keypairs, in the following format: "Filters":{"Tags":["TAGKEY=TAGVALUE"]}.
      --tag-key strings       The keys of the tags associated with the keypairs.
      --tag-value strings     The values of the tags associated with the keypairs.
      --type strings          The types of the keypairs (ssh-rsa, ssh-ed25519, ecdsa-sha2-nistp256, ecdsa-sha2-nistp384, or ecdsa-sha2-nistp521).
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

* [octl iaas keypair](octl_iaas_keypair.md)	 - keypair commands

