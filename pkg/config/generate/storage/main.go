package main

import (
	"os"
	"strings"

	"github.com/goccy/go-yaml"
	configbuilder "github.com/outscale/octl/pkg/builder/config"
	"github.com/outscale/octl/pkg/config"
	"github.com/outscale/octl/pkg/flags"
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
	if base.API == nil {
		base.API = map[string]config.APICall{}
	}
	if base.Entities == nil {
		base.Entities = map[string]config.Entity{}
	}

	cfg := configbuilder.Config{
		AllowedNumOut: []int{1, 2},
		SkipAlias: []string{
			"GetBucketAcl",
			"GetBucketCors",
			"GetBucketEncryption",
			"GetBucketLifecycleConfiguration",
			"GetBucketLocation",
			"GetBucketPolicy",
			"GetBucketVersioning",
			"GetBucketWebsite",
			"GetObjectLockConfiguration",
			"GetObjectRetention",
			"GetObjectTagging",
			"GetObject",
			"GetObjectAcl",
			"GetBucketWebsite",
			"ListObjectVersions",
			"PutObjectTagging",
			"PutBucketLifecycleConfiguration",
			"PutObjectLockConfiguration",
			"PutBucketWebsite",
			"PutBucketPolicy",
			"PutObjectAcl",
			"PutObjectRetention",
			"PutObjectTagging",
			"PutBucketAcl",
			"DeleteBucketLifecycle",
			"DeleteBucketWebsite",
			"DeleteBucketPolicy",
			"DeleteObjectTagging",
			"DeleteBucket",
		},
		InputStructSuffix: "Input",
		SkipFlagsPrefixes: map[string][]string{
			"*": {
				"ContinuationToken", "BucketRegion", "Marker", "RequestPayer", "MaxKeys", "CreateBucketConfiguration.",
				"SSE", "StorageClass", "ContentMD5", "Checksum",
			},
		},
		PriorityFields: []string{},
		FlagOverrides: map[string]config.Flag{
			"Policy": {
				CustomValue: flags.FileOrJSON,
			},
		},
		RequiredFromComment: func(s string) bool {
			return strings.HasSuffix(s, "This member is required.")
		},
	}

	sb := configbuilder.NewSpecBuilder(cfg)
	spec := sb.BuildSpec(&base, "github.com/outscale/osc-sdk-go/v3/pkg/oos", "github.com/aws/aws-sdk-go-v2/service/s3", "github.com/aws/aws-sdk-go-v2/service/s3/types")

	b := configbuilder.NewClientBuilder[*oos.Client](cfg)
	b.BuildFor(&base, spec)

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
