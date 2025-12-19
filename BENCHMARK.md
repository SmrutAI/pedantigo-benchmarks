---
sidebar_position: 99
title: Benchmarks
---

# Benchmark Results

Generated: 2025-12-19 19:38:37 UTC

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

## New
_Validator creation overhead_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Simple | 11.86 µs (110 allocs) | unsupported | unsupported | 30.67 µs (255 allocs) | 26.73 µs (305 allocs) | 6.67 µs (72 allocs) |
| Complex | 28.66 µs (270 allocs) | unsupported | unsupported | 74.70 µs (515 allocs) | 7.52 µs (75 allocs) | 23.33 µs (243 allocs) |

## OpenAPI
_OpenAPI-compatible schema generation_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Uncached | 23.59 µs (204 allocs) | unsupported | unsupported | 30.77 µs (255 allocs) | unsupported | unsupported |
| Cached | unsupported | unsupported | unsupported | 654 ns (6 allocs) | unsupported | unsupported |

## Schema
_JSON Schema generation_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Uncached | 22.94 µs (202 allocs) | unsupported | unsupported | 31.46 µs (255 allocs) | unsupported | unsupported |
| Cached | 20 ns (0 allocs) | unsupported | unsupported | 652 ns (6 allocs) | unsupported | unsupported |

## UnmarshalDirect
_json.Unmarshal + Validate (no intermediate map)_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Simple | 3.46 µs (19 allocs) | unsupported | unsupported | unsupported | unsupported | unsupported |
| Complex | 9.95 µs (39 allocs) | unsupported | unsupported | unsupported | unsupported | unsupported |

## UnmarshalMap
_JSON → map → validate (Pedantigo validates and outputs struct, Huma only validates the map)_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Simple | 5.01 µs (39 allocs) | unsupported | unsupported | 3.57 µs (26 allocs) | unsupported | 5.40 µs (46 allocs) |
| Complex | 18.28 µs (135 allocs) | unsupported | unsupported | 10.43 µs (78 allocs) | unsupported | 17.29 µs (153 allocs) |

## Validate
_Validate existing struct (no JSON parsing)_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Simple | 1.40 µs (10 allocs) | unsupported | 12.74 µs (43 allocs) | unsupported | 6.01 µs (48 allocs) | unsupported |
| Complex | 2.24 µs (15 allocs) | unsupported | 12.47 µs (139 allocs) | unsupported | 13.83 µs (120 allocs) | unsupported |
| Large | 1.59 µs (22 allocs) | unsupported | 47.60 µs (254 allocs) | unsupported | 14.73 µs (126 allocs) | unsupported |

---

## Summary

### Validate_Simple (struct validation)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 1.40 µs | 10 | baseline |
| Playground | - | - | - |
| Ozzo | 12.74 µs | 43 | 9.09x slower |
| Huma | - | - | - |
| Godantic | 6.01 µs | 48 | 4.29x slower |
| Godasse | - | - | - |

### Validate_Complex (nested structs)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 2.24 µs | 15 | baseline |
| Playground | - | - | - |
| Ozzo | 12.47 µs | 139 | 5.57x slower |
| Huma | - | - | - |
| Godantic | 13.83 µs | 120 | 6.17x slower |
| Godasse | - | - | - |

### UnmarshalMap_Simple (JSON → struct + validate)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 5.01 µs | 39 | baseline |
| Playground | - | - | - |
| Ozzo | - | - | - |
| Huma | 3.57 µs | 26 | 1.40x faster |
| Godantic | - | - | - |
| Godasse | 5.40 µs | 46 | 1.08x slower |

### UnmarshalMap_Complex (nested JSON)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 18.28 µs | 135 | baseline |
| Playground | - | - | - |
| Ozzo | - | - | - |
| Huma | 10.43 µs | 78 | 1.75x faster |
| Godantic | - | - | - |
| Godasse | 17.29 µs | 153 | 1.06x faster |

### Schema_Uncached (first-time generation)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 22.94 µs | 202 | baseline |
| Playground | - | - | - |
| Ozzo | - | - | - |
| Huma | 31.46 µs | 255 | 1.37x slower |
| Godantic | - | - | - |
| Godasse | - | - | - |

### Schema_Cached (cached lookup)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 20 ns | 0 | baseline |
| Playground | - | - | - |
| Ozzo | - | - | - |
| Huma | 652 ns | 6 | 32.29x slower |
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
