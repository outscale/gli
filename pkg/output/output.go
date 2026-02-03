package output

import "context"

type Filter interface {
	Output(ctx context.Context, v any) error
}
