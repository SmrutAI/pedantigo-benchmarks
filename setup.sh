#!/bin/bash
set -e

THIRD_PARTY="third_party"

# Create third_party directory
mkdir -p "$THIRD_PARTY"

# Clone Pedantigo only if not already present
if [ -d "$THIRD_PARTY/pedantigo" ]; then
    echo "pedantigo already cloned at $THIRD_PARTY/pedantigo (delete to re-clone)"
else
    echo "Cloning pedantigo..."
    git clone --depth 1 https://github.com/SmrutAI/pedantigo.git "$THIRD_PARTY/pedantigo"
fi

# Only run go mod tidy if vendor doesn't exist
if [ ! -d "vendor" ]; then
    echo "Running go mod tidy..."
    go mod tidy
else
    echo "Using vendored dependencies"
fi

echo "Setup complete!"
