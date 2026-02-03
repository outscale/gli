/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>

SPDX-License-Identifier: BSD-3-Clause
*/
package update

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"runtime"
	"strings"

	"github.com/google/go-github/v82/github"
	"github.com/minio/selfupdate"
	"github.com/outscale/gli/pkg/debug"
)

func latestRelease(ctx context.Context) (*github.RepositoryRelease, error) {
	client := github.NewClient(nil)
	rel, _, err := client.Repositories.GetLatestRelease(ctx, "outscale", "gli")
	if err != nil {
		return nil, fmt.Errorf("github: %w", err)
	}
	if rel == nil || rel.TagName == nil {
		return nil, nil
	}
	return rel, nil
}

func LatestRelease(ctx context.Context) (string, error) {
	rel, err := latestRelease(ctx)
	if rel == nil || err != nil {
		return "", err
	}
	return *rel.TagName, nil
}

func Update(ctx context.Context) error {
	rel, err := latestRelease(ctx)
	if err != nil {
		return err
	}
	if rel == nil {
		return errors.New("no new version found")
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
		if a.Name == nil || a.BrowserDownloadURL == nil {
			continue
		}
		if strings.HasSuffix(strings.ToLower(*a.Name), suffix) {
			debug.Println("found", a.GetName(), a.GetBrowserDownloadURL())
			return update(ctx, rel.GetTagName(), a)
		}
	}
	return nil
}

func update(ctx context.Context, v string, a *github.ReleaseAsset) error {
	fmt.Println("⬇️ Downloading file", a.GetName())
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, a.GetBrowserDownloadURL(), nil)
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
	return nil
}
