package flags

import (
	"encoding/base64"
	"os"

	"github.com/outscale/octl/pkg/debug"
)

// Base64FileValue adapts File.File for use as a flag.
type Base64FileValue struct {
	content []byte
}

func NewBase64FileValue() *Base64FileValue {
	return &Base64FileValue{}
}

// Set sets the value based on a file content.
func (v *Base64FileValue) Set(s string) error {
	buf, err := os.ReadFile(s) //nolint:gosec
	if err != nil {
		return err
	}
	v.content = buf
	return nil
}

// Type name for File.File flags.
func (v *Base64FileValue) Type() string {
	return "base64File"
}

func (v *Base64FileValue) String() string {
	_, err := base64.StdEncoding.DecodeString(string(v.content))
	if err == nil {
		debug.Println("value is already base 64 encoded")
		return string(v.content)
	}
	return base64.StdEncoding.EncodeToString(v.content)
}

func (v *Base64FileValue) Value() (string, bool) {
	return v.String(), true
}
