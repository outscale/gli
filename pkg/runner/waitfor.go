package runner

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/itchyny/gojq"
	"github.com/outscale/octl/pkg/config"
	"github.com/outscale/octl/pkg/messages"
	"github.com/outscale/octl/pkg/output"
	"github.com/spf13/cobra"
)

func waitfor[Client any, Error error](cmd *cobra.Command, args []string, cl Client, cfg config.Config) error {
	expr, _ := cmd.Flags().GetString("waitfor")

	// force JSON output
	f := cmd.Flags().Lookup("output")
	_ = f.Value.Set("json")

	query, err := gojq.Parse(expr)
	if err != nil {
		return fmt.Errorf("parse waitfor condition: %w", err)
	}

	tmout, _ := cmd.Flags().GetDuration("waitfor-timeout")
	ctx, cancel := context.WithTimeout(cmd.Context(), tmout)
	defer cancel()
	interval, _ := cmd.Flags().GetDuration("waitfor-interval")
	for {
		select {
		case <-ctx.Done():
			return errors.New("timeout waiting for condition to succeed")
		default:
		}
		buf := &bytes.Buffer{}
		output.InjectOutput(buf)
		err := doRun[Client, Error](cmd, args, cl, cfg)
		if err != nil {
			return err
		}

		var raw any
		err = json.Unmarshal(buf.Bytes(), &raw)
		if err != nil {
			return fmt.Errorf("parse JSON: %w", err)
		}
		iter := query.RunWithContext(ctx, raw)
		// we expect a single result, no need for a loop
		v, ok := iter.Next()
		if !ok {
			return errors.New("no result from condition")
		}
		// gojq returned an error
		if err, ok := v.(error); ok {
			return fmt.Errorf("jq error: %w", err)
		}
		switch v {
		case false, "false", 0, "0", "", nil:
			messages.Info("⏱️ Waiting for condition to succeed")
		default:
			messages.Success("Condition reached successfully")
			return nil
		}
		time.Sleep(interval)
	}
}
