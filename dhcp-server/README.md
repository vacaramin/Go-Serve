# DHCP Server

This repository contains a DHCP (Dynamic Host Configuration Protocol) server implementation written in Go.

## Overview

DHCP is a network protocol that automatically assigns IP addresses and other network configuration information to devices on a network. This server is designed to automate this process, making it easier to manage and configure network resources.

## Features

- Automatic assignment of IP addresses
- Configuration of subnet masks, default gateways, and DNS servers
- Lease management for IP addresses
- Support for both IPv4 and IPv6

## Prerequisites

Before running the DHCP server, ensure you have the following:

- Docker
- Docker-Compose
- Understanding of networking concepts and DHCP protocol

### Configuration

You can configure various DHCP server settings such as IP range, lease duration, subnet mask, default gateway, and DNS servers in the config/config.go file.
The server will start listening for DHCP requests on the specified interface and port.
