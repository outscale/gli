/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>

SPDX-License-Identifier: BSD-3-Clause
*/
package update

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
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

func Update(ctx context.Context) error {
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
	for _, a := range rel.Assets {
		if a.Name == "" || a.BrowserDownloadURL == "" {
			continue
		}
		if strings.HasSuffix(strings.ToLower(a.Name), suffix) {
			debug.Println("found", rel.TagName, a.BrowserDownloadURL)
			return update(ctx, rel.TagName, a, changelog(version.Version, rel.TagName, rel.Body))
		}
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

func update(ctx context.Context, v string, a ReleaseAsset, changelog string) error {
	fmt.Println("⬇️ Downloading file", a.Name)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, a.BrowserDownloadURL, nil)
	if err != nil {
		return fmt.Errorf("http get: %w", err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("http get: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("http get: invalid status code %d", resp.StatusCode)
	}
	defer resp.Body.Close() //nolint
	fmt.Println("📦 Updating binary to", v)
	err = selfupdate.Apply(resp.Body, selfupdate.Options{})
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
