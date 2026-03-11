/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>

SPDX-License-Identifier: BSD-3-Clause
*/
package update_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/outscale/octl/pkg/update"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newTestdataServer(t *testing.T) *httptest.Server {
	srv := httptest.NewServer(http.FileServer(http.Dir("testdata/")))
	t.Cleanup(srv.Close)
	return srv
}

func installGHRedirect(t *testing.T, rel update.RepositoryRelease) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_ = json.NewEncoder(w).Encode(rel)
	}))

	real := update.GhURL
	update.GhURL = srv.URL

	t.Cleanup(func() {
		srv.Close()
		update.GhURL = real
	})
}

func TestUpdatePolicyCheck(t *testing.T) {
	t.Run("default policy is valid", func(t *testing.T) {
		// Neither flag set → signature + digest both required → valid
		err := (&update.UpdatePolicy{}).Check()
		assert.NoError(t, err)
	})

	t.Run("ignoreSignature only is valid", func(t *testing.T) {
		policy := &update.UpdatePolicy{}
		update.WithIgnoreSignature()(policy)
		assert.NoError(t, policy.Check())
	})

	t.Run("ignoreDigest only is invalid", func(t *testing.T) {
		policy := &update.UpdatePolicy{}
		update.WithIgnoreDigest()(policy)
		err := policy.Check()
		require.Error(t, err)
		assert.Contains(t, err.Error(), "cannot validate signature without digest")
	})

	t.Run("ignoreSignature + ignoreDigest is valid", func(t *testing.T) {
		policy := &update.UpdatePolicy{}
		update.WithIgnoreSignature()(policy)
		update.WithIgnoreDigest()(policy)
		assert.NoError(t, policy.Check())
	})
}

func TestFindAndCheckAssetDigest(t *testing.T) {
	const (
		assetName   = "octl_linux_x86_64"
		correctHash = "abc123def456"
	)

	goodAsset := update.ReleaseAsset{
		Name:   assetName,
		Digest: "sha256:" + correctHash,
	}
	checksums := map[string]string{
		assetName: correctHash,
	}

	t.Run("matching digest succeeds", func(t *testing.T) {
		digest, err := update.FindAndCheckAssetDigest(goodAsset, checksums)
		require.NoError(t, err)
		assert.Equal(t, correctHash, digest)
	})

	t.Run("asset not in checksums map", func(t *testing.T) {
		_, err := update.FindAndCheckAssetDigest(goodAsset, map[string]string{})
		require.Error(t, err)
		assert.Contains(t, err.Error(), "could not find digest for asset")
		assert.Contains(t, err.Error(), assetName)
	})

	t.Run("asset has empty digest field", func(t *testing.T) {
		noDigest := update.ReleaseAsset{Name: assetName, Digest: ""}
		_, err := update.FindAndCheckAssetDigest(noDigest, checksums)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "has no digest set, cannot verify")
	})

	t.Run("digest mismatch is caught", func(t *testing.T) {
		corrupted := update.ReleaseAsset{
			Name:   assetName,
			Digest: "sha256:000000000000", // wrong hash
		}
		_, err := update.FindAndCheckAssetDigest(corrupted, checksums)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "asset digest mismatch")
	})
}

func TestUpdateMissingRequiredAssets(t *testing.T) {
	for _, tc := range []struct {
		name    string
		assets  []update.ReleaseAsset
		options []update.UpdateOption
		wantErr string
	}{
		{
			name:    "no assets at all",
			assets:  nil,
			options: []update.UpdateOption{update.WithIgnoreSignature(), update.WithIgnoreDigest()},
			wantErr: "could not find required assets",
		},
		{
			name: "binary missing, checksums present",
			assets: []update.ReleaseAsset{
				{Name: "octl_checksums.txt", BrowserDownloadURL: "http://example.com/checksums"},
			},
			options: []update.UpdateOption{update.WithIgnoreSignature(), update.WithIgnoreDigest()},
			wantErr: "could not find required assets",
		},
		{
			name: "bundle missing, signature required",
			assets: []update.ReleaseAsset{
				{Name: "octl_linux_x86_64", BrowserDownloadURL: "http://example.com/binary"},
				{Name: "octl_checksums.txt", BrowserDownloadURL: "http://example.com/checksums"},
				// no .sigstore.json
			},
			options: []update.UpdateOption{}, // signature NOT ignored
			wantErr: "could not find required assets",
		},
		{
			name: "checksums missing, digest required",
			assets: []update.ReleaseAsset{
				{Name: "octl_linux_x86_64", BrowserDownloadURL: "http://example.com/binary"},
				// no _checksums.txt, but signature ignored
			},
			options: []update.UpdateOption{update.WithIgnoreSignature()},
			wantErr: "could not find required assets",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			rel := update.RepositoryRelease{
				TagName: "v99.0.0", // always newer than any real version
				Assets:  tc.assets,
			}

			installGHRedirect(t, rel)

			err := update.Update(t.Context(), tc.options...)
			require.Error(t, err)
			assert.Contains(t, err.Error(), tc.wantErr)
		})
	}
}

// testdataRelease builds a RepositoryRelease whose asset download URLs point at
// assetSrv (a file-serving httptest.Server rooted at testdata/).
// The digest values are the real sha256 digests from the v0.0.17 GitHub release.
func testdataRelease(assetSrvURL string) update.RepositoryRelease {
	base := assetSrvURL
	return update.RepositoryRelease{
		TagName: "v0.0.17",
		Assets: []update.ReleaseAsset{
			{
				Name:               "octl_Linux_x86_64",
				BrowserDownloadURL: base + "/octl_Linux_x86_64",
				Digest:             "sha256:37c4f1531af17efdbbe0621a519b34f4aab0b97ef8055c204ef05f667a8cca38",
			},
			{
				Name:               "octl_0.0.17_checksums.txt",
				BrowserDownloadURL: base + "/octl_0.0.17_checksums.txt",
				Digest:             "sha256:10ed869427aeed50ad221ec41b75e73bc5a5e7a0be6c033cb22117b7e6a188a8",
			},
			{
				Name:               "octl_0.0.17_checksums.txt.sigstore.json",
				BrowserDownloadURL: base + "/octl_0.0.17_checksums.txt.sigstore.json",
				Digest:             "sha256:22cf27a66e183fdab480192f18d7b972c207a2c1076473f03c3eda88475ae3bc",
			},
		},
	}
}

func TestUpdateRealReleaseMutation(t *testing.T) {
	for _, tc := range []struct {
		name    string
		options []update.UpdateOption
		mutate  func(srv string, rel *update.RepositoryRelease)
		wantErr string
	}{
		{
			name:    "happy path: full verification passes",
			options: []update.UpdateOption{update.WithDryRun()},
			mutate:  nil,
			wantErr: "",
		},
		{
			name:    "tampered bundle is rejected",
			options: []update.UpdateOption{update.WithDryRun()},
			mutate: func(srv string, rel *update.RepositoryRelease) {
				rel.Assets[2].BrowserDownloadURL = srv + "/octl_0.0.17_corrupted_checksums.txt.sigstore.json"
				rel.Assets[2].Digest = ""
			},
			wantErr: "could not verify signature",
		},
		{
			name:    "tampered bundle is accpted (ignore signature)",
			options: []update.UpdateOption{update.WithIgnoreSignature(), update.WithDryRun()},
			mutate: func(srv string, rel *update.RepositoryRelease) {
				rel.Assets[2].BrowserDownloadURL = srv + "/octl_0.0.17_corrupted_checksums.txt.sigstore.json"
				rel.Assets[2].Digest = ""
			},
			wantErr: "",
		},
		{
			name:    "tampered checksums file is rejected",
			options: []update.UpdateOption{update.WithDryRun()},
			mutate: func(srv string, rel *update.RepositoryRelease) {
				rel.Assets[1].BrowserDownloadURL = srv + "/octl_0.0.17_corrupted_checksums.txt"
				rel.Assets[1].Digest = ""
			},
			wantErr: "could not verify signature",
		},
		{
			name:    "tampered checksums file is accepted (ignore signature and digest)",
			options: []update.UpdateOption{update.WithIgnoreDigest(), update.WithIgnoreSignature(), update.WithDryRun()},
			mutate: func(srv string, rel *update.RepositoryRelease) {
				rel.Assets[1].BrowserDownloadURL = srv + "/octl_0.0.17_corrupted_checksums.txt"
				rel.Assets[1].Digest = ""
			},
			wantErr: "",
		},
		{
			name:    "tampered release checksum is rejected",
			options: []update.UpdateOption{update.WithDryRun()},
			mutate: func(srv string, rel *update.RepositoryRelease) {
				rel.Assets[0].Digest = "sha256:wronghash"
			},
			wantErr: "asset digest mismatch",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			assetSrv := newTestdataServer(t)
			rel := testdataRelease(assetSrv.URL)
			if tc.mutate != nil {
				tc.mutate(assetSrv.URL, &rel)
			}
			installGHRedirect(t, rel)

			err := update.Update(t.Context(), tc.options...)
			if tc.wantErr == "" {
				assert.Equal(t, nil, err)
			} else {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tc.wantErr)
			}
		})
	}
}
