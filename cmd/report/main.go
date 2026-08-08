package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"time"
)

// BenchmarkResult holds parsed benchmark data
type BenchmarkResult struct {
	Library  string
	Feature  string
	Struct   string
	NsPerOp  float64
	BytesOp  int64
	AllocsOp int64
	Runs     int
}

// Key returns a unique key for grouping
func (b BenchmarkResult) Key() string {
	return b.Feature + "_" + b.Struct
}

func main() {
	results := parseBenchmarks(os.Stdin)
	generateMarkdown(results)
}

func parseBenchmarks(input *os.File) []BenchmarkResult {
	var results []BenchmarkResult
	scanner := bufio.NewScanner(input)

	// Regex to parse benchmark output lines
	// Format: Benchmark_Library_Feature_Struct-8  runs  ns/op  bytes/op  allocs/op
	// Example: Benchmark_Pedantigo_Validate_Simple-8  1234567  573.2 ns/op  100 B/op  10 allocs/op
	benchRegex := regexp.MustCompile(`^Benchmark_(\w+)_(\w+)_(\w+)-\d+\s+(\d+)\s+([\d.]+)\s+ns/op\s+(\d+)\s+B/op\s+(\d+)\s+allocs/op`)

	for scanner.Scan() {
		line := scanner.Text()
		matches := benchRegex.FindStringSubmatch(line)
		if matches == nil {
			continue
		}

		nsPerOp, _ := strconv.ParseFloat(matches[5], 64)
		bytesOp, _ := strconv.ParseInt(matches[6], 10, 64)
		allocsOp, _ := strconv.ParseInt(matches[7], 10, 64)
		runs, _ := strconv.Atoi(matches[4])

		results = append(results, BenchmarkResult{
			Library:  matches[1],
			Feature:  matches[2],
			Struct:   matches[3],
			NsPerOp:  nsPerOp,
			BytesOp:  bytesOp,
			AllocsOp: allocsOp,
			Runs:     runs,
		})
	}

	return results
}

func generateMarkdown(results []BenchmarkResult) {
	// Group results by feature
	byFeature := make(map[string][]BenchmarkResult)
	for _, r := range results {
		byFeature[r.Feature] = append(byFeature[r.Feature], r)
	}

	// Get ordered list of features. Unmarshal must sit immediately after
	// Marshal (single-call decode+validate is the natural counterpart to
	// validate+encode), not wherever plain alphabetical sort would place it.
	features := orderedFeatures(byFeature)

	// Print Docusaurus frontmatter (required for docs site)
	fmt.Println("---")
	fmt.Println("sidebar_position: 99")
	fmt.Println("title: Benchmarks")
	fmt.Println("---")
	fmt.Println()

	// Print header
	fmt.Println("# Benchmark Results")
	fmt.Println()
	fmt.Printf("Generated: %s\n", time.Now().UTC().Format("2006-01-02 15:04:05 UTC"))
	fmt.Println()
	fmt.Println("If you're interested in diving deeper, check out our [benchmark repository](https://github.com/smrutAI/pedantigo-benchmarks).")
	fmt.Println()

	// Print library notes
	printLibraryNotes()

	// Explain why New() is expensive and how to amortize it
	printUsageRecommendation()

	// Feature descriptions
	featureDesc := map[string]string{
		"Validate":     "Validate existing struct (no JSON parsing)",
		"JSONValidate": "JSON bytes → struct, then a separate validate step",
		"Unmarshal":    "JSON bytes → validated struct in a single call",
		"New":          "Validator creation overhead",
		"Schema":       "JSON Schema generation",
		"OpenAPI":      "OpenAPI-compatible schema generation",
		"Marshal":      "Validate + JSON marshal",
	}

	for _, feature := range features {
		featureResults := byFeature[feature]

		// Get all libraries and structs for this feature
		libraries := getUniqueLibraries(featureResults)
		structs := getUniqueStructs(featureResults)

		// Skip features with only skipped benchmarks
		if len(libraries) == 0 {
			continue
		}

		desc := featureDesc[feature]
		if desc == "" {
			desc = feature
		}

		fmt.Printf("## %s\n", feature)
		fmt.Printf("_%s_\n\n", desc)

		// Build table header
		header := "| Struct |"
		separator := "|--------|"
		for _, lib := range libraries {
			header += fmt.Sprintf(" %s |", lib)
			separator += "--------|"
		}
		fmt.Println(header)
		fmt.Println(separator)

		// Build table rows
		for _, s := range structs {
			row := fmt.Sprintf("| %s |", s)
			for _, lib := range libraries {
				result := findResult(featureResults, lib, s)
				if result != nil {
					row += fmt.Sprintf(" %s |", formatResult(result))
				} else {
					row += " unsupported |"
				}
			}
			fmt.Println(row)
		}
		fmt.Println()
	}

	// Print summary
	printSummary(results)
}

// orderedFeatures returns feature names in a fixed reading order rather than
// plain alphabetical sort, so Unmarshal (single-call decode+validate) sits
// immediately after Marshal (validate+encode) instead of wherever alphabetical
// order would otherwise place it. Any feature not in the preferred list
// (e.g. a future addition) is appended alphabetically at the end.
func orderedFeatures(byFeature map[string][]BenchmarkResult) []string {
	preferred := []string{"Validate", "JSONValidate", "Marshal", "Unmarshal", "New", "Schema", "OpenAPI"}

	seen := make(map[string]bool, len(byFeature))
	features := make([]string, 0, len(byFeature))
	for _, f := range preferred {
		if _, ok := byFeature[f]; ok {
			features = append(features, f)
			seen[f] = true
		}
	}

	var leftover []string
	for f := range byFeature {
		if !seen[f] {
			leftover = append(leftover, f)
		}
	}
	sort.Strings(leftover)

	return append(features, leftover...)
}

// allLibraries is the fixed list of all libraries to show in every table
var allLibraries = []string{"Pedantigo", "Playground", "Ozzo", "Huma", "Godantic", "Godasse"}

func getUniqueLibraries(results []BenchmarkResult) []string {
	// Always return all libraries for consistent tables
	return allLibraries
}

func getUniqueStructs(results []BenchmarkResult) []string {
	seen := make(map[string]bool)
	var structs []string
	// Preferred order
	order := []string{"Simple", "Complex", "Large", "Uncached", "Cached"}

	for _, r := range results {
		if !seen[r.Struct] {
			seen[r.Struct] = true
		}
	}

	// Add in preferred order
	for _, s := range order {
		if seen[s] {
			structs = append(structs, s)
			delete(seen, s)
		}
	}

	// Add any remaining
	for s := range seen {
		structs = append(structs, s)
	}

	return structs
}

func findResult(results []BenchmarkResult, library, structName string) *BenchmarkResult {
	for i := range results {
		if results[i].Library == library && results[i].Struct == structName {
			return &results[i]
		}
	}
	return nil
}

func formatResult(r *BenchmarkResult) string {
	ns := formatNs(r.NsPerOp)
	return fmt.Sprintf("%s (%d allocs)", ns, r.AllocsOp)
}

func formatNs(ns float64) string {
	if ns >= 1_000_000 {
		return fmt.Sprintf("%.2f ms", ns/1_000_000)
	}
	if ns >= 1_000 {
		return fmt.Sprintf("%.2f µs", ns/1_000)
	}
	return fmt.Sprintf("%.0f ns", ns)
}

// printUsageRecommendation explains why New()/deserializer construction is
// expensive (see the New section below) and how to amortize that cost, so the
// New numbers aren't misread as "this is how fast every call is."
func printUsageRecommendation() {
	fmt.Println("## Getting the Best Performance")
	fmt.Println()
	fmt.Println("`New[T]()` does the expensive work once: it walks the struct via reflection, " +
		"resolves every constraint tag, and builds an internal field-constraint cache " +
		"(plus the JSON field deserializers). That one-time cost is what the `New` section " +
		"below measures (microsecond range). Every other operation - `Validate`, `Unmarshal`, " +
		"`Marshal`, `Schema` - reuses that precomputed cache and runs in the hundreds-of-ns to " +
		"low-µs range, which is why those numbers consistently beat libraries that re-resolve " +
		"constraints or re-walk structs on every call.")
	fmt.Println()
	fmt.Println("This only pays off if the `*Validator[T]` returned by `New` is built once and " +
		"reused - not recreated per request. Two ways to do that:")
	fmt.Println()
	fmt.Println("**Module-level variable** (sufficient to call the validator directly):")
	fmt.Println()
	fmt.Println("```go")
	fmt.Println("var userValidator = validator.New[User]()")
	fmt.Println()
	fmt.Println("func handleCreateUser(body []byte) (*User, error) {")
	fmt.Println("\treturn userValidator.Unmarshal(body) // reuses the cached field constraints")
	fmt.Println("}")
	fmt.Println("```")
	fmt.Println()
	fmt.Println("**`Register`** (needed in addition, only if a framework integration - e.g. the " +
		"Echo Binder plugin, or `UnmarshalInto` - must find the validator for a type it only " +
		"knows via `reflect.Type` at runtime, not through your module-level variable):")
	fmt.Println()
	fmt.Println("```go")
	fmt.Println("var _ = validator.Register(validator.New[User]())")
	fmt.Println("```")
	fmt.Println()
	fmt.Println("`Register[T]` may be called exactly once per type - a second call for the same " +
		"type panics, by design. A type could have multiple differently-configured validators " +
		"(different `Options`), and pedantigo has no way to guess which one a framework plugin " +
		"should resolve to, so it refuses to silently pick one. Call `Register` from exactly one " +
		"package-level `var` declaration per type.")
	fmt.Println()
	fmt.Println("---")
	fmt.Println()
}

func printLibraryNotes() {
	fmt.Println("## Library Notes")
	fmt.Println()
	fmt.Println("### Feature Comparison")
	fmt.Println()
	fmt.Println("| Feature | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |")
	fmt.Println("|---------|-----------|------------|------|------|----------|---------|")
	fmt.Println("| Declarative constraints | ✅ tags | ✅ tags | ✅ rules | ✅ tags | ✅ methods | ❌ hand-written |")
	fmt.Println("| JSON Schema generation | ✅ | ❌ | ❌ | ✅ | ✅ | ❌ |")
	fmt.Println("| Default values | ✅ | ❌ | ❌ | ❌ | ✅ | ✅ |")
	fmt.Println("| Unmarshal + validate | ✅ | ❌ | ❌ | ✅ | ✅ | ✅ |")
	fmt.Println("| Validate existing struct | ✅ | ✅ | ✅ | ❌ | ✅ | ❌* |")
	fmt.Println()
	fmt.Println("_*Godasse requires hand-written `Validate()` methods_")
	fmt.Println()
	fmt.Println("### Library Descriptions")
	fmt.Println()
	fmt.Println("1. **Pedantigo** - Struct tag-based validation (`validate:\"required,email,min=5\"`). JSON Schema generation with caching.")
	fmt.Println()
	fmt.Println("2. **Playground** (go-playground/validator) - Struct tag-based validation. Rich constraint library, no JSON Schema.")
	fmt.Println()
	fmt.Println("3. **Ozzo** (ozzo-validation) - Rule builder API (`validation.Field(&u.Name, validation.Required, validation.Length(2,100))`). No struct tags.")
	fmt.Println()
	fmt.Println("4. **Huma** - OpenAPI-focused. Validates `map[string]any` against schemas, not structs directly.")
	fmt.Println()
	fmt.Println("5. **Godantic** - Method-based constraints (`FieldName() FieldOptions[T]`). JSON Schema, defaults, streaming partial JSON.")
	fmt.Println()
	fmt.Println("6. **Godasse** - Deserializer with `default:` tag. All constraint validation requires hand-written `Validate()` methods.")
	fmt.Println()
	fmt.Println("---")
	fmt.Println()
}

func printSummary(results []BenchmarkResult) {
	fmt.Println("---")
	fmt.Println()
	fmt.Println("## Summary")
	fmt.Println()

	// Print comparison for each key benchmark
	summaryBenchmarks := []struct {
		feature string
		struct_ string
		title   string
	}{
		{"Validate", "Simple", "Validate_Simple (struct validation)"},
		{"Validate", "Complex", "Validate_Complex (nested structs)"},
		{"JSONValidate", "Simple", "JSONValidate_Simple (JSON → struct, then validate)"},
		{"JSONValidate", "Complex", "JSONValidate_Complex (nested JSON)"},
		{"Unmarshal", "Simple", "Unmarshal_Simple (JSON → validated struct, single call)"},
		{"Unmarshal", "Complex", "Unmarshal_Complex (nested JSON, single call)"},
		{"Schema", "Uncached", "Schema_Uncached (first-time generation)"},
		{"Schema", "Cached", "Schema_Cached (cached lookup)"},
	}

	for _, bench := range summaryBenchmarks {
		printComparisonTable(results, bench.feature, bench.struct_, bench.title)
	}

	printLegend()
}

func printComparisonTable(results []BenchmarkResult, feature, struct_, title string) {
	// Find Pedantigo baseline
	var baseline *BenchmarkResult
	for i := range results {
		if results[i].Library == "Pedantigo" && results[i].Feature == feature && results[i].Struct == struct_ {
			baseline = &results[i]
			break
		}
	}

	if baseline == nil {
		return // Skip if no Pedantigo baseline
	}

	fmt.Printf("### %s\n", title)
	fmt.Println()
	fmt.Printf("| Library | ns/op | allocs | vs Pedantigo |\n")
	fmt.Printf("|---------|-------|--------|-------------|\n")

	for _, lib := range allLibraries {
		found := false
		for _, r := range results {
			if r.Library == lib && r.Feature == feature && r.Struct == struct_ {
				ratio := r.NsPerOp / baseline.NsPerOp
				var comparison string
				if lib == "Pedantigo" {
					comparison = "baseline"
				} else if ratio < 1.0 {
					comparison = fmt.Sprintf("%.2fx faster", 1.0/ratio)
				} else {
					comparison = fmt.Sprintf("%.2fx slower", ratio)
				}
				fmt.Printf("| %s | %s | %d | %s |\n", lib, formatNs(r.NsPerOp), r.AllocsOp, comparison)
				found = true
				break
			}
		}
		if !found {
			fmt.Printf("| %s | - | - | - |\n", lib)
		}
	}
	fmt.Println()
}

func printLegend() {
	fmt.Println("---")
	fmt.Println()
	fmt.Println("_Generated by pedantigo-benchmarks_")
	fmt.Println()
	fmt.Println("<details>")
	fmt.Println("<summary>Benchmark naming convention</summary>")
	fmt.Println()
	fmt.Println("```")
	fmt.Println("Benchmark_<Library>_<Feature>_<Struct>")
	fmt.Println()
	fmt.Println("Libraries: Pedantigo, Playground, Ozzo, Huma, Godantic, Godasse")
	fmt.Println("Features: Validate, JSONValidate, Marshal, Unmarshal, New, Schema, OpenAPI")
	fmt.Println("Structs: Simple (5 fields), Complex (nested), Large (20+ fields)")
	fmt.Println("```")
	fmt.Println("</details>")
}
