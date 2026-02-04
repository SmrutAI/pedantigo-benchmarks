---
sidebar_position: 99
title: Benchmarks
---

# Benchmark Results

Generated: 2026-02-04 10:48:27 UTC

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
| Simple | 3.40 µs (19 allocs) | 4.17 µs (16 allocs) | unsupported | 3.46 µs (26 allocs) | unsupported | 5.32 µs (46 allocs) |
| Complex | 9.64 µs (39 allocs) | 10.86 µs (33 allocs) | unsupported | 10.19 µs (78 allocs) | unsupported | 17.28 µs (153 allocs) |

## Marshal
_Validate + JSON marshal_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Simple | 1.87 µs (11 allocs) | 2.69 µs (9 allocs) | unsupported | unsupported | unsupported | unsupported |

## New
_Validator creation overhead_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Simple | 11.87 µs (110 allocs) | 16.87 µs (187 allocs) | unsupported | 30.49 µs (255 allocs) | 26.59 µs (305 allocs) | 6.63 µs (72 allocs) |
| Complex | 28.71 µs (270 allocs) | unsupported | unsupported | 75.21 µs (515 allocs) | 7.36 µs (75 allocs) | 23.12 µs (243 allocs) |

## OpenAPI
_OpenAPI-compatible schema generation_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Uncached | 23.97 µs (204 allocs) | unsupported | unsupported | 30.83 µs (255 allocs) | unsupported | unsupported |
| Cached | 18 ns (0 allocs) | unsupported | unsupported | 630 ns (6 allocs) | unsupported | unsupported |

## Schema
_JSON Schema generation_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Uncached | 23.27 µs (202 allocs) | unsupported | unsupported | 30.73 µs (255 allocs) | unsupported | unsupported |
| Cached | 18 ns (0 allocs) | unsupported | unsupported | 632 ns (6 allocs) | unsupported | unsupported |

## Validate
_Validate existing struct (no JSON parsing)_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Simple | 1.42 µs (10 allocs) | 2.15 µs (7 allocs) | 12.83 µs (43 allocs) | unsupported | 5.98 µs (48 allocs) | unsupported |
| Complex | 2.29 µs (15 allocs) | 3.49 µs (9 allocs) | 12.27 µs (139 allocs) | unsupported | 13.83 µs (120 allocs) | unsupported |
| Large | 1.58 µs (22 allocs) | 1.84 µs (3 allocs) | 47.38 µs (254 allocs) | unsupported | 14.52 µs (126 allocs) | unsupported |

---

## Summary

### Validate_Simple (struct validation)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 1.42 µs | 10 | baseline |
| Playground | 2.15 µs | 7 | 1.52x slower |
| Ozzo | 12.83 µs | 43 | 9.05x slower |
| Huma | - | - | - |
| Godantic | 5.98 µs | 48 | 4.22x slower |
| Godasse | - | - | - |

### Validate_Complex (nested structs)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 2.29 µs | 15 | baseline |
| Playground | 3.49 µs | 9 | 1.52x slower |
| Ozzo | 12.27 µs | 139 | 5.36x slower |
| Huma | - | - | - |
| Godantic | 13.83 µs | 120 | 6.03x slower |
| Godasse | - | - | - |

### JSONValidate_Simple (JSON → struct + validate)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 3.40 µs | 19 | baseline |
| Playground | 4.17 µs | 16 | 1.23x slower |
| Ozzo | - | - | - |
| Huma | 3.46 µs | 26 | 1.02x slower |
| Godantic | - | - | - |
| Godasse | 5.32 µs | 46 | 1.57x slower |

### JSONValidate_Complex (nested JSON)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 9.64 µs | 39 | baseline |
| Playground | 10.86 µs | 33 | 1.13x slower |
| Ozzo | - | - | - |
| Huma | 10.19 µs | 78 | 1.06x slower |
| Godantic | - | - | - |
| Godasse | 17.28 µs | 153 | 1.79x slower |

### Schema_Uncached (first-time generation)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 23.27 µs | 202 | baseline |
| Playground | - | - | - |
| Ozzo | - | - | - |
| Huma | 30.73 µs | 255 | 1.32x slower |
| Godantic | - | - | - |
| Godasse | - | - | - |

### Schema_Cached (cached lookup)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 18 ns | 0 | baseline |
| Playground | - | - | - |
| Ozzo | - | - | - |
| Huma | 632 ns | 6 | 34.42x slower |
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
