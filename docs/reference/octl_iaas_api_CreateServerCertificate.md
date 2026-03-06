## octl iaas api CreateServerCertificate

Creates a server certificate and its matching private key.

### Synopsis

Creates a server certificate and its matching private key.

These elements can be used with other services (for example, to configure SSL termination on load balancers).

You can also specify the chain of intermediate certification authorities if your certificate is not directly signed by a root one. You can specify multiple intermediate certification authorities in the `CertificateChain` parameter. To do so, concatenate all certificates in the correct order (the first certificate must be the authority of your certificate, the second must be the authority of the first one, and so on).

The private key must be a RSA key in PKCS1 form. To check this, open the PEM file and ensure its header reads as follows: BEGIN RSA PRIVATE KEY.

[IMPORTANT]

This private key must not be protected by a password or a passphrase.

For more information, see [About Server Certificates in EIM](https://docs.outscale.com/en/userguide/About-Server-Certificates-in-EIM.html).

```
octl iaas api CreateServerCertificate [flags]
```

### Options

```
      --Body string         The PEM-encoded X509 certificate.
      --Chain string        The PEM-encoded intermediate certification authorities.
      --DryRun              If true, checks whether you have the required permissions to perform the action.
      --Name string         A unique name for the certificate.
      --Path string         The path to the server certificate, set to a slash (/) if not specified.
      --PrivateKey string   The PEM-encoded private key matching the certificate.
  -h, --help                help for CreateServerCertificate
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

* [octl iaas api](octl_iaas_api.md)	 - iaas api calls

