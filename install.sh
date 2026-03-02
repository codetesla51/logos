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

    # Download with a small spinner/loader
    curl -fsSL "$url" -o "${tmp_dir}/logos.tar.gz" &
    dl_pid=$!

    spinner_watch "$dl_pid" "Downloading..."

    wait "$dl_pid"
    dl_rc=$?
    if [ "$dl_rc" -ne 0 ]; then
        echo "\nError: Download failed"
        echo "Make sure a release exists for ${target}"
        exit 1
    fi

    # Extract with loader
    (tar -xzf "${tmp_dir}/logos.tar.gz" -C "$tmp_dir") &
    tar_pid=$!
    spinner_watch "$tar_pid" "Extracting..."
    wait "$tar_pid"
    tar_rc=$?
    if [ "$tar_rc" -ne 0 ]; then
        echo "\nError: Failed to extract archive"
        exit 1
    fi

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

spinner_watch() {
    # $1 = pid to watch
    # $2 = message prefix
    pid="$1"
    msg="$2"
    spin='-\\|/'
    i=0
    printf "%s " "$msg"
    while kill -0 "$pid" 2>/dev/null; do
        printf "\b%s" "${spin:i%${#spin}:1}"
        sleep 0.1
        i=$((i+1))
    done
    printf "\b"  # remove spinner char
}

main
