---
sidebar_position: 99
title: Benchmarks
---

# Benchmark Results

Generated: 2026-02-04 08:57:42 UTC

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
| Simple | 3.49 µs (19 allocs) | 4.24 µs (16 allocs) | unsupported | 3.49 µs (26 allocs) | unsupported | 5.31 µs (46 allocs) |
| Complex | 11.66 µs (39 allocs) | 10.95 µs (33 allocs) | unsupported | 10.39 µs (78 allocs) | unsupported | 17.07 µs (153 allocs) |

## Marshal
_Validate + JSON marshal_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Simple | 1.87 µs (11 allocs) | 2.76 µs (9 allocs) | unsupported | unsupported | unsupported | unsupported |

## New
_Validator creation overhead_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Simple | 11.87 µs (110 allocs) | 16.82 µs (187 allocs) | unsupported | 30.59 µs (255 allocs) | 26.49 µs (305 allocs) | 6.61 µs (72 allocs) |
| Complex | 28.56 µs (270 allocs) | unsupported | unsupported | 75.39 µs (515 allocs) | 7.38 µs (75 allocs) | 23.17 µs (243 allocs) |

## OpenAPI
_OpenAPI-compatible schema generation_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Uncached | 23.73 µs (204 allocs) | unsupported | unsupported | 30.80 µs (255 allocs) | unsupported | unsupported |
| Cached | 18 ns (0 allocs) | unsupported | unsupported | 628 ns (6 allocs) | unsupported | unsupported |

## Schema
_JSON Schema generation_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Uncached | 22.89 µs (202 allocs) | unsupported | unsupported | 30.86 µs (255 allocs) | unsupported | unsupported |
| Cached | 19 ns (0 allocs) | unsupported | unsupported | 627 ns (6 allocs) | unsupported | unsupported |

## Validate
_Validate existing struct (no JSON parsing)_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Simple | 1.43 µs (10 allocs) | 2.23 µs (7 allocs) | 12.96 µs (43 allocs) | unsupported | 5.99 µs (48 allocs) | unsupported |
| Complex | 2.31 µs (15 allocs) | 3.57 µs (9 allocs) | 12.26 µs (139 allocs) | unsupported | 14.34 µs (120 allocs) | unsupported |
| Large | 1.59 µs (22 allocs) | 1.85 µs (3 allocs) | 47.26 µs (254 allocs) | unsupported | 14.84 µs (126 allocs) | unsupported |

---

## Summary

### Validate_Simple (struct validation)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 1.43 µs | 10 | baseline |
| Playground | 2.23 µs | 7 | 1.56x slower |
| Ozzo | 12.96 µs | 43 | 9.09x slower |
| Huma | - | - | - |
| Godantic | 5.99 µs | 48 | 4.20x slower |
| Godasse | - | - | - |

### Validate_Complex (nested structs)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 2.31 µs | 15 | baseline |
| Playground | 3.57 µs | 9 | 1.55x slower |
| Ozzo | 12.26 µs | 139 | 5.31x slower |
| Huma | - | - | - |
| Godantic | 14.34 µs | 120 | 6.21x slower |
| Godasse | - | - | - |

### JSONValidate_Simple (JSON → struct + validate)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 3.49 µs | 19 | baseline |
| Playground | 4.24 µs | 16 | 1.21x slower |
| Ozzo | - | - | - |
| Huma | 3.49 µs | 26 | 1.00x faster |
| Godantic | - | - | - |
| Godasse | 5.31 µs | 46 | 1.52x slower |

### JSONValidate_Complex (nested JSON)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 11.66 µs | 39 | baseline |
| Playground | 10.95 µs | 33 | 1.06x faster |
| Ozzo | - | - | - |
| Huma | 10.39 µs | 78 | 1.12x faster |
| Godantic | - | - | - |
| Godasse | 17.07 µs | 153 | 1.46x slower |

### Schema_Uncached (first-time generation)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 22.89 µs | 202 | baseline |
| Playground | - | - | - |
| Ozzo | - | - | - |
| Huma | 30.86 µs | 255 | 1.35x slower |
| Godantic | - | - | - |
| Godasse | - | - | - |

### Schema_Cached (cached lookup)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 19 ns | 0 | baseline |
| Playground | - | - | - |
| Ozzo | - | - | - |
| Huma | 627 ns | 6 | 33.03x slower |
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
