package flags

import (
	"os"
)

// FileValue sets a flag with a file content.
type FileValue struct {
	content []byte
}

func NewFileValue() *FileValue {
	return &FileValue{}
}

// Set sets the value based on a file content.
func (v *FileValue) Set(s string) error {
	buf, err := os.ReadFile(s) //nolint:gosec
	if err != nil {
		return err
	}
	v.content = buf
	return nil
}

// Type name for File.File flags.
func (v *FileValue) Type() string {
	return "File"
}

func (v *FileValue) String() string {
	return string(v.content)
}

func (v *FileValue) Value() (string, bool) {
	return v.String(), true
}
