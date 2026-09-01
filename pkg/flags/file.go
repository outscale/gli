package flags

import (
	"encoding/json"
	"errors"
	"os"
)

const FileOrJSON = "fileOrJson"

var ErrInvalidFileOrJSON = errors.New("value is neither a file nor a JSON document")

// FileOrJSONValue sets a flag with either a file content.
type FileOrJSONValue struct {
	content []byte
}

func NewFileOrJSONValue() *FileOrJSONValue {
	return &FileOrJSONValue{}
}

// Set sets the value based on a file content.
func (v *FileOrJSONValue) Set(s string) error {
	if _, err := os.Stat(s); err == nil {
		v.content, err = os.ReadFile(s) //nolint:gosec
		return err
	}
	var tmp any
	err := json.Unmarshal([]byte(s), &tmp)
	if err != nil {
		return ErrInvalidFileOrJSON
	}
	v.content = []byte(s)
	return nil
}

// Type name for File.File flags.
func (v *FileOrJSONValue) Type() string {
	return "fileOrJson"
}

func (v *FileOrJSONValue) String() string {
	return string(v.content)
}

func (v *FileOrJSONValue) Value() (string, bool) {
	return v.String(), true
}
