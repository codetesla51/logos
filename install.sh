#!/bin/sh
set -e

# Logos install script
# Usage: curl -fsSL https://install.logos-lang.dev | sh

REPO="codetesla51/logos"
INSTALL_DIR="/usr/local/bin"
BINARY_NAME="lgs"

main() {
    platform=$(detect_platform)
    arch=$(detect_arch)
    
    if [ -z "$platform" ] || [ -z "$arch" ]; then
        echo "Error: Unsupported platform or architecture"
        exit 1
    fi

    target="${platform}-${arch}"
    echo "Installing Logos for ${target}..."

    # Get latest release tag
    latest=$(curl -fsSL "https://api.github.com/repos/${REPO}/releases/latest" | grep '"tag_name"' | cut -d'"' -f4)
    
    if [ -z "$latest" ]; then
        echo "Error: Could not fetch latest release"
        exit 1
    fi

    echo "Latest version: ${latest}"

    # Download URL
    url="https://github.com/${REPO}/releases/download/${latest}/logos-${target}.tar.gz"
    
    # Create temp directory
    tmp_dir=$(mktemp -d)
    trap 'rm -rf "$tmp_dir"' EXIT

    echo "Downloading from ${url}..."
    
    if ! curl -fsSL "$url" -o "${tmp_dir}/logos.tar.gz"; then
        echo "Error: Download failed"
        echo "Make sure a release exists for ${target}"
        exit 1
    fi

    # Extract
    tar -xzf "${tmp_dir}/logos.tar.gz" -C "$tmp_dir"

    # Install
    if [ -w "$INSTALL_DIR" ]; then
        mv "${tmp_dir}/${BINARY_NAME}" "${INSTALL_DIR}/${BINARY_NAME}"
    else
        echo "Installing to ${INSTALL_DIR} (requires sudo)..."
        sudo mv "${tmp_dir}/${BINARY_NAME}" "${INSTALL_DIR}/${BINARY_NAME}"
    fi

    chmod +x "${INSTALL_DIR}/${BINARY_NAME}"

    echo ""
    echo "Logos installed successfully!"
    echo "Run 'lgs' to start the REPL"
}

detect_platform() {
    case "$(uname -s)" in
        Linux*)  echo "linux" ;;
        Darwin*) echo "darwin" ;;
        *)       echo "" ;;
    esac
}

detect_arch() {
    case "$(uname -m)" in
        x86_64)  echo "amd64" ;;
        amd64)   echo "amd64" ;;
        arm64)   echo "arm64" ;;
        aarch64) echo "arm64" ;;
        *)       echo "" ;;
    esac
}

main
