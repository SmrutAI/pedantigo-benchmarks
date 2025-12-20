---
sidebar_position: 99
title: Benchmarks
---

# Benchmark Results

Generated: 2025-12-20 07:31:05 UTC

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
| Simple | 1.89 µs (11 allocs) | 2.68 µs (9 allocs) | unsupported | unsupported | unsupported | unsupported |

## New
_Validator creation overhead_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Simple | 11.76 µs (110 allocs) | 16.00 µs (187 allocs) | unsupported | 30.03 µs (255 allocs) | 26.63 µs (305 allocs) | 6.67 µs (72 allocs) |
| Complex | 28.70 µs (270 allocs) | unsupported | unsupported | 73.34 µs (515 allocs) | 7.53 µs (75 allocs) | 23.24 µs (243 allocs) |

## OpenAPI
_OpenAPI-compatible schema generation_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Uncached | 23.46 µs (204 allocs) | unsupported | unsupported | 30.29 µs (255 allocs) | unsupported | unsupported |
| Cached | 18 ns (0 allocs) | unsupported | unsupported | 642 ns (6 allocs) | unsupported | unsupported |

## Schema
_JSON Schema generation_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Uncached | 22.64 µs (202 allocs) | unsupported | unsupported | 30.16 µs (255 allocs) | unsupported | unsupported |
| Cached | 18 ns (0 allocs) | unsupported | unsupported | 642 ns (6 allocs) | unsupported | unsupported |

## UnmarshalDirect
_json.Unmarshal + Validate (no intermediate map)_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Simple | 3.49 µs (19 allocs) | 4.24 µs (16 allocs) | unsupported | unsupported | unsupported | unsupported |
| Complex | 9.81 µs (39 allocs) | 11.05 µs (33 allocs) | unsupported | unsupported | unsupported | unsupported |

## UnmarshalMap
_JSON → map → validate (Pedantigo validates and outputs struct, Huma only validates the map)_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Simple | 4.97 µs (39 allocs) | unsupported | unsupported | 3.54 µs (26 allocs) | unsupported | 5.38 µs (46 allocs) |
| Complex | 18.20 µs (135 allocs) | unsupported | unsupported | 10.33 µs (78 allocs) | unsupported | 17.17 µs (153 allocs) |

## Validate
_Validate existing struct (no JSON parsing)_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Simple | 1.41 µs (10 allocs) | 2.15 µs (7 allocs) | 12.69 µs (43 allocs) | unsupported | 6.03 µs (48 allocs) | unsupported |
| Complex | 2.25 µs (15 allocs) | 3.42 µs (9 allocs) | 12.30 µs (139 allocs) | unsupported | 13.91 µs (120 allocs) | unsupported |
| Large | 1.58 µs (22 allocs) | 1.85 µs (3 allocs) | 46.67 µs (254 allocs) | unsupported | 14.73 µs (126 allocs) | unsupported |

---

## Summary

### Validate_Simple (struct validation)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 1.41 µs | 10 | baseline |
| Playground | 2.15 µs | 7 | 1.52x slower |
| Ozzo | 12.69 µs | 43 | 8.99x slower |
| Huma | - | - | - |
| Godantic | 6.03 µs | 48 | 4.28x slower |
| Godasse | - | - | - |

### Validate_Complex (nested structs)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 2.25 µs | 15 | baseline |
| Playground | 3.42 µs | 9 | 1.52x slower |
| Ozzo | 12.30 µs | 139 | 5.48x slower |
| Huma | - | - | - |
| Godantic | 13.91 µs | 120 | 6.20x slower |
| Godasse | - | - | - |

### UnmarshalMap_Simple (JSON → struct + validate)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 4.97 µs | 39 | baseline |
| Playground | - | - | - |
| Ozzo | - | - | - |
| Huma | 3.54 µs | 26 | 1.41x faster |
| Godantic | - | - | - |
| Godasse | 5.38 µs | 46 | 1.08x slower |

### UnmarshalMap_Complex (nested JSON)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 18.20 µs | 135 | baseline |
| Playground | - | - | - |
| Ozzo | - | - | - |
| Huma | 10.33 µs | 78 | 1.76x faster |
| Godantic | - | - | - |
| Godasse | 17.17 µs | 153 | 1.06x faster |

### Schema_Uncached (first-time generation)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 22.64 µs | 202 | baseline |
| Playground | - | - | - |
| Ozzo | - | - | - |
| Huma | 30.16 µs | 255 | 1.33x slower |
| Godantic | - | - | - |
| Godasse | - | - | - |

### Schema_Cached (cached lookup)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 18 ns | 0 | baseline |
| Playground | - | - | - |
| Ozzo | - | - | - |
| Huma | 642 ns | 6 | 34.90x slower |
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
