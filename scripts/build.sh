#!/bin/bash

echo "===================================="
echo "  Theme Switcher - Build Script"
echo "===================================="
echo

# Get current directory
SCRIPT_DIR="$(dirname "$0")"
PROJECT_DIR="$SCRIPT_DIR/.."

# Create output directory
mkdir -p "$PROJECT_DIR/bin"

# Change to project root directory
cd "$PROJECT_DIR"
echo "Current directory: $(pwd)"
echo

# Run full build directly
echo "🛠️  Full Build Mode..."
echo

# 1. Clean old builds
echo "1. Cleaning old builds..."
rm -f bin/*
echo "✅ Clean completed"

# 2. Download dependencies
echo ""
echo "2. Downloading dependencies..."
go mod download
if [ $? -ne 0 ]; then
    echo "❌ Dependency download failed"
    exit 1
fi
echo "✅ Dependencies downloaded"

# 3. Build CLI version
echo ""
echo "3. Building CLI version..."
go build -o bin/theme-cli theme-cli.go
if [ $? -ne 0 ]; then
    echo "❌ CLI build failed"
    exit 1
fi
echo "✅ theme-cli built"

echo ""
echo "===================================="
echo "📦 Full build completed!"
echo ""
echo "🚀 Available versions:"
echo "  - bin/theme-cli      (CLI version)"
echo "===================================="
