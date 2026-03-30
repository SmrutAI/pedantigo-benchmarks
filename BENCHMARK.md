---
sidebar_position: 99
title: Benchmarks
---

# Benchmark Results

Generated: 2026-03-30 07:25:04 UTC

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
| Simple | 3.20 µs (19 allocs) | 3.96 µs (16 allocs) | unsupported | 3.33 µs (26 allocs) | unsupported | 5.17 µs (46 allocs) |
| Complex | 8.94 µs (39 allocs) | 9.85 µs (33 allocs) | unsupported | 9.41 µs (78 allocs) | unsupported | 16.68 µs (153 allocs) |

## Marshal
_Validate + JSON marshal_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Simple | 1.81 µs (11 allocs) | 2.47 µs (9 allocs) | unsupported | unsupported | unsupported | unsupported |

## New
_Validator creation overhead_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Simple | 12.16 µs (110 allocs) | 15.78 µs (187 allocs) | unsupported | 30.43 µs (255 allocs) | 26.32 µs (305 allocs) | 6.56 µs (72 allocs) |
| Complex | 29.04 µs (270 allocs) | unsupported | unsupported | 75.94 µs (515 allocs) | 7.45 µs (75 allocs) | 23.16 µs (243 allocs) |

## OpenAPI
_OpenAPI-compatible schema generation_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Uncached | 23.88 µs (204 allocs) | unsupported | unsupported | 30.53 µs (255 allocs) | unsupported | unsupported |
| Cached | 25 ns (0 allocs) | unsupported | unsupported | 644 ns (6 allocs) | unsupported | unsupported |

## Schema
_JSON Schema generation_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Uncached | 22.87 µs (202 allocs) | unsupported | unsupported | 30.84 µs (255 allocs) | unsupported | unsupported |
| Cached | 24 ns (0 allocs) | unsupported | unsupported | 648 ns (6 allocs) | unsupported | unsupported |

## Validate
_Validate existing struct (no JSON parsing)_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Simple | 1.37 µs (10 allocs) | 1.99 µs (7 allocs) | 11.89 µs (43 allocs) | unsupported | 6.28 µs (48 allocs) | unsupported |
| Complex | 2.22 µs (15 allocs) | 3.19 µs (9 allocs) | 12.14 µs (139 allocs) | unsupported | 13.59 µs (120 allocs) | unsupported |
| Large | 1.48 µs (22 allocs) | 1.75 µs (3 allocs) | 44.36 µs (254 allocs) | unsupported | 15.07 µs (126 allocs) | unsupported |

---

## Summary

### Validate_Simple (struct validation)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 1.37 µs | 10 | baseline |
| Playground | 1.99 µs | 7 | 1.46x slower |
| Ozzo | 11.89 µs | 43 | 8.70x slower |
| Huma | - | - | - |
| Godantic | 6.28 µs | 48 | 4.59x slower |
| Godasse | - | - | - |

### Validate_Complex (nested structs)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 2.22 µs | 15 | baseline |
| Playground | 3.19 µs | 9 | 1.44x slower |
| Ozzo | 12.14 µs | 139 | 5.47x slower |
| Huma | - | - | - |
| Godantic | 13.59 µs | 120 | 6.13x slower |
| Godasse | - | - | - |

### JSONValidate_Simple (JSON → struct + validate)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 3.20 µs | 19 | baseline |
| Playground | 3.96 µs | 16 | 1.24x slower |
| Ozzo | - | - | - |
| Huma | 3.33 µs | 26 | 1.04x slower |
| Godantic | - | - | - |
| Godasse | 5.17 µs | 46 | 1.62x slower |

### JSONValidate_Complex (nested JSON)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 8.94 µs | 39 | baseline |
| Playground | 9.85 µs | 33 | 1.10x slower |
| Ozzo | - | - | - |
| Huma | 9.41 µs | 78 | 1.05x slower |
| Godantic | - | - | - |
| Godasse | 16.68 µs | 153 | 1.86x slower |

### Schema_Uncached (first-time generation)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 22.87 µs | 202 | baseline |
| Playground | - | - | - |
| Ozzo | - | - | - |
| Huma | 30.84 µs | 255 | 1.35x slower |
| Godantic | - | - | - |
| Godasse | - | - | - |

### Schema_Cached (cached lookup)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 24 ns | 0 | baseline |
| Playground | - | - | - |
| Ozzo | - | - | - |
| Huma | 648 ns | 6 | 26.46x slower |
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
