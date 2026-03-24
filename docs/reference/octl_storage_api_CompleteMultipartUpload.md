## octl storage api CompleteMultipartUpload



```
octl storage api CompleteMultipartUpload [flags]
```

### Options

```
      --Bucket string                                   Name of the bucket to which the multipart upload was initiated.
      --ChecksumCRC32 string                            This header can be used as a data integrity check to verify that the data received is the same data that was originally sent.
      --ChecksumCRC32C string                           This header can be used as a data integrity check to verify that the data received is the same data that was originally sent.
      --ChecksumSHA1 string                             This header can be used as a data integrity check to verify that the data received is the same data that was originally sent.
      --ChecksumSHA256 string                           This header can be used as a data integrity check to verify that the data received is the same data that was originally sent.
      --ExpectedBucketOwner string                      The account ID of the expected bucket owner.
      --IfMatch string                                  Uploads the object only if the ETag (entity tag) value provided during the WRITE operation matches the ETag of the object in S3.
      --IfNoneMatch string                              Uploads the object only if the object key name does not already exist in the bucket specified.
      --Key string                                      Object key for which the multipart upload was initiated.
      --MultipartUpload.Parts.0.ChecksumCRC32 string    
      --MultipartUpload.Parts.0.ChecksumCRC32C string   
      --MultipartUpload.Parts.0.ChecksumSHA1 string     
      --MultipartUpload.Parts.0.ChecksumSHA256 string   
      --MultipartUpload.Parts.0.ETag string             
      --MultipartUpload.Parts.0.PartNumber int32        
      --RequestPayer string                             Confirms that the requester knows that they will be charged for the request.
      --SSECustomerAlgorithm string                     The server-side encryption (SSE) algorithm used to encrypt the object.
      --SSECustomerKey string                           The server-side encryption (SSE) customer managed key.
      --SSECustomerKeyMD5 string                        The MD5 server-side encryption (SSE) customer managed key.
      --UploadId string                                 ID for the initiated multipart upload.
  -h, --help                                            help for CompleteMultipartUpload
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

