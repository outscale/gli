package runner

import (
	"reflect"
)

type Hook func(arg reflect.Value)

func applyHooks(arg reflect.Value, hooks ...Hook) {
	for _, hook := range hooks {
		hook(arg)
	}
}
