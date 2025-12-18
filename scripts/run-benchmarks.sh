#!/bin/bash
# Run benchmarks and generate report
# Usage: ./scripts/run-benchmarks.sh [count] [pedantigo_path]
# Output: benchmark-output.txt and BENCHMARK.md
#
# Examples:
#   ./scripts/run-benchmarks.sh 50                    # CI mode (uses cloned pedantigo-src)
#   ./scripts/run-benchmarks.sh 50 ../Pedantigo       # Local development

set -e

COUNT="${1:-50}"
PEDANTIGO_PATH="${2:-./pedantigo-src}"

# Set up replace directive for pedantigo
if [ -d "$PEDANTIGO_PATH" ]; then
    echo "Using pedantigo from: $PEDANTIGO_PATH"
    go mod edit -replace "github.com/SmrutAI/pedantigo=$PEDANTIGO_PATH"
    go mod tidy
else
    echo "Error: pedantigo not found at $PEDANTIGO_PATH"
    echo "Usage: $0 [count] [pedantigo_path]"
    exit 1
fi

echo "Running benchmarks with -count=$COUNT..."
go test -bench=. -benchmem -count="$COUNT" ./... 2>&1 | tee benchmark-output.txt

echo "Generating report..."
go run ./cmd/report/main.go < benchmark-output.txt > BENCHMARK.md

echo "Done! Generated:"
echo "  - benchmark-output.txt (raw output)"
echo "  - BENCHMARK.md (formatted report)"
