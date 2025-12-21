---
sidebar_position: 99
title: Benchmarks
---

# Benchmark Results

Generated: 2025-12-21 16:18:48 UTC

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
| Simple | 1.54 µs (19 allocs) | 1.75 µs (16 allocs) | unsupported | 1.71 µs (26 allocs) | unsupported | 2.35 µs (46 allocs) |
| Complex | 4.37 µs (39 allocs) | 5.01 µs (33 allocs) | unsupported | 4.95 µs (78 allocs) | unsupported | 8.81 µs (153 allocs) |

## Marshal
_Validate + JSON marshal_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Simple | 768 ns (11 allocs) | 1.08 µs (9 allocs) | unsupported | unsupported | unsupported | unsupported |

## New
_Validator creation overhead_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Simple | 5.41 µs (110 allocs) | 8.06 µs (187 allocs) | unsupported | 12.96 µs (247 allocs) | 11.33 µs (291 allocs) | 2.90 µs (72 allocs) |
| Complex | 13.94 µs (270 allocs) | unsupported | unsupported | 31.98 µs (515 allocs) | 3.19 µs (65 allocs) | 10.26 µs (243 allocs) |

## OpenAPI
_OpenAPI-compatible schema generation_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Uncached | 10.88 µs (204 allocs) | unsupported | unsupported | 12.96 µs (247 allocs) | unsupported | unsupported |
| Cached | 8 ns (0 allocs) | unsupported | unsupported | 331 ns (6 allocs) | unsupported | unsupported |

## Schema
_JSON Schema generation_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Uncached | 10.65 µs (202 allocs) | unsupported | unsupported | 13.07 µs (247 allocs) | unsupported | unsupported |
| Cached | 9 ns (0 allocs) | unsupported | unsupported | 306 ns (6 allocs) | unsupported | unsupported |

## Validate
_Validate existing struct (no JSON parsing)_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Simple | 599 ns (10 allocs) | 890 ns (7 allocs) | 6.45 µs (43 allocs) | unsupported | 2.65 µs (48 allocs) | unsupported |
| Complex | 1.03 µs (15 allocs) | 1.44 µs (9 allocs) | 5.31 µs (139 allocs) | unsupported | 5.84 µs (120 allocs) | unsupported |
| Large | 741 ns (22 allocs) | 863 ns (3 allocs) | 22.70 µs (254 allocs) | unsupported | 6.48 µs (126 allocs) | unsupported |

---

## Summary

### Validate_Simple (struct validation)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 599 ns | 10 | baseline |
| Playground | 890 ns | 7 | 1.49x slower |
| Ozzo | 6.45 µs | 43 | 10.77x slower |
| Huma | - | - | - |
| Godantic | 2.65 µs | 48 | 4.43x slower |
| Godasse | - | - | - |

### Validate_Complex (nested structs)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 1.03 µs | 15 | baseline |
| Playground | 1.44 µs | 9 | 1.40x slower |
| Ozzo | 5.31 µs | 139 | 5.14x slower |
| Huma | - | - | - |
| Godantic | 5.84 µs | 120 | 5.66x slower |
| Godasse | - | - | - |

### JSONValidate_Simple (JSON → struct + validate)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 1.54 µs | 19 | baseline |
| Playground | 1.75 µs | 16 | 1.14x slower |
| Ozzo | - | - | - |
| Huma | 1.71 µs | 26 | 1.11x slower |
| Godantic | - | - | - |
| Godasse | 2.35 µs | 46 | 1.53x slower |

### JSONValidate_Complex (nested JSON)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 4.37 µs | 39 | baseline |
| Playground | 5.01 µs | 33 | 1.15x slower |
| Ozzo | - | - | - |
| Huma | 4.95 µs | 78 | 1.13x slower |
| Godantic | - | - | - |
| Godasse | 8.81 µs | 153 | 2.02x slower |

### Schema_Uncached (first-time generation)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 10.65 µs | 202 | baseline |
| Playground | - | - | - |
| Ozzo | - | - | - |
| Huma | 13.07 µs | 247 | 1.23x slower |
| Godantic | - | - | - |
| Godasse | - | - | - |

### Schema_Cached (cached lookup)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 9 ns | 0 | baseline |
| Playground | - | - | - |
| Ozzo | - | - | - |
| Huma | 306 ns | 6 | 33.55x slower |
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
