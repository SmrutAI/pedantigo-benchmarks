---
sidebar_position: 99
title: Benchmarks
---

# Benchmark Results

Generated: 2026-09-30 13:42:40 UTC

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
| Simple | 1.34 µs (10 allocs) | 2.12 µs (7 allocs) | 12.07 µs (43 allocs) | unsupported | 5.41 µs (48 allocs) | unsupported |
| Complex | 2.46 µs (15 allocs) | 3.38 µs (9 allocs) | 11.52 µs (139 allocs) | unsupported | 12.95 µs (120 allocs) | unsupported |
| Large | 1.48 µs (22 allocs) | 1.95 µs (3 allocs) | 42.88 µs (254 allocs) | unsupported | 13.31 µs (126 allocs) | unsupported |

## JSONValidate
_JSON bytes → struct, then a separate validate step_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Simple | 3.31 µs (19 allocs) | 4.11 µs (16 allocs) | unsupported | 3.34 µs (26 allocs) | unsupported | unsupported |
| Complex | 9.38 µs (39 allocs) | 10.30 µs (33 allocs) | unsupported | 9.91 µs (78 allocs) | unsupported | unsupported |

## Marshal
_Validate + JSON marshal_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Simple | 1.81 µs (11 allocs) | 2.64 µs (9 allocs) | unsupported | unsupported | unsupported | unsupported |

## Unmarshal
_JSON bytes → validated struct in a single call_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Simple | 4.75 µs (41 allocs) | unsupported | unsupported | unsupported | 10.76 µs (81 allocs) | 4.40 µs (42 allocs) |
| Complex | 13.58 µs (122 allocs) | unsupported | unsupported | unsupported | 45.70 µs (285 allocs) | 15.67 µs (149 allocs) |

## New
_Validator creation overhead_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Simple | 15.01 µs (121 allocs) | 16.01 µs (187 allocs) | unsupported | 29.97 µs (255 allocs) | 25.58 µs (305 allocs) | 6.26 µs (72 allocs) |
| Complex | 44.58 µs (359 allocs) | unsupported | unsupported | 74.05 µs (515 allocs) | 7.16 µs (75 allocs) | 21.97 µs (243 allocs) |

## Schema
_JSON Schema generation_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Uncached | 26.14 µs (219 allocs) | unsupported | unsupported | 29.85 µs (255 allocs) | unsupported | unsupported |
| Cached | 21 ns (0 allocs) | unsupported | unsupported | 575 ns (6 allocs) | unsupported | unsupported |

## OpenAPI
_OpenAPI-compatible schema generation_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Uncached | 27.06 µs (221 allocs) | unsupported | unsupported | 29.80 µs (255 allocs) | unsupported | unsupported |
| Cached | 21 ns (0 allocs) | unsupported | unsupported | 568 ns (6 allocs) | unsupported | unsupported |

---

## Summary

### Validate_Simple (struct validation)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 1.34 µs | 10 | baseline |
| Playground | 2.12 µs | 7 | 1.57x slower |
| Ozzo | 12.07 µs | 43 | 8.99x slower |
| Huma | - | - | - |
| Godantic | 5.41 µs | 48 | 4.03x slower |
| Godasse | - | - | - |

### Validate_Complex (nested structs)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 2.46 µs | 15 | baseline |
| Playground | 3.38 µs | 9 | 1.38x slower |
| Ozzo | 11.52 µs | 139 | 4.69x slower |
| Huma | - | - | - |
| Godantic | 12.95 µs | 120 | 5.28x slower |
| Godasse | - | - | - |

### JSONValidate_Simple (JSON → struct, then validate)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 3.31 µs | 19 | baseline |
| Playground | 4.11 µs | 16 | 1.24x slower |
| Ozzo | - | - | - |
| Huma | 3.34 µs | 26 | 1.01x slower |
| Godantic | - | - | - |
| Godasse | - | - | - |

### JSONValidate_Complex (nested JSON)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 9.38 µs | 39 | baseline |
| Playground | 10.30 µs | 33 | 1.10x slower |
| Ozzo | - | - | - |
| Huma | 9.91 µs | 78 | 1.06x slower |
| Godantic | - | - | - |
| Godasse | - | - | - |

### Unmarshal_Simple (JSON → validated struct, single call)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 4.75 µs | 41 | baseline |
| Playground | - | - | - |
| Ozzo | - | - | - |
| Huma | - | - | - |
| Godantic | 10.76 µs | 81 | 2.26x slower |
| Godasse | 4.40 µs | 42 | 1.08x faster |

### Unmarshal_Complex (nested JSON, single call)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 13.58 µs | 122 | baseline |
| Playground | - | - | - |
| Ozzo | - | - | - |
| Huma | - | - | - |
| Godantic | 45.70 µs | 285 | 3.37x slower |
| Godasse | 15.67 µs | 149 | 1.15x slower |

### Schema_Uncached (first-time generation)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 26.14 µs | 219 | baseline |
| Playground | - | - | - |
| Ozzo | - | - | - |
| Huma | 29.85 µs | 255 | 1.14x slower |
| Godantic | - | - | - |
| Godasse | - | - | - |

### Schema_Cached (cached lookup)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 21 ns | 0 | baseline |
| Playground | - | - | - |
| Ozzo | - | - | - |
| Huma | 575 ns | 6 | 27.03x slower |
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
