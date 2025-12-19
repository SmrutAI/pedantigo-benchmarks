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

	// Get sorted list of features
	features := make([]string, 0, len(byFeature))
	for f := range byFeature {
		features = append(features, f)
	}
	sort.Strings(features)

	// Print header
	fmt.Println("# Benchmark Results")
	fmt.Println()
	fmt.Printf("Generated: %s\n", time.Now().UTC().Format("2006-01-02 15:04:05 UTC"))
	fmt.Println()

	// Feature descriptions
	featureDesc := map[string]string{
		"Validate":        "Validate existing struct (no JSON parsing)",
		"UnmarshalMap":    "JSON → map → struct + validate",
		"UnmarshalDirect": "json.Unmarshal + Validate (bypass, no map conversion)",
		"New":             "Validator creation overhead",
		"Schema":          "JSON Schema generation",
		"Marshal":         "Validate + JSON marshal",
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
					row += " - |"
				}
			}
			fmt.Println(row)
		}
		fmt.Println()
	}

	// Print summary
	printSummary(results)
}

func getUniqueLibraries(results []BenchmarkResult) []string {
	seen := make(map[string]bool)
	var libraries []string
	// Preferred order
	order := []string{"Pedantigo", "Playground", "Ozzo", "Huma", "godantic", "godasse"}

	for _, r := range results {
		if !seen[r.Library] {
			seen[r.Library] = true
		}
	}

	// Add in preferred order
	for _, lib := range order {
		if seen[lib] {
			libraries = append(libraries, lib)
			delete(seen, lib)
		}
	}

	// Add any remaining
	for lib := range seen {
		libraries = append(libraries, lib)
	}

	return libraries
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

func printSummary(results []BenchmarkResult) {
	fmt.Println("---")
	fmt.Println()
	fmt.Println("## Summary")
	fmt.Println()

	// Group by library
	byLibrary := make(map[string][]BenchmarkResult)
	for _, r := range results {
		byLibrary[r.Library] = append(byLibrary[r.Library], r)
	}

	// Find Pedantigo baseline for Validate_Simple
	var pedantigoValidateSimple *BenchmarkResult
	for _, r := range results {
		if r.Library == "Pedantigo" && r.Feature == "Validate" && r.Struct == "Simple" {
			pedantigoValidateSimple = &r
			break
		}
	}

	if pedantigoValidateSimple == nil {
		fmt.Println("_No Pedantigo baseline found_")
		return
	}

	fmt.Println("### Validate_Simple Comparison")
	fmt.Println()
	fmt.Printf("| Library | ns/op | vs Pedantigo |\n")
	fmt.Printf("|---------|-------|-------------|\n")

	for _, lib := range []string{"Pedantigo", "Playground", "Ozzo", "Huma", "godantic", "godasse"} {
		for _, r := range results {
			if r.Library == lib && r.Feature == "Validate" && r.Struct == "Simple" {
				ratio := r.NsPerOp / pedantigoValidateSimple.NsPerOp
				var comparison string
				if ratio < 1.0 {
					comparison = fmt.Sprintf("%.2fx faster", 1.0/ratio)
				} else if ratio > 1.0 {
					comparison = fmt.Sprintf("%.2fx slower", ratio)
				} else {
					comparison = "baseline"
				}
				fmt.Printf("| %s | %s | %s |\n", lib, formatNs(r.NsPerOp), comparison)
				break
			}
		}
	}
	fmt.Println()

	// Print allocations summary
	fmt.Println("### Allocations (Validate_Simple)")
	fmt.Println()
	fmt.Printf("| Library | B/op | allocs/op |\n")
	fmt.Printf("|---------|------|----------|\n")

	for _, lib := range []string{"Pedantigo", "Playground", "Ozzo", "Huma", "godantic", "godasse"} {
		for _, r := range results {
			if r.Library == lib && r.Feature == "Validate" && r.Struct == "Simple" {
				fmt.Printf("| %s | %d | %d |\n", lib, r.BytesOp, r.AllocsOp)
				break
			}
		}
	}
	fmt.Println()

	// Legend
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
	fmt.Println("Libraries: Pedantigo, Playground, Ozzo, Huma, godantic, godasse")
	fmt.Println("Features: Validate, UnmarshalMap, UnmarshalDirect, New, Schema, Marshal")
	fmt.Println("Structs: Simple (5 fields), Complex (nested), Large (20+ fields)")
	fmt.Println("```")
	fmt.Println("</details>")
}
