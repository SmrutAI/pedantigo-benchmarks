---
sidebar_position: 99
title: Benchmarks
---

# Benchmark Results

Generated: 2026-01-01 20:12:06 UTC

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
| Simple | 3.49 µs (19 allocs) | 4.30 µs (16 allocs) | unsupported | 3.63 µs (26 allocs) | unsupported | 5.51 µs (46 allocs) |
| Complex | 10.00 µs (39 allocs) | 11.21 µs (33 allocs) | unsupported | 10.54 µs (78 allocs) | unsupported | 17.74 µs (153 allocs) |

## Marshal
_Validate + JSON marshal_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Simple | 1.90 µs (11 allocs) | 2.76 µs (9 allocs) | unsupported | unsupported | unsupported | unsupported |

## New
_Validator creation overhead_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Simple | 11.94 µs (110 allocs) | 16.27 µs (187 allocs) | unsupported | 30.83 µs (255 allocs) | 27.28 µs (305 allocs) | 6.75 µs (72 allocs) |
| Complex | 28.64 µs (270 allocs) | unsupported | unsupported | 75.49 µs (515 allocs) | 7.61 µs (75 allocs) | 23.65 µs (243 allocs) |

## OpenAPI
_OpenAPI-compatible schema generation_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Uncached | 23.65 µs (204 allocs) | unsupported | unsupported | 30.80 µs (255 allocs) | unsupported | unsupported |
| Cached | 18 ns (0 allocs) | unsupported | unsupported | 647 ns (6 allocs) | unsupported | unsupported |

## Schema
_JSON Schema generation_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Uncached | 22.84 µs (202 allocs) | unsupported | unsupported | 30.73 µs (255 allocs) | unsupported | unsupported |
| Cached | 18 ns (0 allocs) | unsupported | unsupported | 647 ns (6 allocs) | unsupported | unsupported |

## Validate
_Validate existing struct (no JSON parsing)_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Simple | 1.44 µs (10 allocs) | 2.19 µs (7 allocs) | 12.85 µs (43 allocs) | unsupported | 6.21 µs (48 allocs) | unsupported |
| Complex | 2.27 µs (15 allocs) | 3.49 µs (9 allocs) | 12.38 µs (139 allocs) | unsupported | 13.95 µs (120 allocs) | unsupported |
| Large | 1.58 µs (22 allocs) | 1.87 µs (3 allocs) | 46.88 µs (254 allocs) | unsupported | 14.91 µs (126 allocs) | unsupported |

---

## Summary

### Validate_Simple (struct validation)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 1.44 µs | 10 | baseline |
| Playground | 2.19 µs | 7 | 1.52x slower |
| Ozzo | 12.85 µs | 43 | 8.91x slower |
| Huma | - | - | - |
| Godantic | 6.21 µs | 48 | 4.31x slower |
| Godasse | - | - | - |

### Validate_Complex (nested structs)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 2.27 µs | 15 | baseline |
| Playground | 3.49 µs | 9 | 1.53x slower |
| Ozzo | 12.38 µs | 139 | 5.45x slower |
| Huma | - | - | - |
| Godantic | 13.95 µs | 120 | 6.14x slower |
| Godasse | - | - | - |

### JSONValidate_Simple (JSON → struct + validate)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 3.49 µs | 19 | baseline |
| Playground | 4.30 µs | 16 | 1.23x slower |
| Ozzo | - | - | - |
| Huma | 3.63 µs | 26 | 1.04x slower |
| Godantic | - | - | - |
| Godasse | 5.51 µs | 46 | 1.58x slower |

### JSONValidate_Complex (nested JSON)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 10.00 µs | 39 | baseline |
| Playground | 11.21 µs | 33 | 1.12x slower |
| Ozzo | - | - | - |
| Huma | 10.54 µs | 78 | 1.05x slower |
| Godantic | - | - | - |
| Godasse | 17.74 µs | 153 | 1.77x slower |

### Schema_Uncached (first-time generation)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 22.84 µs | 202 | baseline |
| Playground | - | - | - |
| Ozzo | - | - | - |
| Huma | 30.73 µs | 255 | 1.35x slower |
| Godantic | - | - | - |
| Godasse | - | - | - |

### Schema_Cached (cached lookup)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 18 ns | 0 | baseline |
| Playground | - | - | - |
| Ozzo | - | - | - |
| Huma | 647 ns | 6 | 35.26x slower |
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
