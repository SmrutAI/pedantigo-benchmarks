---
sidebar_position: 99
title: Benchmarks
---

# Benchmark Results

Generated: 2026-04-08 11:02:30 UTC

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
| Simple | 3.26 µs (19 allocs) | 4.12 µs (16 allocs) | unsupported | 3.34 µs (26 allocs) | unsupported | 4.99 µs (46 allocs) |
| Complex | 8.95 µs (39 allocs) | 10.27 µs (33 allocs) | unsupported | 9.80 µs (78 allocs) | unsupported | 16.39 µs (153 allocs) |

## Marshal
_Validate + JSON marshal_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Simple | 1.78 µs (11 allocs) | 2.64 µs (9 allocs) | unsupported | unsupported | unsupported | unsupported |

## New
_Validator creation overhead_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Simple | 11.16 µs (110 allocs) | 15.09 µs (187 allocs) | unsupported | 29.54 µs (255 allocs) | 25.52 µs (305 allocs) | 6.27 µs (72 allocs) |
| Complex | 26.74 µs (270 allocs) | unsupported | unsupported | 73.53 µs (515 allocs) | 7.06 µs (75 allocs) | 22.06 µs (243 allocs) |

## OpenAPI
_OpenAPI-compatible schema generation_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Uncached | 22.11 µs (204 allocs) | unsupported | unsupported | 29.89 µs (255 allocs) | unsupported | unsupported |
| Cached | 19 ns (0 allocs) | unsupported | unsupported | 591 ns (6 allocs) | unsupported | unsupported |

## Schema
_JSON Schema generation_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Uncached | 21.69 µs (202 allocs) | unsupported | unsupported | 29.75 µs (255 allocs) | unsupported | unsupported |
| Cached | 19 ns (0 allocs) | unsupported | unsupported | 606 ns (6 allocs) | unsupported | unsupported |

## Validate
_Validate existing struct (no JSON parsing)_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Simple | 1.36 µs (10 allocs) | 2.10 µs (7 allocs) | 11.81 µs (43 allocs) | unsupported | 5.42 µs (48 allocs) | unsupported |
| Complex | 2.15 µs (15 allocs) | 3.36 µs (9 allocs) | 11.38 µs (139 allocs) | unsupported | 12.39 µs (120 allocs) | unsupported |
| Large | 1.44 µs (22 allocs) | 1.65 µs (3 allocs) | 42.50 µs (254 allocs) | unsupported | 13.32 µs (126 allocs) | unsupported |

---

## Summary

### Validate_Simple (struct validation)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 1.36 µs | 10 | baseline |
| Playground | 2.10 µs | 7 | 1.55x slower |
| Ozzo | 11.81 µs | 43 | 8.69x slower |
| Huma | - | - | - |
| Godantic | 5.42 µs | 48 | 3.99x slower |
| Godasse | - | - | - |

### Validate_Complex (nested structs)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 2.15 µs | 15 | baseline |
| Playground | 3.36 µs | 9 | 1.56x slower |
| Ozzo | 11.38 µs | 139 | 5.30x slower |
| Huma | - | - | - |
| Godantic | 12.39 µs | 120 | 5.77x slower |
| Godasse | - | - | - |

### JSONValidate_Simple (JSON → struct + validate)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 3.26 µs | 19 | baseline |
| Playground | 4.12 µs | 16 | 1.26x slower |
| Ozzo | - | - | - |
| Huma | 3.34 µs | 26 | 1.02x slower |
| Godantic | - | - | - |
| Godasse | 4.99 µs | 46 | 1.53x slower |

### JSONValidate_Complex (nested JSON)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 8.95 µs | 39 | baseline |
| Playground | 10.27 µs | 33 | 1.15x slower |
| Ozzo | - | - | - |
| Huma | 9.80 µs | 78 | 1.09x slower |
| Godantic | - | - | - |
| Godasse | 16.39 µs | 153 | 1.83x slower |

### Schema_Uncached (first-time generation)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 21.69 µs | 202 | baseline |
| Playground | - | - | - |
| Ozzo | - | - | - |
| Huma | 29.75 µs | 255 | 1.37x slower |
| Godantic | - | - | - |
| Godasse | - | - | - |

### Schema_Cached (cached lookup)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 19 ns | 0 | baseline |
| Playground | - | - | - |
| Ozzo | - | - | - |
| Huma | 606 ns | 6 | 31.73x slower |
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
