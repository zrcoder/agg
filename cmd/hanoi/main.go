package main

import (
	hanoi "github.com/zrcoder/agg/internal/hanoi"
)

//go:generate igop export -outdir ./internal/exported ./internal/export.go

func main() {
	hanoi.Run()
}
