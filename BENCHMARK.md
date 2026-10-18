---
sidebar_position: 99
title: Benchmarks
---

# Benchmark Results

Generated: 2026-10-18 14:09:55 UTC

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

## Getting the Best Performance

`New[T]()` does the expensive work once: it walks the struct via reflection, resolves every constraint tag, and builds an internal field-constraint cache (plus the JSON field deserializers). That one-time cost is what the `New` section below measures (microsecond range). Every other operation - `Validate`, `Unmarshal`, `Marshal`, `Schema` - reuses that precomputed cache and runs in the hundreds-of-ns to low-µs range, which is why those numbers consistently beat libraries that re-resolve constraints or re-walk structs on every call.

This only pays off if the `*Validator[T]` returned by `New` is built once and reused - not recreated per request. Two ways to do that:

**Module-level variable** (sufficient to call the validator directly):

```go
var userValidator = validator.New[User]()

func handleCreateUser(body []byte) (*User, error) {
	return userValidator.Unmarshal(body) // reuses the cached field constraints
}
```

**`Register`** (needed in addition, only if a framework integration - e.g. the Echo Binder plugin, or `UnmarshalInto` - must find the validator for a type it only knows via `reflect.Type` at runtime, not through your module-level variable):

```go
var _ = validator.Register(validator.New[User]())
```

`Register[T]` may be called exactly once per type - a second call for the same type panics, by design. A type could have multiple differently-configured validators (different `Options`), and pedantigo has no way to guess which one a framework plugin should resolve to, so it refuses to silently pick one. Call `Register` from exactly one package-level `var` declaration per type.

---

## Validate
_Validate existing struct (no JSON parsing)_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Simple | 1.13 µs (10 allocs) | 1.63 µs (7 allocs) | 10.39 µs (43 allocs) | unsupported | 4.71 µs (48 allocs) | unsupported |
| Complex | 1.98 µs (15 allocs) | 2.57 µs (9 allocs) | 10.06 µs (139 allocs) | unsupported | 10.80 µs (120 allocs) | unsupported |
| Large | 1.28 µs (22 allocs) | 1.47 µs (3 allocs) | 37.30 µs (254 allocs) | unsupported | 11.68 µs (126 allocs) | unsupported |

## JSONValidate
_JSON bytes → struct, then a separate validate step_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Simple | 2.68 µs (19 allocs) | 3.16 µs (16 allocs) | unsupported | 2.72 µs (26 allocs) | unsupported | unsupported |
| Complex | 7.75 µs (39 allocs) | 8.28 µs (33 allocs) | unsupported | 8.53 µs (78 allocs) | unsupported | unsupported |

## Marshal
_Validate + JSON marshal_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Simple | 1.48 µs (11 allocs) | 2.02 µs (9 allocs) | unsupported | unsupported | unsupported | unsupported |

## Unmarshal
_JSON bytes → validated struct in a single call_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Simple | 3.86 µs (41 allocs) | unsupported | unsupported | unsupported | 8.68 µs (81 allocs) | 3.65 µs (42 allocs) |
| Complex | 10.90 µs (122 allocs) | unsupported | unsupported | unsupported | 36.96 µs (285 allocs) | 12.87 µs (149 allocs) |

## New
_Validator creation overhead_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Simple | 13.72 µs (121 allocs) | 12.31 µs (187 allocs) | unsupported | 26.11 µs (255 allocs) | 22.04 µs (305 allocs) | 5.43 µs (72 allocs) |
| Complex | 40.89 µs (359 allocs) | unsupported | unsupported | 64.86 µs (515 allocs) | 6.03 µs (75 allocs) | 19.00 µs (243 allocs) |

## Schema
_JSON Schema generation_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Uncached | 24.00 µs (219 allocs) | unsupported | unsupported | 25.98 µs (255 allocs) | unsupported | unsupported |
| Cached | 21 ns (0 allocs) | unsupported | unsupported | 560 ns (6 allocs) | unsupported | unsupported |

## OpenAPI
_OpenAPI-compatible schema generation_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Uncached | 24.72 µs (221 allocs) | unsupported | unsupported | 26.13 µs (255 allocs) | unsupported | unsupported |
| Cached | 23 ns (0 allocs) | unsupported | unsupported | 553 ns (6 allocs) | unsupported | unsupported |

---

## Summary

### Validate_Simple (struct validation)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 1.13 µs | 10 | baseline |
| Playground | 1.63 µs | 7 | 1.45x slower |
| Ozzo | 10.39 µs | 43 | 9.21x slower |
| Huma | - | - | - |
| Godantic | 4.71 µs | 48 | 4.18x slower |
| Godasse | - | - | - |

### Validate_Complex (nested structs)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 1.98 µs | 15 | baseline |
| Playground | 2.57 µs | 9 | 1.30x slower |
| Ozzo | 10.06 µs | 139 | 5.08x slower |
| Huma | - | - | - |
| Godantic | 10.80 µs | 120 | 5.46x slower |
| Godasse | - | - | - |

### JSONValidate_Simple (JSON → struct, then validate)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 2.68 µs | 19 | baseline |
| Playground | 3.16 µs | 16 | 1.18x slower |
| Ozzo | - | - | - |
| Huma | 2.72 µs | 26 | 1.02x slower |
| Godantic | - | - | - |
| Godasse | - | - | - |

### JSONValidate_Complex (nested JSON)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 7.75 µs | 39 | baseline |
| Playground | 8.28 µs | 33 | 1.07x slower |
| Ozzo | - | - | - |
| Huma | 8.53 µs | 78 | 1.10x slower |
| Godantic | - | - | - |
| Godasse | - | - | - |

### Unmarshal_Simple (JSON → validated struct, single call)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 3.86 µs | 41 | baseline |
| Playground | - | - | - |
| Ozzo | - | - | - |
| Huma | - | - | - |
| Godantic | 8.68 µs | 81 | 2.25x slower |
| Godasse | 3.65 µs | 42 | 1.06x faster |

### Unmarshal_Complex (nested JSON, single call)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 10.90 µs | 122 | baseline |
| Playground | - | - | - |
| Ozzo | - | - | - |
| Huma | - | - | - |
| Godantic | 36.96 µs | 285 | 3.39x slower |
| Godasse | 12.87 µs | 149 | 1.18x slower |

### Schema_Uncached (first-time generation)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 24.00 µs | 219 | baseline |
| Playground | - | - | - |
| Ozzo | - | - | - |
| Huma | 25.98 µs | 255 | 1.08x slower |
| Godantic | - | - | - |
| Godasse | - | - | - |

### Schema_Cached (cached lookup)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 21 ns | 0 | baseline |
| Playground | - | - | - |
| Ozzo | - | - | - |
| Huma | 560 ns | 6 | 27.30x slower |
| Godantic | - | - | - |
| Godasse | - | - | - |

---

_Generated by pedantigo-benchmarks_

<details>
<summary>Benchmark naming convention</summary>

```
Benchmark_<Library>_<Feature>_<Struct>

Libraries: Pedantigo, Playground, Ozzo, Huma, Godantic, Godasse
Features: Validate, JSONValidate, Marshal, Unmarshal, New, Schema, OpenAPI
Structs: Simple (5 fields), Complex (nested), Large (20+ fields)
```
</details>
