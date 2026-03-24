## octl iaas imageexporttask create

alias for api CreateImageExportTask

### Synopsis

> *alias for api CreateImageExportTask*

Exports an OUTSCALE machine image (OMI) to an OUTSCALE Object Storage (OOS) bucket.

This enables you to copy an OMI between OUTSCALE accounts in different Regions.

This action creates the necessary snapshots and manifest file in the bucket. The OMI can then be imported to another account using a pre-signed URL of its manifest file. For more information, see [Creating a Pre-Signed URL](https://docs.outscale.com/en/userguide/Creating-a-Pre-Signed-URL.html).

To copy an OMI in the same Region, you can also use the [CreateImage](#createimage) method.

**[IMPORTANT]**

* You cannot export a shared or public OMI, as they do not belong to you. To do so, you must first copy it to your account. The copy then belongs to you and you can export it.

* Export tasks can only be canceled while in the `pending/queued` state.

For more information, see [About OMIs](https://docs.outscale.com/en/userguide/About-OMIs.html).

```
octl iaas imageexporttask create [flags]
```

### Options

```
  -h, --help                                       help for create
      --image-id string                            The ID of the OMI to export.
      --osu-export-disk-image-format string        The format of the export disk (qcow2 | raw).
      --osu-export-osu-api-key-api-key-id string   The API key of the OOS account that enables you to access the bucket.
      --osu-export-osu-api-key-secret-key string   The secret key of the OOS account that enables you to access the bucket.
      --osu-export-osu-bucket string               The name of the OOS bucket where you want to export the object.
      --osu-export-osu-manifest-url string         The URL of the manifest file.
      --osu-export-osu-prefix string               The prefix for the key of the OOS object.
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

* [octl iaas imageexporttask](octl_iaas_imageexporttask.md)	 - imageexporttask commands

