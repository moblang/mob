#!/bin/bash

set -e

BINARY_NAME="mob"
INSTALL_DIR="${HOME}/.local/bin"
REPO="moblang/mob"
VERSION="${VERSION:-v0.0.1}"

echo "🚀 Installing moblang..."

OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

case $ARCH in
    x86_64)
        ARCH="amd64"
        ;;
    aarch64|arm64)
        ARCH="arm64"
        ;;
    *)
        echo "❌ Unsupported architecture: $ARCH"
        echo "Currently supported: x86_64 (amd64), aarch64 (arm64)"
        exit 1
        ;;
esac

if [ "$OS" != "linux" ]; then
    echo "❌ moblang currently supports Linux only"
    exit 1
fi

mkdir -p "$INSTALL_DIR"

echo "📦 Downloading mob ${VERSION} for ${OS}/${ARCH}..."

DOWNLOAD_URL="https://github.com/${REPO}/releases/download/${VERSION}/mob-${OS}-${ARCH}"

if command -v wget >/dev/null 2>&1; then
    wget -qO "${INSTALL_DIR}/${BINARY_NAME}" "$DOWNLOAD_URL"
elif command -v curl >/dev/null 2>&1; then
    curl -fsSL "$DOWNLOAD_URL" -o "${INSTALL_DIR}/${BINARY_NAME}"
else
    echo "❌ Neither wget nor curl found. Please install one of them."
    exit 1
fi

chmod +x "${INSTALL_DIR}/${BINARY_NAME}"

echo ""
echo "✅ moblang installed successfully!"
echo ""
echo "Location: ${INSTALL_DIR}/${BINARY_NAME}"
echo ""
echo "Quick test:"
echo "  ${BINARY_NAME} version"
echo ""
echo "Make sure ${INSTALL_DIR} is in your PATH:"
echo "  export PATH=\"\$PATH:${INSTALL_DIR}\""
echo ""
