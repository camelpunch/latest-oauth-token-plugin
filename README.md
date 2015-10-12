# Latest OAuth Token CF CLI plugin

A [CF CLI](https://github.com/cloudfoundry/cli) plugin that uses the CLI's
AccessToken function to update (in your `~/.cf/config.json`) and print your
oauth token.

## Installation

```
go build
cf install-plugin ./latest-oauth-token-plugin
```

## Usage

```
cf latest-oauth-token
```

If https://github.com/cloudfoundry/cli/pull/623 has been merged, you should see changes to your `~/.cf/config.json` each time this command is
run.

## Uninstallation

```
cf uninstall-plugin LatestOauthTokenPlugin
```
