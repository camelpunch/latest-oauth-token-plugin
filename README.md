# Hello, World! CF CLI plugin

A minimal [CF CLI](https://github.com/cloudfoundry/cli) plugin that prints out "Hello, World!".

## Installation

```
go build
cf install-plugin ./hello-world-plugin
```

## Usage

```
cf hello-world
```

## Uninstallation

```
cf uninstall-plugin HelloWorldPlugin # bask in the asymmetry
```
