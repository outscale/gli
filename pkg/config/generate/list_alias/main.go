package main

import (
	"fmt"

	"github.com/outscale/octl/pkg/config"
	"github.com/samber/lo"
)

func main() {
	for _, provider := range []string{"iaas", "storage"} {
		fmt.Println("***", provider, "***")
		cfg := config.For(provider)
		for call := range cfg.Calls {
			if lo.ContainsBy(cfg.Aliases, func(a config.Alias) bool {
				return a.AliasTo == call
			}) {
				continue
			}
			fmt.Println(call, "is missing")
		}
	}
}
