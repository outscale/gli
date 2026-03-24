## octl iaas api UpdateAccount

Updates the OUTSCALE account information for the account that sends the request.

### Synopsis

Updates the OUTSCALE account information for the account that sends the request.

```
octl iaas api UpdateAccount [flags]
```

### Options

```
      --AdditionalEmails strings   One or more additional email addresses for the account.
      --City string                The new city of the account owner.
      --CompanyName string         The new name of the company for the account.
      --Country string             The new country of the account owner.
      --DryRun                     If true, checks whether you have the required permissions to perform the action.
      --Email string               The main email address for the account.
      --FirstName string           The new first name of the account owner.
      --JobTitle string            The new job title of the account owner.
      --LastName string            The new last name of the account owner.
      --MobileNumber string        The new mobile phone number of the account owner.
      --PhoneNumber string         The new landline phone number of the account owner.
      --StateProvince string       The new state/province of the account owner.
      --VatNumber string           The new value added tax (VAT) number for the account.
      --ZipCode string             The new ZIP code of the city.
  -h, --help                       help for UpdateAccount
```

### Options inherited from parent commands

```
  -c, --columns string              columns to display - [+]<title>:<jq query for content>||<title>:<jq query for content>
      --config string               Path of profile file (by default, ~/.osc/config.json)
      --filter strings              comma separated list of filters for results - name:value,name:value, alias for jq filter 'select(.name | test("value"))'
      --jq string                   jq filter
      --no-upgrade                  do not check for new versions
  -O, --out-file string             redirect output to file
  -o, --output string               output format (raw, json, yaml, table, csv, none, base64) (default "raw")
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

