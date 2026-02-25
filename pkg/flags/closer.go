package flags

import (
	"fmt"

	"github.com/spf13/pflag"
)

type Closer interface {
	Close() error
}

func CloseAll(fs *pflag.FlagSet) error {
	var err error
	fs.VisitAll(func(f *pflag.Flag) {
		v := f.Value
		if closer, ok := v.(Closer); ok {
			cerr := closer.Close()
			if cerr != nil {
				err = fmt.Errorf("closing %s: %v", f.Name, cerr)
			}
		}
	})
	return err
}
