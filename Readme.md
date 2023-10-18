# IP Address CLI Tool

This is a command-line tool for tracking and gathering information about IP addresses. It allows you to retrieve data about a specific IP address, including details like city, region, country, location, timezone, and postal code.

## Installation

To install the IP Address CLI Tool, follow these steps:

1. Make sure you have Go (Golang) installed on your system.

2. Use `go get` to install the CLI tool and its dependencies:

```shell
go get github.com/nitintf/ipaddress-cli
```

3. Once the installation is complete, you can use the `ipaddress` command in your terminal.

## Usage

### Get Your IP Address

Running the `ipaddress` command without any arguments will display both your public and private IP addresses with public address info.

```shell
ipaddress
```

### Track IP Address

Use the `track` command to gather information about a specific IP address. Here's the syntax:

```shell
ipaddress track [IP Address]
```
