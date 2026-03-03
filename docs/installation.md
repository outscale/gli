# Installation

Download the latest binary from:
- <https://github.com/outscale/octl/releases>

## Signature Verification

All release binaries are cryptographically signed using Sigstore.

- For verification details, see [signature-verification.md](signature-verification.md)

## Shell completion

### Bash

```sh
octl completion bash > octl-completion.bash
sudo cp octl-completion.bash /etc/bash_completion.d/
source ~/.bashrc
```

### Zsh

```zsh
octl completion zsh > _octl
sudo mkdir -p /usr/local/share/zsh/site-functions
sudo cp _octl /usr/local/share/zsh/site-functions/
source ~/.zshrc
```

### Fish

```fish
octl completion fish > octl.fish
sudo cp octl.fish /usr/share/fish/completions/
source ~/.config/fish/config.fish
```

### Self-updating
```sh
octl update
```

> This requires write access to the binary. If `octl update` does not work, download the binary from the [latest release](https://github.com/outscale/octl/releases/latest).
