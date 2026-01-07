---
sidebar_position: 99
title: Benchmarks
---

# Benchmark Results

Generated: 2026-01-07 21:52:26 UTC

If you're interested in diving deeper, check out our [benchmark repository](https://github.com/smrutAI/pedantigo-benchmarks).

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

## JSONValidate
_JSON bytes → struct + validate_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Simple | 3.58 µs (19 allocs) | 4.38 µs (16 allocs) | unsupported | 3.64 µs (26 allocs) | unsupported | 5.63 µs (46 allocs) |
| Complex | 10.15 µs (39 allocs) | 11.43 µs (33 allocs) | unsupported | 10.64 µs (78 allocs) | unsupported | 18.08 µs (153 allocs) |

## Marshal
_Validate + JSON marshal_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Simple | 1.93 µs (11 allocs) | 2.77 µs (9 allocs) | unsupported | unsupported | unsupported | unsupported |

## New
_Validator creation overhead_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Simple | 12.31 µs (110 allocs) | 17.84 µs (187 allocs) | unsupported | 32.14 µs (255 allocs) | 27.88 µs (305 allocs) | 6.98 µs (72 allocs) |
| Complex | 29.02 µs (270 allocs) | unsupported | unsupported | 77.73 µs (515 allocs) | 7.88 µs (75 allocs) | 24.75 µs (243 allocs) |

## OpenAPI
_OpenAPI-compatible schema generation_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Uncached | 24.77 µs (204 allocs) | unsupported | unsupported | 32.56 µs (255 allocs) | unsupported | unsupported |
| Cached | 20 ns (0 allocs) | unsupported | unsupported | 650 ns (6 allocs) | unsupported | unsupported |

## Schema
_JSON Schema generation_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Uncached | 23.73 µs (202 allocs) | unsupported | unsupported | 32.69 µs (255 allocs) | unsupported | unsupported |
| Cached | 20 ns (0 allocs) | unsupported | unsupported | 653 ns (6 allocs) | unsupported | unsupported |

## Validate
_Validate existing struct (no JSON parsing)_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Simple | 1.46 µs (10 allocs) | 2.20 µs (7 allocs) | 13.10 µs (43 allocs) | unsupported | 6.21 µs (48 allocs) | unsupported |
| Complex | 2.28 µs (15 allocs) | 3.51 µs (9 allocs) | 12.62 µs (139 allocs) | unsupported | 14.30 µs (120 allocs) | unsupported |
| Large | 1.62 µs (22 allocs) | 1.90 µs (3 allocs) | 48.27 µs (254 allocs) | unsupported | 15.14 µs (126 allocs) | unsupported |

---

## Summary

### Validate_Simple (struct validation)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 1.46 µs | 10 | baseline |
| Playground | 2.20 µs | 7 | 1.51x slower |
| Ozzo | 13.10 µs | 43 | 8.97x slower |
| Huma | - | - | - |
| Godantic | 6.21 µs | 48 | 4.25x slower |
| Godasse | - | - | - |

### Validate_Complex (nested structs)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 2.28 µs | 15 | baseline |
| Playground | 3.51 µs | 9 | 1.54x slower |
| Ozzo | 12.62 µs | 139 | 5.53x slower |
| Huma | - | - | - |
| Godantic | 14.30 µs | 120 | 6.27x slower |
| Godasse | - | - | - |

### JSONValidate_Simple (JSON → struct + validate)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 3.58 µs | 19 | baseline |
| Playground | 4.38 µs | 16 | 1.22x slower |
| Ozzo | - | - | - |
| Huma | 3.64 µs | 26 | 1.01x slower |
| Godantic | - | - | - |
| Godasse | 5.63 µs | 46 | 1.57x slower |

### JSONValidate_Complex (nested JSON)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 10.15 µs | 39 | baseline |
| Playground | 11.43 µs | 33 | 1.13x slower |
| Ozzo | - | - | - |
| Huma | 10.64 µs | 78 | 1.05x slower |
| Godantic | - | - | - |
| Godasse | 18.08 µs | 153 | 1.78x slower |

### Schema_Uncached (first-time generation)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 23.73 µs | 202 | baseline |
| Playground | - | - | - |
| Ozzo | - | - | - |
| Huma | 32.69 µs | 255 | 1.38x slower |
| Godantic | - | - | - |
| Godasse | - | - | - |

### Schema_Cached (cached lookup)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 20 ns | 0 | baseline |
| Playground | - | - | - |
| Ozzo | - | - | - |
| Huma | 653 ns | 6 | 32.66x slower |
| Godantic | - | - | - |
| Godasse | - | - | - |

---

_Generated by pedantigo-benchmarks_

<details>
<summary>Benchmark naming convention</summary>

```
Benchmark_<Library>_<Feature>_<Struct>

Libraries: Pedantigo, Playground, Ozzo, Huma, Godantic, Godasse
Features: Validate, JSONValidate, New, Schema, OpenAPI, Marshal
Structs: Simple (5 fields), Complex (nested), Large (20+ fields)
```
</details>
