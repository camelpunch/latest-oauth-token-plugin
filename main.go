package main

import (
	"fmt"

	"github.com/cloudfoundry/cli/plugin"
)

type LatestOauthToken struct {
}

func (c *LatestOauthToken) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "LatestOauthTokenPlugin",
		Version: plugin.VersionType{
			Major: 0,
			Minor: 0,
			Build: 1,
		},
		MinCliVersion: plugin.VersionType{
			Major: 6,
			Minor: 12,
			Build: 4,
		},
		Commands: []plugin.Command{
			{
				Name:     "latest-oauth-token",
				HelpText: "Command to update and print the latest oauth token",
				UsageDetails: plugin.Usage{
					Usage: "cf latest-oauth-token",
				},
			},
		},
	}
}

func (c *LatestOauthToken) Run(cliConnection plugin.CliConnection, args []string) {
	token, _ := cliConnection.AccessToken()
	fmt.Println(token)
}

func main() {
	plugin.Start(new(LatestOauthToken))
}
