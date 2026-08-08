---
sidebar_position: 99
title: Benchmarks
---

# Benchmark Results

Generated: 2026-08-08 01:43:13 UTC

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
| Simple | 547 ns (10 allocs) | 814 ns (7 allocs) | 5.72 µs (43 allocs) | unsupported | 2.27 µs (48 allocs) | unsupported |
| Complex | 944 ns (15 allocs) | 1.43 µs (9 allocs) | 5.12 µs (139 allocs) | unsupported | 8.26 µs (120 allocs) | unsupported |
| Large | 649 ns (22 allocs) | 790 ns (3 allocs) | 20.97 µs (254 allocs) | unsupported | 5.79 µs (126 allocs) | unsupported |

## JSONValidate
_JSON bytes → struct, then a separate validate step_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Simple | 1.37 µs (19 allocs) | 1.72 µs (16 allocs) | unsupported | 1.40 µs (26 allocs) | unsupported | unsupported |
| Complex | 4.24 µs (39 allocs) | 4.79 µs (33 allocs) | unsupported | 4.67 µs (78 allocs) | unsupported | unsupported |

## Marshal
_Validate + JSON marshal_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Simple | 705 ns (11 allocs) | 1.02 µs (9 allocs) | unsupported | unsupported | unsupported | unsupported |

## Unmarshal
_JSON bytes → validated struct in a single call_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Simple | 1.93 µs (39 allocs) | unsupported | unsupported | unsupported | 4.56 µs (81 allocs) | 1.85 µs (42 allocs) |
| Complex | 6.38 µs (122 allocs) | unsupported | unsupported | unsupported | 19.25 µs (285 allocs) | 6.93 µs (149 allocs) |

## New
_Validator creation overhead_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Simple | 6.59 µs (129 allocs) | 7.61 µs (187 allocs) | unsupported | 11.41 µs (243 allocs) | 10.29 µs (291 allocs) | 2.56 µs (72 allocs) |
| Complex | 15.88 µs (299 allocs) | unsupported | unsupported | 27.88 µs (499 allocs) | 2.78 µs (65 allocs) | 9.05 µs (243 allocs) |

## Schema
_JSON Schema generation_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Uncached | 11.20 µs (227 allocs) | unsupported | unsupported | 11.45 µs (243 allocs) | unsupported | unsupported |
| Cached | 9 ns (0 allocs) | unsupported | unsupported | 253 ns (6 allocs) | unsupported | unsupported |

## OpenAPI
_OpenAPI-compatible schema generation_

| Struct | Pedantigo | Playground | Ozzo | Huma | Godantic | Godasse |
|--------|--------|--------|--------|--------|--------|--------|
| Uncached | 11.61 µs (229 allocs) | unsupported | unsupported | 11.51 µs (243 allocs) | unsupported | unsupported |
| Cached | 9 ns (0 allocs) | unsupported | unsupported | 260 ns (6 allocs) | unsupported | unsupported |

---

## Summary

### Validate_Simple (struct validation)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 547 ns | 10 | baseline |
| Playground | 814 ns | 7 | 1.49x slower |
| Ozzo | 5.72 µs | 43 | 10.46x slower |
| Huma | - | - | - |
| Godantic | 2.27 µs | 48 | 4.16x slower |
| Godasse | - | - | - |

### Validate_Complex (nested structs)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 944 ns | 15 | baseline |
| Playground | 1.43 µs | 9 | 1.52x slower |
| Ozzo | 5.12 µs | 139 | 5.43x slower |
| Huma | - | - | - |
| Godantic | 8.26 µs | 120 | 8.75x slower |
| Godasse | - | - | - |

### JSONValidate_Simple (JSON → struct, then validate)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 1.37 µs | 19 | baseline |
| Playground | 1.72 µs | 16 | 1.25x slower |
| Ozzo | - | - | - |
| Huma | 1.40 µs | 26 | 1.02x slower |
| Godantic | - | - | - |
| Godasse | - | - | - |

### JSONValidate_Complex (nested JSON)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 4.24 µs | 39 | baseline |
| Playground | 4.79 µs | 33 | 1.13x slower |
| Ozzo | - | - | - |
| Huma | 4.67 µs | 78 | 1.10x slower |
| Godantic | - | - | - |
| Godasse | - | - | - |

### Unmarshal_Simple (JSON → validated struct, single call)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 1.93 µs | 39 | baseline |
| Playground | - | - | - |
| Ozzo | - | - | - |
| Huma | - | - | - |
| Godantic | 4.56 µs | 81 | 2.37x slower |
| Godasse | 1.85 µs | 42 | 1.04x faster |

### Unmarshal_Complex (nested JSON, single call)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 6.38 µs | 122 | baseline |
| Playground | - | - | - |
| Ozzo | - | - | - |
| Huma | - | - | - |
| Godantic | 19.25 µs | 285 | 3.01x slower |
| Godasse | 6.93 µs | 149 | 1.09x slower |

### Schema_Uncached (first-time generation)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 11.20 µs | 227 | baseline |
| Playground | - | - | - |
| Ozzo | - | - | - |
| Huma | 11.45 µs | 243 | 1.02x slower |
| Godantic | - | - | - |
| Godasse | - | - | - |

### Schema_Cached (cached lookup)

| Library | ns/op | allocs | vs Pedantigo |
|---------|-------|--------|-------------|
| Pedantigo | 9 ns | 0 | baseline |
| Playground | - | - | - |
| Ozzo | - | - | - |
| Huma | 253 ns | 6 | 28.50x slower |
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
