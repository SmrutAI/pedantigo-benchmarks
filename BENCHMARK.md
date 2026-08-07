---
sidebar_position: 99
title: Benchmarks
---

# Benchmark Results

Generated: 2026-08-07 20:57:26 UTC

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
| Simple | 3.48 µs (19 allocs) | 4.29 µs (16 allocs) | unsupported | 3.58 µs (26 allocs) | unsupported | 5.59 µs (46 allocs) |
| Complex | 9.85 µs (39 allocs) | 11.18 µs (33 allocs) | unsupported | 10.46 µs (78 allocs) | unsupported | 18.06 µs (153 allocs) |

## Marshal
_Validate + JSON marshal_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Simple | 1.94 µs (11 allocs) | 2.78 µs (9 allocs) | unsupported | unsupported | unsupported | unsupported |

## New
_Validator creation overhead_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Simple | 12.38 µs (110 allocs) | 17.33 µs (187 allocs) | unsupported | 31.88 µs (255 allocs) | 28.66 µs (305 allocs) | 6.92 µs (72 allocs) |
| Complex | 29.79 µs (270 allocs) | unsupported | unsupported | 78.77 µs (515 allocs) | 8.06 µs (75 allocs) | 24.97 µs (243 allocs) |

## OpenAPI
_OpenAPI-compatible schema generation_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Uncached | 25.80 µs (204 allocs) | unsupported | unsupported | 32.48 µs (255 allocs) | unsupported | unsupported |
| Cached | 19 ns (0 allocs) | unsupported | unsupported | 659 ns (6 allocs) | unsupported | unsupported |

## Schema
_JSON Schema generation_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Uncached | 23.85 µs (202 allocs) | unsupported | unsupported | 31.84 µs (255 allocs) | unsupported | unsupported |
| Cached | 18 ns (0 allocs) | unsupported | unsupported | 649 ns (6 allocs) | unsupported | unsupported |

## Validate
_Validate existing struct (no JSON parsing)_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Simple | 1.49 µs (10 allocs) | 2.23 µs (7 allocs) | 13.08 µs (43 allocs) | unsupported | 6.23 µs (48 allocs) | unsupported |
| Complex | 2.34 µs (15 allocs) | 3.64 µs (9 allocs) | 12.62 µs (139 allocs) | unsupported | 14.40 µs (120 allocs) | unsupported |
| Large | 1.61 µs (22 allocs) | 1.91 µs (3 allocs) | 48.31 µs (254 allocs) | unsupported | 15.26 µs (126 allocs) | unsupported |

---

## Summary

### Validate_Simple (struct validation)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 1.49 µs | 10 | baseline |
| Playground | 2.23 µs | 7 | 1.50x slower |
| Ozzo | 13.08 µs | 43 | 8.80x slower |
| Huma | - | - | - |
| Godantic | 6.23 µs | 48 | 4.19x slower |
| Godasse | - | - | - |

### Validate_Complex (nested structs)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 2.34 µs | 15 | baseline |
| Playground | 3.64 µs | 9 | 1.56x slower |
| Ozzo | 12.62 µs | 139 | 5.40x slower |
| Huma | - | - | - |
| Godantic | 14.40 µs | 120 | 6.17x slower |
| Godasse | - | - | - |

### JSONValidate_Simple (JSON → struct + validate)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 3.48 µs | 19 | baseline |
| Playground | 4.29 µs | 16 | 1.23x slower |
| Ozzo | - | - | - |
| Huma | 3.58 µs | 26 | 1.03x slower |
| Godantic | - | - | - |
| Godasse | 5.59 µs | 46 | 1.61x slower |

### JSONValidate_Complex (nested JSON)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 9.85 µs | 39 | baseline |
| Playground | 11.18 µs | 33 | 1.13x slower |
| Ozzo | - | - | - |
| Huma | 10.46 µs | 78 | 1.06x slower |
| Godantic | - | - | - |
| Godasse | 18.06 µs | 153 | 1.83x slower |

### Schema_Uncached (first-time generation)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 23.85 µs | 202 | baseline |
| Playground | - | - | - |
| Ozzo | - | - | - |
| Huma | 31.84 µs | 255 | 1.34x slower |
| Godantic | - | - | - |
| Godasse | - | - | - |

### Schema_Cached (cached lookup)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 18 ns | 0 | baseline |
| Playground | - | - | - |
| Ozzo | - | - | - |
| Huma | 649 ns | 6 | 35.24x slower |
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
