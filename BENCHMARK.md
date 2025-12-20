---
sidebar_position: 99
title: Benchmarks
---

# Benchmark Results

Generated: 2025-12-20 18:06:46 UTC

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
| Simple | 1.87 µs (11 allocs) | 2.69 µs (9 allocs) | unsupported | unsupported | unsupported | unsupported |

## New
_Validator creation overhead_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Simple | 11.73 µs (110 allocs) | 16.84 µs (187 allocs) | unsupported | 29.97 µs (255 allocs) | 26.52 µs (305 allocs) | 6.69 µs (72 allocs) |
| Complex | 28.39 µs (270 allocs) | unsupported | unsupported | 73.13 µs (515 allocs) | 7.50 µs (75 allocs) | 23.27 µs (243 allocs) |

## OpenAPI
_OpenAPI-compatible schema generation_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Uncached | 23.39 µs (204 allocs) | unsupported | unsupported | 30.26 µs (255 allocs) | unsupported | unsupported |
| Cached | 18 ns (0 allocs) | unsupported | unsupported | 644 ns (6 allocs) | unsupported | unsupported |

## Schema
_JSON Schema generation_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Uncached | 22.62 µs (202 allocs) | unsupported | unsupported | 29.93 µs (255 allocs) | unsupported | unsupported |
| Cached | 18 ns (0 allocs) | unsupported | unsupported | 634 ns (6 allocs) | unsupported | unsupported |

## UnmarshalDirect
_json.Unmarshal + Validate (no intermediate map)_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Simple | 3.42 µs (19 allocs) | 4.25 µs (16 allocs) | unsupported | unsupported | unsupported | unsupported |
| Complex | 9.81 µs (39 allocs) | 11.01 µs (33 allocs) | unsupported | unsupported | unsupported | unsupported |

## UnmarshalMap
_JSON → map → validate (Pedantigo validates and outputs struct, Huma only validates the map)_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Simple | 4.99 µs (39 allocs) | unsupported | unsupported | 3.53 µs (26 allocs) | unsupported | 5.40 µs (46 allocs) |
| Complex | 18.27 µs (135 allocs) | unsupported | unsupported | 10.35 µs (78 allocs) | unsupported | 17.24 µs (153 allocs) |

## Validate
_Validate existing struct (no JSON parsing)_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Simple | 1.39 µs (10 allocs) | 2.15 µs (7 allocs) | 12.77 µs (43 allocs) | unsupported | 5.99 µs (48 allocs) | unsupported |
| Complex | 2.27 µs (15 allocs) | 3.41 µs (9 allocs) | 12.29 µs (139 allocs) | unsupported | 13.70 µs (120 allocs) | unsupported |
| Large | 1.59 µs (22 allocs) | 1.85 µs (3 allocs) | 46.82 µs (254 allocs) | unsupported | 14.68 µs (126 allocs) | unsupported |

---

## Summary

### Validate_Simple (struct validation)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 1.39 µs | 10 | baseline |
| Playground | 2.15 µs | 7 | 1.54x slower |
| Ozzo | 12.77 µs | 43 | 9.18x slower |
| Huma | - | - | - |
| Godantic | 5.99 µs | 48 | 4.31x slower |
| Godasse | - | - | - |

### Validate_Complex (nested structs)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 2.27 µs | 15 | baseline |
| Playground | 3.41 µs | 9 | 1.51x slower |
| Ozzo | 12.29 µs | 139 | 5.42x slower |
| Huma | - | - | - |
| Godantic | 13.70 µs | 120 | 6.04x slower |
| Godasse | - | - | - |

### UnmarshalMap_Simple (JSON → struct + validate)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 4.99 µs | 39 | baseline |
| Playground | - | - | - |
| Ozzo | - | - | - |
| Huma | 3.53 µs | 26 | 1.41x faster |
| Godantic | - | - | - |
| Godasse | 5.40 µs | 46 | 1.08x slower |

### UnmarshalMap_Complex (nested JSON)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 18.27 µs | 135 | baseline |
| Playground | - | - | - |
| Ozzo | - | - | - |
| Huma | 10.35 µs | 78 | 1.76x faster |
| Godantic | - | - | - |
| Godasse | 17.24 µs | 153 | 1.06x faster |

### Schema_Uncached (first-time generation)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 22.62 µs | 202 | baseline |
| Playground | - | - | - |
| Ozzo | - | - | - |
| Huma | 29.93 µs | 255 | 1.32x slower |
| Godantic | - | - | - |
| Godasse | - | - | - |

### Schema_Cached (cached lookup)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 18 ns | 0 | baseline |
| Playground | - | - | - |
| Ozzo | - | - | - |
| Huma | 634 ns | 6 | 34.68x slower |
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
