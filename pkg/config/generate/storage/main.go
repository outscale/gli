package main

import (
	"os"
	"strings"

	"github.com/goccy/go-yaml"
	"github.com/outscale/octl/pkg/config"
	"github.com/outscale/octl/pkg/config/generate/builder"
	"github.com/outscale/octl/pkg/messages"
	"github.com/outscale/osc-sdk-go/v3/pkg/oos"
)

func main() {
	src := os.Args[1]
	dst := os.Args[2]
	var base config.Config
	data, err := os.ReadFile(src) //nolint:gosec
	if err != nil {
		messages.ExitErr(err)
	}
	err = yaml.Unmarshal(data, &base)
	if err != nil {
		messages.ExitErr(err)
	}
	if base.Calls == nil {
		base.Calls = map[string]config.Call{}
	}
	if base.Entities == nil {
		base.Entities = map[string]config.Entity{}
	}
	if base.Spec.Calls == nil {
		base.Spec.Calls = map[string]config.SpecCall{}
	}
	if base.Spec.Attributes == nil {
		base.Spec.Attributes = map[string]config.SpecAttribute{}
	}

	cfg := builder.Config{
		InputStructSuffix: "Input",
		SkipFlagsPrefixes: []string{"ContinuationToken", "BucketRegion", "Marker", "RequestPayer", "MaxKeys", "CreateBucketConfiguration.", "SSE", "StorageClass", "ContentMD5", "Checksum"},
		PriorityFields:    []string{},
		FlagOverrides:     map[string]config.Flag{},
		RequiredFromComment: func(s string) bool {
			return strings.HasSuffix(s, "This member is required.")
		},
	}

	sb := builder.NewSpecBuilder(cfg)
	sb.BuildSpec(&base, "github.com/outscale/osc-sdk-go/v3/pkg/oos", "github.com/aws/aws-sdk-go-v2/service/s3", "github.com/aws/aws-sdk-go-v2/service/s3/types")

	b := builder.NewClientBuilder[*oos.Client](cfg)
	b.BuildFor(&base)

	fd, err := os.Create(dst) //nolint:gosec
	if err != nil {
		messages.ExitErr(err)
	}
	err = yaml.NewEncoder(fd, yaml.UseSingleQuote(true), yaml.UseLiteralStyleIfMultiline(true)).Encode(base)
	if err != nil {
		messages.ExitErr(err)
	}
	err = fd.Close()
	if err != nil {
		messages.ExitErr(err)
	}
}
