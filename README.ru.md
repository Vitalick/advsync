# advsync

[![Go Report Card](https://goreportcard.com/badge/github.com/vitalick/adv-sync)](https://goreportcard.com/report/github.com/vitalick/adv-sync)
[![GoDoc](https://godoc.org/github.com/vitalick/adv-sync?status.svg)](https://godoc.org/github.com/vitalick/adv-sync)

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
