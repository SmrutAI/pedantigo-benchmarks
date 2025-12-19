# Pedantigo Benchmarks Makefile
# Usage: make bench      # Run full benchmarks (count=20)
#        make bench-quick # Quick test (count=3)
#        make setup      # Clone pedantigo to third_party/
#        make vendor     # Vendor dependencies
#        make report     # Generate report from existing output

COUNT ?= 20

.PHONY: bench bench-quick setup vendor report clean help

# Setup: clone pedantigo
setup:
	./setup.sh

# Vendor dependencies (run after setup)
vendor: setup
	go mod vendor
	@echo "Vendor complete. Commit vendor/ to repo."

# Run full benchmarks
bench: setup
	./scripts/run-benchmarks.sh $(COUNT)

# Quick test with count=3
bench-quick: setup
	./scripts/run-benchmarks.sh 3

# Generate report from existing benchmark-output.txt
report:
	go run ./cmd/report/main.go < benchmark-output.txt > BENCHMARK.md
	@echo "Generated BENCHMARK.md"

# Clean generated files
clean:
	rm -f benchmark-output.txt BENCHMARK.md
	rm -rf third_party/

# Show help
help:
	@echo "Pedantigo Benchmarks"
	@echo ""
	@echo "Usage:"
	@echo "  make setup          Clone pedantigo to third_party/"
	@echo "  make vendor         Vendor all dependencies"
	@echo "  make bench          Run full benchmarks (count=20)"
	@echo "  make bench-quick    Quick test (count=3)"
	@echo "  make report         Generate report from existing output"
	@echo "  make clean          Remove generated files"
	@echo ""
	@echo "Variables:"
	@echo "  COUNT               Benchmark iterations (default: 20)"
