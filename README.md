# DMS REST API 2.0

A complete rewrite of the DMS API in Golang. While the existing API (written in Python) was good enough, it was never
designed to scale.

This rewrite aims to solve the issue of handling concurrent requests while also being a much more stable and robust code
base to maintain.

TLDR; Version 1.0 was a personal project scaled too far. This is meant to be a production grade API server.

### Note regarding authentication

Due to the vague concern by MUJ officials, this API does not accept user-name and password for authentication of routes
other than login. An HTTP session is created upon login which is cached in a `Redis` backend. The caveat of this
approach is the need of re-auth every 20 minutes of inactivity.

# Configuration

The API server can be optionally configured with custom options. The configuration file `config.toml` lies in the root
directory.

The API supports following custom configurations:

```toml
# The development environment, change to `prod` when shipping
env = "dev"

[api]
# Use the optional redis server to improve performance?
# If `true` a redis server must be running on your machine
redis = false

# Turn this on if you want to monitor physical data store using a separate go routine
monitorDataStore = true

# Host and port to listen on
host = "localhost"
port = 8000
```

# How to run

DMS REST API requires [Golang], [Mux], [Zap] and [Viper] to run.

To install the dependencies and start the server, run:

```sh
$ cd dms-api
$ go mod tidy
$ go run main.go
```

# Usage

The API supports the following routes:

### Mess Menu:

Type: `GET` - URL: `/mess_menu` - Return type:

```swift
struct MessMenu {
    let last_updated_at: String?,
    let last_updated_meal: String?,
    let breakfast: [String]?,
    let lunch: [String]?,
    let high_tea: [String]?,
    let dinner: [String]?
}
```

### Faculty Contacts:

Type: `GET` - URL: `/contacts` - Return type:

```swift
struct FacultyContacts {
    let id: Int
    let name: String
    let designation: String
    let department: String
    let email: String
    let phone: String
    let image: String
}
```

# Tech

This API uses a number of open source projects to work properly:

* [Golang] - An open source programming language that makes it easy to build simple, reliable, and efficient software.
* [Mux] - A powerful HTTP router and URL matcher for building Go web servers with 🦍
* [Zap] - Blazing fast, structured, leveled logging in Go.
* [Viper] - Go configuration with fangs

# Contributing

Would like to help? You are more than welcome to do so. Just make sure that the pull request you submit **meets the
following requirements**:

- Code should be well documented in the Go Doc format
- Code should be properly formatted
- Relevant variable names should be used

Failure to comply to above will lead in your PR being trashed.

# Contributors

This API would never have been possible without the support of:

- [Sidharth Soni](https://github.com/sid-sun) - For everything.
- [black-dragon74](https://github.com/black-dragon74) - For writing this software and maintaining it.

***Another hobby project by Nick ;=)***

<!-- LINKS USED IN THIS MARKDOWN FILE -->

[Golang]: <https://golang.org/>

[Mux]: <https://github.com/gorilla/mux>

[Zap]: <https://github.com/uber-go/zap>

[Viper]: <https://github.com/spf13/viper>
