---
sidebar_position: 99
title: Benchmarks
---

# Benchmark Results

Generated: 2026-03-23 11:25:09 UTC

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
| Simple | 3.18 µs (19 allocs) | 3.89 µs (16 allocs) | unsupported | 3.29 µs (26 allocs) | unsupported | 5.08 µs (46 allocs) |
| Complex | 8.81 µs (39 allocs) | 9.65 µs (33 allocs) | unsupported | 9.39 µs (78 allocs) | unsupported | 16.21 µs (153 allocs) |

## Marshal
_Validate + JSON marshal_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Simple | 1.81 µs (11 allocs) | 2.46 µs (9 allocs) | unsupported | unsupported | unsupported | unsupported |

## New
_Validator creation overhead_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Simple | 12.15 µs (110 allocs) | 15.82 µs (187 allocs) | unsupported | 29.87 µs (255 allocs) | 25.78 µs (305 allocs) | 6.52 µs (72 allocs) |
| Complex | 29.49 µs (270 allocs) | unsupported | unsupported | 75.08 µs (515 allocs) | 7.42 µs (75 allocs) | 22.87 µs (243 allocs) |

## OpenAPI
_OpenAPI-compatible schema generation_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Uncached | 23.27 µs (204 allocs) | unsupported | unsupported | 30.84 µs (255 allocs) | unsupported | unsupported |
| Cached | 24 ns (0 allocs) | unsupported | unsupported | 661 ns (6 allocs) | unsupported | unsupported |

## Schema
_JSON Schema generation_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Uncached | 23.01 µs (202 allocs) | unsupported | unsupported | 29.97 µs (255 allocs) | unsupported | unsupported |
| Cached | 24 ns (0 allocs) | unsupported | unsupported | 640 ns (6 allocs) | unsupported | unsupported |

## Validate
_Validate existing struct (no JSON parsing)_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Simple | 1.37 µs (10 allocs) | 1.96 µs (7 allocs) | 11.90 µs (43 allocs) | unsupported | 5.91 µs (48 allocs) | unsupported |
| Complex | 2.22 µs (15 allocs) | 3.16 µs (9 allocs) | 12.33 µs (139 allocs) | unsupported | 13.44 µs (120 allocs) | unsupported |
| Large | 1.47 µs (22 allocs) | 1.71 µs (3 allocs) | 44.16 µs (254 allocs) | unsupported | 14.55 µs (126 allocs) | unsupported |

---

## Summary

### Validate_Simple (struct validation)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 1.37 µs | 10 | baseline |
| Playground | 1.96 µs | 7 | 1.43x slower |
| Ozzo | 11.90 µs | 43 | 8.68x slower |
| Huma | - | - | - |
| Godantic | 5.91 µs | 48 | 4.31x slower |
| Godasse | - | - | - |

### Validate_Complex (nested structs)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 2.22 µs | 15 | baseline |
| Playground | 3.16 µs | 9 | 1.43x slower |
| Ozzo | 12.33 µs | 139 | 5.56x slower |
| Huma | - | - | - |
| Godantic | 13.44 µs | 120 | 6.06x slower |
| Godasse | - | - | - |

### JSONValidate_Simple (JSON → struct + validate)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 3.18 µs | 19 | baseline |
| Playground | 3.89 µs | 16 | 1.22x slower |
| Ozzo | - | - | - |
| Huma | 3.29 µs | 26 | 1.03x slower |
| Godantic | - | - | - |
| Godasse | 5.08 µs | 46 | 1.60x slower |

### JSONValidate_Complex (nested JSON)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 8.81 µs | 39 | baseline |
| Playground | 9.65 µs | 33 | 1.10x slower |
| Ozzo | - | - | - |
| Huma | 9.39 µs | 78 | 1.07x slower |
| Godantic | - | - | - |
| Godasse | 16.21 µs | 153 | 1.84x slower |

### Schema_Uncached (first-time generation)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 23.01 µs | 202 | baseline |
| Playground | - | - | - |
| Ozzo | - | - | - |
| Huma | 29.97 µs | 255 | 1.30x slower |
| Godantic | - | - | - |
| Godasse | - | - | - |

### Schema_Cached (cached lookup)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 24 ns | 0 | baseline |
| Playground | - | - | - |
| Ozzo | - | - | - |
| Huma | 640 ns | 6 | 26.20x slower |
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
