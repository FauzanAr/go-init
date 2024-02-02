# Project Title

Go-Init.

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Prerequisites](#prerequisites)
- [Getting Started](#getting-started)
- [Configuration](#configuration)
- [Contributing](#contributing)
- [License](#license)

## Introduction

This project was using SOLID principles that might be more easier to scale up or to add more component inside it and not depends on other packages. also the purpose of this project was to make initial project using golang more easier.

## Features

- Database Connection.
- Trace Logging
- REST API
- More features coming soon :)

## Prerequisites

What we're using:
- Go v1.19.13.
- MySql v8.0
- Echo V4
- Golang Migrate

## Getting Started

For now you can simply.

```bash
# Clone the repository
git clone https://github.com/yourusername/your-repo.git

# Change into the project directory
cd your-repo

# Running migration - See https://github.com/golang-migrate/migrate
migrate -database "{driver}://{username}:{password}@tcp({host})/{database_name}" -p {path} up

# Run the application
go run main.go

```

## Contirbuting

Thank you for considering contributing to this project! Contributions are welcome from everyone, regardless of your level of experience :).
