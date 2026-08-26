## octl iaas keypair create

alias for api CreateKeypair

### Synopsis

> alias for api CreateKeypair

Creates a keypair to use with your virtual machines (VMs).

You can use this method in two different ways:
* **Creating a keypair**: In that case, 3DS OUTSCALE creates a 2048-bit RSA keypair, stores its public key in your OUTSCALE account, and returns its private key in the response of the call so that you can save it in a file.

When you save the returned private key, make sure you replace the `\n` escape sequences with real line breaks.
* **Importing a keypair created locally**: If you already have a keypair that you have created locally with a third-party tool, you can import its public key in your account. The following types of key can be imported: RSA (2048 bits or preferably 4096 bits), Ed25519, and ECDSA (256 bits, 384 bits, or 521 bits). The following formats can be used: PEM, PKCS8, RFC4716, and OpenSSH.

For more information, see [About Keypairs](https://docs.outscale.com/en/userguide/About-Keypairs.html).

```
octl iaas keypair create [flags]
```

### Options

```
  -h, --help                    help for create
      --name string             [REQUIRED] A unique name for the keypair, with a maximum length of 255 [ASCII printable characters](https://en.wikipedia.org/wiki/ASCII#Printable_characters).
      --public-key base64File   The file storing the public key to import in your account, if you are importing an existing keypair.
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

* [octl iaas keypair](octl_iaas_keypair.md)	 - keypair commands

