# Pedantigo Benchmarks Makefile
# Usage: make bench      # Run full benchmarks (count=50)
#        make bench-quick # Quick test (count=5)
#        make report     # Generate report from existing output

PEDANTIGO_PATH ?= ../Pedantigo
COUNT ?= 50

.PHONY: bench bench-quick report clean help

# Run full benchmarks with count=50
bench:
	./scripts/run-benchmarks.sh $(COUNT) $(PEDANTIGO_PATH)

# Quick test with count=5
bench-quick:
	./scripts/run-benchmarks.sh 5 $(PEDANTIGO_PATH)

# Generate report from existing benchmark-output.txt
report:
	go run ./cmd/report/main.go < benchmark-output.txt > BENCHMARK.md
	@echo "Generated BENCHMARK.md"

# Clean generated files
clean:
	rm -f benchmark-output.txt BENCHMARK.md
	go mod edit -dropreplace github.com/SmrutAI/pedantigo 2>/dev/null || true

# Show help
help:
	@echo "Pedantigo Benchmarks"
	@echo ""
	@echo "Usage:"
	@echo "  make bench          Run full benchmarks (count=50)"
	@echo "  make bench-quick    Quick test (count=5)"
	@echo "  make report         Generate report from existing output"
	@echo "  make clean          Remove generated files"
	@echo ""
	@echo "Variables:"
	@echo "  PEDANTIGO_PATH      Path to pedantigo (default: ../Pedantigo)"
	@echo "  COUNT               Benchmark iterations (default: 50)"
	@echo ""
	@echo "Examples:"
	@echo "  make bench COUNT=10"
	@echo "  make bench PEDANTIGO_PATH=/path/to/pedantigo"
