# packet-cli

[![GitHub release](https://img.shields.io/github/release/packethost/packet-cli/all.svg?style=flat-square)](https://github.com/packethost/packet-cli/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/packethost/packet-cli)](https://goreportcard.com/report/github.com/packethost/packet-cli)
[![Slack](https://slack.packet.com/badge.svg)](https://slack.packet.com)
[![Twitter Follow](https://img.shields.io/twitter/follow/packethost.svg?style=social&label=Follow)](https://twitter.com/intent/follow?screen_name=packethost)

## Table of Contents

* [Packet Command Line Interface](#packet-command-line-interface)
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

## Packet Command Line Interface

[Packet](https://www.packet.net/) provides an API-driven bare metal platform that combines the automation experience of the cloud with the benefits of physical, bare metal servers.

The Packet CLI wraps the [Packet Go SDK](https://github.com/packethost/packngo) allowing interaction with Packet platform from a command-line interface.

## Requirements

* Packet authentication token.
* Packet CLI [binaries](https://github.com/packethost/packet-cli/releases).

## Supported Platforms

The [Packet CLI binaries](https://github.com/packethost/packet-cli/releases) are available for Linux, Windows, and Mac OS X for various architectures including ARM on Linux.

## Installation

Download the appropriate Packet CLI binaries for your platform to the desired location and rename it to `packet`. If the directory is not already in your `PATH` environment variable, then it will need to be added.

### Configuring your Path

<details>
  <summary>Configure your path in Linux</summary>

## Linux

If you plan to run the Packet CLI in a shell on Linux and placed the binary in /home/YOUR-USER-NAME/packet-cli/, then type the following into your terminal:

```
export PATH=$PATH:/home/$USER/packet-cli
```

You can view the current value of $PATH by running:

```
echo $PATH
```
</details>

<details>
  <summary>Configure your path in Mac OS/X</summary>

### Mac OS X

If you plan to run the Packet CLI in a shell on a Mac, download the `darwin` binary and placed the it in /Users/YOUR-USER-NAME/packet-cli/, then type the following into your terminal.

```
export PATH=$PATH:/Users/$USER/packet-cli
```

You can view the current value of $PATH by running:

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

If you plan to run the Packet CLI in PowerShell on Windows and placed the binary in c:\packet-cli, then type the following into PowerShell:

```
$env:Path += ";c:\packet-cli"
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

The Packet authentication token can be stored in the `$PACKET_TOKEN` environment variable or in JSON or YAML configuration files. The configuration file path can be overridden with the `--config` flag.

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

After installing Packet CLI, verify the installation by executing `packet` or `packet.exe`. You should see the default output:


```bash
$ packet
Command line interface for Packet Host

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
* Contact the [Packet Community Slack](http://slack.packet.net) or on Freenode IRC in the #packethost channel.
