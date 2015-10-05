# Latest OAuth Token CF CLI plugin

A [CF CLI](https://github.com/cloudfoundry/cli) plugin that uses the CLI's AccessToken function to update (in your `~/.cf/config.json`) and print your oauth token.

## Installation

```
go build
cf install-plugin ./latest-oauth-token-plugin
```

## Usage

```
cf latest-oauth-token
```

## Uninstallation

```
cf uninstall-plugin LatestOauthTokenPlugin
```
