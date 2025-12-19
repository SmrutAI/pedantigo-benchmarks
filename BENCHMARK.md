---
sidebar_position: 99
title: Benchmarks
---

# Benchmark Results

Generated: 2025-12-19 13:23:15 UTC

## Library Notes

### Feature Comparison

| Feature | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|---------|-----------|------------|------|------|----------|---------|
| Declarative constraints | ✅ tags | ✅ tags | ✅ rules | ✅ tags | ✅ methods | ❌ hand-written |
| JSON Schema generation | ✅ | ❌ | ❌ | ✅ | ✅ | ❌ |
| Default values | ✅ | ❌ | ❌ | ❌ | ✅ | ✅ |
| Unmarshal + validate | ✅ | ❌ | ❌ | ✅ | ✅ | ✅ |
| Validate existing struct | ✅ | ✅ | ✅ | ❌ | ✅ | ❌* |

_*Godasse requires hand-written `Validate()` methods_

### Library Descriptions

1. **Pedantigo** - Struct tag-based validation (`validate:"required,email,min=5"`). JSON Schema generation with caching.

2. **Playground** (go-playground/validator) - Struct tag-based validation. Rich constraint library, no JSON Schema.

3. **Ozzo** (ozzo-validation) - Rule builder API (`validation.Field(&u.Name, validation.Required, validation.Length(2,100))`). No struct tags.

4. **Huma** - OpenAPI-focused. Validates `map[string]any` against schemas, not structs directly.

5. **Godantic** - Method-based constraints (`FieldName() FieldOptions[T]`). JSON Schema, defaults, streaming partial JSON.

6. **Godasse** - Deserializer with `default:` tag. All constraint validation requires hand-written `Validate()` methods.

---

## Marshal
_Validate + JSON marshal_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Simple | 765 ns (11 allocs) | 1.04 µs (9 allocs) | unsupported | unsupported | unsupported | unsupported |

## New
_Validator creation overhead_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Simple | 5.44 µs (110 allocs) | 7.64 µs (183 allocs) | unsupported | 12.20 µs (247 allocs) | 10.50 µs (291 allocs) | 2.68 µs (72 allocs) |
| Complex | 13.02 µs (270 allocs) | unsupported | unsupported | 32.02 µs (515 allocs) | 2.87 µs (65 allocs) | 9.37 µs (243 allocs) |

## OpenAPI
_OpenAPI-compatible schema generation_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Uncached | 10.24 µs (204 allocs) | unsupported | unsupported | 12.40 µs (247 allocs) | unsupported | unsupported |
| Cached | 8 ns (0 allocs) | unsupported | unsupported | 281 ns (6 allocs) | unsupported | unsupported |

## Schema
_JSON Schema generation_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Uncached | 9.80 µs (202 allocs) | unsupported | unsupported | 12.83 µs (247 allocs) | unsupported | unsupported |
| Cached | 8 ns (0 allocs) | unsupported | unsupported | 292 ns (6 allocs) | unsupported | unsupported |

## UnmarshalDirect
_json.Unmarshal + Validate (no intermediate map)_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Simple | 1.42 µs (19 allocs) | 1.70 µs (16 allocs) | unsupported | unsupported | unsupported | unsupported |
| Complex | 4.47 µs (39 allocs) | 4.76 µs (33 allocs) | unsupported | unsupported | unsupported | unsupported |

## UnmarshalMap
_JSON → map → validate (Pedantigo validates and outputs struct, Huma only validates the map)_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Simple | 2.11 µs (39 allocs) | unsupported | unsupported | 1.49 µs (26 allocs) | unsupported | 2.09 µs (46 allocs) |
| Complex | 8.57 µs (135 allocs) | unsupported | unsupported | 5.21 µs (78 allocs) | unsupported | 7.47 µs (153 allocs) |

## Validate
_Validate existing struct (no JSON parsing)_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Simple | 604 ns (10 allocs) | 864 ns (7 allocs) | 6.06 µs (43 allocs) | unsupported | 2.37 µs (48 allocs) | unsupported |
| Complex | 1.04 µs (15 allocs) | 1.41 µs (9 allocs) | 5.65 µs (139 allocs) | unsupported | 5.49 µs (120 allocs) | unsupported |
| Large | 714 ns (22 allocs) | 841 ns (3 allocs) | 21.23 µs (254 allocs) | unsupported | 5.74 µs (126 allocs) | unsupported |

---

## Summary

### Validate_Simple (struct validation)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 604 ns | 10 | baseline |
| Playground | 864 ns | 7 | 1.43x slower |
| Ozzo | 6.06 µs | 43 | 10.03x slower |
| Huma | - | - | - |
| Godantic | 2.37 µs | 48 | 3.92x slower |
| Godasse | - | - | - |

### Validate_Complex (nested structs)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 1.04 µs | 15 | baseline |
| Playground | 1.41 µs | 9 | 1.35x slower |
| Ozzo | 5.65 µs | 139 | 5.41x slower |
| Huma | - | - | - |
| Godantic | 5.49 µs | 120 | 5.26x slower |
| Godasse | - | - | - |

### UnmarshalMap_Simple (JSON → struct + validate)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 2.11 µs | 39 | baseline |
| Playground | - | - | - |
| Ozzo | - | - | - |
| Huma | 1.49 µs | 26 | 1.41x faster |
| Godantic | - | - | - |
| Godasse | 2.09 µs | 46 | 1.01x faster |

### UnmarshalMap_Complex (nested JSON)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 8.57 µs | 135 | baseline |
| Playground | - | - | - |
| Ozzo | - | - | - |
| Huma | 5.21 µs | 78 | 1.64x faster |
| Godantic | - | - | - |
| Godasse | 7.47 µs | 153 | 1.15x faster |

### Schema_Uncached (first-time generation)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 9.80 µs | 202 | baseline |
| Playground | - | - | - |
| Ozzo | - | - | - |
| Huma | 12.83 µs | 247 | 1.31x slower |
| Godantic | - | - | - |
| Godasse | - | - | - |

### Schema_Cached (cached lookup)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 8 ns | 0 | baseline |
| Playground | - | - | - |
| Ozzo | - | - | - |
| Huma | 292 ns | 6 | 35.28x slower |
| Godantic | - | - | - |
| Godasse | - | - | - |

---

_Generated by pedantigo-benchmarks_

<details>
<summary>Benchmark naming convention</summary>

```
Benchmark_<Library>_<Feature>_<Struct>

Libraries: Pedantigo, Playground, Ozzo, Huma, Godantic, Godasse
Features: Validate, UnmarshalMap, UnmarshalDirect, New, Schema, OpenAPI, Marshal
Structs: Simple (5 fields), Complex (nested), Large (20+ fields)
```
</details>
