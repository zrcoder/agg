package main

import (
	"errors"
	"strings"

	_ "github.com/zrcoder/agg/internal/exported/github.com/zrcoder/agg/pkg/api/hanoi"
	"github.com/zrcoder/agg/internal/hanoi"

	"github.com/goplus/igop"
	_ "github.com/goplus/igop/gopbuild"
	_ "github.com/goplus/igop/gopbuild/pkg"
)

const (
	gopfileName = "main.gop"

	preCodes = `
	import (
		. "github.com/zrcoder/agg/pkg/api/hanoi"
	)

	var (
		a = A
		b = B
		c = C
	)

	type Pile = func()
	`
)

func runCode(code string) error {
	if strings.TrimSpace(code) == "" {
		return errors.New("code is empty")
	}
	hanoi.Hanoi.Reset()
	code = preCodes + code
	_, err := igop.RunFile(gopfileName, code, nil, 0)
	return err
}

func main() {
	hanoi.Run(runCode)
}
