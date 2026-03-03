## **OCTL**

[![Project Stage](https://docs.outscale.com/fr/userguide/_images/Project-Sandbox-yellow.svg)](https://docs.outscale.com/en/userguide/Open-Source-Projects.html)
[![](https://dcbadge.limes.pink/api/server/HUVtY5gT6s?style=flat&theme=default-inverted)](https://discord.gg/HUVtY5gT6s)

---

## 🌐 Links

- Documentation (this repo): [docs/README.md](docs/README.md)
- OUTSCALE documentation: [docs](https://docs.outscale.com/en/userguide/Home.html)
- Releases: <https://github.com/outscale/octl/releases>
- Join our community on [Discord](https://discord.gg/HUVtY5gT6s)

---

## 📄 Table of Contents

- [🌐 Links](#-links)
- [🧭 Overview](#-overview)
- [✅ Requirements](#-requirements)
- [⚙ Installation](#-installation)
- [🛠 Configuration](#-configuration)
- [🚀 Quickstart](#-quickstart)
- [📚 Documentation](#-documentation)
- [📜 License](#-license)
- [🤝 Contributing](#-contributing)

---

## 🧭 Overview

**octl** is a modern CLI for the Outscale APIs, written in Go.

It supports:
- installation via a single static binary,
- direct flags to all request fields (no JSON required),
- autocompletion support for all API calls, flags, and flag values,
- jq-style output filters,
- command chaining,
- syntax highlighting of output,
- auto-update to the latest version.

It manages:
- IaaS resources (nets, vms, ...),
- OOS storage (buckets, objects).

And includes a preliminary support for OKS Kubernetes clusters.

---

## ✅ Requirements

- Access to the OUTSCALE API (with appropriate credentials)

---

## ⚙ Installation

Download the latest binary from the [Releases page](https://github.com/outscale/octl/releases).

Autocompletion setup is documented here:
- [docs/installation.md](docs/installation.md)

---

## 🛠 Configuration

Configuration can be provided either via environment variables or a profile file:
- [docs/configuration.md](docs/configuration.md)

Quick example (env vars):

```bash
export OSC_ACCESS_KEY="..."
export OSC_SECRET_KEY="..."
export OSC_REGION="eu-west-2"
```

## 🚀 Quickstart

List volumes using the high-level command:
```bash
octl iaas volume list
```
Call an API operation directly:
```bash
octl iaas api ReadVms --Filters.VmStateNames running
```
More examples:

- [docs/usage/overview.md](./docs/usage/overview.md)
- [docs/usage/iaas.md](./docs/usage/iaas.md)

## 📚 Documentation

- Start here: [docs/README.md](./docs/README.md)
- Installation + completion: [docs/installation.md](./docs/installation.md)
- Configuration: [docs/configuration.md](./docs/configuration.md)
- Usage overview: [docs/usage/overview.md](./docs/usage/overview.md)
- Signature verification: [docs/signature-verification.md](./docs/signature-verification.md)
- Troubleshooting: [docs/troubleshooting.md](./docs/troubleshooting.md)
- Command reference: [docs/reference/commands.md](./docs/reference/commands.md)

## 📜 License

octl is released under the BSD 3-Clause license.

© 2026 Outscale SAS

See [LICENSE](./LICENSE) for full details.

## 🤝 Contributing

We welcome contributions!

Please read our [Contributing Guidelines](./CONTRIBUTING.md) and [Code of Conduct](./CODE_OF_CONDUCT.md) before submitting a pull request.
