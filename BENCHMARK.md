---
sidebar_position: 99
title: Benchmarks
---

# Benchmark Results

Generated: 2026-06-21 16:30:32 UTC

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
| Simple | 3.23 µs (19 allocs) | 4.10 µs (16 allocs) | unsupported | 3.32 µs (26 allocs) | unsupported | 4.99 µs (46 allocs) |
| Complex | 8.96 µs (39 allocs) | 10.21 µs (33 allocs) | unsupported | 9.87 µs (78 allocs) | unsupported | 16.25 µs (153 allocs) |

## Marshal
_Validate + JSON marshal_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Simple | 1.75 µs (11 allocs) | 2.63 µs (9 allocs) | unsupported | unsupported | unsupported | unsupported |

## New
_Validator creation overhead_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Simple | 11.24 µs (110 allocs) | 15.07 µs (187 allocs) | unsupported | 29.52 µs (255 allocs) | 25.55 µs (305 allocs) | 6.26 µs (72 allocs) |
| Complex | 26.75 µs (270 allocs) | unsupported | unsupported | 73.43 µs (515 allocs) | 7.12 µs (75 allocs) | 21.83 µs (243 allocs) |

## OpenAPI
_OpenAPI-compatible schema generation_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Uncached | 22.18 µs (204 allocs) | unsupported | unsupported | 29.81 µs (255 allocs) | unsupported | unsupported |
| Cached | 19 ns (0 allocs) | unsupported | unsupported | 587 ns (6 allocs) | unsupported | unsupported |

## Schema
_JSON Schema generation_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Uncached | 21.36 µs (202 allocs) | unsupported | unsupported | 29.79 µs (255 allocs) | unsupported | unsupported |
| Cached | 19 ns (0 allocs) | unsupported | unsupported | 596 ns (6 allocs) | unsupported | unsupported |

## Validate
_Validate existing struct (no JSON parsing)_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Simple | 1.33 µs (10 allocs) | 2.10 µs (7 allocs) | 11.82 µs (43 allocs) | unsupported | 5.44 µs (48 allocs) | unsupported |
| Complex | 2.20 µs (15 allocs) | 3.38 µs (9 allocs) | 11.40 µs (139 allocs) | unsupported | 12.47 µs (120 allocs) | unsupported |
| Large | 1.43 µs (22 allocs) | 1.66 µs (3 allocs) | 42.77 µs (254 allocs) | unsupported | 13.39 µs (126 allocs) | unsupported |

---

## Summary

### Validate_Simple (struct validation)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 1.33 µs | 10 | baseline |
| Playground | 2.10 µs | 7 | 1.58x slower |
| Ozzo | 11.82 µs | 43 | 8.91x slower |
| Huma | - | - | - |
| Godantic | 5.44 µs | 48 | 4.10x slower |
| Godasse | - | - | - |

### Validate_Complex (nested structs)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 2.20 µs | 15 | baseline |
| Playground | 3.38 µs | 9 | 1.54x slower |
| Ozzo | 11.40 µs | 139 | 5.19x slower |
| Huma | - | - | - |
| Godantic | 12.47 µs | 120 | 5.68x slower |
| Godasse | - | - | - |

### JSONValidate_Simple (JSON → struct + validate)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 3.23 µs | 19 | baseline |
| Playground | 4.10 µs | 16 | 1.27x slower |
| Ozzo | - | - | - |
| Huma | 3.32 µs | 26 | 1.03x slower |
| Godantic | - | - | - |
| Godasse | 4.99 µs | 46 | 1.55x slower |

### JSONValidate_Complex (nested JSON)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 8.96 µs | 39 | baseline |
| Playground | 10.21 µs | 33 | 1.14x slower |
| Ozzo | - | - | - |
| Huma | 9.87 µs | 78 | 1.10x slower |
| Godantic | - | - | - |
| Godasse | 16.25 µs | 153 | 1.81x slower |

### Schema_Uncached (first-time generation)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 21.36 µs | 202 | baseline |
| Playground | - | - | - |
| Ozzo | - | - | - |
| Huma | 29.79 µs | 255 | 1.40x slower |
| Godantic | - | - | - |
| Godasse | - | - | - |

### Schema_Cached (cached lookup)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 19 ns | 0 | baseline |
| Playground | - | - | - |
| Ozzo | - | - | - |
| Huma | 596 ns | 6 | 31.16x slower |
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
