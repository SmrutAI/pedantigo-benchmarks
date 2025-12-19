#!/bin/bash
# Run benchmarks and generate report
# Usage: ./scripts/run-benchmarks.sh [count]
# Output: benchmark-output.txt and BENCHMARK.md

set -e

COUNT="${1:-50}"

# Run setup to clone pedantigo
echo "Running setup..."
./setup.sh

echo "Running benchmarks with -count=$COUNT..."
go test -bench=. -benchmem -count="$COUNT" ./... 2>&1 | tee benchmark-output.txt

echo "Generating report..."
go run ./cmd/report/main.go < benchmark-output.txt > BENCHMARK.md

echo "Done! Generated:"
echo "  - benchmark-output.txt (raw output)"
echo "  - BENCHMARK.md (formatted report)"
