package main

import (
	"archive/zip"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/outscale/octl/pkg/messages"
)

func main() {
	path := os.Args[1]

	archive, err := os.Create(filepath.Join(path, "defaults.zip")) //nolint:gosec
	if err != nil {
		messages.ExitErr(err)
	}
	w := zip.NewWriter(archive)
	dir := os.DirFS(path)
	ls, err := fs.ReadDir(dir, ".")
	if err != nil {
		messages.ExitErr(err)
	}
	for _, d := range ls {
		if !strings.HasSuffix(d.Name(), ".yaml") {
			continue
		}
		f, err := w.Create(d.Name())
		if err != nil {
			messages.ExitErr(err)
		}
		body, err := os.ReadFile(d.Name()) //nolint:gosec
		if err != nil {
			messages.ExitErr(err)
		}
		_, err = f.Write(body)
		if err != nil {
			messages.ExitErr(err)
		}
	}
	if err == nil {
		err = w.Close()
	}
	if err != nil {
		messages.ExitErr(err)
	}
}
