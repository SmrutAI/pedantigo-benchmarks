#!/bin/bash
# Run benchmarks and generate report
# Usage: ./scripts/run-benchmarks.sh [count]
# Output: benchmark-output.txt and BENCHMARK.md

set -e

COUNT="${1:-50}"

# Run setup to clone pedantigo
echo "Running setup..."
./setup.sh

# Use vendor if it exists
MOD_FLAG=""
if [ -d "vendor" ]; then
    MOD_FLAG="-mod=vendor"
fi

echo "Running benchmarks with -count=$COUNT..."
go test $MOD_FLAG -bench=. -benchmem -count="$COUNT" ./... 2>&1 | tee benchmark-output.txt

echo "Generating report..."
go run $MOD_FLAG ./cmd/report/main.go < benchmark-output.txt > BENCHMARK.md

echo "Done! Generated:"
echo "  - benchmark-output.txt (raw output)"
echo "  - BENCHMARK.md (formatted report)"
