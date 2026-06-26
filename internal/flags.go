package internal

import "github.com/urfave/cli/v2"

// AppIdentifierFlag returns the CLI flag for authenticating as a GitHub App.
// GitHub recommends using the client ID as the JWT iss claim, but also accepts the application ID.
// See: https://docs.github.com/en/apps/creating-github-apps/authenticating-with-a-github-app/generating-a-json-web-token-jwt-for-a-github-app
func AppIdentifierFlag() cli.Flag {
	return &cli.StringFlag{
		Name:     "client-id",
		Usage:    "GitHub App client ID (preferred). --app-id is also accepted; either value is used as the JWT iss claim",
		Required: true,
		Aliases:  []string{"app-id", "app_id", "client_id", "i"},
	}
}

// AppIdentifier returns the GitHub App identifier from the CLI context.
func AppIdentifier(c *cli.Context) string {
	return c.String("client-id")
}
