# packet-cli

[![GitHub release](https://img.shields.io/github/release/packethost/packet-cli/all.svg?style=flat-square)](https://github.com/packethost/packet-cli/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/packethost/packet-cli)](https://goreportcard.com/report/github.com/packethost/packet-cli)
[![Slack](https://slack.equinixmetal.com/badge.svg)](https://slack.equinixmetal.com)
[![Twitter Follow](https://img.shields.io/twitter/follow/packethost.svg?style=social&label=Follow)](https://twitter.com/intent/follow?screen_name=equinixmetal)
![](https://img.shields.io/badge/Stability-Maintained-green.svg)

This repository is [Maintained](https://github.com/packethost/standards/blob/master/maintained-statement.md) meaning that this software is supported by Equinix Metal and its community - available to use in production environments.

## Table of Contents

* [Equinix Metal Command Line Interface](#packet-command-line-interface)
* [Requirements](#requirements)
* [Supported Platforms](#supported-platforms)
* [Installation](#installation)
  * [Linux](#linux)
  * [Mac OS X](#mac-os-x)
  * [Windows](#windows)
  * [Shell Completion](#shell-completion)
* [Authentication](#authentication)
* [Reference](#reference)
* [Example Syntax](#example-syntax)
* [Support](#support)

## Equinix Metal Command Line Interface

[Packet is now Equinix Metal!](https://blog.equinix.com/blog/2020/10/06/equinix-metal-metal-and-more/), keep an eye on this project for future project and command line name changes.

[Equinix Metal](https://metal.equinix.com/) provides an API-driven bare metal platform that combines the automation experience of the cloud with the benefits of physical, bare metal servers.

The Equinix Metal CLI wraps the [Equinix Metal Go SDK](https://github.com/packethost/packngo) allowing interaction with Equinix Metal platform from a command-line interface.

## Requirements

* Equinix Metal authentication token.
* Equinix Metal CLI [binaries](https://github.com/packethost/packet-cli/releases).

## Supported Platforms

The [Equinix Metal CLI binaries](https://github.com/packethost/packet-cli/releases) are available for Linux, Windows, and Mac OS X for various architectures including ARM on Linux.

## Installation

### Install binary from Source

Download the appropriate Equinix Metal CLI package, build and install them. Type the following in your terminal:

```
GO111MODULE=on go get github.com/packethost/packet-cli
```

You can find the installed executable/binary in either `$GOPATH/bin` or `$HOME/go/bin` folder.

Rename binary from `packet-cli` to `packet`. 

### Install binary from Release Download

Download the appropriate Equinix Metal CLI binaries for your platform to the desired location,`chmod` it and rename it to `packet`.

### Configuring your Path

If the directory where your binaries were installed is not already in your `PATH` environment variable, then it will need to be added.
Choose the steps to follow for your platform to add directory to `PATH`.

<details>
  <summary>Configure your path in Linux</summary>

## Linux

If you plan to run the Equinix Metal CLI in a shell on Linux and placed the binary in `/home/YOUR-USER-NAME/packet-cli/`, then type the following into your terminal:

```
export PATH=$PATH:/home/$USER/packet-cli
```
If you plan to run the Equinix Metal CLI in a shell on Linux and your binary is in `$GOPATH/bin` or `$HOME/go/bin`, then type the following into your terminal:

``
export PATH=$PATH:$GOPATH/bin
```

or:

```
export PATH=$PATH:$HOME/go/bin
```

You can view the current value of `$PATH` by running:

```
echo $PATH
```

</details>

<details>
  <summary>Configure your path in Mac OS/X</summary>

### Mac OS X

If you plan to run the Equinix Metal CLI in a shell on a Mac, download the `darwin` binary and placed the it in `/Users/YOUR-USER-NAME/packet-cli/`, then type the following into your terminal.

```
export PATH=$PATH:/Users/$USER/packet-cli
```

If you plan to run the Equinix Metal CLI in a shell on a Mac and your binary is in `$GOPATH/bin` or `$HOME/go/bin`, then type the following into your terminal:

```
export PATH=$PATH:$GOPATH/bin
```

or:

```
export PATH=$PATH:$HOME/go/bin
```

You can view the current value of `$PATH` by running:

```
echo $PATH
```

When running the downloaded binary on a Mac, you may be prompted with the following message:

> "packet" cannot be opened because the developer cannot be verified

The binary can be trusted by enabling "App Store and identified developers" in "System Preferences -> Security & Privacy -> General".  Any blocked apps will appear in the bottom of this window, where they can be authorized.

</details>

<details>
  <summary>Configure your path in Windows</summary>

### Windows

If you plan to run the Equinix Metal CLI in PowerShell on Windows and placed the binary in `c:\packet-cli`, then type the following into PowerShell:

```
$env:Path += ";c:\packet-cli"
```
If you plan to run the Equinix Metal CLI in PowerShell on Windows and your binary is in `$GOPATH/bin` or `$HOME/go/bin`, then type the following into PowerShell:

```
$setx PATH "$($env:path);$GOPATH\bin"
```

or:

```
$setx PATH "$($env:path);$HOME\go\bin"
```

The path can be viewed by running:


```
echo $env:Path
```
</details>

### Shell Completion

Once installed, shell completion can be enabled (in Bash) with `source <(packet completion bash)` (or for some versions of Bash, `eval "$(packet completion bash)").

Check `packet completion -h` for instructions to use in other shells.

## Authentication

The Equinix Metal authentication token can be stored in the `$PACKET_TOKEN` environment variable or in JSON or YAML configuration files. The configuration file path can be overridden with the `--config` flag.

Environment variable:

```bash
export PACKET_TOKEN=[PACKET_TOKEN]
```

YAML configuration file - `$HOME/.packet-cli.yaml`:

```yaml
---
token: PACKET_TOKEN

```

JSON configuration file - `$HOME/.packet-cli.json`:

```json
{
  "token": "PACKET_TOKEN"
}
```

After installing Equinix Metal CLI, verify the installation by executing `packet` or `packet.exe`. You should see the default output:


```bash
$ packet
Command line interface for Equinix Metal

Usage:
  packet [command]

Available Commands:
  device            Device operations
  facilities        Facility operations
  help              Help about any command
  ip                IP operations
  operating-systems Operating system operations
  organization      Organization operations
  plan              Plan operations
  project           Project operations
  ssh-key           SSH key operations
  user              User operations
  virtual-network   Virtual network operations
  volume            Volume operations

Flags:
      --config string   Path to JSON or YAML configuration file
  -h, --help            help for packet
      --version         version for packet

Use "packet [command] --help" for more information about a command.
```

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
packet devices get --project-id $ID --yaml --exclude=ssh_keys,plan --include=project
```

These arguments are available in any command that returns a response document.
The included and excluded fields requested from the API may differ based on the
output format, for example, for historic reasons, `packet projects get --yaml`
includes the details of all project members. In the table output format, no
member related fields are displayed and so `packet projects get` will exclude
the member resource.

Excluding fields needed for the table output will result in an error. Mixing
includes and excludes affecting the same top-level field is not supported.

## Reference

The full CLI documentation can be found [here](docs/packet.md) or by clicking the links below.

* [Device operations](docs/packet_device.md)
* [Facility operations](docs/packet_facilities.md)
* [IP operations](docs/packet_ip.md)
* [Operating system operations](docs/packet_operating-systems.md)
* [Organization operations](docs/packet_organization.md)
* [Plan operations](docs/packet_plan.md)
* [Project operations](docs/packet_project.md)
* [SSH key operations](docs/packet_ssh-key.md)
* [User operations](docs/packet_user.md)
* [Virtual network operations](docs/packet_virtual-network.md)
* [Volume operations](docs/packet_volume.md)
* [VPN operations](docs/packet_vpn.md)

## Example Syntax

### Create a device

```
packet device create --hostname [hostname] --plan [plan] --facility [facility_code] --operating-system [operating_system] --project-id [project_UUID]
```

### Create a volume

```
packet volume create --size [size_in_GB] --plan [plan_UUID] --project-id [project_UUID] --facility [facility_code]
```

### Attach a volume

```
packet volume attach --id [volume_UUID] --device-id [device_UUID]
```

### Get a device

```
packet device get --id [device_UUID]
```

### Get a volume

```
packet volume get --id [volume_UUID]
```

### List projects

```
packet project get
```

### Get a project

```
packet project get -i [project_UUID]
```

Details on all available commands can be found by visiting the reference [pages](docs/packet.md) or typing `packet [command] --help` for more information about the specific command.

## Support

For help with this package:

* Open up a GitHub issue [here](https://github.com/packethost/packet-cli/issues).
* Contact the [Equinix Metal Community Slack](http://slack.equinixmetal.net) or on Freenode IRC in the #equinixmetal channel.
