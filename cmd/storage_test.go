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
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStorageCRUD(t *testing.T) {
	sum := sha1.Sum([]byte(t.TempDir()))
	bucket := hex.EncodeToString(sum[:])

	object := "object.txt"
	path := filepath.Join(t.TempDir(), object)
	err := os.WriteFile(path, []byte("Hello world !"), 0600)
	require.NoError(t, err)
	t.Run("Create/Update/Delete works", func(t *testing.T) {
		_ = run(t, []string{"storage", "bucket", "create", "--bucket", bucket}, nil)

		var res s3.PutObjectOutput
		runJSON(t, []string{"storage", "object", "put", object, "--bucket", bucket, "--body", path, "--output", "json"}, nil, &res)
		assert.NotNil(t, res.ETag)

		var lres s3.ListObjectsV2Output
		runJSON(t, []string{"storage", "object", "list", "--bucket", bucket, "-o", "raw"}, nil, &lres)
		require.Len(t, lres.Contents, 1)
		require.NotNil(t, object, lres.Contents[0].Key)
		assert.Equal(t, object, *lres.Contents[0].Key)

		_ = run(t, []string{"storage", "object", "del", object, "--bucket", bucket, "-y"}, nil)
		_ = run(t, []string{"storage", "bucket", "del", bucket, "-y"}, nil)
	})
}
