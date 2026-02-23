/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>

SPDX-License-Identifier: BSD-3-Clause
*/
package cmd_test

import (
	"path/filepath"
	"testing"

	"github.com/outscale/osc-sdk-go/v3/pkg/profile"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProfileAdd(t *testing.T) {
	t.Run("A profile can be added to a new file", func(t *testing.T) {
		file := filepath.Join(t.TempDir(), "add.json")
		_ = run(t, []string{"profile", "add", "add", "--ak", "foo", "--sk", "bar", "--config", file}, nil)

		cf, err := profile.LoadConfigFile(file)
		require.NoError(t, err)
		require.Len(t, cf.Profiles, 1)
		assert.Equal(t, "foo", cf.Profiles["add"].AccessKey)
		assert.Equal(t, "bar", cf.Profiles["add"].SecretKey)
	})
	t.Run("A profile can be added to an existing file", func(t *testing.T) {
		file := filepath.Join(t.TempDir(), "add.json")
		cf := &profile.ConfigFile{
			Path: file,
			Profiles: map[string]profile.Profile{
				"bar": {},
			},
		}
		err := cf.Save()
		require.NoError(t, err)
		_ = run(t, []string{"profile", "add", "add", "--ak", "foo", "--sk", "bar", "--config", file}, nil)

		cf, err = profile.LoadConfigFile(file)
		require.NoError(t, err)
		require.Len(t, cf.Profiles, 2)
		assert.Equal(t, "foo", cf.Profiles["add"].AccessKey)
		assert.Equal(t, "bar", cf.Profiles["add"].SecretKey)
	})
}

func TestProfileDelete(t *testing.T) {
	t.Run("A profile can be removed", func(t *testing.T) {
		file := filepath.Join(t.TempDir(), "delete.json")
		cf := &profile.ConfigFile{
			Path: file,
			Profiles: map[string]profile.Profile{
				"foo": {},
			},
		}
		err := cf.Save()
		require.NoError(t, err)
		_ = run(t, []string{"profile", "delete", "foo", "--config", file}, nil)

		cf, err = profile.LoadConfigFile(file)
		require.NoError(t, err)
		require.Empty(t, cf.Profiles)
	})
}

func TestProfileList(t *testing.T) {
	t.Run("Profiles can be listed", func(t *testing.T) {
		file := filepath.Join(t.TempDir(), "list.json")
		cf := &profile.ConfigFile{
			Path: file,
			Profiles: map[string]profile.Profile{
				"foo": {Default: true},
				"bar": {},
			},
		}
		err := cf.Save()
		require.NoError(t, err)
		var lst []any
		runJSON(t, []string{"profile", "list", "--config", file, "-o", "json"}, nil, &lst)
		require.Len(t, lst, 2)
		assert.Equal(t, "bar", lst[0].(map[string]any)["Name"])
		assert.False(t, lst[0].(map[string]any)["Default"].(bool))
		assert.Equal(t, "foo", lst[1].(map[string]any)["Name"])
		assert.True(t, lst[1].(map[string]any)["Default"].(bool))
	})
	t.Run("Profiles can be listed, even if no default profile is configured", func(t *testing.T) {
		file := filepath.Join(t.TempDir(), "list.json")
		cf := &profile.ConfigFile{
			Path: file,
			Profiles: map[string]profile.Profile{
				"foo": {},
			},
		}
		err := cf.Save()
		require.NoError(t, err)
		var lst []any
		runJSON(t, []string{"profile", "list", "--config", file, "-o", "json"}, nil, &lst)
		require.Len(t, lst, 1)
	})
}

func TestProfileSetDefault(t *testing.T) {
	t.Run("The default can be changed", func(t *testing.T) {
		file := filepath.Join(t.TempDir(), "use.json")
		cf := &profile.ConfigFile{
			Path: file,
			Profiles: map[string]profile.Profile{
				"foo": {Default: true},
				"bar": {},
			},
		}
		err := cf.Save()
		require.NoError(t, err)
		_ = run(t, []string{"profile", "use", "bar", "--config", file}, nil)
		var lst []any
		runJSON(t, []string{"profile", "list", "--config", file, "-o", "json"}, nil, &lst)
		require.Len(t, lst, 2)
		assert.Equal(t, "bar", lst[0].(map[string]any)["Name"])
		assert.True(t, lst[0].(map[string]any)["Default"].(bool))
		assert.Equal(t, "foo", lst[1].(map[string]any)["Name"])
		assert.False(t, lst[1].(map[string]any)["Default"].(bool))
	})
}

func TestProfilePriority(t *testing.T) {
	_ = run(t, []string{"profile", "add", "test_priority", "--ak", "profile_ak", "--sk", "profile_sk", "--region", "profile_region", "--default"}, nil)
	defer func() {
		_ = run(t, []string{"profile", "del", "test_priority"}, nil)
	}()
	t.Setenv("OSC_ACCESS_KEY", "env_ak")
	t.Setenv("OSC_SECRET_KEY", "env_sk")
	t.Setenv("OSC_REGION", "env_region")
	cfg, err := profile.DefaultConfigPath()
	require.NoError(t, err)
	t.Run("By default, the env is loaded", func(t *testing.T) {
		var p profile.Profile
		runJSON(t, []string{"profile", "current", "-o", "json"}, nil, &p)
		assert.Equal(t, "env_ak", p.AccessKey)
		assert.Equal(t, "env_sk", p.SecretKey)
		assert.Equal(t, "env_region", p.Region)
	})
	t.Run("If --profile is set, profile is loaded", func(t *testing.T) {
		var p profile.Profile
		runJSON(t, []string{"profile", "current", "-o", "json", "--profile", "test_priority"}, nil, &p)
		assert.Equal(t, "profile_ak", p.AccessKey)
		assert.Equal(t, "profile_sk", p.SecretKey)
		assert.Equal(t, "profile_region", p.Region)
	})
	t.Run("If --config is set, profile is loaded", func(t *testing.T) {
		var p profile.Profile
		runJSON(t, []string{"profile", "current", "-o", "json", "--config", cfg}, nil, &p)
		assert.Equal(t, "profile_ak", p.AccessKey)
		assert.Equal(t, "profile_sk", p.SecretKey)
		assert.Equal(t, "profile_region", p.Region)
	})
	t.Run("If --config and --profile are set, profile is loaded", func(t *testing.T) {
		var p profile.Profile
		runJSON(t, []string{"profile", "current", "-o", "json", "--profile", "test_priority", "--config", cfg}, nil, &p)
		assert.Equal(t, "profile_ak", p.AccessKey)
		assert.Equal(t, "profile_sk", p.SecretKey)
		assert.Equal(t, "profile_region", p.Region)
	})
	t.Run("An alternate config file can be used", func(t *testing.T) {
		cfg := filepath.Join(t.TempDir(), "priority.json")
		cf := &profile.ConfigFile{
			Path: cfg,
			Profiles: map[string]profile.Profile{
				"foo": {
					AccessKey: "priority_ak",
					SecretKey: "priority_sk",
					Region:    "priority_region",
					Default:   true,
				},
			},
		}
		err := cf.Save()
		require.NoError(t, err)
		var p profile.Profile
		runJSON(t, []string{"profile", "current", "-o", "json", "--config", cfg}, nil, &p)
		assert.Equal(t, "priority_ak", p.AccessKey)
		assert.Equal(t, "priority_sk", p.SecretKey)
		assert.Equal(t, "priority_region", p.Region)
	})
}
