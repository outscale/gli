package preferences

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"dario.cat/mergo"
	"github.com/goccy/go-yaml"
	"github.com/outscale/octl/pkg/debug"
	"github.com/outscale/octl/pkg/messages"
)

type Kube struct {
	DefaultProject string `yaml:"default_project,omitempty"`
}

type Preferences struct {
	Kube Kube `yaml:"kube,omitzero"`
}

type File struct {
	Global     Preferences            `yaml:"global,omitzero"`
	PerProfile map[string]Preferences `yaml:"per_profile,omitzero"`
}

func Load() (File, error) {
	root, err := os.UserConfigDir()
	if err != nil {
		debug.Println("Unable to compute user preferences dir", err)
		return File{}, err
	}
	path := filepath.Join(root, "octl", "preferences.yaml")
	fd, err := os.Open(path) //nolint:gosec
	if errors.Is(err, fs.ErrNotExist) {
		debug.Println("no user config file found", path)
		return File{}, nil
	}
	if err != nil {
		messages.Err("Unable to open preferences path: %w", err)
		return File{}, err
	}
	debug.Println("loading user config from", path)
	var pf File
	err = yaml.NewDecoder(fd).Decode(&pf)
	if err != nil {
		return File{}, err
	}
	return pf, nil
}

func Get(profile string) (Preferences, error) {
	pf, err := Load()
	if err != nil {
		return Preferences{}, fmt.Errorf("unable to load preferences: %w", err)
	}
	if pf.PerProfile == nil {
		return pf.Global, nil
	}
	prefs := pf.PerProfile[profile]
	err = mergo.Merge(&prefs, pf.Global)
	if err != nil {
		messages.Err("Unable to process preferences: %w", err)
	}
	return prefs, nil
}

func Set(profile string, fn func(prefs *Preferences)) error {
	pf, err := Load()
	if err != nil {
		return fmt.Errorf("unable to load preferences: %w", err)
	}
	if pf.PerProfile == nil {
		pf.PerProfile = map[string]Preferences{}
	}
	pref := pf.PerProfile[profile]
	fn(&pref)
	pf.PerProfile[profile] = pref
	return pf.Save()
}

func SetGlobal(fn func(prefs *Preferences)) error {
	pf, err := Load()
	if err != nil {
		return fmt.Errorf("unable to load preferences: %w", err)
	}
	fn(&pf.Global)
	return pf.Save()
}

func (pf *File) Save() (err error) {
	root, err := os.UserConfigDir()
	if err != nil {
		debug.Println("Unable to compute user preferences dir", err)
		return err
	}
	dir := filepath.Join(root, "octl")
	err = os.MkdirAll(dir, 0o700)
	if err != nil {
		return err
	}
	path := filepath.Join(dir, "preferences.yaml")
	fd, err := os.Create(path) //nolint:gosec
	if err != nil {
		return err
	}
	defer func() {
		cerr := fd.Close()
		if err == nil {
			err = cerr
		}
	}()
	debug.Println("saving user preferences to", path)
	return yaml.NewEncoder(fd).Encode(pf)
}
