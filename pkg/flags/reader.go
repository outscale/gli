package flags

import (
	"io"
	"os"
)

// ReaderValue sets an io.Reader value from a flag.
type ReaderValue struct {
	path string
	r    io.ReadCloser
}

func NewReaderValue() *ReaderValue {
	return &ReaderValue{}
}

// Set sets the value based on a file content.
func (v *ReaderValue) Set(s string) error {
	r, err := os.Open(s)
	if err != nil {
		return err
	}
	v.path = s
	v.r = r
	return nil
}

// Type name for File.File flags.
func (v *ReaderValue) Type() string {
	return "streamedFile"
}

func (v *ReaderValue) String() string {
	return v.path
}

func (v *ReaderValue) Value() (io.Reader, bool) {
	return v.r, true
}

func (v *ReaderValue) Close() error {
	return v.r.Close()
}
