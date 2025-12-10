#!/bin/sh
set -e

# Detect OS and architecture
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

case $ARCH in
    x86_64) ARCH="amd64" ;;
    aarch64|arm64) ARCH="arm64" ;;
esac

REPO="shivgitcode/gimi"
BINARY="gimi"

echo "Detecting platform: ${OS}_${ARCH}"

# Get latest release version
echo "Fetching latest release..."
VERSION=$(curl -s "https://api.github.com/repos/$REPO/releases/latest" | grep '"tag_name"' | sed -E 's/.*"([^"]+)".*/\1/')

if [ -z "$VERSION" ]; then
    echo "Error: Could not fetch latest version"
    exit 1
fi

echo "Installing $BINARY $VERSION..."

# Download binary
TMP_DIR=$(mktemp -d)
cd "$TMP_DIR"

FILE="${BINARY}_${VERSION#v}_${OS}_${ARCH}.tar.gz"
URL="https://github.com/$REPO/releases/download/$VERSION/$FILE"

echo "Downloading from $URL"
curl -sL "$URL" -o "$FILE"

# Extract
tar -xzf "$FILE"

# Install
INSTALL_DIR="${INSTALL_DIR:-/usr/local/bin}"

if [ -w "$INSTALL_DIR" ]; then
    mv "$BINARY" "$INSTALL_DIR/"
    chmod +x "$INSTALL_DIR/$BINARY"
else
    echo "Installing to $INSTALL_DIR (requires sudo)..."
    sudo mv "$BINARY" "$INSTALL_DIR/"
    sudo chmod +x "$INSTALL_DIR/$BINARY"
fi

# Cleanup
cd -
rm -rf "$TMP_DIR"

echo "✓ $BINARY installed successfully!"
echo ""
echo "Run 'gimi --version' to verify"