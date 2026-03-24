package read

import (
	"encoding"
	"io"
	"reflect"
	"unicode/utf8"

	"github.com/gabriel-vasile/mimetype"
	"github.com/outscale/octl/pkg/debug"
	"github.com/outscale/octl/pkg/structs"
)

const contentPreviewSize = 100

// PreviewReader is a wrapper around the object body, as returned by GetObject.
type PreviewReader struct {
	io.ReadCloser
}

// MarshalText allows previewing the body in JSON/YAML dumps.
// Note: this only works if the reader has not been consumed.
func (o PreviewReader) MarshalText() ([]byte, error) {
	defer o.Close() //nolint
	debug.Println("read preview")
	// ensure we have enough for contentPreviewSize UTF-8 chars
	buf, err := io.ReadAll(io.LimitReader(o, contentPreviewSize*4))
	if err != nil {
		return nil, err
	}
	mime := mimetype.Detect(buf)
	if mime.Is("text/plain") {
		str := string(buf)
		str = str[:min(contentPreviewSize, len(str))]
		// the content could still be binary data
		if utf8.ValidString(str) && len(str) == contentPreviewSize {
			return []byte(str + " ..."), nil
		}
	}
	return []byte(mime.String() + " data"), nil
}

var _ encoding.TextMarshaler = PreviewReader{}

func addPreview(v reflect.Value) {
	if v.Kind() != reflect.Struct {
		return
	}
	v, found := structs.FindFieldByType[io.ReadCloser](v)
	if found && v.CanSet() && v.CanInterface() {
		if r, ok := v.Interface().(io.ReadCloser); ok {
			debug.Println("add preview")
			pv := PreviewReader{r}
			v.Set(reflect.ValueOf(pv))
			return
		}
	}
}
