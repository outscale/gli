package prerun

import (
	"context"

	"github.com/outscale/octl/pkg/messages"
	"github.com/outscale/octl/pkg/preferences"
	"github.com/outscale/osc-sdk-go/v3/pkg/profile"
	"github.com/spf13/cobra"
)

type preferenceKey struct{}

func contextWithPreferences(ctx context.Context, p preferences.Preferences) context.Context {
	return context.WithValue(ctx, preferenceKey{}, p)
}

func PreferencesFrom(ctx context.Context) preferences.Preferences {
	p := ctx.Value(preferenceKey{})
	if p == nil {
		return preferences.Preferences{}
	}
	return p.(preferences.Preferences)
}

func LoadPreferences(cmd *cobra.Command) {
	prof, _ := cmd.Flags().GetString("profile")
	if prof == "" {
		prof = profile.DefaultProfile
	}
	prefs, err := preferences.Get(prof)
	if err != nil {
		messages.Err(err.Error())
	}
	cmd.SetContext(contextWithPreferences(cmd.Context(), prefs))
}
