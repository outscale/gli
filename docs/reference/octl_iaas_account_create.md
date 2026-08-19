## octl iaas account create

alias for api CreateAccount

### Synopsis

> *alias for api CreateAccount*

Creates an OUTSCALE account.

**[IMPORTANT]**

* You need OUTSCALE credentials and the appropriate quotas to create an account via API. To get quotas, you can send an email to sales@outscale.com.

For more information, see [About Your Account](https://docs.outscale.com/en/userguide/About-Your-OUTSCALE-Account.html).

```
octl iaas account create [flags]
```

### Options

```
      --additional-email strings   One or more additional email addresses for the account.
      --city string                [REQUIRED] The city of the account owner.
      --company-name string        [REQUIRED] The name of the company for the account.
      --country string             [REQUIRED] The country of the account owner.
      --customer-id string         [REQUIRED] The ID of the customer.
      --email string               [REQUIRED] The main email address for the account.
      --first-name string          [REQUIRED] The first name of the account owner.
  -h, --help                       help for create
      --job-title string           The job title of the account owner.
      --last-name string           [REQUIRED] The last name of the account owner.
      --mobile-number string       The mobile phone number of the account owner.
      --phone-number string        The landline phone number of the account owner.
      --state-province string      The state/province of the account.
      --vat-number string          The value added tax (VAT) number for the account.
      --zip-code string            [REQUIRED] The ZIP code of the city.
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

* [octl iaas account](octl_iaas_account.md)	 - account commands

