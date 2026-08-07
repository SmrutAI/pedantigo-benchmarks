---
sidebar_position: 99
title: Benchmarks
---

# Benchmark Results

Generated: 2026-08-07 22:38:05 UTC

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
| Simple | 1.83 µs (19 allocs) | 2.38 µs (16 allocs) | unsupported | 2.05 µs (26 allocs) | unsupported | 2.98 µs (46 allocs) |
| Complex | 4.97 µs (39 allocs) | 5.70 µs (33 allocs) | unsupported | 5.64 µs (78 allocs) | unsupported | 9.57 µs (153 allocs) |

## Marshal
_Validate + JSON marshal_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Simple | 1.02 µs (11 allocs) | 1.45 µs (9 allocs) | unsupported | unsupported | unsupported | unsupported |

## New
_Validator creation overhead_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Simple | 6.55 µs (110 allocs) | 8.43 µs (187 allocs) | unsupported | 17.47 µs (255 allocs) | 15.25 µs (305 allocs) | 3.91 µs (72 allocs) |
| Complex | 16.19 µs (270 allocs) | unsupported | unsupported | 43.53 µs (515 allocs) | 4.42 µs (75 allocs) | 13.62 µs (243 allocs) |

## OpenAPI
_OpenAPI-compatible schema generation_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Uncached | 13.01 µs (204 allocs) | unsupported | unsupported | 18.34 µs (255 allocs) | unsupported | unsupported |
| Cached | 17 ns (0 allocs) | unsupported | unsupported | 354 ns (6 allocs) | unsupported | unsupported |

## Schema
_JSON Schema generation_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Uncached | 12.48 µs (202 allocs) | unsupported | unsupported | 17.44 µs (255 allocs) | unsupported | unsupported |
| Cached | 17 ns (0 allocs) | unsupported | unsupported | 355 ns (6 allocs) | unsupported | unsupported |

## Validate
_Validate existing struct (no JSON parsing)_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Simple | 772 ns (10 allocs) | 1.09 µs (7 allocs) | 6.66 µs (43 allocs) | unsupported | 3.31 µs (48 allocs) | unsupported |
| Complex | 1.22 µs (15 allocs) | 1.79 µs (9 allocs) | 6.77 µs (139 allocs) | unsupported | 7.77 µs (120 allocs) | unsupported |
| Large | 828 ns (22 allocs) | 939 ns (3 allocs) | 24.53 µs (254 allocs) | unsupported | 8.06 µs (126 allocs) | unsupported |

---

## Summary

### Validate_Simple (struct validation)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 772 ns | 10 | baseline |
| Playground | 1.09 µs | 7 | 1.41x slower |
| Ozzo | 6.66 µs | 43 | 8.63x slower |
| Huma | - | - | - |
| Godantic | 3.31 µs | 48 | 4.28x slower |
| Godasse | - | - | - |

### Validate_Complex (nested structs)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 1.22 µs | 15 | baseline |
| Playground | 1.79 µs | 9 | 1.47x slower |
| Ozzo | 6.77 µs | 139 | 5.56x slower |
| Huma | - | - | - |
| Godantic | 7.77 µs | 120 | 6.38x slower |
| Godasse | - | - | - |

### JSONValidate_Simple (JSON → struct + validate)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 1.83 µs | 19 | baseline |
| Playground | 2.38 µs | 16 | 1.30x slower |
| Ozzo | - | - | - |
| Huma | 2.05 µs | 26 | 1.12x slower |
| Godantic | - | - | - |
| Godasse | 2.98 µs | 46 | 1.63x slower |

### JSONValidate_Complex (nested JSON)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 4.97 µs | 39 | baseline |
| Playground | 5.70 µs | 33 | 1.15x slower |
| Ozzo | - | - | - |
| Huma | 5.64 µs | 78 | 1.13x slower |
| Godantic | - | - | - |
| Godasse | 9.57 µs | 153 | 1.92x slower |

### Schema_Uncached (first-time generation)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 12.48 µs | 202 | baseline |
| Playground | - | - | - |
| Ozzo | - | - | - |
| Huma | 17.44 µs | 255 | 1.40x slower |
| Godantic | - | - | - |
| Godasse | - | - | - |

### Schema_Cached (cached lookup)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 17 ns | 0 | baseline |
| Playground | - | - | - |
| Ozzo | - | - | - |
| Huma | 355 ns | 6 | 21.28x slower |
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
