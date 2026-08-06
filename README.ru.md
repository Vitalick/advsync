# advsync

[![CI](https://github.com/vitalick/advsync/actions/workflows/ci.yml/badge.svg)](https://github.com/vitalick/advsync/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/vitalick/advsync.svg)](https://pkg.go.dev/github.com/vitalick/advsync)
[![Go Report Card](https://goreportcard.com/badge/github.com/vitalick/advsync)](https://goreportcard.com/report/github.com/vitalick/advsync)
[![Go Version](https://img.shields.io/github/go-mod/go-version/vitalick/advsync)](go.mod)
[![License](https://img.shields.io/badge/license-GPL--2.0--only-blue.svg)](LICENSE)

[English version](README.md)

`advsync` — небольшая Go-библиотека с синхронизационными примитивами по ключу.
В ней есть mutex, read/write mutex и семафоры, реализованные через обычную map
или `xsync.Map`.

## Установка

```bash
go get github.com/vitalick/advsync
```

## Использование

```go
package main

import "github.com/vitalick/advsync"

func main() {
	semaphore := advsync.NewSemaphore(2)
	semaphore.Acquire()
	defer func() { _ = semaphore.Release() }()

	// Работа не более чем для двух одновременных вызовов.
}
```

## Примитивы

- `NamedMutex` и `NamedMutexSM` создают независимый mutex для каждого ключа.
- `NamedRWMutex` и `NamedRWMutexSM` предоставляют mutex чтения/записи по ключу.
- `Semaphore` построен на `sync.Cond`, а `SemaphoreChan` — на буферизированном канале.
- `NamedSemaphore`, `NamedSemaphoreSM`, `NamedSemaphoreChan` и
  `NamedSemaphoreChanSM` — варианты семафоров по ключу.

Для получения и освобождения именованного примитива используйте один и тот же
ключ. Как и у стандартных примитивов синхронизации Go, освобождать блокировку
должен код, который её удерживает.

## Бенчмарки

Запустите набор бенчмарков:

```bash
go test -run '^$' -bench . -benchmem
```

Следующие результаты для сценариев без конкуренции получены в Windows 11, на
AMD Ryzen 9 9950X3D, Go 1.26.5, `windows/amd64`.

| Бенчмарк | ns/op | B/op | allocs/op |
| --- | ---: | ---: | ---: |
| `NamedMutex` | 28.69 | 0 | 0 |
| `NamedMutexSM` | 31.59 | 16 | 2 |
| `NamedRWMutex` | 36.91 | 0 | 0 |
| `NamedRWMutexSM` | 42.27 | 48 | 2 |
| `Semaphore` | 17.67 | 0 | 0 |
| `SemaphoreChan` | 24.09 | 0 | 0 |

Это только локальная базовая линия: результаты зависят от CPU, операционной
системы, версии Go, нагрузки и уровня конкуренции.

## Тесты и линтинг

Запустите тесты:

```bash
go test ./...
```

Отформатируйте и проверьте код локально:

```bash
golangci-lint fmt
golangci-lint run
```

GitHub Actions запускает тесты и golangci-lint при каждом push и pull request.

## Лицензия

Проект распространяется по [лицензии GPL-2.0](LICENSE).
