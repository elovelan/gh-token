package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func TestAppIdentifierFlagAliases(t *testing.T) {
	tests := []struct {
		name string
		args []string
	}{
		{
			name: "client-id flag",
			args: []string{"--client-id", "Iv1.abc123"},
		},
		{
			name: "app-id alias",
			args: []string{"--app-id", "123456"},
		},
		{
			name: "client_id alias",
			args: []string{"--client_id", "Iv1.abc123"},
		},
		{
			name: "app_id alias",
			args: []string{"--app_id", "123456"},
		},
		{
			name: "short i alias",
			args: []string{"-i", "123456"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got string
			app := &cli.App{
				Flags: []cli.Flag{AppIdentifierFlag()},
				Action: func(c *cli.Context) error {
					got = AppIdentifier(c)
					return nil
				},
			}

			err := app.Run(append([]string{"test"}, tt.args...))
			assert.NoError(t, err)
			assert.Equal(t, tt.args[len(tt.args)-1], got)
		})
	}
}
