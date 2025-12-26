---
sidebar_position: 99
title: Benchmarks
---

# Benchmark Results

Generated: 2025-12-26 10:57:03 UTC

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
| Simple | 3.21 µs (19 allocs) | 3.86 µs (16 allocs) | unsupported | 3.30 µs (26 allocs) | unsupported | 5.26 µs (46 allocs) |
| Complex | 8.90 µs (39 allocs) | 9.75 µs (33 allocs) | unsupported | 9.39 µs (78 allocs) | unsupported | 16.55 µs (153 allocs) |

## Marshal
_Validate + JSON marshal_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Simple | 1.80 µs (11 allocs) | 2.43 µs (9 allocs) | unsupported | unsupported | unsupported | unsupported |

## New
_Validator creation overhead_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Simple | 11.65 µs (110 allocs) | 15.28 µs (187 allocs) | unsupported | 29.63 µs (255 allocs) | 25.65 µs (305 allocs) | 6.56 µs (72 allocs) |
| Complex | 27.73 µs (270 allocs) | unsupported | unsupported | 73.58 µs (515 allocs) | 7.38 µs (75 allocs) | 23.03 µs (243 allocs) |

## OpenAPI
_OpenAPI-compatible schema generation_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Uncached | 22.79 µs (204 allocs) | unsupported | unsupported | 29.56 µs (255 allocs) | unsupported | unsupported |
| Cached | 24 ns (0 allocs) | unsupported | unsupported | 625 ns (6 allocs) | unsupported | unsupported |

## Schema
_JSON Schema generation_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Uncached | 22.20 µs (202 allocs) | unsupported | unsupported | 29.62 µs (255 allocs) | unsupported | unsupported |
| Cached | 24 ns (0 allocs) | unsupported | unsupported | 640 ns (6 allocs) | unsupported | unsupported |

## Validate
_Validate existing struct (no JSON parsing)_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Simple | 1.34 µs (10 allocs) | 1.94 µs (7 allocs) | 11.41 µs (43 allocs) | unsupported | 6.01 µs (48 allocs) | unsupported |
| Complex | 2.16 µs (15 allocs) | 3.10 µs (9 allocs) | 11.91 µs (139 allocs) | unsupported | 13.76 µs (120 allocs) | unsupported |
| Large | 1.48 µs (22 allocs) | 1.72 µs (3 allocs) | 42.91 µs (254 allocs) | unsupported | 14.68 µs (126 allocs) | unsupported |

---

## Summary

### Validate_Simple (struct validation)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 1.34 µs | 10 | baseline |
| Playground | 1.94 µs | 7 | 1.45x slower |
| Ozzo | 11.41 µs | 43 | 8.52x slower |
| Huma | - | - | - |
| Godantic | 6.01 µs | 48 | 4.49x slower |
| Godasse | - | - | - |

### Validate_Complex (nested structs)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 2.16 µs | 15 | baseline |
| Playground | 3.10 µs | 9 | 1.44x slower |
| Ozzo | 11.91 µs | 139 | 5.52x slower |
| Huma | - | - | - |
| Godantic | 13.76 µs | 120 | 6.38x slower |
| Godasse | - | - | - |

### JSONValidate_Simple (JSON → struct + validate)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 3.21 µs | 19 | baseline |
| Playground | 3.86 µs | 16 | 1.20x slower |
| Ozzo | - | - | - |
| Huma | 3.30 µs | 26 | 1.03x slower |
| Godantic | - | - | - |
| Godasse | 5.26 µs | 46 | 1.64x slower |

### JSONValidate_Complex (nested JSON)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 8.90 µs | 39 | baseline |
| Playground | 9.75 µs | 33 | 1.10x slower |
| Ozzo | - | - | - |
| Huma | 9.39 µs | 78 | 1.05x slower |
| Godantic | - | - | - |
| Godasse | 16.55 µs | 153 | 1.86x slower |

### Schema_Uncached (first-time generation)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 22.20 µs | 202 | baseline |
| Playground | - | - | - |
| Ozzo | - | - | - |
| Huma | 29.62 µs | 255 | 1.33x slower |
| Godantic | - | - | - |
| Godasse | - | - | - |

### Schema_Cached (cached lookup)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 24 ns | 0 | baseline |
| Playground | - | - | - |
| Ozzo | - | - | - |
| Huma | 640 ns | 6 | 26.54x slower |
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
