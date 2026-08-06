# advsync

[![Go Report Card](https://goreportcard.com/badge/github.com/vitalick/advsync)](https://goreportcard.com/report/github.com/vitalick/advsync)
[![GoDoc](https://godoc.org/github.com/vitalick/advsync?status.svg)](https://godoc.org/github.com/vitalick/advsync)

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
