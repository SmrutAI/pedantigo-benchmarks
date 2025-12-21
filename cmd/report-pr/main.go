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

func main() {
	results := parseBenchmarks(os.Stdin)
	generatePRReport(results)
}

func parseBenchmarks(input *os.File) []BenchmarkResult {
	var results []BenchmarkResult
	scanner := bufio.NewScanner(input)

	// Regex to parse benchmark output lines
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

var allLibraries = []string{"Pedantigo", "Playground", "Ozzo", "Huma", "Godantic", "Godasse"}

func generatePRReport(results []BenchmarkResult) {
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
	fmt.Println("## Benchmark Results")
	fmt.Println()
	fmt.Printf("_Generated: %s_\n", time.Now().UTC().Format("2006-01-02 15:04:05 UTC"))
	fmt.Println()

	// Feature descriptions (short)
	featureDesc := map[string]string{
		"Validate":     "Validate struct",
		"JSONValidate": "JSON validate",
		"New":          "Validator creation",
		"Schema":       "Schema generation",
		"OpenAPI":      "OpenAPI schema",
		"Marshal":      "Validate + marshal",
	}

	for _, feature := range features {
		featureResults := byFeature[feature]
		structs := getUniqueStructs(featureResults)

		if len(structs) == 0 {
			continue
		}

		desc := featureDesc[feature]
		if desc == "" {
			desc = feature
		}

		fmt.Printf("### %s\n", desc)
		fmt.Println()

		// Build table header
		header := "| Struct |"
		separator := "|:-------|"
		for _, lib := range allLibraries {
			header += fmt.Sprintf(" %s |", lib)
			separator += ":-------:|"
		}
		fmt.Println(header)
		fmt.Println(separator)

		// Build table rows
		for _, s := range structs {
			row := fmt.Sprintf("| %s |", s)
			for _, lib := range allLibraries {
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

	// Print quick comparison
	printQuickComparison(results)
}

func getUniqueStructs(results []BenchmarkResult) []string {
	seen := make(map[string]bool)
	var structs []string
	order := []string{"Simple", "Complex", "Large", "Uncached", "Cached"}

	for _, r := range results {
		seen[r.Struct] = true
	}

	for _, s := range order {
		if seen[s] {
			structs = append(structs, s)
			delete(seen, s)
		}
	}

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
	return fmt.Sprintf("%s / %d", formatNs(r.NsPerOp), r.AllocsOp)
}

func formatNs(ns float64) string {
	if ns >= 1_000_000 {
		return fmt.Sprintf("%.1fms", ns/1_000_000)
	}
	if ns >= 1_000 {
		return fmt.Sprintf("%.1fµs", ns/1_000)
	}
	return fmt.Sprintf("%.0fns", ns)
}

func printQuickComparison(results []BenchmarkResult) {
	fmt.Println("---")
	fmt.Println()
	fmt.Println("**Legend:** `time / allocs` • `-` = not supported")
}