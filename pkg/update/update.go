/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>

SPDX-License-Identifier: BSD-3-Clause
*/
package update

import (
	"bufio"
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"runtime"
	"strings"
	"time"

	"github.com/minio/selfupdate"
	"github.com/outscale/octl/pkg/debug"
	"github.com/outscale/octl/pkg/markdown"
	"github.com/outscale/octl/pkg/messages"
	"github.com/outscale/octl/pkg/version"
	"github.com/sigstore/sigstore-go/pkg/bundle"
	"github.com/sigstore/sigstore-go/pkg/root"
	"github.com/sigstore/sigstore-go/pkg/tuf"
	"github.com/sigstore/sigstore-go/pkg/verify"
	"golang.org/x/mod/semver"
)

// RepositoryRelease represents a GitHub release in a repository.
type RepositoryRelease struct {
	TagName string         `json:"tag_name,omitempty"`
	Body    string         `json:"body,omitempty"`
	Assets  []ReleaseAsset `json:"assets,omitempty"`
}

// ReleaseAsset represents a GitHub release asset in a repository.
type ReleaseAsset struct {
	Name               string `json:"name,omitempty"`
	BrowserDownloadURL string `json:"browser_download_url,omitempty"`
	Digest             string `json:"digest,omitempty"`
}

const ghURL = "https://api.github.com/repos/outscale/octl/releases/latest"

func latestRelease(ctx context.Context) *RepositoryRelease {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, ghURL, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		debug.Println("github error", err)
		return nil
	}
	defer resp.Body.Close() //nolint
	if resp.StatusCode != http.StatusOK {
		debug.Println("github error", resp.Status)
		return nil
	}
	rel := &RepositoryRelease{}
	err = json.NewDecoder(resp.Body).Decode(rel)
	if err != nil {
		debug.Println("github decoding error", err)
		return nil
	}
	if rel.TagName == "" {
		return nil
	}
	return rel
}

func LatestRelease(ctx context.Context) string {
	rel := latestRelease(ctx)
	if rel == nil {
		return ""
	}
	return rel.TagName
}

type UpdatePolicy struct {
	ignoreSignature bool
	ignoreDigest    bool
}

type UpdateOption func(*UpdatePolicy)

// WithIgnoreSignature creates an option to ignore signature verification during update.
func WithIgnoreSignature() UpdateOption {
	return func(p *UpdatePolicy) {
		p.ignoreSignature = true
	}
}

// WithIgnoreDigest creates
func WithIgnoreDigest() UpdateOption {
	return func(p *UpdatePolicy) {
		p.ignoreDigest = true
	}
}

func (up *UpdatePolicy) Check() error {
	if !up.ignoreSignature && up.ignoreDigest {
		return errors.New("cannot validate signature without digest")
	}
	return nil
}

func Update(ctx context.Context, options ...UpdateOption) error {
	var policy UpdatePolicy
	for _, o := range options {
		o(&policy)
	}

	if err := policy.Check(); err != nil {
		return err
	}

	rel := latestRelease(ctx)
	if rel == nil {
		return errors.New("no new version found")
	}
	if semver.Compare(version.Version, rel.TagName) >= 0 {
		messages.Warn("Already using the latest version")
		return nil
	}

	debug.Println("OS, arch:", runtime.GOOS, runtime.GOARCH)
	suffix := runtime.GOOS + "_"
	switch runtime.GOARCH {
	case "amd64":
		suffix += "x86_64"
	default:
		suffix += runtime.GOARCH
	}
	if runtime.GOOS == "windows" {
		suffix += ".exe"
	}

	var assetToDownload, checksums, bundleAsset *ReleaseAsset
	for _, a := range rel.Assets {
		if a.Name == "" || a.BrowserDownloadURL == "" {
			continue
		}
		if strings.HasSuffix(strings.ToLower(a.Name), suffix) {
			debug.Println("found asset", a.Name, a.BrowserDownloadURL)
			assetToDownload = &a
		} else if strings.HasSuffix(strings.ToLower(a.Name), ".sigstore.json") {
			debug.Println("found bundle", a.Name, a.BrowserDownloadURL)
			bundleAsset = &a
		} else if strings.HasSuffix(strings.ToLower(a.Name), "_checksums.txt") {
			debug.Println("found checksums", a.Name, a.BrowserDownloadURL)
			checksums = &a
		}
	}

	if (bundleAsset == nil && !policy.ignoreSignature) || (checksums == nil && !policy.ignoreDigest) || assetToDownload == nil {
		return fmt.Errorf("could not find required assets in release %s", rel.TagName)
	}

	var bundle *bundle.Bundle
	if !policy.ignoreSignature {

		b, err := downloadBundle(ctx, *bundleAsset)
		if err != nil {
			return err
		}

		bundle = b
	} else {
		fmt.Println("⚠️ Skipping signature verification")
	}

	var digest string
	if !policy.ignoreDigest {

		cs, err := downloadAndVerifyChecksum(ctx, *checksums, bundle)
		if err != nil {
			return err
		}

		digest, err = findAndCheckAssetDigest(*assetToDownload, cs)
		if err != nil {
			return err
		}
	} else {
		fmt.Println("⚠️ Skipping digest verification")
	}

	if err := update(ctx, rel.TagName, *assetToDownload, digest, changelog(version.Version, rel.TagName, rel.Body)); err != nil {
		return err
	}

	return nil
}

var reFullChangelog = regexp.MustCompile(`\**Full Changelog.*`)

func changelog(from, to string, body string) string {
	full := fmt.Sprintf("**Full Changelog**: https://github.com/outscale/octl/compare/%s...%s", from, to)
	if body == "" {
		return full
	}
	txt := reFullChangelog.ReplaceAllString(body, "")
	return txt + full
}

func findAndCheckAssetDigest(a ReleaseAsset, cs map[string]string) (string, error) {
	digest, ok := cs[a.Name]
	if !ok {
		return "", fmt.Errorf("could not find digest for asset %s in checksums", a.Name)
	}

	if a.Digest == "" {
		return "", fmt.Errorf("asset %s has no digest set, cannot verify", a.Name)
	}

	if !strings.HasSuffix(a.Digest, digest) {
		return "", fmt.Errorf("asset digest mismatch: expected %s, got %s", digest, a.Digest)
	}

	return digest, nil
}

func downloadFile(ctx context.Context, a ReleaseAsset) (io.ReadCloser, error) {
	fmt.Println("⬇️ Downloading file", a.Name)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, a.BrowserDownloadURL, nil)
	if err != nil {
		return nil, fmt.Errorf("http get: %w", err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http get: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close() //nolint
		return nil, fmt.Errorf("http get: invalid status code %d", resp.StatusCode)
	}
	return resp.Body, nil
}

func update(ctx context.Context, v string, a ReleaseAsset, digest, changelog string) error {
	resp, err := downloadFile(ctx, a)
	if err != nil {
		return err
	}

	defer resp.Close() //nolint
	fmt.Println("📦 Updating binary to", v)

	var opts selfupdate.Options
	if digest != "" {
		h, err := hex.DecodeString(digest)
		if err != nil {
			return fmt.Errorf("invalid digest: %w", err)
		}

		opts.Checksum = h
	}
	err = selfupdate.Apply(resp, opts)
	if err != nil {
		return fmt.Errorf("apply update: %w", err)
	}
	fmt.Println("✅ Done")
	if changelog != "" {
		fmt.Println("📝 Changes ⤵")
		md := markdown.NewRenderer()
		out, err := md.Render(changelog)
		if err == nil {
			fmt.Println(out)
		} else {
			fmt.Println(changelog)
		}
	}
	return nil
}

func downloadBundle(ctx context.Context, a ReleaseAsset) (*bundle.Bundle, error) {
	resp, err := downloadFile(ctx, a)
	if err != nil {
		return nil, err
	}
	defer resp.Close() //nolint

	buf, err := io.ReadAll(resp)
	if err != nil {
		return nil, fmt.Errorf("could not read bundle: %w", err)
	}

	var bundle bundle.Bundle
	if err := bundle.UnmarshalJSON(buf); err != nil {
		return nil, fmt.Errorf("could not unmarshal bundle: %w", err)
	}

	return &bundle, nil
}

func downloadAndVerifyChecksum(ctx context.Context, a ReleaseAsset, b *bundle.Bundle) (map[string]string, error) {
	checksums := make(map[string]string)
	resp, err := downloadFile(ctx, a)
	if err != nil {
		return checksums, err
	}
	defer resp.Close() //nolint

	buf, err := io.ReadAll(resp)
	if err != nil {
		return checksums, fmt.Errorf("could not read checksums: %w", err)
	}
	if b != nil {
		if err := checkBundle(bytes.NewBuffer(buf), b); err != nil {
			return checksums, fmt.Errorf("could not verify singuature: %w", err)
		}
	}
	scanner := bufio.NewScanner(bytes.NewBuffer(buf))
	for scanner.Scan() {
		line := scanner.Text()
		hash, fileName, found := strings.Cut(line, "  ")
		if !found {
			continue
		}

		checksums[fileName] = hash
	}
	if err := scanner.Err(); err != nil {
		return checksums, fmt.Errorf("could not scan checksums: %w", err)
	}

	return checksums, nil
}

func checkBundle(a *bytes.Buffer, b *bundle.Bundle) error {
	opts := tuf.DefaultOptions()
	client, err := tuf.New(opts)
	if err != nil {
		return err
	}

	trustedMaterial, err := root.GetTrustedRoot(client)
	if err != nil {
		return err
	}

	sev, err := verify.NewVerifier(trustedMaterial,
		verify.WithSignedCertificateTimestamps(1),
		verify.WithTransparencyLog(1),
		verify.WithObserverTimestamps(1))
	if err != nil {
		return err
	}

	certID, err := verify.NewShortCertificateIdentity(
		"https://token.actions.githubusercontent.com",
		"",
		"",
		"^https://github.com/outscale/octl/",
	)
	if err != nil {
		return err
	}

	_, err = sev.Verify(b, verify.NewPolicy(verify.WithArtifact(a), verify.WithCertificateIdentity(certID)))
	if err != nil {
		return err
	}

	fmt.Println("🔐 Signature verified successfully")
	return nil
}
