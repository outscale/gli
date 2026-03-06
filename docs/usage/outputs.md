# Output formats and output-related options

## Output-related options

| Option | Default | Allowed values | Description |
| ------ | ------- | -------------- | ----------- |
| `--jq` | | | jq-like filter |
| `--filter` | | | content filter |
| `--template` | | | JSON template for query body |
| `--output` | | `raw`, `json`, `yaml`, `table`, `csv`, `base64`, `none` | output format |
| `--columns` | | | columns to display in a table |
| `--single` | | | convert single entry lists to a single object |

## Output formats

- `raw` is the raw JSON format, as returned by the API.
- `json` displays the content in JSON format, without response context.
- `yaml` displays the content in YAML format, without response context.
- `table` displays the content in CSV format, based on columns defined by the `--columns` flag.
- `csv` displays the content in a text table, based on columns defined by the `--columns` flag.
- `base64` decodes base64-encoded strings or lists of strings.
- `none` disables output.

> Please note that `raw` output returns the raw payload whereas the other formats only output the content
(e.g. a list of VMs when listing VMs instead of an object with a `Vms` attribute storing the list).