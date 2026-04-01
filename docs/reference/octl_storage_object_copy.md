## octl storage object copy

Copy an object, alias for api CopyObject --CopySource bucket_src/key_src --Key key_dst

### Synopsis

> *Copy an object, alias for api CopyObject --CopySource bucket_src/key_src --Key key_dst*



```
octl storage object copy bucket_src/key_src key_dst [flags]
```

### Options

```
      --bucket string                           The name of the destination bucket.
      --cache-control string                    Specifies the caching behavior along the request/reply chain.
      --content-disposition string              Specifies presentational information for the object.
      --content-encoding string                 Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field.
      --content-language string                 The language the content is in.
      --content-type string                     A standard MIME type that describes the format of the object data.
      --expires osctime                         The date and time at which the object is no longer cacheable.
      --grant-full-control string               Gives the grantee READ, READ_ACP, and WRITE_ACP permissions on the object.
      --grant-read string                       Allows grantee to read the object data and its metadata.
      --grant-read-acp string                   Allows grantee to read the object ACL.
      --grant-write-acp string                  Allows grantee to write the ACL for the applicable object.
  -h, --help                                    help for copy
      --object-lock-mode string                 The Object Lock mode that you want to apply to the object copy.
      --object-lock-retain-until-date osctime   The date and time when you want the Object Lock of the object copy to expire.
      --server-side-encryption string           The server-side encryption algorithm used when storing this object in Amazon S3.
      --tagging string                          The tag-set for the object copy in the destination bucket.
      --tagging-directive string                Specifies whether the object tag-set is copied from the source object or replaced with the tag-set that's provided in the request.
      --website-redirect-location string        If the destination bucket is configured as a website, redirects requests for this object copy to another object in the same bucket or to an external URL.
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

* [octl storage object](octl_storage_object.md)	 - object commands

