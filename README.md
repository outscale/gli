## **GLI, An Experimental CLI for Outscale, written in Go**

[![Project Stage](https://docs.outscale.com/fr/userguide/_images/Project-Sandbox-yellow.svg)](https://docs.outscale.com/en/userguide/Open-Source-Projects.html) [![](https://dcbadge.limes.pink/api/server/HUVtY5gT6s?style=flat&theme=default-inverted)](https://discord.gg/HUVtY5gT6s)

<p align="center">
  <img alt="<Project Logo or Icon>" src="<Logo URL or Placeholder>" width="100px">
</p>

---

## 🌐 Links

- Documentation: <https://docs.outscale.com/en/>
- Project website: <https://github.com/outscale/<project-name>>
- Join our community on [Discord](https://discord.gg/HUVtY5gT6s)
- Related tools or community: <<https://example.com>> *(optional)*

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

**CLI** is an experimental CLI for the Outscale APIs, written in Go.

---

## ✅ Requirements

- Access to the OUTSCALE API (with appropriate credentials)

---

## ⚙ Installation

Download the latest binary from the [Releases page](https://github.com/outscale/gli/releases).

---

## 🛠 Configuration


The tool expects either environment variables or a configuration file.

### Environment variables

The tool will try to read the following environment variables:
* `OSC_ACCESS_KEY`
* `OSC_SECRET_KEY`
* `OSC_REGION`

### Profile file

If no environment variables are defined, the tool will read `~/.osc/config.json` and look for the `default` profile.

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

### ReadVms

List all VMs in the `running` state:
```bash
gli oapi ReadVms --Filters.VmStateNames running
```

### Using jq filters

```bash
gli oapi ReadVms --Filters.VmStateNames running --jq ".Vms[].VmId"
```

### Self updating

```bash
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
