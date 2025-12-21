# Pedantigo Benchmarks

Performance benchmarks comparing [Pedantigo](https://github.com/SmrutAI/pedantigo) against other Go validation libraries.

**View the formatted results at [pedantigo.dev/docs/benchmarks](https://pedantigo.dev/docs/benchmarks)**

## Libraries Compared

- **Pedantigo** - Tag-based validation with JSON Schema generation
- **Playground** (go-playground/validator) - Tag-based validation
- **Ozzo** (ozzo-validation) - Rule builder API
- **Huma** - OpenAPI-focused validation
- **Godantic** - Method-based constraints
- **Godasse** - Deserializer with defaults

## Running Locally

```bash
# Setup (clones pedantigo to third_party/)
make setup

# Run benchmarks (default: 50 iterations)
make bench

# Run with custom iteration count
make bench COUNT=7
```

## How It Works

1. `scripts/run-benchmarks.sh` runs Go benchmarks across all libraries
2. `cmd/report/main.go` generates `BENCHMARK.md` for the docs website
3. `cmd/report-pr/main.go` generates compact reports for PR comments
4. GitHub Actions automatically updates benchmarks on push to main

## License

MIT