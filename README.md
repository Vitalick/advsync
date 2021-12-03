# Advanched sync
[![Go Report Card](https://goreportcard.com/badge/github.com/Vitalick/adv-sync)](https://goreportcard.com/report/github.com/Vitalick/adv-sync)
[![GoDoc](https://godoc.org/github.com/Vitalick/adv-sync?status.svg)](https://godoc.org/github.com/Vitalick/adv-sync)

Advanched sync package for Golang.

## NamedMutex
It's a multiple mutexes with lock and unlock by name implemented as `interface{}`. Uses `sync.RWMutex+map`.

## NamedRWMutex
It's a multiple read/write mutexes with lock, unlock, rw lock and rw unlock by name implemented as `interface{}`. Uses `sync.RWMutex+map`.

## NamedMutexSM
It's a multiple mutexes with lock and unlock by name implemented as `interface{}`. Uses `sync.Map`.

## NamedRWMutexSM
It's a multiple read/write mutexes with lock, unlock, rw lock and rw unlock by name implemented as `interface{}`. Uses `sync.Map`.

## Semaphore
It's semaphore primitive based by `sync.Cond`.
