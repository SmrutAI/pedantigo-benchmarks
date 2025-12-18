#!/bin/bash
# Benchmark Report Generator
# Usage: go test -bench=. -benchmem ./... 2>&1 | ./report.sh > report.csv
#
# Parses Go benchmark output and generates a CSV report.
# Only includes benchmarks that actually ran (skipped benchmarks are excluded).

echo "Library,Feature,ns/op,B/op,allocs/op"

while read -r line; do
    # Match lines like: Benchmark_Pedantigo_Validate_Simple-12    2087373    573.2 ns/op    277 B/op    10 allocs/op
    # The pattern handles:
    #   - Benchmark function name with CPU count suffix (-12)
    #   - Iteration count
    #   - ns/op (can be decimal like 573.2)
    #   - B/op
    #   - allocs/op
    if [[ $line =~ ^Benchmark_([^_]+)_([^-]+)-[0-9]+[[:space:]]+[0-9]+[[:space:]]+([0-9.]+)[[:space:]]ns/op[[:space:]]+([0-9]+)[[:space:]]B/op[[:space:]]+([0-9]+)[[:space:]]allocs/op ]]; then
        library="${BASH_REMATCH[1]}"
        feature="${BASH_REMATCH[2]}"
        ns_op="${BASH_REMATCH[3]}"
        b_op="${BASH_REMATCH[4]}"
        allocs="${BASH_REMATCH[5]}"
        echo "${library},${feature},${ns_op},${b_op},${allocs}"
    fi
done
