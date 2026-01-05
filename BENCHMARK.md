---
sidebar_position: 99
title: Benchmarks
---

# Benchmark Results

Generated: 2026-01-05 20:16:31 UTC

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
| Simple | 3.47 µs (19 allocs) | 4.28 µs (16 allocs) | unsupported | 3.52 µs (26 allocs) | unsupported | 5.42 µs (46 allocs) |
| Complex | 9.84 µs (39 allocs) | 11.06 µs (33 allocs) | unsupported | 10.30 µs (78 allocs) | unsupported | 17.24 µs (153 allocs) |

## Marshal
_Validate + JSON marshal_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Simple | 1.89 µs (11 allocs) | 2.71 µs (9 allocs) | unsupported | unsupported | unsupported | unsupported |

## New
_Validator creation overhead_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Simple | 11.73 µs (110 allocs) | 16.86 µs (187 allocs) | unsupported | 29.97 µs (255 allocs) | 27.07 µs (305 allocs) | 6.68 µs (72 allocs) |
| Complex | 28.52 µs (270 allocs) | unsupported | unsupported | 73.56 µs (515 allocs) | 7.50 µs (75 allocs) | 23.25 µs (243 allocs) |

## OpenAPI
_OpenAPI-compatible schema generation_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Uncached | 23.51 µs (204 allocs) | unsupported | unsupported | 30.27 µs (255 allocs) | unsupported | unsupported |
| Cached | 18 ns (0 allocs) | unsupported | unsupported | 646 ns (6 allocs) | unsupported | unsupported |

## Schema
_JSON Schema generation_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Uncached | 22.64 µs (202 allocs) | unsupported | unsupported | 30.00 µs (255 allocs) | unsupported | unsupported |
| Cached | 18 ns (0 allocs) | unsupported | unsupported | 634 ns (6 allocs) | unsupported | unsupported |

## Validate
_Validate existing struct (no JSON parsing)_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Simple | 1.43 µs (10 allocs) | 2.16 µs (7 allocs) | 12.70 µs (43 allocs) | unsupported | 5.96 µs (48 allocs) | unsupported |
| Complex | 2.24 µs (15 allocs) | 3.46 µs (9 allocs) | 12.32 µs (139 allocs) | unsupported | 13.75 µs (120 allocs) | unsupported |
| Large | 1.56 µs (22 allocs) | 1.85 µs (3 allocs) | 46.57 µs (254 allocs) | unsupported | 14.56 µs (126 allocs) | unsupported |

---

## Summary

### Validate_Simple (struct validation)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 1.43 µs | 10 | baseline |
| Playground | 2.16 µs | 7 | 1.52x slower |
| Ozzo | 12.70 µs | 43 | 8.90x slower |
| Huma | - | - | - |
| Godantic | 5.96 µs | 48 | 4.18x slower |
| Godasse | - | - | - |

### Validate_Complex (nested structs)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 2.24 µs | 15 | baseline |
| Playground | 3.46 µs | 9 | 1.55x slower |
| Ozzo | 12.32 µs | 139 | 5.51x slower |
| Huma | - | - | - |
| Godantic | 13.75 µs | 120 | 6.15x slower |
| Godasse | - | - | - |

### JSONValidate_Simple (JSON → struct + validate)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 3.47 µs | 19 | baseline |
| Playground | 4.28 µs | 16 | 1.23x slower |
| Ozzo | - | - | - |
| Huma | 3.52 µs | 26 | 1.01x slower |
| Godantic | - | - | - |
| Godasse | 5.42 µs | 46 | 1.56x slower |

### JSONValidate_Complex (nested JSON)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 9.84 µs | 39 | baseline |
| Playground | 11.06 µs | 33 | 1.12x slower |
| Ozzo | - | - | - |
| Huma | 10.30 µs | 78 | 1.05x slower |
| Godantic | - | - | - |
| Godasse | 17.24 µs | 153 | 1.75x slower |

### Schema_Uncached (first-time generation)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 22.64 µs | 202 | baseline |
| Playground | - | - | - |
| Ozzo | - | - | - |
| Huma | 30.00 µs | 255 | 1.32x slower |
| Godantic | - | - | - |
| Godasse | - | - | - |

### Schema_Cached (cached lookup)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 18 ns | 0 | baseline |
| Playground | - | - | - |
| Ozzo | - | - | - |
| Huma | 634 ns | 6 | 34.61x slower |
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
