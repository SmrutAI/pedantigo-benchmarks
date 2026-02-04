---
sidebar_position: 99
title: Benchmarks
---

# Benchmark Results

Generated: 2026-02-04 11:01:32 UTC

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
| Simple | 3.49 µs (19 allocs) | 4.25 µs (16 allocs) | unsupported | 3.57 µs (26 allocs) | unsupported | 5.50 µs (46 allocs) |
| Complex | 9.90 µs (39 allocs) | 11.04 µs (33 allocs) | unsupported | 10.48 µs (78 allocs) | unsupported | 17.83 µs (153 allocs) |

## Marshal
_Validate + JSON marshal_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Simple | 1.90 µs (11 allocs) | 2.82 µs (9 allocs) | unsupported | unsupported | unsupported | unsupported |

## New
_Validator creation overhead_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Simple | 12.01 µs (110 allocs) | 16.29 µs (187 allocs) | unsupported | 32.29 µs (255 allocs) | 27.51 µs (305 allocs) | 6.94 µs (72 allocs) |
| Complex | 28.90 µs (270 allocs) | unsupported | unsupported | 77.85 µs (515 allocs) | 7.80 µs (75 allocs) | 23.96 µs (243 allocs) |

## OpenAPI
_OpenAPI-compatible schema generation_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Uncached | 24.20 µs (204 allocs) | unsupported | unsupported | 32.04 µs (255 allocs) | unsupported | unsupported |
| Cached | 20 ns (0 allocs) | unsupported | unsupported | 640 ns (6 allocs) | unsupported | unsupported |

## Schema
_JSON Schema generation_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Uncached | 23.28 µs (202 allocs) | unsupported | unsupported | 32.41 µs (255 allocs) | unsupported | unsupported |
| Cached | 21 ns (0 allocs) | unsupported | unsupported | 640 ns (6 allocs) | unsupported | unsupported |

## Validate
_Validate existing struct (no JSON parsing)_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Simple | 1.55 µs (10 allocs) | 2.18 µs (7 allocs) | 13.15 µs (43 allocs) | unsupported | 6.05 µs (48 allocs) | unsupported |
| Complex | 2.38 µs (15 allocs) | 3.57 µs (9 allocs) | 12.52 µs (139 allocs) | unsupported | 14.25 µs (120 allocs) | unsupported |
| Large | 1.65 µs (22 allocs) | 1.89 µs (3 allocs) | 48.01 µs (254 allocs) | unsupported | 14.78 µs (126 allocs) | unsupported |

---

## Summary

### Validate_Simple (struct validation)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 1.55 µs | 10 | baseline |
| Playground | 2.18 µs | 7 | 1.41x slower |
| Ozzo | 13.15 µs | 43 | 8.49x slower |
| Huma | - | - | - |
| Godantic | 6.05 µs | 48 | 3.91x slower |
| Godasse | - | - | - |

### Validate_Complex (nested structs)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 2.38 µs | 15 | baseline |
| Playground | 3.57 µs | 9 | 1.50x slower |
| Ozzo | 12.52 µs | 139 | 5.27x slower |
| Huma | - | - | - |
| Godantic | 14.25 µs | 120 | 5.99x slower |
| Godasse | - | - | - |

### JSONValidate_Simple (JSON → struct + validate)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 3.49 µs | 19 | baseline |
| Playground | 4.25 µs | 16 | 1.22x slower |
| Ozzo | - | - | - |
| Huma | 3.57 µs | 26 | 1.02x slower |
| Godantic | - | - | - |
| Godasse | 5.50 µs | 46 | 1.58x slower |

### JSONValidate_Complex (nested JSON)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 9.90 µs | 39 | baseline |
| Playground | 11.04 µs | 33 | 1.12x slower |
| Ozzo | - | - | - |
| Huma | 10.48 µs | 78 | 1.06x slower |
| Godantic | - | - | - |
| Godasse | 17.83 µs | 153 | 1.80x slower |

### Schema_Uncached (first-time generation)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 23.28 µs | 202 | baseline |
| Playground | - | - | - |
| Ozzo | - | - | - |
| Huma | 32.41 µs | 255 | 1.39x slower |
| Godantic | - | - | - |
| Godasse | - | - | - |

### Schema_Cached (cached lookup)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 21 ns | 0 | baseline |
| Playground | - | - | - |
| Ozzo | - | - | - |
| Huma | 640 ns | 6 | 31.01x slower |
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
