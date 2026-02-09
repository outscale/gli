## **GLI, An Experimental CLI for Outscale, written in Go**

[![Project Stage](https://docs.outscale.com/fr/userguide/_images/Project-Sandbox-yellow.svg)](https://docs.outscale.com/en/userguide/Open-Source-Projects.html) [![](https://dcbadge.limes.pink/api/server/HUVtY5gT6s?style=flat&theme=default-inverted)](https://discord.gg/HUVtY5gT6s)

---

## 🌐 Links

- Documentation: <https://docs.outscale.com/en/>
- Project website: <https://github.com/outscale/gli>
- Join our community on [Discord](https://discord.gg/HUVtY5gT6s)

---

## 📄 Table of Contents

- [Overview](#-overview)
- [Requirements](#-requirements)
- [Installation](#-installation)
- [Configuration](#-configuration)
- [Usage](#-usage)
- [Examples](#-examples)
- [License](#-license)
- [Contributing](#-contributing)

---

## 🧭 Overview

**GLI** is an experimental CLI for the Outscale APIs, written in Go.

It supports:
* installation via a single static binary,
* direct flags to all request fields, without JSON,
* autocompletion support for all API calls, flags, and flag values,
* jq-style output filters,
* syntax highlighting of output,
* auto-update to latest version.

It currenty only supports OAPI, but other Outscale APIs are planned.

---

## ✅ Requirements

- Access to the OUTSCALE API (with appropriate credentials)

---

## ⚙ Installation

Download the latest binary from the [Releases page](https://github.com/outscale/gli/releases).

### Autocompletion configuration

#### Bash

```shell
gli completion bash > gli-completion.bash
sudo cp gli-completion.bash /etc/bash_completion.d/
source ~/.bashrc
```

#### Zsh

```shell
gli completion zsh > _gli
sudo mkdir -p /usr/local/share/zsh/site-functions
sudo cp _gli /usr/local/share/zsh/site-functions/
source ~/.zshrc
```

#### Fish

```shell
gli completion fish > gli.fish
sudo cp gli.fish /usr/share/fish/completions/
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
gli <command> <api call>
```

### Commands

| Command | Description |
| ------- | ----------- |
| `oapi` | Call OAPI |
| `update` | Update to the latest version |
| `version` | Display version |

### Options

| Option | Description |
| ------ | ----------- |
| `-v, --verbose` | Dump HTTP request and response |
| `-h, --help` | Help about a command |
| `--jq` | jq-like output filter |

---

## 💡 Examples

### Querying oapi

List all VMs in the `running` state:
```shell
gli oapi ReadVms --Filters.VmStateNames running
```

The flag syntax is:
* list of values are comma-separated: `--Filters.VmStateNames running,stopped`
* boolean flags can be set to false by setting: `--TmpEnabled=false`

### Multiple values

Lists of embedded objects (e.g. `Nics` or `BlockDeviceMappings` in `CreateVms`) can be configured using indexes: `--BlockDeviceMappings.0.Bsu.VolumeType`.

### Chaining

Commands may be chained, and attributes returned by a command can be reinjected in a subsequent command, using Go template syntax:

```shell
gli oapi CreateNic --SubnetId subnet-foo | gli oapi LinkNic -v --NicId {{.Nic.NicId}} --VmId i-foo --DeviceNumber 7
```

### Sending raw JSON

```shell
echo '{"SubnetId":"subnet-foo"}' | gli oapi CreateNic
```

### Using jq filters

```shell
gli oapi ReadVms --Filters.VmStateNames running --jq ".Vms[].VmId"
```

### Self updating

```shell
gli update
```

> This requires write access to the binary. If `gli update` does not work, you will need to download the binary from the [latest release](./releases/latest).

---

## 📜 License

**GLI** is released under the BSD 3-Clause license.

© 2026 Outscale SAS

See [LICENSE](./LICENSE) for full details.

---

## 🤝 Contributing

We welcome contributions!

Please read our [Contributing Guidelines](CONTRIBUTING.md) and [Code of Conduct](CODE_OF_CONDUCT.md) before submitting a pull request.
