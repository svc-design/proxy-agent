# About

Proxy Agent - Cross-Platform Proxy Management
Proxy Agent is a cross-platform application built with Go (Fyne) that allows users to manage system-wide or browser-specific proxy settings. It supports Linux, Windows, and macOS, providing an intuitive UI with features like:

- Server Subscription: Import and export proxy configurations from server subscriptions.
- One-click Proxy: Quickly enable or disable proxies with a single click.
- System Tray Support: View and control proxy status (active/inactive/error) from the system tray.
- Auto Updates: Includes regular updates to integrate the latest core versions.
- Real-Time Configuration: Hot-reload proxy configurations without restarting the app.

Easily manage your proxy settings with a clean, user-friendly interface and robust functionality. Perfect for users who need quick access to network configurations across different platforms.

# Build 

## Install Go for MacOS

brew install go gtk+3 pkg-config git 
go version
export GO111MODULE=on
export GOPROXY=https://goproxy.cn
go mod init fyne-project
go get fyne.io/fyne/v2
go mod tidy
go run main.go
go build -o proxy-agent


# Technical Stack

- Programming Language:
Go (Golang): Utilizes Go for its simplicity and performance.
Go Official Website

- GUI Framework:
Fyne: A cross-platform GUI toolkit for building applications on Linux, Windows, and macOS.
Fyne GitHub Repository

- Proxy Protocol:
XTLS VLESS: Implements the XTLS VLESS protocol for secure proxy management.
XTLS VLESS GitHub Repository

- CI/CD:
GitHub Actions: Automates building, testing, and artifact management.
GitHub Actions Documentation

# Features Todo List:
- Cross-platform support (Linux, Windows, macOS)
- User-friendly interface for managing proxy settings
- Support for importing/exporting server configurations
- System tray notifications for proxy status
- Automatic updates for core dependencies

