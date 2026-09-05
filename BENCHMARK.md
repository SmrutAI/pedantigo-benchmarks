---
sidebar_position: 99
title: Benchmarks
---

# Benchmark Results

Generated: 2026-09-05 04:58:51 UTC

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
| Simple | 935 ns (10 allocs) | 1.35 µs (7 allocs) | 10.02 µs (43 allocs) | unsupported | 3.89 µs (48 allocs) | unsupported |
| Complex | 1.92 µs (15 allocs) | 2.13 µs (9 allocs) | 8.50 µs (139 allocs) | unsupported | 9.00 µs (120 allocs) | unsupported |
| Large | 1.08 µs (22 allocs) | 1.24 µs (3 allocs) | 31.59 µs (254 allocs) | unsupported | 9.68 µs (126 allocs) | unsupported |

## JSONValidate
_JSON bytes → struct, then a separate validate step_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Simple | 2.20 µs (19 allocs) | 2.62 µs (16 allocs) | unsupported | 2.26 µs (26 allocs) | unsupported | unsupported |
| Complex | 6.44 µs (39 allocs) | 6.89 µs (33 allocs) | unsupported | 6.84 µs (78 allocs) | unsupported | unsupported |

## Marshal
_Validate + JSON marshal_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Simple | 1.22 µs (11 allocs) | 1.68 µs (9 allocs) | unsupported | unsupported | unsupported | unsupported |

## Unmarshal
_JSON bytes → validated struct in a single call_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Simple | 3.21 µs (41 allocs) | unsupported | unsupported | unsupported | 7.23 µs (81 allocs) | 3.03 µs (42 allocs) |
| Complex | 9.22 µs (122 allocs) | unsupported | unsupported | unsupported | 31.06 µs (285 allocs) | 10.84 µs (149 allocs) |

## New
_Validator creation overhead_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Simple | 12.58 µs (121 allocs) | 11.11 µs (187 allocs) | unsupported | 21.86 µs (255 allocs) | 18.82 µs (305 allocs) | 4.51 µs (72 allocs) |
| Complex | 34.33 µs (359 allocs) | unsupported | unsupported | 53.97 µs (515 allocs) | 4.91 µs (75 allocs) | 15.85 µs (243 allocs) |

## Schema
_JSON Schema generation_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Uncached | 20.03 µs (219 allocs) | unsupported | unsupported | 21.86 µs (255 allocs) | unsupported | unsupported |
| Cached | 17 ns (0 allocs) | unsupported | unsupported | 474 ns (6 allocs) | unsupported | unsupported |

## OpenAPI
_OpenAPI-compatible schema generation_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Uncached | 20.87 µs (221 allocs) | unsupported | unsupported | 21.85 µs (255 allocs) | unsupported | unsupported |
| Cached | 17 ns (0 allocs) | unsupported | unsupported | 460 ns (6 allocs) | unsupported | unsupported |

---

## Summary

### Validate_Simple (struct validation)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 935 ns | 10 | baseline |
| Playground | 1.35 µs | 7 | 1.44x slower |
| Ozzo | 10.02 µs | 43 | 10.72x slower |
| Huma | - | - | - |
| Godantic | 3.89 µs | 48 | 4.16x slower |
| Godasse | - | - | - |

### Validate_Complex (nested structs)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 1.92 µs | 15 | baseline |
| Playground | 2.13 µs | 9 | 1.11x slower |
| Ozzo | 8.50 µs | 139 | 4.43x slower |
| Huma | - | - | - |
| Godantic | 9.00 µs | 120 | 4.69x slower |
| Godasse | - | - | - |

### JSONValidate_Simple (JSON → struct, then validate)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 2.20 µs | 19 | baseline |
| Playground | 2.62 µs | 16 | 1.19x slower |
| Ozzo | - | - | - |
| Huma | 2.26 µs | 26 | 1.03x slower |
| Godantic | - | - | - |
| Godasse | - | - | - |

### JSONValidate_Complex (nested JSON)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 6.44 µs | 39 | baseline |
| Playground | 6.89 µs | 33 | 1.07x slower |
| Ozzo | - | - | - |
| Huma | 6.84 µs | 78 | 1.06x slower |
| Godantic | - | - | - |
| Godasse | - | - | - |

### Unmarshal_Simple (JSON → validated struct, single call)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 3.21 µs | 41 | baseline |
| Playground | - | - | - |
| Ozzo | - | - | - |
| Huma | - | - | - |
| Godantic | 7.23 µs | 81 | 2.26x slower |
| Godasse | 3.03 µs | 42 | 1.06x faster |

### Unmarshal_Complex (nested JSON, single call)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 9.22 µs | 122 | baseline |
| Playground | - | - | - |
| Ozzo | - | - | - |
| Huma | - | - | - |
| Godantic | 31.06 µs | 285 | 3.37x slower |
| Godasse | 10.84 µs | 149 | 1.17x slower |

### Schema_Uncached (first-time generation)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 20.03 µs | 219 | baseline |
| Playground | - | - | - |
| Ozzo | - | - | - |
| Huma | 21.86 µs | 255 | 1.09x slower |
| Godantic | - | - | - |
| Godasse | - | - | - |

### Schema_Cached (cached lookup)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 17 ns | 0 | baseline |
| Playground | - | - | - |
| Ozzo | - | - | - |
| Huma | 474 ns | 6 | 28.48x slower |
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
