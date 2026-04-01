## octl storage api PutBucketWebsite



```
octl storage api PutBucketWebsite [flags]
```

### Options

```
      --Bucket string                                                                      The bucket name.
      --ChecksumAlgorithm string                                                           Indicates the algorithm used to create the checksum for the object when you use the SDK.
      --ContentMD5 string                                                                  The base64-encoded 128-bit MD5 digest of the data.
      --ExpectedBucketOwner string                                                         The account ID of the expected bucket owner.
      --WebsiteConfiguration.ErrorDocument.Key string                                      The object key name to use when a 4XX class error occurs.
      --WebsiteConfiguration.IndexDocument.Suffix string                                   A suffix that is appended to a request that is for a directory on the website endpoint.
      --WebsiteConfiguration.RedirectAllRequestsTo.HostName string                         Name of the host where requests are redirected.
      --WebsiteConfiguration.RedirectAllRequestsTo.Protocol string                         to use when redirecting requests.
      --WebsiteConfiguration.RoutingRules.0.Condition.HttpErrorCodeReturnedEquals string   The HTTP error code when the redirect is applied.
      --WebsiteConfiguration.RoutingRules.0.Condition.KeyPrefixEquals string               The object key name prefix when the redirect is applied.
      --WebsiteConfiguration.RoutingRules.0.Redirect.HostName string                       The host name to use in the redirect request.
      --WebsiteConfiguration.RoutingRules.0.Redirect.HttpRedirectCode string               The HTTP redirect code to use on the response.
      --WebsiteConfiguration.RoutingRules.0.Redirect.Protocol string                       to use when redirecting requests.
      --WebsiteConfiguration.RoutingRules.0.Redirect.ReplaceKeyPrefixWith string           The object key prefix to use in the redirect request.
      --WebsiteConfiguration.RoutingRules.0.Redirect.ReplaceKeyWith string                 The specific object key to use in the redirect request.
  -h, --help                                                                               help for PutBucketWebsite
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

* [octl storage api](octl_storage_api.md)	 - storage api calls

