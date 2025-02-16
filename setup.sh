#!/bin/bash

set -e  # Exit immediately if a command exits with a non-zero status

# 1️⃣ Detect OS
echo "🔍 Detecting operating system..."
OS="$(uname)"
if [[ "$OS" == "Linux" ]]; then
    echo "✅ Running on Linux"
    PKG_MANAGER=$(command -v apt || command -v yum || command -v dnf || command -v pacman)
elif [[ "$OS" == "Darwin" ]]; then
    echo "✅ Running on macOS"
    PKG_MANAGER="brew"
else
    echo "❌ Unsupported OS: $OS"
    exit 1
fi

# 2️⃣ Check if Go is installed, otherwise install it
if ! command -v go &> /dev/null; then
    echo "⚠️ Go is not installed. Installing..."
    if [[ "$OS" == "Linux" ]]; then
        curl -OL https://go.dev/dl/go1.21.6.linux-amd64.tar.gz
        sudo tar -C /usr/local -xzf go1.21.6.linux-amd64.tar.gz
        echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
        source ~/.bashrc
    elif [[ "$OS" == "Darwin" ]]; then
        brew install go
    fi
else
    echo "✅ Go is already installed: $(go version)"
fi

# 3️⃣ Check if PostgreSQL is installed, otherwise install it
if ! command -v psql &> /dev/null; then
    echo "⚠️ PostgreSQL is not installed. Installing..."
    if [[ "$OS" == "Linux" ]]; then
        if [[ "$PKG_MANAGER" == *"apt"* ]]; then
            sudo apt update && sudo apt install -y postgresql postgresql-contrib
        elif [[ "$PKG_MANAGER" == *"yum"* || "$PKG_MANAGER" == *"dnf"* ]]; then
            sudo yum install -y postgresql-server postgresql-contrib
        elif [[ "$PKG_MANAGER" == *"pacman"* ]]; then
            sudo pacman -Sy postgresql
        fi
    elif [[ "$OS" == "Darwin" ]]; then
        brew install postgresql
    fi
else
    echo "✅ PostgreSQL is already installed: $(psql --version)"
fi
