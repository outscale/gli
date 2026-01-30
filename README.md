## **< Project Name >**

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

**< Project Name >** is a <short description of what the project does, e.g., "CLI tool to manage...">.

Key features:
- <Feature 1>
- <Feature 2>
- <Feature 3>

---

## ✅ Requirements

- <Dependency 1> (e.g., Rust, Go, Python 3.11+)
- <Dependency 2> (e.g., Git)
- Access to the OUTSCALE API (with appropriate credentials)

---

## ⚙ Installation

### Option 1: Download from Releases

Download the latest binary from the [Releases page](https://github.com/outscale/<project-name>/releases).

### Option 2: Install from source

```bash
git clone https://github.com/outscale/<project-name>.git
cd <project-name>
<build or install command>
````

Example (for Go projects):

```bash
go install github.com/outscale/<project-name>@latest
```

---

## 🛠 Configuration

\<Explain where the credentials or config are stored, e.g.:>

The tool expects a configuration file at `~/.osc/config.json`.

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

Use the `--profile` flag to select another profile.

---

## 🚀 Usage

```bash
<command> [OPTIONS]
```

### Options

| Option             | Description                            |
| ------------------ | -------------------------------------- |
| `-f, --flag`       | What this flag does                    |
| `-c, --count <N>`  | Run N times                            |
| `--profile <name>` | Use a specific profile from the config |
| `-v, --version`    | Print version and exit                 |

---

## 💡 Examples

### Basic usage

```bash
<command>
```

### With options

```bash
<command> --flag value --profile test
```

### Using `jq` to filter JSON output

```bash
jq '.[] | select(.ResponseStatusCode != 200)' logs.json
```

---

## 📜 License

**< Project Name >** is released under the < License Name > license.

© < Year > Outscale SAS

See [LICENSE](./LICENSE) for full details.

---

## 🤝 Contributing

We welcome contributions!

Please read our [Contributing Guidelines](CONTRIBUTING.md) and [Code of Conduct](CODE_OF_CONDUCT.md) before submitting a pull request.

---

### Notes for reuse:
- Replace all `<...>` placeholders with your content.
- You can prefill the `Project Stage` badge with values like:
  - `Project-Incubating-blue.svg`
  - `Project-Graduated-green.svg`
- You may include platform-specific instructions (macOS/Linux/Windows) in collapsible `<details>` blocks if needed.

>Labels are centrally managed in the outscale/.github repository (labels.yml).
>This repo includes a workflow (.github/workflows/sync-labels.yml) that syncs labels from there.
