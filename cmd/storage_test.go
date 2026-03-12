/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>

SPDX-License-Identifier: BSD-3-Clause
*/
package cmd_test

import (
	"crypto/sha1"
	"encoding/hex"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
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
