# metal-cli

[![GitHub release](https://img.shields.io/github/release/equinix/metal-cli/all.svg?style=flat-square)](https://github.com/equinix/metal-cli/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/equinix/metal-cli)](https://goreportcard.com/report/github.com/equinix/metal-cli)
[![Slack](https://slack.equinixmetal.com/badge.svg)](https://slack.equinixmetal.com)
[![Twitter Follow](https://img.shields.io/twitter/follow/equinixmetal.svg?style=social&label=Follow)](https://twitter.com/intent/follow?screen_name=equinixmetal)
[![Stability: Maintained](https://img.shields.io/badge/Stability-Maintained-green.svg)](https://github.com/packethost/standards/blob/master/maintained-statement.md)

This repository is [Maintained](https://github.com/packethost/standards/blob/master/maintained-statement.md) meaning that this software is supported by Equinix Metal and its community - available to use in production environments.

## Table of Contents

* [Equinix Metal Command Line Interface](#metal-command-line-interface)
* [Requirements](#requirements)
* [Supported Platforms](#supported-platforms)
* [Installation](#installation)
  * [Install binary from Source](#install-binary-from-source)
  * [Install binary from Release Download](#install-binary-from-release-download)
  * [Install via Homebrew](#install-via-homebrew)
  * [Configuring your Path](#configuring-your-path)
  * [Shell Completion](#shell-completion)
* [Authentication](#authentication)
* [Reference](#reference)
* [Example Syntax](#example-syntax)
* [Support](#support)

## Equinix Metal Command Line Interface

[Equinix Metal](https://metal.equinix.com/) provides an API-driven bare metal platform that combines the automation experience of the cloud with the benefits of physical, bare metal servers.

The Equinix Metal CLI wraps the [Equinix Metal Go SDK](https://github.com/packethost/packngo) allowing interaction with Equinix Metal platform from a command-line interface.

## Requirements

* Equinix Metal authentication token.
* Equinix Metal CLI [binaries](https://github.com/equinix/metal-cli/releases).

## Supported Platforms

The [Equinix Metal CLI binaries](https://github.com/equinix/metal-cli/releases) are available for Linux, Windows, and Mac OS X for various architectures including ARM on Linux.

## Installation

### Install binary from Source

If you have `go` 1.16 or later installed, you can build and install the latest version with:

```sh
go install github.com/equinix/metal-cli/cmd/metal@latest
```

You can find the installed executable/binary in either `$GOPATH/bin` or `$HOME/go/bin` folder.

### Install binary from Release Download

Download the appropriate Equinix Metal CLI binaries for your platform to the desired location,`chmod` it and rename it to `metal`.


### Install via Homebrew

If you prefer installing via Homebrew, you can run the following:

```bash
brew tap equinix/homebrew-tap
brew install metal-cli
```
### Configuring your Path

If the directory where your binaries were installed is not already in your `PATH` environment variable, then it will need to be added.
Choose the steps to follow for your platform to add directory to `PATH`.

<details>
  <summary>Configure your path in Linux</summary>

## Linux

If you plan to run the Equinix Metal CLI in a shell on Linux and placed the binary in `/home/YOUR-USER-NAME/metal-cli/`, then type the following into your terminal:

```sh
export PATH=$PATH:/home/$USER/metal-cli
```

If you plan to run the Equinix Metal CLI in a shell on Linux and your binary is in `$GOPATH/bin` or `$HOME/go/bin`, then type the following into your terminal:

```sh
export PATH=$PATH:$GOPATH/bin
```

or:

```sh
export PATH=$PATH:$HOME/go/bin
```

You can view the current value of `$PATH` by running:

```sh
echo $PATH
```

</details>

<details>
  <summary>Configure your path in Mac OS/X</summary>

### Mac OS X

If you plan to run the Equinix Metal CLI in a shell on a Mac, download the `darwin` binary and placed the it in `/Users/YOUR-USER-NAME/metal-cli/`, then type the following into your terminal.

```sh
export PATH=$PATH:/Users/$USER/metal-cli
```

If you plan to run the Equinix Metal CLI in a shell on a Mac and your binary is in `$GOPATH/bin` or `$HOME/go/bin`, then type the following into your terminal:

```sh
export PATH=$PATH:$GOPATH/bin
```

or:

```sh
export PATH=$PATH:$HOME/go/bin
```

You can view the current value of `$PATH` by running:

```sh
echo $PATH
```

When running the downloaded binary on a Mac, you may be prompted with the following message:

> "metal" cannot be opened because the developer cannot be verified

The binary can be trusted by enabling "App Store and identified developers" in "System Preferences -> Security & Privacy -> General".  Any blocked apps will appear in the bottom of this window, where they can be authorized.

</details>

<details>
  <summary>Configure your path in Windows</summary>

### Windows

If you plan to run the Equinix Metal CLI in PowerShell on Windows and placed the binary in `c:\metal-cli`, then type the following into PowerShell:

```powershell
$env:Path += ";c:\metal-cli"
```

If you plan to run the Equinix Metal CLI in PowerShell on Windows and your binary is in `$GOPATH/bin` or `$HOME/go/bin`, then type the following into PowerShell:

```powershell
$setx PATH "$($env:path);$GOPATH\bin"
```

or:

```powershell
$setx PATH "$($env:path);$HOME\go\bin"
```

The path can be viewed by running:

```sh
echo $env:Path
```
</details>

### Shell Completion

Once installed, shell completion can be enabled (in Bash) with `source <(metal completion bash)` (or for some versions of Bash, `eval "$(metal completion bash)").

Check `metal completion -h` for instructions to use in other shells.

## Authentication

After installing Equinix Metal CLI, configure your account using `metal init`:

```bash
$ metal init
Equinix Metal API Tokens can be obtained through the portal at https://console.equinix.com/.
See https://metal.equinix.com/developers/docs/accounts/users/ for more details.

Token (hidden): 
Organization ID []: 87e62b5c-7e4e-4a29-ac40-d5df9535868f
Project ID []: a4e48c3a-6819-485a-822f-81b3845d5aa5

Writing /Users/username/.config/equinix/metal.yaml
```

The Equinix Metal authentication token can be stored in the `$METAL_AUTH_TOKEN` environment variable or in JSON or YAML configuration files. The configuration file path can be overridden with the `--config` flag.  The default configuration path is "$HOME/config/equinix/metal.*" (any supported filetype).

## Includes and Excludes

Equinix Metal API resource responses may have related resources. These related
resources can be embedded in the result or referred. Referred resources will
only include a `Href` value, which includes the unique ID of the resource.
Embedded resources will be represented with their full API value, which may
contain additional embedded or referred resources.

The resources that you want embedded can be _included_ in results using
`--include`.  The resources that you want referred can be _excluded_ with
`--exclude`.  By excluding some of the embedded-by-default resources, you can
speed up and reduce the size of responses.  By including referred-by-default
resources, you can avoid the round trip of subsequent calls.

```sh
metal devices get --project-id $ID --yaml --exclude=ssh_keys,plan --include=project
```

These arguments are available in any command that returns a response document.
The included and excluded fields requested from the API may differ based on the
output format, for example, for historic reasons, `metal projects get --yaml`
includes the details of all project members. In the table output format, no
member related fields are displayed and so `metal projects get` will exclude
the member resource.

Excluding fields needed for the table output will result in an error. Mixing
includes and excludes affecting the same top-level field is not supported.

## Reference

The full CLI documentation can be found [here](docs/metal.md) or by clicking the links below.

## Example Syntax

### Create a device

```sh
metal device create --hostname [hostname] --plan [plan] --facility [facility_code] --operating-system [operating_system] --project-id [project_UUID]
```

### Get a device

```sh
metal device get --id [device_UUID]
```

### List projects

```sh
metal project get
```

### Get a project

```sh
metal project get -i [project_UUID]
```

Details on all available commands can be found by visiting the reference [pages](docs/metal.md) or typing `metal [command] --help` for more information about the specific command.

## Support

For help with this package:

* Open up a GitHub issue [here](https://github.com/equinix/metal-cli/issues).
* Contact the [Equinix Metal Community Slack](http://slack.equinixmetal.net) or on Freenode IRC in the #equinixmetal channel.
