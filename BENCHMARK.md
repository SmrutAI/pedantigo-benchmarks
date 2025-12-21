---
sidebar_position: 99
title: Benchmarks
---

# Benchmark Results

Generated: 2025-12-21 17:13:21 UTC

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
| Simple | 3.50 µs (19 allocs) | 4.30 µs (16 allocs) | unsupported | 3.54 µs (26 allocs) | unsupported | 5.44 µs (46 allocs) |
| Complex | 9.88 µs (39 allocs) | 11.14 µs (33 allocs) | unsupported | 10.32 µs (78 allocs) | unsupported | 17.33 µs (153 allocs) |

## Marshal
_Validate + JSON marshal_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Simple | 1.91 µs (11 allocs) | 2.71 µs (9 allocs) | unsupported | unsupported | unsupported | unsupported |

## New
_Validator creation overhead_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Simple | 11.88 µs (110 allocs) | 16.82 µs (187 allocs) | unsupported | 30.02 µs (255 allocs) | 26.61 µs (305 allocs) | 6.67 µs (72 allocs) |
| Complex | 28.64 µs (270 allocs) | unsupported | unsupported | 73.48 µs (515 allocs) | 7.49 µs (75 allocs) | 23.31 µs (243 allocs) |

## OpenAPI
_OpenAPI-compatible schema generation_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Uncached | 23.72 µs (204 allocs) | unsupported | unsupported | 30.30 µs (255 allocs) | unsupported | unsupported |
| Cached | 18 ns (0 allocs) | unsupported | unsupported | 647 ns (6 allocs) | unsupported | unsupported |

## Schema
_JSON Schema generation_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Uncached | 22.97 µs (202 allocs) | unsupported | unsupported | 30.18 µs (255 allocs) | unsupported | unsupported |
| Cached | 18 ns (0 allocs) | unsupported | unsupported | 639 ns (6 allocs) | unsupported | unsupported |

## Validate
_Validate existing struct (no JSON parsing)_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Simple | 1.44 µs (10 allocs) | 2.17 µs (7 allocs) | 12.77 µs (43 allocs) | unsupported | 6.12 µs (48 allocs) | unsupported |
| Complex | 2.25 µs (15 allocs) | 3.50 µs (9 allocs) | 12.44 µs (139 allocs) | unsupported | 13.85 µs (120 allocs) | unsupported |
| Large | 1.58 µs (22 allocs) | 1.87 µs (3 allocs) | 46.85 µs (254 allocs) | unsupported | 14.69 µs (126 allocs) | unsupported |

---

## Summary

### Validate_Simple (struct validation)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 1.44 µs | 10 | baseline |
| Playground | 2.17 µs | 7 | 1.51x slower |
| Ozzo | 12.77 µs | 43 | 8.84x slower |
| Huma | - | - | - |
| Godantic | 6.12 µs | 48 | 4.24x slower |
| Godasse | - | - | - |

### Validate_Complex (nested structs)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 2.25 µs | 15 | baseline |
| Playground | 3.50 µs | 9 | 1.56x slower |
| Ozzo | 12.44 µs | 139 | 5.54x slower |
| Huma | - | - | - |
| Godantic | 13.85 µs | 120 | 6.17x slower |
| Godasse | - | - | - |

### JSONValidate_Simple (JSON → struct + validate)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 3.50 µs | 19 | baseline |
| Playground | 4.30 µs | 16 | 1.23x slower |
| Ozzo | - | - | - |
| Huma | 3.54 µs | 26 | 1.01x slower |
| Godantic | - | - | - |
| Godasse | 5.44 µs | 46 | 1.56x slower |

### JSONValidate_Complex (nested JSON)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 9.88 µs | 39 | baseline |
| Playground | 11.14 µs | 33 | 1.13x slower |
| Ozzo | - | - | - |
| Huma | 10.32 µs | 78 | 1.05x slower |
| Godantic | - | - | - |
| Godasse | 17.33 µs | 153 | 1.75x slower |

### Schema_Uncached (first-time generation)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 22.97 µs | 202 | baseline |
| Playground | - | - | - |
| Ozzo | - | - | - |
| Huma | 30.18 µs | 255 | 1.31x slower |
| Godantic | - | - | - |
| Godasse | - | - | - |

### Schema_Cached (cached lookup)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 18 ns | 0 | baseline |
| Playground | - | - | - |
| Ozzo | - | - | - |
| Huma | 639 ns | 6 | 34.71x slower |
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
