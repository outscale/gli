package output

import "github.com/spf13/pflag"

func NewFromFlags(fs *pflag.FlagSet) (Filter, error) {
	jq, _ := fs.GetString("jq")
	if jq != "" {
		return NewJQFilter(jq)
	}
	return &Default{}, nil
}
