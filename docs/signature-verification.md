# Signature Verification

All `octl` release binaries are signed using [Sigstore](https://www.sigstore.dev/) and can be verified cryptographically. This ensures that the binaries you download are authentic and have not been tampered with.

---

## Table of Contents

- [Overview](#overview)
- [Automatic Verification (Self-Update)](#automatic-verification-self-update)
- [Manual Verification](#manual-verification)
- [Installing Sigstore Tools](#installing-sigstore-tools)

---

## Overview

When you download `octl` from the [GitHub Releases page](https://github.com/outscale/octl/releases), each binary comes with:

- A **signature bundle** (`.sigstore.json`) containing the cryptographic signature and verification metadata
- A **checksums file** (`.checksums.txt`) with SHA-256 hashes of all release artifacts

The signature uses the Sigstore infrastructure:
- **Fulcio** for certificate issuance (OIDC-based short-lived certificates)
- **Rekor** for transparency log entries
- **GitHub Actions OIDC** for identity verification

---

## Automatic Verification (Self-Update)

When using the built-in update command:

```sh
octl update
```

The signature verification is performed **automatically**:

1. Downloads the signature bundle from the GitHub release
2. Verifies the certificate identity through Sigstore
3. Validates the checksum file against the signature
4. Verifies the binary against the checksum before applying the update

If any verification fails, the update is aborted and an error message is displayed.

**Note:** This requires network access to Sigstore infrastructure (Sigstore TUF repository, Fulcio, and Rekor).

---

## Manual Verification

If you download the binary manually, you can verify the signature using the `cosign` CLI tool.

### Step 1: Download Required Files

Download all three files for your platform from the [releases page](https://github.com/outscale/octl/releases):

- `octl_<version>_<os>_<arch>.tar.gz` (or `.zip` for Windows)
- `octl_<version>_<os>_<arch>_checksums.txt`
- `octl_<version>_checksums.txt.sigstore.json`

Example for Linux x86_64:

```sh
VERSION="v0.1.0"
ASSET="octl_${VERSION}_linux_x86_64.tar.gz"
curl -LO "https://github.com/outscale/octl/releases/download/${VERSION}/${ASSET}"
curl -LO "https://github.com/outscale/octl/releases/download/${VERSION}/${VERSION}_checksums.txt"
curl -LO "https://github.com/outscale/octl/releases/download/${VERSION}/${VERSION}_checksums.txt.sigstore.json"
```

### Step 2: Verify Using Cosign

The signature bundle can be verified using Cosign v2.0 or later:

```sh
# Verify the checksum file against the Sigstore bundle
cosign verify-blob \
  --bundle <bundle-file>.sigstore.json \
  --certificate-oidc-issuer "https://token.actions.githubusercontent.com" \
  --certificate-identity-regexp "^https://github.com/outscale/octl/" \
  <checksums-file>.txt
```

Expected output:

```
Verified OK
tlog entry index: <number>
```

### Step 3: Verify Binary Against Checksums

After verifying the signature on the checksums file, you can validate the binary:

```sh
sha256sum -c <checksums-file>.txt --ignore-missing
```

This ensures the binary hash matches the signed checksums file.

### Certificate Identity Details

When verifying, the certificate identity follows these patterns:

| Field | Value |
|-------|-------|
| Issuer | `https://token.actions.githubusercontent.com` |
| Subject | Starts with `https://github.com/outscale/octl/` |

This confirms the binary was built by the official `outscale/octl` GitHub repository's CI/CD pipeline.

---

## Installing Sigstore Tools

### Cosign

Cosign is the recommended tool for verifying Sigstore signatures:

```sh
# macOS with Homebrew
brew install cosign

# Linux (via package manager)
# See: https://docs.sigstore.dev/cosign/system_config/installation/

# Download binary directly
curl -O -L "https://github.com/sigstore/cosign/releases/latest/download/cosign-$(uname -m)-pc-linux-gnu"
chmod +x cosign-$(uname -m)-pc-linux-gnu
sudo mv cosign-$(uname -m)-pc-linux-gnu /usr/local/bin/cosign
```

For more options, see the [Cosign installation guide](https://docs.sigstore.dev/cosign/system_config/installation/).

---

## Further Reading

- [Sigstore Documentation](https://docs.sigstore.dev/)
- [Cosign Documentation](https://docs.sigstore.dev/cosign/overview/)
- [Sigstore-go Library](https://github.com/sigstore/sigstore-go) (used internally by `octl`)
- [GitHub OIDC Token](https://docs.github.com/en/actions/security-for-github-actions/security-hardening-your-deployments/about-security-hardening-with-openid-connect)
