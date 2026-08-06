# advsync

[![tests](https://github.com/vitalick/advsync/actions/workflows/test.yml/badge.svg)](https://github.com/vitalick/advsync/actions/workflows/test.yml)
[![golangci-lint](https://github.com/vitalick/advsync/actions/workflows/lint.yml/badge.svg)](https://github.com/vitalick/advsync/actions/workflows/lint.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/vitalick/advsync.svg)](https://pkg.go.dev/github.com/vitalick/advsync)
[![Go Version](https://img.shields.io/github/go-mod/go-version/vitalick/advsync)](go.mod)
[![License](https://img.shields.io/badge/license-GPL--2.0--only-blue.svg)](LICENSE)

[Русская версия](README.ru.md)

`advsync` provides small, keyed synchronization primitives for Go. It offers
mutexes, read/write mutexes, and semaphores backed by either a regular map or
`xsync.Map`.

## Installation

```bash
go get github.com/vitalick/advsync
```

## Usage

```go
package main

import "github.com/vitalick/advsync"

func main() {
	semaphore := advsync.NewSemaphore(2)
	semaphore.Acquire()
	defer func() { _ = semaphore.Release() }()

	// Work with at most two concurrent callers.
}
```

## Primitives

- `NamedMutex` and `NamedMutexSM` manage an independent mutex for every key.
- `NamedRWMutex` and `NamedRWMutexSM` provide a keyed read/write mutex.
- `Semaphore` uses `sync.Cond`; `SemaphoreChan` uses a buffered channel.
- `NamedSemaphore`, `NamedSemaphoreSM`, `NamedSemaphoreChan`, and
  `NamedSemaphoreChanSM` provide keyed semaphore variants.

Use the same key to acquire and release a named primitive. As with Go's
standard synchronization types, a lock must be released by code that holds it.

## Benchmarks

Run the benchmark suite with:

```bash
go test -run '^$' -bench . -benchmem
```

The following uncontended results were measured on Windows 11, AMD Ryzen 9
9950X3D, Go 1.26.5, `windows/amd64`.

| Benchmark | ns/op | B/op | allocs/op |
| --- | ---: | ---: | ---: |
| `NamedMutex` | 28.69 | 0 | 0 |
| `NamedMutexSM` | 31.59 | 16 | 2 |
| `NamedRWMutex` | 36.91 | 0 | 0 |
| `NamedRWMutexSM` | 42.27 | 48 | 2 |
| `Semaphore` | 17.67 | 0 | 0 |
| `SemaphoreChan` | 24.09 | 0 | 0 |

Results are a local baseline only. They vary by CPU, operating system, Go
version, workload, and contention level.

## Tests and linting

Run the test suite with:

```bash
go test ./...
```

Format and lint locally with:

```bash
golangci-lint fmt
golangci-lint run
```

GitHub Actions runs tests and golangci-lint on every push and pull request.

## License

This project is distributed under the [GPL-2.0 license](LICENSE).
