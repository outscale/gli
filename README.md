## **octl, An Experimental CLI for Outscale, written in Go**

[![Project Stage](https://docs.outscale.com/fr/userguide/_images/Project-Sandbox-yellow.svg)](https://docs.outscale.com/en/userguide/Open-Source-Projects.html) [![](https://dcbadge.limes.pink/api/server/HUVtY5gT6s?style=flat&theme=default-inverted)](https://discord.gg/HUVtY5gT6s)

---

## 🌐 Links

- Documentation: <https://docs.outscale.com/en/>
- Project website: <https://github.com/outscale/octl>
- Join our community on [Discord](https://discord.gg/HUVtY5gT6s)

---

## 📄 Table of Contents

- [Overview](#-overview)
- [Requirements](#-requirements)
- [Installation](#-installation)
- [Configuration](#-configuration)
- [Usage](#-usage)
- [License](#-license)
- [Contributing](#-contributing)

---

## 🧭 Overview

**octl** is an experimental CLI for the Outscale APIs, written in Go.

It supports:
* installation via a single static binary,
* direct flags to all request fields, without JSON,
* autocompletion support for all API calls, flags, and flag values,
* jq-style output filters,
* syntax highlighting of output,
* auto-update to latest version.

It currenty only supports iaas api, but other Outscale APIs are planned.

---

## ✅ Requirements

- Access to the OUTSCALE API (with appropriate credentials)

---

## ⚙ Installation

Download the latest binary from the [Releases page](https://github.com/outscale/octl/releases).

### Autocompletion configuration

#### Bash

```shell
octl completion bash > octl-completion.bash
sudo cp octl-completion.bash /etc/bash_completion.d/
source ~/.bashrc
```

#### Zsh

```shell
octl completion zsh > _octl
sudo mkdir -p /usr/local/share/zsh/site-functions
sudo cp _octl /usr/local/share/zsh/site-functions/
source ~/.zshrc
```

#### Fish

```shell
octl completion fish > octl.fish
sudo cp octl.fish /usr/share/fish/completions/
source ~/.config/fish/config.fish
```

---

## 🛠 Configuration

The tool expects either environment variables or a configuration file.

### Environment variables

The tool will try to read the following environment variables:
* `OSC_ACCESS_KEY`
* `OSC_SECRET_KEY`
* `OSC_REGION`

### Profile file

If no environment variables are defined, the tool expects to find a profile in a profile file.

The default profile file path is `~/.osc/config.json` and can be set with the `--config` flag or the `OSC_CONFIG_FILE` environment variable.

The default profile name path is `default` and can be set with the `--profile` flag or the `OSC_PROFILE` environment variable.

### Example

```json
{
  "default": {
    "access_key": "MyAccessKey",
    "secret_key": "MySecretKey",
    "region": "eu-west-2"
  }
}
```

Use `OSC_CONFIG_FILE` to define an alternate config file and `OSC_PROFILE` an alternate profile.

---

## 🚀 Usage

```bash
octl <command> <command>
```

### Commands

| Command | Description |
| ------- | ----------- |
| `iaas` | Core IaaS API |
| `update` | Update to the latest version |
| `completion` | Generate completion shell script |

### Options

| Option | Default | Allowed values | Description |
| ------ | ------- | -------------- | ----------- |
| `--version` | | | Display octl version |
| `-v, --verbose` | | | Dump HTTP request and response |
| `-h, --help` | | | Help about a command |
| `--jq` | | | jq-like output filter |
| `--template` | | | JSON template for query body |
| `--config` | `~/.osc/config.json` | | config file path |
| `--profile` | `default` | | profile name |
| `--output` | | `raw`, `json`, `yaml`, `table` | output format |
| `--columns` | | | columns to display in a table |

### Output formats

* `raw` is the raw JSON format, as returned by the API.
* `json` displays the content in JSON format, without response context,
* `yaml` displays the content in YAML format, without response context,
* `table` displays the content in a text table, based on columns defined by the `--columns` flag.

### High level command

High level commands are available:

* `octl iaas <entity> list` lists all entities using the `table` format:
```shell
octl iaas volume list
┌──────────────┬──────────┬───────────┬───────────────┬──────┬──────┐
│      ID      │   Type   │   State   │ SubregionName │ Size │ Iops │
├──────────────┼──────────┼───────────┼───────────────┼──────┼──────┤
│ vol-foo      │ io1      │ in-use    │ eu-west-2a    │ 300  │ 5000 │
│ vol-bar      │ standard │ in-use    │ eu-west-2a    │ 20   │ 0    │
│ vol-baz      │ gp2      │ available │ eu-west-2a    │ 4    │ 100  │
└──────────────┴──────────┴───────────┴───────────────┴──────┴──────┘
```
* `octl iaas <entity> describe <id>` displays an entity using the `yaml` format:
```shell
octl iaas volume describe vol-foo
CreationDate: '2024-12-17T11:07:58.757Z'
Iops: 5000
LinkedVolumes:
- DeviceName: /dev/sda1
  State: attached
  VmId: i-foo
  VolumeId: vol-foo
Size: 300
SnapshotId: snap-foo
State: in-use
SubregionName: eu-west-2a
Tags: []
VolumeId: vol-foo
VolumeType: io1
```

Columns can be replaced:
```shell
octl iaas vm list --columns "ID:VmId|DNS:PrivateDnsName"
┌────────────┬───────────────────────────────────────────┐
│     ID     │                    DNS                    │
├────────────┼───────────────────────────────────────────┤
│ i-foo      │ ip-10-1-112-23.eu-west-2.compute.internal │
│ i-bar      │ ip-10-9-35-211.eu-west-2.compute.internal │
│ i-baz      │ ip-10-0-4-143.eu-west-2.compute.internal  │
└────────────┴───────────────────────────────────────────┘
```

Columns can be added to the standard columns:
```shell
octl iaas vm list --columns +DNS:PrivateDnsName
```

Column content is defined with the [expr language](https://expr-lang.org/docs/language-definition). To display a tag value:
```shell
octl iaas vm list --columns "+tag:find(Tags, #?.Key == \"Name\")?.Value"
```


### API access

The API can be directly called, with a `raw` output:
```shell
octl iaas api ReadVms --Filters.VmStateNames running
```

### Flag syntax

The flag syntax is:
* list of values are comma-separated: `--Filters.VmStateNames running,stopped`,
* boolean flags can be set to false by setting: `--TmpEnabled=false`,
* lists of embedded objects (e.g. `Nics` or `BlockDeviceMappings` in `CreateVms`) can be configured using indexes: `--BlockDeviceMappings.0.Bsu.VolumeType`,
* time flag values can be set:
  * using the RFC3339 format (e.g. `2026-02-10T14:52:30Z`),
  * as a duration offset with a `+` or `-` prefix (e.g. `+10m`, `-1h`),
  * as a day/month/year offset with a `+` or `-` prefix (e.g. `+1mo`, `-1y`).

### Chaining

Commands may be chained, and attributes returned by a command can be reinjected in a subsequent command, using Go template syntax:

```shell
octl iaas api CreateNic --SubnetId subnet-foo | octl iaas api LinkNic -v --NicId {{.Nic.NicId}} --VmId i-foo --DeviceNumber 7
```

### Sending raw JSON

```shell
echo '{"SubnetId":"subnet-foo"}' | octl iaas api CreateNic
```

### Templating

A JSON document can be used as a template, with additional config using flags.

Either from stdin:
```shell
echo '{"NetId":"vpc-foo"}' | octl iaas api CreateSubnet --IpRange 10.0.1.0/24
```
Or from a file:
```shell
octl iaas api CreateSubnet --IpRange 10.0.1.0/24 --template subnet.json
```

### Using jq filters

```shell
octl iaas api ReadVms --Filters.VmStateNames running --jq ".Vms[].VmId"
```

> Note: `--jq` is not currently compatible with `--output`

### Self updating

```shell
octl update
```

> This requires write access to the binary. If `octl update` does not work, you will need to download the binary from the [latest release](./releases/latest).

---

## 📜 License

**octl** is released under the BSD 3-Clause license.

© 2026 Outscale SAS

See [LICENSE](./LICENSE) for full details.

---

## 🤝 Contributing

We welcome contributions!

Please read our [Contributing Guidelines](CONTRIBUTING.md) and [Code of Conduct](CODE_OF_CONDUCT.md) before submitting a pull request.
