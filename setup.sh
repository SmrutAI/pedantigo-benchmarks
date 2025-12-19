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

echo "Running go mod tidy..."
go mod tidy

echo "Setup complete!"
