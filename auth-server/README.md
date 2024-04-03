# Auth Server in Go

This repository contains a simple authentication server written in Go. It serves as a microservice dedicated to handling authentication requests.

## Features

- User registration
- User login
- Token generation
- Token verification

## Prerequisites

Before running the server, ensure you have the following installed:

- Docker
- Docker-compose

The server will start running on localhost:8080 by default. You can change the port in the configuration file if needed.

Use the following endpoints for authentication:

/register: Register a new user.
/login: Log in with existing credentials.
/logout: Log out and invalidate the token.
/refresh: Refresh the token.
/verify: Verify the token.

### Configuration

You can configure the server by modifying the settings in config/config.go. This includes database connection details, JWT secret, and server settings.