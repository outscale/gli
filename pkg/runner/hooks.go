package runner

import "reflect"

type Hook func(arg reflect.Value)

var registeredHooks map[string]Hook

func init() {
	registeredHooks = map[string]Hook{}
}

func RegisterHook(name string, hook Hook) {
	registeredHooks[name] = hook
}

func applyHooks(arg reflect.Value, hooks []string) {
	for _, name := range hooks {
		if hook, ok := registeredHooks[name]; ok {
			hook(arg)
		}
	}
}
