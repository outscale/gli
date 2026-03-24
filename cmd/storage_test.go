/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>

SPDX-License-Identifier: BSD-3-Clause
*/
package cmd_test

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/outscale/goutils/sdk/ptr"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStorageCRUD(t *testing.T) {
	sum := sha1.Sum([]byte(t.TempDir()))
	bucket := hex.EncodeToString(sum[:])

	object := "object.txt"
	path := filepath.Join(t.TempDir(), object)
	err := os.WriteFile(path, []byte("Hello world !"), 0o600)
	require.NoError(t, err)
	t.Run("Create/Update/Delete works", func(t *testing.T) {
		_ = run(t, []string{"storage", "bucket", "create", "--bucket", bucket}, nil)
		defer func() {
			_ = run(t, []string{"storage", "bucket", "del", bucket, "-y"}, nil)
		}()

		var res s3.PutObjectOutput
		runJSON(t, []string{"storage", "object", "put", object, "--bucket", bucket, "--body", path, "--output", "json"}, nil, &res)
		defer func() {
			_ = run(t, []string{"storage", "object", "del", object, "--bucket", bucket, "-y"}, nil)
		}()
		assert.NotNil(t, res.ETag)

		var lres s3.ListObjectsV2Output
		runJSON(t, []string{"storage", "object", "list", "--bucket", bucket, "-o", "raw"}, nil, &lres)
		require.Len(t, lres.Contents, 1)
		require.NotNil(t, object, lres.Contents[0].Key)
		assert.Equal(t, object, *lres.Contents[0].Key)
	})
}

const hello = "Hello world !"

func TestObjectDownload(t *testing.T) {
	sum := sha1.Sum([]byte(t.TempDir()))
	bucket := hex.EncodeToString(sum[:])

	_ = run(t, []string{"storage", "bucket", "create", "--bucket", bucket}, nil)
	defer func() {
		_ = run(t, []string{"storage", "bucket", "del", bucket, "-y"}, nil)
	}()

	t.Run("An small object can be downloaded", func(t *testing.T) {
		payload := hello
		small := "small.txt"
		smallPath := filepath.Join(t.TempDir(), small)
		err := os.WriteFile(smallPath, []byte(hello), 0o600)
		require.NoError(t, err)
		_ = run(t, []string{"storage", "object", "put", small, "--bucket", bucket, "--body", smallPath, "--output", "json"}, nil)
		defer func() {
			_ = run(t, []string{"storage", "object", "del", small, "--bucket", bucket, "-y"}, nil)
		}()

		content := run(t, []string{"storage", "object", "download", small, "--bucket", bucket}, nil)
		assert.Equal(t, payload, string(content))
		_ = run(t, []string{"storage", "object", "download", small, "--bucket", bucket, "-O", smallPath}, nil)
		content, err = os.ReadFile(smallPath)
		require.NoError(t, err)
		assert.Equal(t, payload, string(content))
	})
	t.Run("An large object can be downloaded", func(t *testing.T) {
		payload := strings.Repeat(hello, 100)
		large := "large.txt"
		largePath := filepath.Join(t.TempDir(), large)
		err := os.WriteFile(largePath, []byte(payload), 0o600)
		require.NoError(t, err)

		_ = run(t, []string{"storage", "object", "put", large, "--bucket", bucket, "--body", largePath, "--output", "json"}, nil)
		defer func() {
			_ = run(t, []string{"storage", "object", "del", large, "--bucket", bucket, "-y"}, nil)
		}()

		content := run(t, []string{"storage", "object", "download", large, "--bucket", bucket}, nil)
		assert.Equal(t, payload, string(content))
		_ = run(t, []string{"storage", "object", "download", large, "--bucket", bucket, "-O", largePath}, nil)
		content, err = os.ReadFile(largePath)
		require.NoError(t, err)
		assert.Equal(t, payload, string(content))
	})
}

func TestBucketVersioning(t *testing.T) {
	sum := sha1.Sum([]byte(t.TempDir()))
	bucket := hex.EncodeToString(sum[:])

	_ = run(t, []string{"storage", "bucket", "create", "--bucket", bucket}, nil)
	defer func() {
		_ = run(t, []string{"storage", "bucket", "del", bucket, "-y"}, nil)
	}()

	t.Run("Versioning can be enabled and disabled", func(t *testing.T) {
		var resp s3.GetBucketVersioningOutput

		runJSON(t, []string{"storage", "bucket", "versioning", "describe", bucket, "-o", "json"}, nil, &resp)
		assert.Equal(t, types.BucketVersioningStatus(""), resp.Status)

		_ = run(t, []string{"storage", "bucket", "versioning", "enable", bucket}, nil)

		runJSON(t, []string{"storage", "bucket", "versioning", "describe", bucket, "-o", "json"}, nil, &resp)
		assert.Equal(t, types.BucketVersioningStatusEnabled, resp.Status)

		_ = run(t, []string{"storage", "bucket", "versioning", "disable", bucket}, nil)

		runJSON(t, []string{"storage", "bucket", "versioning", "describe", bucket, "-o", "json"}, nil, &resp)
		assert.Equal(t, types.BucketVersioningStatusSuspended, resp.Status)
	})
}

func TestBucketEncryption(t *testing.T) {
	sum := sha1.Sum([]byte(t.TempDir()))
	bucket := hex.EncodeToString(sum[:])

	_ = run(t, []string{"storage", "bucket", "create", "--bucket", bucket}, nil)
	defer func() {
		_ = run(t, []string{"storage", "bucket", "del", bucket, "-y"}, nil)
	}()

	t.Run("Encryption can be enabled and disabled", func(t *testing.T) {
		_ = run(t, []string{"storage", "bucket", "encryption", "enable", bucket}, nil)

		var resp s3.GetBucketEncryptionOutput
		runJSON(t, []string{"storage", "bucket", "encryption", "describe", bucket, "-o", "json"}, nil, &resp)
		require.Len(t, resp.ServerSideEncryptionConfiguration.Rules, 1)
		assert.Equal(t, types.ServerSideEncryptionAes256, resp.ServerSideEncryptionConfiguration.Rules[0].ApplyServerSideEncryptionByDefault.SSEAlgorithm)

		_ = run(t, []string{"storage", "bucket", "encryption", "disable", bucket}, nil)
	})
}

func TestBucketCORS(t *testing.T) {
	sum := sha1.Sum([]byte(t.TempDir()))
	bucket := hex.EncodeToString(sum[:])

	_ = run(t, []string{"storage", "bucket", "create", "--bucket", bucket}, nil)
	defer func() {
		_ = run(t, []string{"storage", "bucket", "del", bucket, "-y"}, nil)
	}()

	t.Run("CORS can be enabled and disabled", func(t *testing.T) {
		runWithError(t, []string{"storage", "bucket", "cors", "describe", bucket, "-o", "json"}, nil)

		_ = run(t, []string{"storage", "bucket", "cors", "configure", bucket, "--from-file", "testdata/storage/cors.json"}, nil)

		var resp s3.GetBucketCorsOutput
		runJSON(t, []string{"storage", "bucket", "cors", "describe", bucket, "-o", "json"}, nil, &resp)
		assert.NotEmpty(t, resp.CORSRules)

		_ = run(t, []string{"storage", "bucket", "cors", "disable", bucket}, nil)

		runWithError(t, []string{"storage", "bucket", "cors", "describe", bucket, "-o", "json"}, nil)
	})
}

func TestBucketPolicy(t *testing.T) {
	sum := sha1.Sum([]byte(t.TempDir()))
	bucket := hex.EncodeToString(sum[:])

	_ = run(t, []string{"storage", "bucket", "create", "--bucket", bucket}, nil)
	defer func() {
		_ = run(t, []string{"storage", "bucket", "del", bucket, "-y"}, nil)
	}()

	file := filepath.Join(t.TempDir(), "policy.json")
	data := `{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Deny",
            "Principal": "*",
            "Action": "s3:PutObject",
            "Resource": "arn:aws:s3:::` + bucket + `/forbidden/*"
        }
    ]
}
`
	err := os.WriteFile(file, []byte(data), 0o600)
	require.NoError(t, err)

	t.Run("A policy can be configured and removed", func(t *testing.T) {
		runWithError(t, []string{"storage", "bucket", "policy", "describe", bucket, "-o", "json"}, nil)

		_ = run(t, []string{"storage", "bucket", "policy", "configure", bucket, "--from-file", file}, nil)

		resp := run(t, []string{"storage", "bucket", "policy", "describe", bucket, "-o", "json"}, nil)
		assert.NotEmpty(t, resp)

		_ = run(t, []string{"storage", "bucket", "policy", "disable", bucket}, nil)

		runWithError(t, []string{"storage", "bucket", "policy", "describe", bucket, "-o", "json"}, nil)
	})
}

func TestBucketWebsite(t *testing.T) {
	sum := sha1.Sum([]byte(t.TempDir()))
	bucket := hex.EncodeToString(sum[:])

	_ = run(t, []string{"storage", "bucket", "create", "--bucket", bucket}, nil)
	defer func() {
		_ = run(t, []string{"storage", "bucket", "del", bucket, "-y"}, nil)
	}()

	t.Run("A website can be configured and disabled", func(t *testing.T) {
		runWithError(t, []string{"storage", "bucket", "website", "describe", bucket, "-o", "json"}, nil)

		_ = run(t, []string{"storage", "bucket", "website", "configure", bucket, "--from-file", "testdata/storage/website.json"}, nil)

		var resp s3.GetBucketWebsiteOutput
		runJSON(t, []string{"storage", "bucket", "website", "describe", bucket, "-o", "json"}, nil, &resp)
		assert.NotNil(t, resp.IndexDocument)
		assert.NotNil(t, resp.IndexDocument.Suffix)
		assert.Equal(t, "index.html", *resp.IndexDocument.Suffix)

		_ = run(t, []string{"storage", "bucket", "website", "disable", bucket}, nil)

		runWithError(t, []string{"storage", "bucket", "website", "describe", bucket, "-o", "json"}, nil)
	})
}

func TestBucketLifecycle(t *testing.T) {
	sum := sha1.Sum([]byte(t.TempDir()))
	bucket := hex.EncodeToString(sum[:])

	_ = run(t, []string{"storage", "bucket", "create", "--bucket", bucket}, nil)
	defer func() {
		_ = run(t, []string{"storage", "bucket", "del", bucket, "-y"}, nil)
	}()

	t.Run("Lifecycle can be configured and disabled", func(t *testing.T) {
		runWithError(t, []string{"storage", "bucket", "lifecycle", "describe", bucket, "-o", "json"}, nil)

		_ = run(t, []string{"storage", "bucket", "lifecycle", "configure", bucket, "--from-file", "testdata/storage/lifecycle.json"}, nil)

		var resp s3.GetBucketLifecycleConfigurationOutput
		runJSON(t, []string{"storage", "bucket", "lifecycle", "describe", bucket, "-o", "json"}, nil, &resp)
		assert.NotEmpty(t, resp.Rules)

		_ = run(t, []string{"storage", "bucket", "lifecycle", "disable", bucket}, nil)

		runWithError(t, []string{"storage", "bucket", "lifecycle", "describe", bucket, "-o", "json"}, nil)
	})
}

func TestBucketACL(t *testing.T) {
	sum := sha1.Sum([]byte(t.TempDir()))
	bucket := hex.EncodeToString(sum[:])

	_ = run(t, []string{"storage", "bucket", "create", "--bucket", bucket}, nil)
	defer func() {
		_ = run(t, []string{"storage", "bucket", "del", bucket, "-y"}, nil)
	}()

	t.Run("ACL can be added", func(t *testing.T) {
		var resp s3.GetBucketAclOutput
		runJSON(t, []string{"storage", "bucket", "acl", "describe", bucket, "-o", "json"}, nil, &resp)
		assert.NotEmpty(t, resp.Grants)
		before := len(resp.Grants)

		resp.Grants = append(resp.Grants, types.Grant{
			Grantee: &types.Grantee{
				Type: types.TypeGroup,
				URI:  ptr.To("http://acs.amazonaws.com/groups/global/AllUsers"),
			},
			Permission: types.PermissionRead,
		})
		file := filepath.Join(t.TempDir(), "acl.json")
		fd, err := os.Create(file)
		require.NoError(t, err)
		err = json.NewEncoder(fd).Encode(resp)
		require.NoError(t, err)

		_ = run(t, []string{"storage", "bucket", "acl", "configure", bucket, "--from-file", file}, nil)

		runJSON(t, []string{"storage", "bucket", "acl", "describe", bucket, "-o", "json"}, nil, &resp)
		assert.Len(t, resp.Grants, before+1)
	})
}
